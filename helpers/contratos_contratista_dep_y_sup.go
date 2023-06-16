package helpers

import (
	_ "encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/models"
)

func ContratosContratistaDependencia(doc_ordenador string, cod_dependencia string, limit int, offset int) (pagos_contratista_cdp_rp []models.PagoContratistaCdpRp, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/ContratosContratistaDependencia", "err": err, "status": "404"}
			panic(outputError)
		}
	}()

	var contrato_dependencia models.ContratoDependencia
	var contratistas []models.InformacionProveedor
	var conteo_offset, conteo_limit = 0, 0
	var contratos_disponibilidad []models.ContratoDisponibilidad
	var respuesta_peticion map[string]interface{}

	t := time.Now()
	fecha_inicio := fmt.Sprintf("%d-%02d",
		t.Year(), t.Month())

	t2 := t.AddDate(-1, 0, 0)
	fecha_final := fmt.Sprintf("%d-%02d",
		t2.Year(), t2.Month())

	if contrato_dependencia, outputError = GetContratosDependenciaFiltro(cod_dependencia, fecha_inicio, fecha_final); outputError != nil {
		return nil, outputError
	}
	for _, cd := range contrato_dependencia.Contratos.Contrato {

		var pagos_mensuales []models.PagoMensual

		if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/pago_mensual/?query=NumeroContrato:"+cd.NumeroContrato+",VigenciaContrato:"+cd.Vigencia+",EstadoPagoMensualId.CodigoAbreviacion:AS,DocumentoResponsableId:"+doc_ordenador, &respuesta_peticion); (err == nil) && (response == 200) {
			pagos_mensuales = []models.PagoMensual{}
			LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
			for v, pm := range pagos_mensuales {
				if conteo_limit == limit && conteo_offset == offset {
					break
				}
				var pago_mensual models.PagoMensual

				if pm.NumeroContrato != "" {
					if conteo_offset == offset && conteo_limit < limit {
						pago_mensual.DocumentoPersonaId = pm.DocumentoPersonaId
						pago_mensual.VigenciaContrato = pm.VigenciaContrato
						pago_mensual.NumeroContrato = pm.NumeroContrato

						if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pago_mensual.DocumentoPersonaId, &contratistas); (err == nil) && (response == 200) {
							for _, contratista := range contratistas {

								var informacion_contrato_contratista models.InformacionContratoContratista
								informacion_contrato_contratista, outputError = GetInformacionContratoContratista(pago_mensual.NumeroContrato, strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64))
								if outputError == nil {

									var contrato models.InformacionContrato
									contrato, outputError = GetContrato(pago_mensual.NumeroContrato, strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64))
									if outputError == nil {

										if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+contrato.Contrato.NumeroContrato+",Vigencia:"+contrato.Contrato.Vigencia, &contratos_disponibilidad); (err == nil) && (response == 200) {

											for _, contrato_disponibilidad := range contratos_disponibilidad {

												var cdprp models.InformacionCdpRp
												cdprp, outputError = GetRP(strconv.Itoa(contrato_disponibilidad.NumeroCdp), strconv.Itoa(contrato_disponibilidad.VigenciaCdp))
												if outputError == nil {
													for _, rp := range cdprp.CdpXRp.CdpRp {
														var pago_contratista_cdp_rp models.PagoContratistaCdpRp
														pago_contratista_cdp_rp.PagoMensual = &pagos_mensuales[v]
														pago_contratista_cdp_rp.NombreDependencia = informacion_contrato_contratista.InformacionContratista.Dependencia
														pago_contratista_cdp_rp.NombrePersona = contratista.NomProveedor
														pago_contratista_cdp_rp.NumeroCdp = strconv.Itoa(contrato_disponibilidad.NumeroCdp)
														pago_contratista_cdp_rp.VigenciaCdp = strconv.Itoa(contrato_disponibilidad.VigenciaCdp)
														pago_contratista_cdp_rp.NumeroRp = rp.RpNumeroRegistro
														pago_contratista_cdp_rp.VigenciaRp = rp.RpVigencia
														pago_contratista_cdp_rp.Rubro = contrato.Contrato.Rubro
														pagos_contratista_cdp_rp = append(pagos_contratista_cdp_rp, pago_contratista_cdp_rp)
													}
												} else {
													return nil, outputError
												}

											}

										} else { // If contrato_disponibilidad get
											logs.Error(err)
											outputError = map[string]interface{}{"funcion": "/ContratosContratistaDependencia", "err": err.Error(), "status": "502"}
											return nil, outputError
										}
									} else {
										return nil, outputError
									}
								} else {
									return nil, outputError
								}

							}
							conteo_limit = conteo_limit + 1
						} else { //If informacion_proveedor get
							logs.Error(err)
							outputError = map[string]interface{}{"funcion": "/ContratosContratistaDependencia", "err": err.Error(), "status": "502"}
							return nil, outputError
						}

					} else {
						conteo_offset = conteo_offset + 1
					}
				}
			}

		} else { //If pago_mensual get
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/ContratosContratistaDependencia", "err": err.Error(), "status": "502"}
			return nil, outputError
		}

	}
	return
}

func ContratosContratistaSupervisor(doc_supervisor string) (pagos_contratista_cdp_rp []models.PagoContratistaCdpRp, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/ContratosContratistaSupervisor", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var pagos_mensuales []models.PagoMensual

	var respuesta_peticion map[string]interface{}

	fmt.Println(beego.AppConfig.String("UrlCrudCumplidos") + "/pago_mensual/?limit=-1&query=EstadoPagoMensualId.CodigoAbreviacion:PRS,DocumentoResponsableId:" + doc_supervisor)
	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/pago_mensual/?limit=-1&query=EstadoPagoMensualId.CodigoAbreviacion:PRS,DocumentoResponsableId:"+doc_supervisor, &respuesta_peticion); (err == nil) && (response == 200) {
		LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
		for _, pago_mensual := range pagos_mensuales {
			var pago_contratista_cdp_rp models.PagoContratistaCdpRp
			var outputError map[string]interface{}
			pago_contratista_cdp_rp, outputError = getInfoPagoMensual(pago_mensual)
			if outputError == nil {
				pagos_contratista_cdp_rp = append(pagos_contratista_cdp_rp, pago_contratista_cdp_rp)
			}
		}

	} else { //If pago_mensual get
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/ContratosContratistaSupervisor1", "err": err, "status": "502"}
		return nil, outputError
	}
	return
}

func getInfoPagoMensual(pago_mensual models.PagoMensual) (pago_info models.PagoContratistaCdpRp, outputError map[string]interface{}) {
	var contratistas []models.InformacionProveedor
	var contratos_disponibilidad []models.ContratoDisponibilidad
	fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/informacion_proveedor/?query=NumDocumento:" + pago_mensual.DocumentoPersonaId)
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pago_mensual.DocumentoPersonaId, &contratistas); (err == nil) && (response == 200) {
		for _, contratista := range contratistas {

			var informacion_contrato_contratista models.InformacionContratoContratista
			informacion_contrato_contratista, outputError = GetInformacionContratoContratista(pago_mensual.NumeroContrato, strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64))
			if outputError == nil {
				var contrato models.InformacionContrato
				contrato, outputError = GetContrato(pago_mensual.NumeroContrato, strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64))
				if outputError == nil {
					fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/contrato_disponibilidad/?query=NumeroContrato:" + contrato.Contrato.NumeroContrato + ",Vigencia:" + contrato.Contrato.Vigencia)
					if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+contrato.Contrato.NumeroContrato+",Vigencia:"+contrato.Contrato.Vigencia, &contratos_disponibilidad); (err == nil) && (response == 200) {

						for _, contrato_disponibilidad := range contratos_disponibilidad {

							var cdprp models.InformacionCdpRp
							cdprp, outputError = GetRP(strconv.Itoa(contrato_disponibilidad.NumeroCdp), strconv.Itoa(contrato_disponibilidad.VigenciaCdp))
							if outputError == nil {
								for _, rp := range cdprp.CdpXRp.CdpRp {

									pago_info.PagoMensual = &pago_mensual
									pago_info.NombreDependencia = informacion_contrato_contratista.InformacionContratista.Dependencia
									pago_info.NombrePersona = contratista.NomProveedor
									pago_info.NumeroCdp = strconv.Itoa(contrato_disponibilidad.NumeroCdp)
									pago_info.VigenciaCdp = strconv.Itoa(contrato_disponibilidad.VigenciaCdp)
									pago_info.NumeroRp = rp.RpNumeroRegistro
									pago_info.VigenciaRp = rp.RpVigencia
									pago_info.Rubro = contrato.Contrato.Rubro
									return pago_info, nil

								}
							} else {
								return pago_info, outputError
							}
						}

					} else { // If contrato_disponibilidad get
						logs.Error(err)
						outputError = map[string]interface{}{"funcion": "/ContratosContratistaSupervisor3", "err": err, "status": "502"}
						return pago_info, outputError
					}
				} else {
					return pago_info, outputError
				}
			} else {
				return pago_info, outputError
			}

		}
		outputError = map[string]interface{}{"funcion": "/ContratosContratistaSupervisor2", "err": err, "status": "502"}
	} else { //If informacion_proveedor get
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/ContratosContratistaSupervisor2", "err": err, "status": "502"}
		return pago_info, outputError
	}
	return pago_info, outputError
}
