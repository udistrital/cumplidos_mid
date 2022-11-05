package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

func ContratosContratista(numero_documento string) (contratos_disponibilidad_rp []models.ContratoDisponibilidadRp, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": "/ContratosContratista", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var contratos_disponibilidad []models.ContratoDisponibilidad
	var novedades_postcontractuales []models.NovedadPostcontractual
	var novedades_novedad []models.NovedadPostcontractual
	var informacion_proveedores []models.InformacionProveedor
	contratos_persona, outputError := GetContratosPersona(numero_documento)
	//fmt.Println("Contratos persona:", contratos_persona)
	if outputError == nil {
		//fmt.Println("outputError==nil")
		//if contratos_persona.ContratosPersonas.ContratoPersona == nil { // Si no tiene contrato
		fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/informacion_proveedor/?query=NumDocumento:" + numero_documento)
		if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+numero_documento, &informacion_proveedores); (err == nil) && (response == 200) {
			//fmt.Println("informacion_proveedor:", informacion_proveedores)
			for _, persona := range informacion_proveedores {
				fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/novedad_postcontractual/?query=Contratista:" + strconv.Itoa(persona.Id) + "&sortby=FechaInicio&order=asc&limit=-1")
				if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/novedad_postcontractual/?query=Contratista:"+strconv.Itoa(persona.Id)+",TipoNovedad:219&sortby=FechaInicio&order=asc&limit=-1", &novedades_postcontractuales); (err == nil) && (response == 200) {
					//fmt.Println("novedades", novedades_postcontractuales)
					for _, novedad := range novedades_postcontractuales {
						var contrato models.InformacionContrato
						contrato, outputError = GetContrato(novedad.NumeroContrato, strconv.Itoa(novedad.Vigencia))
						if novedad.FechaFin.Before(time.Now().AddDate(1, 0, 0)) && novedad.FechaInicio.Before(time.Now()) {
							if outputError == nil {
								var informacion_contrato_contratista models.InformacionContratoContratista
								informacion_contrato_contratista, outputError = GetInformacionContratoContratista(novedad.NumeroContrato, strconv.Itoa(novedad.Vigencia))
								if outputError == nil {
									//LLenar el registro del contrato base
									if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+contrato.Contrato.NumeroContrato+",Vigencia:"+contrato.Contrato.Vigencia, &contratos_disponibilidad); (err == nil) && (response == 200) {
										for _, contrato_disponibilidad := range contratos_disponibilidad {
											var cdprp models.InformacionCdpRp
											cdprp, outputError = GetRP(strconv.Itoa(contrato_disponibilidad.NumeroCdp), strconv.Itoa(contrato_disponibilidad.VigenciaCdp))
											if outputError == nil {
												for _, rp := range cdprp.CdpXRp.CdpRp {
													var contrato_disponibilidad_rp models.ContratoDisponibilidadRp
													contrato_disponibilidad_rp.NumeroContratoSuscrito = novedad.NumeroContrato
													contrato_disponibilidad_rp.Vigencia = strconv.Itoa(novedad.Vigencia)
													contrato_disponibilidad_rp.NumeroCdp = strconv.Itoa(contrato_disponibilidad.NumeroCdp)
													contrato_disponibilidad_rp.VigenciaCdp = strconv.Itoa(contrato_disponibilidad.VigenciaCdp)
													contrato_disponibilidad_rp.NumeroRp = rp.RpNumeroRegistro
													contrato_disponibilidad_rp.VigenciaRp = rp.RpVigencia
													contrato_disponibilidad_rp.NombreDependencia = informacion_contrato_contratista.InformacionContratista.Dependencia
													contrato_disponibilidad_rp.NumDocumentoSupervisor = contrato.Contrato.Supervisor.DocumentoIdentificacion
													var actasInicio []models.ActaInicio
													if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/acta_inicio/?query=NumeroContrato:"+contrato_disponibilidad.NumeroContrato+",Vigencia:"+strconv.Itoa(contrato_disponibilidad.Vigencia), &actasInicio); (err == nil) && (response == 200) {
														for _, actaInicio := range actasInicio {
															contrato_disponibilidad_rp.FechaInicio = novedad.FechaInicio
															contrato_disponibilidad_rp.FechaFin = actaInicio.FechaFin
														}
													}
													contratos_disponibilidad_rp = append(contratos_disponibilidad_rp, contrato_disponibilidad_rp)
												}
											} else {
												return nil, outputError
											}
										}
									} else { // If contrato_disponibilidad get
										logs.Error(err)
										outputError = map[string]interface{}{"funcion": "/contratosContratista", "err": err, "status": "502"}
										return nil, outputError
									}
									fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/novedad_postcontractual/?query=NumeroContrato:" + novedad.NumeroContrato + ",Vigencia:" + strconv.Itoa(novedad.Vigencia) + "&sortby=FechaInicio&order=asc&limit=-1")
									if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/novedad_postcontractual/?query=NumeroContrato:"+novedad.NumeroContrato+",Vigencia:"+strconv.Itoa(novedad.Vigencia)+"&sortby=FechaInicio&order=asc&limit=-1", &novedades_novedad); (err == nil) && (response == 200) {
										//fmt.Println("Novedades de nuevo", novedades_novedad)
										for _, novedad_novedad := range novedades_novedad {
											//Se recorren las novedades del contrato de la cesion
											if novedad_novedad.Id != novedad.Id && novedad_novedad.FechaInicio.After(novedad.FechaInicio) {
												//Se evaluan las novedades que tiene despues de la cesion
												if novedad_novedad.TipoNovedad == 219 { // si es una cesión
													contratos_disponibilidad_rp[len(contratos_disponibilidad_rp)-1].FechaFin = novedad_novedad.FechaInicio.AddDate(0, 0, -1)
													break
												} else { // si no es una cesión
													var cdprp models.InformacionCdpRp
													if novedad_novedad.TipoNovedad == 220 { //Novedad de otro si
														//Registro Otro Si
														cdprp, outputError = GetRP(strconv.Itoa(novedad_novedad.NumeroCdp), strconv.Itoa(novedad_novedad.VigenciaCdp))
														if outputError == nil {
															for _, rp := range cdprp.CdpXRp.CdpRp {
																var contrato_disponibilidad_rp models.ContratoDisponibilidadRp
																contrato_disponibilidad_rp.NumeroContratoSuscrito = novedad_novedad.NumeroContrato
																contrato_disponibilidad_rp.Vigencia = strconv.Itoa(novedad_novedad.Vigencia)
																contrato_disponibilidad_rp.NumeroCdp = strconv.Itoa(novedad_novedad.NumeroCdp)
																contrato_disponibilidad_rp.VigenciaCdp = strconv.Itoa(novedad_novedad.VigenciaCdp)
																contrato_disponibilidad_rp.NumeroRp = rp.RpNumeroRegistro
																contrato_disponibilidad_rp.VigenciaRp = rp.RpVigencia
																contrato_disponibilidad_rp.NombreDependencia = informacion_contrato_contratista.InformacionContratista.Dependencia
																contrato_disponibilidad_rp.NumDocumentoSupervisor = contrato.Contrato.Supervisor.DocumentoIdentificacion
																contrato_disponibilidad_rp.FechaInicio = novedad_novedad.FechaInicio
																contrato_disponibilidad_rp.FechaFin = novedad_novedad.FechaFin
																contratos_disponibilidad_rp = append(contratos_disponibilidad_rp, contrato_disponibilidad_rp)
															}
														} else {
															return nil, outputError
														}
													} else {
														if novedad_novedad.TipoNovedad == 218 { //Novedad de Terminacion anticipada
															//Registro
															for i, contrato := range contratos_disponibilidad_rp {
																if contrato.FechaFin.After(novedad_novedad.FechaFin) && contrato.NumeroContratoSuscrito == novedad_novedad.NumeroContrato {
																	contratos_disponibilidad_rp[i].FechaFin = novedad_novedad.FechaFin
																}
															}
														} else {
															if novedad_novedad.TipoNovedad == 216 { //Novedad de suspencion
																//Registro
																var days int
																days = int(novedad_novedad.FechaFin.Sub(novedad_novedad.FechaInicio).Hours() / 24)
																//fmt.Println("dias de diferencia en suspencion " + strconv.Itoa(days))
																for i, contrato := range contratos_disponibilidad_rp {
																	if contrato.FechaInicio.Before(novedad_novedad.FechaInicio) && contrato.FechaFin.After(novedad_novedad.FechaInicio) && contrato.NumeroContratoSuscrito == novedad.NumeroContrato {
																		contratos_disponibilidad_rp[i].FechaFin = contratos_disponibilidad_rp[i].FechaFin.AddDate(0, 0, days)
																	}
																}
															}
														}
													}
												}
											} else {
												if novedad_novedad.TipoNovedad == 220 && novedad_novedad.FechaInicio.Before(novedad.FechaInicio) {
													//Cesion de un otro Si
													for i, contrato := range contratos_disponibilidad_rp {
														if contrato.NumeroContratoSuscrito == novedad_novedad.NumeroContrato {
															contratos_disponibilidad_rp[i].FechaFin = novedad_novedad.FechaFin
															contratos_disponibilidad_rp[i].NumeroCdp = strconv.Itoa(novedad_novedad.NumeroCdp)
															contratos_disponibilidad_rp[i].VigenciaCdp = strconv.Itoa(novedad_novedad.VigenciaCdp)
															var cdprp models.InformacionCdpRp
															cdprp, outputError = GetRP(strconv.Itoa(novedad_novedad.NumeroCdp), strconv.Itoa(novedad_novedad.VigenciaCdp))
															if outputError == nil {
																for _, rp := range cdprp.CdpXRp.CdpRp {
																	contratos_disponibilidad_rp[i].NumeroRp = rp.RpNumeroRegistro
																	contratos_disponibilidad_rp[i].VigenciaRp = rp.RpVigencia
																}
															}
														}
													}
												}
											}
										} //fin for novedad novedad

									} else { // If novedad_postcontractual get
										logs.Error(err)
										outputError = map[string]interface{}{"funcion": "/contratosContratista2", "err": err, "status": "502"}
										return nil, outputError
									}
								} else {
									return nil, outputError
								}
							} else {
								return nil, outputError
							}
						} else {
							//fmt.Println("fuera de rango")
						}
					}

				} else { // If novedad_postcontractual get
					logs.Error(err)
					outputError = map[string]interface{}{"funcion": "/contratosContratista3", "err": err, "status": "502"}
					return nil, outputError
				}
			}

		} else { // If informacion_proveedor get
			//fmt.Println(err)
			//fmt.Println(err)
			//fmt.Println(response)
			outputError = map[string]interface{}{"funcion": "/contratosContratista4", "err": err, "status": "502"}
			return nil, outputError
		}

		//} else { // si tiene contrato
		//fmt.Println("contratos disponibilidad", contratos_disponibilidad_rp)
		//fmt.Println("for contrato persona")
		for _, contrato_persona := range contratos_persona.ContratosPersonas.ContratoPersona {
			var contrato models.InformacionContrato
			contrato, outputError = GetContrato(contrato_persona.NumeroContrato, contrato_persona.Vigencia)
			var informacion_contrato_contratista models.InformacionContratoContratista
			informacion_contrato_contratista, outputError = GetInformacionContratoContratista(contrato_persona.NumeroContrato, contrato_persona.Vigencia)
			// se llena el contrato original en el indice 0
			if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+contrato.Contrato.NumeroContrato+",Vigencia:"+contrato.Contrato.Vigencia, &contratos_disponibilidad); (err == nil) && (response == 200) {
				for _, contrato_disponibilidad := range contratos_disponibilidad {
					var cdprp models.InformacionCdpRp
					cdprp, outputError = GetRP(strconv.Itoa(contrato_disponibilidad.NumeroCdp), strconv.Itoa(contrato_disponibilidad.VigenciaCdp))
					if outputError == nil {
						for _, rp := range cdprp.CdpXRp.CdpRp {
							var contrato_disponibilidad_rp models.ContratoDisponibilidadRp
							contrato_disponibilidad_rp.NumeroContratoSuscrito = contrato_persona.NumeroContrato
							contrato_disponibilidad_rp.Vigencia = contrato_persona.Vigencia
							contrato_disponibilidad_rp.NumeroCdp = strconv.Itoa(contrato_disponibilidad.NumeroCdp)
							contrato_disponibilidad_rp.VigenciaCdp = strconv.Itoa(contrato_disponibilidad.VigenciaCdp)
							contrato_disponibilidad_rp.NumeroRp = rp.RpNumeroRegistro
							contrato_disponibilidad_rp.VigenciaRp = rp.RpVigencia
							contrato_disponibilidad_rp.NombreDependencia = informacion_contrato_contratista.InformacionContratista.Dependencia
							contrato_disponibilidad_rp.NumDocumentoSupervisor = contrato.Contrato.Supervisor.DocumentoIdentificacion
							var actasInicio []models.ActaInicio
							if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/acta_inicio/?query=NumeroContrato:"+contrato_disponibilidad.NumeroContrato+",Vigencia:"+strconv.Itoa(contrato_disponibilidad.Vigencia), &actasInicio); (err == nil) && (response == 200) {
								for _, actaInicio := range actasInicio {
									contrato_disponibilidad_rp.FechaInicio = actaInicio.FechaInicio
									contrato_disponibilidad_rp.FechaFin = actaInicio.FechaFin
								}
							}
							contratos_disponibilidad_rp = append(contratos_disponibilidad_rp, contrato_disponibilidad_rp)
						}
					} else {
						return nil, outputError
					}
				}
			} else { // If contrato_disponibilidad get
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/contratosContratista", "err": err, "status": "502"}
				return nil, outputError
			}
			if outputError == nil {
				// Se actua respecto a las novedades encontradas
				fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/novedad_postcontractual/?query=NumeroContrato:" + contrato_persona.NumeroContrato + ",Vigencia:" + contrato_persona.Vigencia + "&sortby=FechaInicio&order=asc&limit=-1")
				//var novedad_postcontractual models.NovedadPostcontractual
				if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/novedad_postcontractual/?query=NumeroContrato:"+contrato_persona.NumeroContrato+",Vigencia:"+contrato_persona.Vigencia+"&sortby=FechaInicio&order=asc&limit=-1", &novedades_postcontractuales); (err == nil) && (response == 200) {
					//var	prueba []models.NovedadPostcontractual
					//	json.NewDecoder(r.Body).Decode(prueba)
					//fmt.Println("Novedades postcontractualese", novedades_postcontractuales)
					//fmt.Println("Informacion contrato contratista", informacion_contrato_contratista)

					if outputError == nil {
						if novedades_postcontractuales != nil { // Si tiene novedades
							for _, novedad := range novedades_postcontractuales {
								if novedad.TipoNovedad == 219 { // si es una cesión
									contratos_disponibilidad_rp[len(contratos_disponibilidad_rp)-1].FechaFin = novedad.FechaInicio.AddDate(0, 0, -1)
									break
								} else { // si no es una cesión
									var cdprp models.InformacionCdpRp
									if novedad.TipoNovedad == 220 { //Novedad de otro si
										//Registro Otro Si
										cdprp, outputError = GetRP(strconv.Itoa(novedad.NumeroCdp), strconv.Itoa(novedad.VigenciaCdp))
										if outputError == nil {
											for _, rp := range cdprp.CdpXRp.CdpRp {
												var contrato_disponibilidad_rp models.ContratoDisponibilidadRp
												contrato_disponibilidad_rp.NumeroContratoSuscrito = novedad.NumeroContrato
												contrato_disponibilidad_rp.Vigencia = strconv.Itoa(novedad.Vigencia)
												contrato_disponibilidad_rp.NumeroCdp = strconv.Itoa(novedad.NumeroCdp)
												contrato_disponibilidad_rp.VigenciaCdp = strconv.Itoa(novedad.VigenciaCdp)
												contrato_disponibilidad_rp.NumeroRp = rp.RpNumeroRegistro
												contrato_disponibilidad_rp.VigenciaRp = rp.RpVigencia
												contrato_disponibilidad_rp.NombreDependencia = informacion_contrato_contratista.InformacionContratista.Dependencia
												contrato_disponibilidad_rp.NumDocumentoSupervisor = contrato.Contrato.Supervisor.DocumentoIdentificacion
												contrato_disponibilidad_rp.FechaInicio = novedad.FechaInicio
												contrato_disponibilidad_rp.FechaFin = novedad.FechaFin
												contratos_disponibilidad_rp = append(contratos_disponibilidad_rp, contrato_disponibilidad_rp)
											}
										} else {
											return nil, outputError
										}
									} else {
										if novedad.TipoNovedad == 218 { //Novedad de Terminacion anticipada
											//Registro
											for i, contrato := range contratos_disponibilidad_rp {
												if contrato.FechaFin.After(novedad.FechaFin) && contrato.NumeroContratoSuscrito == novedad.NumeroContrato {
													contratos_disponibilidad_rp[i].FechaFin = novedad.FechaFin
												}
											}
										} else {
											if novedad.TipoNovedad == 216 { //Novedad de suspencion
												//Registro
												var days int
												days = int(novedad.FechaFin.Sub(novedad.FechaInicio).Hours() / 24)
												//fmt.Println("dias de diferencia en suspencion " + strconv.Itoa(days))
												for i, contrato := range contratos_disponibilidad_rp {
													if contrato.FechaInicio.Before(novedad.FechaInicio) && contrato.FechaFin.After(novedad.FechaInicio) && contrato.NumeroContratoSuscrito == novedad.NumeroContrato {
														contratos_disponibilidad_rp[i].FechaFin = contratos_disponibilidad_rp[i].FechaFin.AddDate(0, 0, days)
													}
												}
											}
										}
									}
								}
							}
							//fmt.Println("contratos disponibilidad si tiene novedades", contratos_disponibilidad_rp)
						} else { // si no tiene novedades
							//fmt.Println("contratos disponibilidad no tiene novedades", contratos_disponibilidad_rp)
						}
					} else {
						return nil, outputError
					}
				} else { // If novedad_postcontractual get
					logs.Error(err)
					outputError = map[string]interface{}{"funcion": "/contratosContratista", "err": err, "status": "502"}
					return nil, outputError
				}
			} else {
				return nil, outputError
			}
		}
		//}
	} else {
		return nil, outputError
	}
	return
}

func GetRP(numero_cdp string, vigencia_cdp string) (rp models.InformacionCdpRp, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetRP0", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var temp map[string]interface{}
	var temp_cdp_rp models.InformacionCdpRp
	fmt.Println(beego.AppConfig.String("UrlFinancieraJBPM") + "/" + "cdprp/" + numero_cdp + "/" + vigencia_cdp + "/01")
	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlFinancieraJBPM")+"/"+"cdprp/"+numero_cdp+"/"+vigencia_cdp+"/01", &temp); (err == nil) && (response == 200) {
		json_cdp_rp, error_json := json.Marshal(temp)

		if error_json == nil {
			if err := json.Unmarshal(json_cdp_rp, &temp_cdp_rp); err == nil {
				rp = temp_cdp_rp
				return rp, nil
			} else {
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/GetRP1", "err": err, "status": "502"}
				return rp, outputError
			}
		} else {
			logs.Error(error_json)
			outputError = map[string]interface{}{"funcion": "/GetRP2", "err": error_json, "status": "502"}
			return rp, outputError
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetRP3", "err": err, "status": "502"}
		return rp, outputError
	}
	return rp, outputError
}

func GetContratosPersona(num_documento string) (contratos_persona models.InformacionContratosPersona, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetContratosPersona", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var temp map[string]interface{}
	var contratos models.InformacionContratosPersona
	fmt.Println(beego.AppConfig.String("UrlAdministrativaJBPM") + "/" + "contratos_persona/" + num_documento)
	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlAdministrativaJBPM")+"/"+"contratos_persona/"+num_documento, &temp); (err == nil) && (response == 200) {
		json_contratos, error_json := json.Marshal(temp)
		if error_json == nil {
			err := json.Unmarshal(json_contratos, &contratos)
			if err == nil {
				contratos_persona = contratos
				//fmt.Println("Contratos personas", contratos_persona)
				return contratos_persona, nil
			} else {
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/GetContratosPersona", "err": err, "status": "502"}
				return contratos_persona, outputError
			}

		} else {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/GetContratosPersona", "err": error_json.Error(), "status": "502"}
			return contratos_persona, outputError
		}

	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetContratosPersona", "err": err, "status": "502"}
		return contratos_persona, outputError
	}

	return contratos_persona, nil

}

func GetContrato(num_contrato_suscrito string, vigencia string) (informacion_contrato models.InformacionContrato, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetContrato", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var temp map[string]interface{}
	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlAdministrativaJBPM")+"/"+"contrato/"+num_contrato_suscrito+"/"+vigencia, &temp); (err == nil) && (response == 200) {
		json_contrato, error_json := json.Marshal(temp)
		if error_json == nil {
			var contrato models.InformacionContrato
			if err := json.Unmarshal(json_contrato, &contrato); err == nil {
				informacion_contrato = contrato
				return informacion_contrato, nil
			} else {
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/GetContrato", "err": err, "status": "502"}
				return informacion_contrato, outputError
			}
		} else {
			logs.Error(error_json.Error())
			outputError = map[string]interface{}{"funcion": "/GetContrato", "err": error_json.Error(), "status": "502"}
			return informacion_contrato, outputError
		}

	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetContrato", "err": err, "status": "502"}
		return informacion_contrato, outputError
	}

	return informacion_contrato, nil
}

func GetInformacionContratoContratista(num_contrato_suscrito string, vigencia string) (informacion_contrato_contratista models.InformacionContratoContratista, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetInformacionContratoContratista", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var temp map[string]interface{}

	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlAdministrativaJBPM")+"/"+"informacion_contrato_contratista/"+num_contrato_suscrito+"/"+vigencia, &temp); (err == nil) && (response == 200) {
		json_contrato, error_json := json.Marshal(temp)
		if error_json == nil {
			var contrato_contratista models.InformacionContratoContratista
			if err := json.Unmarshal(json_contrato, &contrato_contratista); err == nil {
				informacion_contrato_contratista = contrato_contratista
				return informacion_contrato_contratista, nil
			} else {
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/GetInformacionContratoContratista", "err": err, "status": "502"}
				return informacion_contrato_contratista, outputError
			}
		} else {
			logs.Error(error_json.Error())
			outputError = map[string]interface{}{"funcion": "/GetInformacionContratoContratista", "err": error_json.Error(), "status": "502"}
			return informacion_contrato_contratista, outputError
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/getInformacionContratosContratista", "err": err, "status": "502"}
		return informacion_contrato_contratista, outputError
	}
	return informacion_contrato_contratista, nil
}
