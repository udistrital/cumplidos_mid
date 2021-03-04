package helpers

import (
	_ "encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

func CertificacionDocumentosAprobados(dependencia string, anio string, mes string) (personas []models.Persona, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/CertificacionDocumentosAprobados3", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var contrato_ordenador_dependencia models.ContratoOrdenadorDependencia
	var pagos_mensuales []models.PagoMensual
	var persona models.Persona
	var vinculaciones_docente []models.VinculacionDocente
	var respuesta_peticion map[string]interface{}
	var mes_cer, _ = strconv.Atoi(mes)

	if mes_cer < 10 {

		mes = "0" + mes

	}

	if contrato_ordenador_dependencia, outputError = GetContratosOrdenadorDependencia(dependencia, anio+"-"+mes, anio+"-"+mes); outputError != nil {
		outputError["funcion"] = "/CertificacionDocumentosAprobados" + fmt.Sprintf("%v", outputError["funcion"])
		return nil, outputError
	}

	for _, contrato := range contrato_ordenador_dependencia.ContratosOrdenadorDependencia.InformacionContratos {

		if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAdmin")+"/vinculacion_docente/?limit=-1&query=NumeroContrato:"+contrato.NumeroContrato+",Vigencia:"+contrato.Vigencia, &vinculaciones_docente); (err == nil) && (response == 200) {

			for _, vinculacion_docente := range vinculaciones_docente {
				if vinculacion_docente.NumeroContrato.Valid == true {
					if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/pago_mensual/?query=EstadoPagoMensualId.CodigoAbreviacion:AP,NumeroContrato:"+contrato.NumeroContrato+",VigenciaContrato:"+contrato.Vigencia+",Mes:"+strconv.Itoa(mes_cer)+",Ano:"+anio, &respuesta_peticion); (err == nil) && (response == 200) {
						pagos_mensuales = []models.PagoMensual{}
						if len(respuesta_peticion["Data"].([]interface{})[0].(map[string]interface{})) != 0 {
							LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
						} else {
							pagos_mensuales = nil
						}
						if pagos_mensuales == nil {
							persona.NumDocumento = contrato.Documento
							persona.Nombre = contrato.NombreContratista
							persona.NumeroContrato = contrato.NumeroContrato
							persona.Vigencia, _ = strconv.Atoi(contrato.Vigencia)
							personas = append(personas, persona)
						}
					} else { //If informacion_proveedor get
						logs.Error(err)
						outputError = map[string]interface{}{"funcion": "/CertificacionDocumentosAprobados1", "err": err.Error(), "status": "502"}
						return nil, outputError
					}
				}
			}
		} else { //If vinculacion_docente get
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/CertificacionDocumentosAprobados2", "err": err.Error(), "status": "502"}
			return nil, outputError
		}
	}
	return
}

func CertificadoVistoBueno(dependencia string, mes string, anio string) (personas []models.Persona, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/CertificadoVistoBueno", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var vinculaciones_docente []models.VinculacionDocente
	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var persona models.Persona
	var actasInicio []models.ActaInicio
	var mes_cer, _ = strconv.Atoi(mes)
	var anio_cer, _ = strconv.Atoi(anio)
	var respuesta_peticion map[string]interface{}
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAdmin")+"/vinculacion_docente/?limit=-1&query=IdProyectoCurricular:"+dependencia, &vinculaciones_docente); (err == nil) && (response == 200) {
		for _, vinculacion_docente := range vinculaciones_docente {
			if vinculacion_docente.NumeroContrato.Valid == true {
				if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/acta_inicio/?query=NumeroContrato:"+vinculacion_docente.NumeroContrato.String+",Vigencia:"+strconv.FormatInt(vinculacion_docente.Vigencia.Int64, 10), &actasInicio); (err == nil) && (response == 200) {
					for _, actaInicio := range actasInicio {
						//If Estado = 4
						if int(actaInicio.FechaInicio.Month()) <= mes_cer && actaInicio.FechaInicio.Year() <= anio_cer && int(actaInicio.FechaFin.Month()) >= mes_cer && actaInicio.FechaFin.Year() >= anio_cer {
							if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/pago_mensual/?query=EstadoPagoMensualId.CodigoAbreviacion.in:PAD|AD|AP,NumeroContrato:"+vinculacion_docente.NumeroContrato.String+",VigenciaContrato:"+strconv.FormatInt(vinculacion_docente.Vigencia.Int64, 10)+",Mes:"+mes+",Ano:"+anio, &respuesta_peticion); (err == nil) && (response == 200) {
								pagos_mensuales = []models.PagoMensual{}
								if len(respuesta_peticion["Data"].([]interface{})[0].(map[string]interface{})) != 0 {
									LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
								} else {
									pagos_mensuales = nil
								}
								if pagos_mensuales == nil {
									if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+vinculacion_docente.IdPersona, &contratistas); (err == nil) && (response == 200) {
										for _, contratista := range contratistas {
											persona.NumDocumento = contratista.NumDocumento
											persona.Nombre = contratista.NomProveedor
											persona.NumeroContrato = actaInicio.NumeroContrato
											persona.Vigencia = actaInicio.Vigencia
											personas = append(personas, persona)
										}
									} else { //If informacion_proveedor get
										logs.Error(err)
										outputError = map[string]interface{}{"funcion": "/CertificadoVistoBueno1", "err": err, "status": "502"}
										return nil, outputError
									}

								}

							} else { //If pago_mensual get
								logs.Error(err)
								outputError = map[string]interface{}{"funcion": "/CertificadoVistoBueno2", "err": err.Error(), "status": "502"}
								return nil, outputError
							}
						}
					}
				} else { //If contrato_estado get
					logs.Error(err)
					outputError = map[string]interface{}{"funcion": "/CertificadoVistoBueno3", "err": err.Error(), "status": "502"}
					return nil, outputError
				}
			}
		}

	} else { //If vinculacion_docente get
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/CertificadoVistoBueno4", "err": err.Error(), "status": "502"}
		return nil, outputError
	}
	return
}
