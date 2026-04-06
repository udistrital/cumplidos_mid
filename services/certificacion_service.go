package services

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
	"github.com/udistrital/cumplidos_mid/models"
)

// ICertificacionService interfaz para servicios de certificación
type ICertificacionService interface {
	GetCertificaciones(fechaInicio string, dependencia string, mes string, anio string) ([]models.Persona, map[string]interface{})
}

// CertificacionService implementación del servicio de certificación
type CertificacionService struct{}

// NewCertificacionService constructor
func NewCertificacionService() ICertificacionService {
	return &CertificacionService{}
}

// GetCertificaciones obtiene las certificaciones de cumplidos aprobados por supervisor
// usando el periodo consultado (mes/anio) para resolver contratos y fechaInicio para filtrar aprobaciones.
func (s *CertificacionService) GetCertificaciones(fechaInicio string, dependencia string, mes string, anio string) ([]models.Persona, map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			panic(map[string]interface{}{"funcion": "GetCertificaciones", "err": fmt.Sprintf("%v", err), "status": "500"})
		}
	}()

	// Validar fechaInicio
	fechaInicio = strings.TrimSpace(fechaInicio)
	if fechaInicio == "" {
		logs.Error("Fecha inválida: empty")
		return nil, map[string]interface{}{"funcion": "GetCertificaciones", "err": "Fecha inválida", "status": "400"}
	}

	contratosDependencia, err := s.getContratosDependencia(dependencia, fechaInicio, mes, anio)
	if err != nil {
		logs.Error("Error getting contratos: %v", err)
		return nil, err
	}
	pagosMensuales, err := s.getPagosMensualesPorContratos(contratosDependencia, mes, anio)
	if err != nil {
		logs.Error("Error getting pagos mensuales: %v", err)
		return nil, err
	}
	// Filtrar pagos válidos (descartar pagos con Id=0 o NumeroContrato vacío)
	var pagosValidos []models.PagoMensual
	for _, pago := range pagosMensuales {
		if pago.Id > 0 && strings.TrimSpace(pago.NumeroContrato) != "" {
			pagosValidos = append(pagosValidos, pago)
		}
	}

	// Si no hay pagos válidos, intentar buscar cambios de estado aprobados directamente
	if len(pagosValidos) == 0 {
		return []models.Persona{}, nil
	}

	// Extraer IDs de pagos filtrados
	pagoMensualIds := s.extractPagoMensualIdsFromPagos(pagosValidos)

	// Obtener cambios de estado aprobados desde fecha para esos pagos
	cambiosEstado, err := s.getCambiosEstadoAprobadosDesdeFechaParaPagos(fechaInicio, pagoMensualIds)
	if err != nil {
		logs.Error("Error getting cambios estado: %v", err)
		return nil, err
	}
	// Extraer IDs de pagos que tienen cambios aprobados desde fecha
	pagoMensualIdsAprobados := s.extractPagoMensualIds(cambiosEstado)

	// Filtrar pagos finales por IDs aprobados
	pagosFinales := s.filtrarPagosPorIds(pagosValidos, pagoMensualIdsAprobados)

	// Obtener personas de los pagos finales
	personas, err := s.getPersonasFromPagos(pagosFinales)
	if err != nil {
		logs.Error("Error getting personas: %v", err)
		return nil, err
	}
	return personas, nil
}

// getPagosMensualesPorContratos obtiene pagos mensuales para los contratos dados
func (s *CertificacionService) getPagosMensualesPorContratos(contratosDependencia models.ContratoDependencia, mes string, anio string) ([]models.PagoMensual, map[string]interface{}) {
	// Extraer contratos únicos
	numsMap := make(map[string]bool)
	vigsMap := make(map[string]bool)
	for _, contrato := range contratosDependencia.Contratos.Contrato {
		numsMap[contrato.NumeroContrato] = true
		vigsMap[contrato.Vigencia] = true
	}

	var nums []string
	for num := range numsMap {
		nums = append(nums, num)
	}
	var vigs []string
	for vig := range vigsMap {
		vigs = append(vigs, vig)
	}

	if len(nums) == 0 || len(vigs) == 0 {
		return []models.PagoMensual{}, nil
	}

	numsQuery := s.joinIds(nums)
	vigsQuery := s.joinIds(vigs)
	query := "NumeroContrato.in:" + numsQuery + ",VigenciaContrato.in:" + vigsQuery + ",EstadoPagoMensualId.CodigoAbreviacion.in:AS|AP"

	if strings.TrimSpace(mes) != "" {
		mesInt, err := strconv.Atoi(strings.TrimSpace(mes))
		if err != nil || mesInt < 1 || mesInt > 12 {
			return nil, map[string]interface{}{"funcion": "getPagosMensualesPorContratos", "err": "Mes inválido", "status": "400"}
		}
		query += ",Mes:" + strconv.Itoa(mesInt)
	}

	if strings.TrimSpace(anio) != "" {
		anioInt, err := strconv.Atoi(strings.TrimSpace(anio))
		if err != nil || anioInt < 1 {
			return nil, map[string]interface{}{"funcion": "getPagosMensualesPorContratos", "err": "Año inválido", "status": "400"}
		}
		query += ",Ano:" + strconv.Itoa(anioInt)
	}

	url := beego.AppConfig.String("UrlCrudCumplidos") + "/pago_mensual/?query=" + query + "&limit=-1"
	var respuestaPeticion map[string]interface{}
	var pagosMensuales []models.PagoMensual

	if response, err := helpers.GetJsonTest(url, &respuestaPeticion); (err == nil) && (response == 200) {
		helpers.LimpiezaRespuestaRefactor(respuestaPeticion, &pagosMensuales)
		return pagosMensuales, nil
	} else {
		logs.Error("Error en getPagosMensualesPorContratos, response: %d, err: %v", response, err)
		errMsg := "Server error"
		if err != nil {
			errMsg = err.Error()
		}
		return nil, map[string]interface{}{"funcion": "getPagosMensualesPorContratos", "err": errMsg, "status": "500"}
	}
}

// extractPagoMensualIdsFromPagos extrae IDs únicos de pagos mensuales
func (s *CertificacionService) extractPagoMensualIdsFromPagos(pagos []models.PagoMensual) []string {
	idsMap := make(map[string]bool)
	var ids []string
	for _, pago := range pagos {
		idStr := strconv.Itoa(pago.Id)
		if !idsMap[idStr] {
			idsMap[idStr] = true
			ids = append(ids, idStr)
		}
	}
	return ids
}

// getCambiosEstadoAprobadosDesdeFechaParaPagos obtiene cambios de estado a aprobado supervisor desde fecha para pagos específicos
func (s *CertificacionService) getCambiosEstadoAprobadosDesdeFechaParaPagos(fechaInicio string, pagoMensualIds []string) ([]models.CambioEstadoPago, map[string]interface{}) {
	parsed, err := time.Parse("2006-01-02", fechaInicio)
	if err != nil {
		parsed, err = time.Parse(time.RFC3339, fechaInicio)
		if err != nil {
			return nil, map[string]interface{}{"funcion": "getCambiosEstadoAprobadosDesdeFechaParaPagos", "err": "Fecha inválida", "status": "400"}
		}
	}
	queryFecha := parsed.UTC().Format("2006-01-02T15:04:05Z")
	var cambiosEstado []models.CambioEstadoPago

	for _, pagoMensualID := range pagoMensualIds {
		query := "PagoMensualId:" + pagoMensualID + ",Activo:true,FechaCreacion__gte:" + queryFecha
		url := beego.AppConfig.String("UrlCrudCumplidos") + "/cambio_estado_pago/?query=" + query + "&limit=-1"

		var respuestaPeticion map[string]interface{}
		if response, err := helpers.GetJsonTest(url, &respuestaPeticion); (err == nil) && (response == 200) {
			data, ok := respuestaPeticion["Data"].([]interface{})
			if !ok || len(data) == 0 {
				continue
			}

			for _, rawCambio := range data {
				cambioMap, ok := rawCambio.(map[string]interface{})
				if !ok || len(cambioMap) == 0 {
					continue
				}

				estadoVal, hasEstado := cambioMap["EstadoPagoMensualId"]
				if hasEstado {
					switch estado := estadoVal.(type) {
					case map[string]interface{}:
						if codigo, ok := estado["CodigoAbreviacion"].(string); ok && codigo != "AS" {
							continue
						}
					case float64:
						if int(estado) != 13 {
							continue
						}
					}
				}

				cambio := models.CambioEstadoPago{}
				cambio.Activo = true
				cambio.PagoMensualId.Id, _ = strconv.Atoi(pagoMensualID)

				if id, ok := cambioMap["Id"].(float64); ok {
					cambio.Id = int(id)
				}
				if documento, ok := cambioMap["DocumentoResponsableId"].(string); ok {
					cambio.DocumentoResponsableId = documento
				}
				if cargo, ok := cambioMap["CargoResponsable"].(string); ok {
					cambio.CargoResponsable = cargo
				}
				if fechaCreacion, ok := cambioMap["FechaCreacion"].(string); ok {
					if parsedFecha, err := time.Parse(time.RFC3339Nano, fechaCreacion); err == nil {
						cambio.FechaCreacion = parsedFecha
					}
				}
				if fechaModificacion, ok := cambioMap["FechaModificacion"].(string); ok {
					if parsedFecha, err := time.Parse(time.RFC3339Nano, fechaModificacion); err == nil {
						cambio.FechaModificacion = parsedFecha
					}
				}

				if pagoMap, ok := cambioMap["PagoMensualId"].(map[string]interface{}); ok {
					if id, ok := pagoMap["Id"].(float64); ok {
						cambio.PagoMensualId.Id = int(id)
					}
					if numeroContrato, ok := pagoMap["NumeroContrato"].(string); ok {
						cambio.PagoMensualId.NumeroContrato = numeroContrato
					}
					if vigencia, ok := pagoMap["VigenciaContrato"].(float64); ok {
						cambio.PagoMensualId.VigenciaContrato = int(vigencia)
					}
					if mes, ok := pagoMap["Mes"].(float64); ok {
						cambio.PagoMensualId.Mes = int(mes)
					}
					if documento, ok := pagoMap["DocumentoPersonaId"].(string); ok {
						cambio.PagoMensualId.DocumentoPersonaId = documento
					}
					if ano, ok := pagoMap["Ano"].(float64); ok {
						cambio.PagoMensualId.Ano = int(ano)
					}
					if numeroCDP, ok := pagoMap["NumeroCDP"].(string); ok {
						cambio.PagoMensualId.NumeroCDP = numeroCDP
					}
					if vigenciaCDP, ok := pagoMap["VigenciaCDP"].(float64); ok {
						cambio.PagoMensualId.VigenciaCDP = int(vigenciaCDP)
					}
				}

				cambiosEstado = append(cambiosEstado, cambio)
			}
		} else {
			logs.Error("Error en getCambiosEstadoAprobadosDesdeFechaParaPagos para pago %s, response: %d, err: %v", pagoMensualID, response, err)
		}
	}
	return cambiosEstado, nil
}

// extractPagoMensualIds extrae IDs únicos de pagos mensuales
func (s *CertificacionService) extractPagoMensualIds(cambios []models.CambioEstadoPago) []string {
	idsMap := make(map[string]bool)
	var ids []string
	for _, cambio := range cambios {
		idStr := strconv.Itoa(cambio.PagoMensualId.Id)
		if !idsMap[idStr] {
			idsMap[idStr] = true
			ids = append(ids, idStr)
		}
	}
	return ids
}

// getPagosMensuales obtiene pagos mensuales por IDs
func (s *CertificacionService) getPagosMensuales(ids []string) ([]models.PagoMensual, map[string]interface{}) {
	if len(ids) == 0 {
		return []models.PagoMensual{}, nil
	}

	query := "Id.in:" + s.joinIds(ids)
	var respuestaPeticion map[string]interface{}
	var pagosMensuales []models.PagoMensual

	if response, err := helpers.GetJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/pago_mensual/?query="+query+"&limit=-1", &respuestaPeticion); (err == nil) && (response == 200) {
		helpers.LimpiezaRespuestaRefactor(respuestaPeticion, &pagosMensuales)
		return pagosMensuales, nil
	} else {
		logs.Error(err)
		return nil, map[string]interface{}{"funcion": "getPagosMensuales", "err": err.Error(), "status": "404"}
	}
}

// joinIds une IDs con |
func (s *CertificacionService) joinIds(ids []string) string {
	result := ""
	for i, id := range ids {
		if i > 0 {
			result += "|"
		}
		result += id
	}
	return result
}

// getPersonasFromPagos obtiene info de personas de los pagos
func (s *CertificacionService) getPersonasFromPagos(pagos []models.PagoMensual) ([]models.Persona, map[string]interface{}) {
	var personas []models.Persona
	for _, pago := range pagos {
		persona, err := s.getPersonaFromPago(pago)
		if err != nil {
			logs.Error("Error getting persona for pago ID %d: %v", pago.Id, err)
			return nil, err
		}
		personas = append(personas, persona)
	}
	return personas, nil
}

// getPersonaFromPago obtiene persona de un pago
func (s *CertificacionService) getPersonaFromPago(pago models.PagoMensual) (models.Persona, map[string]interface{}) {
	var contratistas []models.InformacionProveedor

	if response, err := helpers.GetJsonTest(beego.AppConfig.String("UrlCrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pago.DocumentoPersonaId, &contratistas); (err == nil) && (response == 200) {
		contrato, err := helpers.GetContrato(pago.NumeroContrato, strconv.FormatFloat(pago.VigenciaContrato, 'f', 0, 64))
		if err != nil {
			logs.Error("Error getting contrato: %v", err)
			return models.Persona{}, err
		}

		for _, contratista := range contratistas {
			persona := models.Persona{
				NumDocumento:   contratista.NumDocumento,
				Nombre:         contratista.NomProveedor,
				NumeroContrato: pago.NumeroContrato,
				Vigencia:       int(pago.VigenciaContrato),
				NumeroCdp:      pago.NumeroCDP,
				Rubro:          contrato.Contrato.Rubro,
			}
			return persona, nil
		}
	} else {
		logs.Error("Error getting contratista, response: %d, err: %v", response, err)
		errMsg := "Server error"
		if err != nil {
			errMsg = err.Error()
		}
		return models.Persona{}, map[string]interface{}{"funcion": "getPersonaFromPago", "err": errMsg, "status": "500"}
	}
	return models.Persona{}, map[string]interface{}{"funcion": "getPersonaFromPago", "err": "No se encontró contratista", "status": "404"}
}

// getContratosDependencia obtiene contratos de una dependencia para el periodo consultado.
func (s *CertificacionService) getContratosDependencia(dependencia string, fechaInicio string, mes string, anio string) (models.ContratoDependencia, map[string]interface{}) {
	var fechaPeriodo time.Time

	if strings.TrimSpace(mes) != "" && strings.TrimSpace(anio) != "" {
		mesInt, err := strconv.Atoi(strings.TrimSpace(mes))
		if err != nil || mesInt < 1 || mesInt > 12 {
			logs.Error("Invalid mes in getContratosDependencia: %s", mes)
			return models.ContratoDependencia{}, map[string]interface{}{"funcion": "getContratosDependencia", "err": "Mes inválido", "status": "400"}
		}

		anioInt, err := strconv.Atoi(strings.TrimSpace(anio))
		if err != nil || anioInt < 1 {
			logs.Error("Invalid anio in getContratosDependencia: %s", anio)
			return models.ContratoDependencia{}, map[string]interface{}{"funcion": "getContratosDependencia", "err": "Año inválido", "status": "400"}
		}

		fechaPeriodo, err = time.Parse("2006-01", fmt.Sprintf("%04d-%02d", anioInt, mesInt))
		if err != nil {
			logs.Error("Invalid periodo in getContratosDependencia: anio=%s, mes=%s", anio, mes)
			return models.ContratoDependencia{}, map[string]interface{}{"funcion": "getContratosDependencia", "err": "Mes o año inválido", "status": "400"}
		}
	} else {
		var err error
		fechaPeriodo, err = time.Parse("2006-01-02", fechaInicio)
		if err != nil {
			fechaPeriodo, err = time.Parse(time.RFC3339, fechaInicio)
			if err != nil {
				logs.Error("Invalid date in getContratosDependencia: %s", fechaInicio)
				return models.ContratoDependencia{}, map[string]interface{}{"funcion": "getContratosDependencia", "err": "Fecha inválida", "status": "400"}
			}
		}
	}

	fechaConsulta := fechaPeriodo.Format("2006-01")

	contratosDependencia, outputError := helpers.GetContratosDependenciaFiltro(dependencia, fechaConsulta, fechaConsulta)
	if outputError != nil {
		logs.Error("Error in GetContratosDependenciaFiltro: %v", outputError)
		return models.ContratoDependencia{}, outputError
	}
	if contratosDependencia.Contratos.Contrato == nil {
		contratosDependencia.Contratos.Contrato = []struct {
			Vigencia       string `json:"vigencia"`
			NumeroContrato string `json:"numero_contrato"`
		}{}
	}
	return contratosDependencia, nil
}

// filtrarPagosPorContratos filtra pagos que pertenezcan a contratos de la dependencia
func (s *CertificacionService) filtrarPagosPorContratos(pagos []models.PagoMensual, contratosDependencia models.ContratoDependencia) []models.PagoMensual {
	var pagosFiltrados []models.PagoMensual
	for _, pago := range pagos {
		if s.contratoExists(strconv.FormatFloat(pago.VigenciaContrato, 'f', 0, 64), pago.NumeroContrato, contratosDependencia.Contratos.Contrato) {
			pagosFiltrados = append(pagosFiltrados, pago)
		}
	}
	return pagosFiltrados
}

// filtrarPagosPorIds filtra pagos por IDs
func (s *CertificacionService) filtrarPagosPorIds(pagos []models.PagoMensual, ids []string) []models.PagoMensual {
	idsMap := make(map[string]bool)
	for _, id := range ids {
		idsMap[id] = true
	}
	var pagosFiltrados []models.PagoMensual
	for _, pago := range pagos {
		idStr := strconv.Itoa(pago.Id)
		if idsMap[idStr] {
			pagosFiltrados = append(pagosFiltrados, pago)
		}
	}
	return pagosFiltrados
}

// contratoExists verifica si un contrato existe en la lista
func (s *CertificacionService) contratoExists(vigencia string, numero string, contratos []struct {
	Vigencia       string `json:"vigencia"`
	NumeroContrato string `json:"numero_contrato"`
}) bool {
	for _, contrato := range contratos {
		if contrato.NumeroContrato == numero && contrato.Vigencia == vigencia {
			return true
		}
	}
	return false
}

// getCambiosEstadoAprobadosDesdeFecha obtiene todos los cambios de estado aprobados desde una fecha (sin filtro de pago específico)
func (s *CertificacionService) getCambiosEstadoAprobadosDesdeFecha(fechaInicio string) ([]models.CambioEstadoPago, map[string]interface{}) {
	parsed, err := time.Parse("2006-01-02", fechaInicio)
	if err != nil {
		parsed, err = time.Parse(time.RFC3339, fechaInicio)
		if err != nil {
			return nil, map[string]interface{}{"funcion": "getCambiosEstadoAprobadosDesdeFecha", "err": "Fecha inválida", "status": "400"}
		}
	}
	queryFecha := parsed.UTC().Format("2006-01-02T15:04:05Z")
	query := "EstadoPagoMensualId.CodigoAbreviacion:AS,Activo:true,FechaCreacion__gte:" + queryFecha
	url := beego.AppConfig.String("UrlCrudCumplidos") + "/cambio_estado_pago/?query=" + query + "&limit=-1"
	var respuestaPeticion map[string]interface{}
	var cambiosEstado []models.CambioEstadoPago

	if response, err := helpers.GetJsonTest(url, &respuestaPeticion); (err == nil) && (response == 200) {
		helpers.LimpiezaRespuestaRefactor(respuestaPeticion, &cambiosEstado)
		return cambiosEstado, nil
	} else {
		logs.Error("[CAMBIOS_DIRECTO] Error en getCambiosEstadoAprobadosDesdeFecha, response: %d, err: %v", response, err)
		errMsg := "Server error"
		if err != nil {
			errMsg = err.Error()
		}
		return nil, map[string]interface{}{"funcion": "getCambiosEstadoAprobadosDesdeFecha", "err": errMsg, "status": "500"}
	}
}

// filtrarCambiosEstadoPorDependencia filtra cambios de estado para que solo incluyan contratos de la dependencia
func (s *CertificacionService) filtrarCambiosEstadoPorDependencia(cambios []models.CambioEstadoPago, contratosDependencia models.ContratoDependencia) []models.CambioEstadoPago {
	var cambiosFiltrados []models.CambioEstadoPago

	for _, cambio := range cambios {
		contrato := cambio.PagoMensualId.NumeroContrato
		vigencia := strconv.Itoa(cambio.PagoMensualId.VigenciaContrato)
		if s.contratoExists(vigencia, contrato, contratosDependencia.Contratos.Contrato) {
			cambiosFiltrados = append(cambiosFiltrados, cambio)
		}
	}
	return cambiosFiltrados
}

// extractPagoMensualIdsFromCambios extrae IDs de pagos mensuales desde cambios de estado
func (s *CertificacionService) extractPagoMensualIdsFromCambios(cambios []models.CambioEstadoPago) []string {
	idsMap := make(map[string]bool)
	var ids []string
	for _, cambio := range cambios {
		idStr := strconv.Itoa(cambio.PagoMensualId.Id)
		if !idsMap[idStr] {
			idsMap[idStr] = true
			ids = append(ids, idStr)
		}
	}
	return ids
}

// getPagosMensualesById obtiene pagos mensuales por sus IDs
func (s *CertificacionService) getPagosMensualesById(ids []string) ([]models.PagoMensual, map[string]interface{}) {
	if len(ids) == 0 {
		return []models.PagoMensual{}, nil
	}

	query := "Id.in:" + s.joinIds(ids)
	url := beego.AppConfig.String("UrlCrudCumplidos") + "/pago_mensual/?query=" + query + "&limit=-1"
	var respuestaPeticion map[string]interface{}
	var pagosMensuales []models.PagoMensual

	if response, err := helpers.GetJsonTest(url, &respuestaPeticion); (err == nil) && (response == 200) {
		helpers.LimpiezaRespuestaRefactor(respuestaPeticion, &pagosMensuales)
		return pagosMensuales, nil
	} else {
		logs.Error("[PAGOS_BY_ID] Error, response: %d, err: %v", response, err)
		errMsg := "Server error"
		if err != nil {
			errMsg = err.Error()
		}
		return nil, map[string]interface{}{"funcion": "getPagosMensualesById", "err": errMsg, "status": "500"}
	}
}
