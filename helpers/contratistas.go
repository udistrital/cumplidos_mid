package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/models"
)

//RFC 45758 Se agrega función para comparar los contratos de la dependencia con la informacion de los cumplidos solicitados
func contratoExists(vigencia string, numero string, contratos []struct {
	Vigencia       string "json:\"vigencia\""
	NumeroContrato string "json:\"numero_contrato\""
}) (result bool) {
	result = false
	for _, contrato := range contratos {
		if contrato.Vigencia == vigencia && contrato.NumeroContrato == numero {
			result = true
			break
		}
	}
	return result
}

func CertificacionCumplidosContratistas(dependencia string, mes string, anio string) (personas []models.Persona, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	//var contrato_dependencia models.ContratoDependencia
	var contratos_dependencia models.ContratoDependencia
	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var persona models.Persona

	var nmes, _ = strconv.Atoi(mes)
	var respuesta_peticion map[string]interface{}
	//traemos los contratos activos para un mes en una dependencia

	//RFC 45758 Se modifica la función que trae los contratos por dependencia a la funcion GetContratosDependenciaFiltro
	if contratos_dependencia, outputError = GetContratosDependenciaFiltro(dependencia, anio+"-"+mes, anio+"-"+mes); outputError != nil {
		return nil, outputError
	}

	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudCumplidos")+"/pago_mensual/?query=EstadoPagoMensualId.CodigoAbreviacion.in:AS|AP,Mes:"+strconv.Itoa(nmes)+",Ano:"+anio+"&limit=-1", &respuesta_peticion); (err == nil) && (response == 200) {

		pagos_mensuales = []models.PagoMensual{}
		LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
		for _, pago_mensual := range pagos_mensuales {

			//RFC 45758 Se modifica la condición para comparar los contratos de la dependencia con los cumpñlidos solicitados en el mes
			if contratoExists(strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64), pago_mensual.NumeroContrato, contratos_dependencia.Contratos.Contrato) {
				if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pago_mensual.DocumentoPersonaId, &contratistas); (err == nil) && (response == 200) {
					var contrato models.InformacionContrato
					contrato, outputError = GetContrato(pago_mensual.NumeroContrato, strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64))
					if outputError == nil {

						for _, contratista := range contratistas {
							persona.NumDocumento = contratista.NumDocumento
							persona.Nombre = contratista.NomProveedor
							persona.NumeroContrato = pago_mensual.NumeroContrato
							persona.Vigencia = int(pago_mensual.VigenciaContrato) //strconv.Atoi(cd.Vigencia)
							persona.NumeroCdp = pago_mensual.NumeroCDP            //RFC 50388
							persona.Rubro = contrato.Contrato.Rubro
							personas = append(personas, persona)
						}
					} else {
						return nil, outputError
					}

				} else { //If informacion_proveedor get
					logs.Error(err)
					outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas", "err": err.Error(), "status": "404"}
					return nil, outputError
				}

			} else {
			}

		}
	} else { //If pago_mensual get
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas", "err": err.Error(), "status": "404"}
		return nil, outputError
	}
	return
}

func AprobacionPagosContratistas(v []models.PagoContratistaCdpRp) (outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/AprobacionPagosContratistas", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var response interface{}
	var pagos_mensuales []*models.PagoMensual
	var pago_mensual *models.PagoMensual

	for _, pm := range v {

		pago_mensual = pm.PagoMensual
		pagos_mensuales = append(pagos_mensuales, pago_mensual)
	}
	if err := sendJson(beego.AppConfig.String("UrlCrudCumplidos")+"/tr_aprobacion_masiva_pagos", "POST", &response, pagos_mensuales); err != nil {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/AprobacionPagosContratistas", "err": err.Error(), "status": "502"}
		return outputError

	}
	return nil
}

func SolicitudesOrdenadorContratistas(doc_ordenador string, limit int, offset int) (pagos_contratista_cdp_rp []models.PagoContratistaCdpRp, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/SolicitudesOrdenadorContratistas0", "err": err, "status": "404"}
			panic(outputError)
		}
	}()

	var pagos_mensuales []models.PagoMensual
	// var contratistas []models.InformacionProveedor

	// var contratos_disponibilidad []models.ContratoDisponibilidad
	var respuesta_peticion map[string]interface{}
	fmt.Println(beego.AppConfig.String("UrlCrudCumplidos") + "/pago_mensual/?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset) + "&query=EstadoPagoMensualId.CodigoAbreviacion:AS,DocumentoResponsableId:" + doc_ordenador)
	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/pago_mensual/?limit="+strconv.Itoa(limit)+"&offset="+strconv.Itoa(offset)+"&query=EstadoPagoMensualId.CodigoAbreviacion:AS,DocumentoResponsableId:"+doc_ordenador, &respuesta_peticion); (err == nil) && (response == 200) {

		pagos_mensuales = []models.PagoMensual{}
		LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
		fmt.Println(pagos_mensuales)
		for _, pago_mensual := range pagos_mensuales {

			fmt.Println(pago_mensual)
			var pago_contratista_cdp_rp models.PagoContratistaCdpRp
			var outputError map[string]interface{}
			pago_contratista_cdp_rp, outputError = getInfoPagoMensual(pago_mensual)
			fmt.Println(pago_contratista_cdp_rp)
			fmt.Println(outputError)
			if outputError == nil {
				pagos_contratista_cdp_rp = append(pagos_contratista_cdp_rp, pago_contratista_cdp_rp)
			}
		}

	} else { //If pago_mensual get
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/SolicitudesOrdenadorContratistas", "err": err, "status": "502"}
		return nil, outputError
	}

	return pagos_contratista_cdp_rp, nil
}

func TraerInfoOrdenador(numero_contrato string, vigencia string) (informacion_ordenador models.InformacionOrdenador, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			outputError = map[string]interface{}{"funcion": "/TraerInfoOrdenador", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var temp map[string]interface{}
	var contrato_elaborado models.ContratoElaborado
	//var informacion_ordenador models.InformacionOrdenador
	var ordenadores []models.Ordenador

	fmt.Println(beego.AppConfig.String("UrlAdministrativaJBPM") + "/" + "contrato_elaborado/" + numero_contrato + "/" + vigencia)
	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlAdministrativaJBPM")+"/"+"contrato_elaborado/"+numero_contrato+"/"+vigencia, &temp); err == nil && temp != nil && response == 200 {
		json_contrato_elaborado, error_json := json.Marshal(temp)
		if error_json == nil {
			if err := json.Unmarshal(json_contrato_elaborado, &contrato_elaborado); err == nil {
				fmt.Println(contrato_elaborado)
				fecha := strings.Split(contrato_elaborado.Contrato.FechaRegistro, "+")
				fecha = strings.Split(fecha[0], "-")

				//RFC 45758 Se consulta el ordenador inmediatamente anterior a la fecha de registro del contrato
				if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/ordenadores/?query=IdOrdenador:"+contrato_elaborado.Contrato.OrdenadorGasto+",FechaInicio__lt:"+fecha[1]+"/"+fecha[2]+"/"+fecha[0]+"&sortby=FechaInicio&order=desc&limit=1", &ordenadores); (err == nil) && (response == 200) {

					for _, ordenador := range ordenadores {

						//RFC 45758 Se consulta el ordenador más reciente vinculado al rol obtenido con la consulta anterior
						if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/ordenadores/?query=RolOrdenador:"+strings.Replace(ordenador.RolOrdenador, " ", "%20", -1)+"&sortby=FechaInicio&order=desc&limit=1", &ordenadores); (err == nil) && (response == 200) {

							for _, ordenador := range ordenadores {
								informacion_ordenador.NumeroDocumento = ordenador.Documento
								informacion_ordenador.Cargo = ordenador.RolOrdenador
								informacion_ordenador.Nombre = ordenador.NombreOrdenador
								//c.Data["json"] = informacion_ordenador
							}
						} else {
							logs.Error(err)
							outputError = map[string]interface{}{"funcion": "/TraerInfoOrdenador/ordenadores", "err": err.Error(), "status": "502"}
							return informacion_ordenador, outputError
						}
					}

				} else {
					logs.Error(err)
					outputError = map[string]interface{}{"funcion": "/TraerInfoOrdenador/ordenadores", "err": err.Error(), "status": "502"}
					return informacion_ordenador, outputError
				}

			} else {
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/TraerInfoOrdenador/Unmarshal", "err": err.Error(), "status": "502"}
				return informacion_ordenador, outputError

			}
		} else {
			fmt.Println(error_json.Error())
			logs.Error(error_json.Error())
			outputError = map[string]interface{}{"funcion": "/TraerInfoOrdenador/Marshal", "err": error_json.Error(), "status": "502"}
			return informacion_ordenador, outputError
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/TraerInfoOrdenador/contrato_elaborado", "err": err.Error(), "status": "502"}
		return informacion_ordenador, outputError
	}
	return
}

func GetCumplidosRevertiblesPorOrdenador(NumDocumentoOrdenador string) (cumplidos_revertibles []models.PagoContratistaCdpRp, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetCumplidosRevertiblesPorOrdenador", "err": err, "status": "404"}
			panic(outputError)
		}
	}()

	var pagos_mensuales []models.PagoMensual
	// var contratistas []models.InformacionProveedor

	// var contratos_disponibilidad []models.ContratoDisponibilidad
	var respuesta_peticion map[string]interface{}
	fmt.Println(beego.AppConfig.String("UrlCrudCumplidos") + "/pago_mensual/?query=EstadoPagoMensualId.CodigoAbreviacion:AP,DocumentoResponsableId:" + NumDocumentoOrdenador)
	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/pago_mensual/?query=EstadoPagoMensualId.CodigoAbreviacion:AP,DocumentoResponsableId:"+NumDocumentoOrdenador, &respuesta_peticion); (err == nil) && (response == 200) {

		pagos_mensuales = []models.PagoMensual{}
		LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
		fmt.Println(pagos_mensuales)
		for _, pago_mensual := range pagos_mensuales {

			fmt.Println(pago_mensual)
			var cumplidos_revertible models.PagoContratistaCdpRp
			var outputError map[string]interface{}
			cumplidos_revertible, outputError = getInfoPagoMensual(pago_mensual)
			fmt.Println(cumplidos_revertible)
			fmt.Println(outputError)
			if outputError == nil {
				cumplidos_revertibles = append(cumplidos_revertibles, cumplidos_revertible)
			}
		}

	} else { //If pago_mensual get
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/SolicitudesOrdenadorContratistas", "err": err, "status": "502"}
		return nil, outputError
	}

	return cumplidos_revertibles, nil
}
