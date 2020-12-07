package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	//"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/models"
)

// ContratosContratistaController operations for contratos_contratista
type ContratosContratistaController struct {
	beego.Controller
}

//URLMapping ...
func (c *ContratosContratistaController) URLMapping() {
	// c.Mapping("ObtenerInfoCoordinador", c.ObtenerInfoCoordinador)
	c.Mapping("GetContratosContratista", c.GetContratosContratista)
	// c.Mapping("ObtenerInfoOrdenador", c.ObtenerInfoOrdenador)
	// c.Mapping("PagoAprobado", c.PagoAprobado)
	// c.Mapping("CertificacionVistoBueno", c.CertificacionVistoBueno)
	// c.Mapping("CertificacionDocumentosAprobados", c.CertificacionDocumentosAprobados)
	// c.Mapping("ObtenerDependenciaOrdenador", c.ObtenerDependenciaOrdenador)

}

// GetContratosContratista ...
// @Title GetContratosContratista
// @Description create ContratosContratista
// @Param numero_documento path string true "Número documento"
// @Success 200 {object} []models.ContratoDisponibilidadRp
// @Failure 404 not found resource
// @router /contratos_contratista/:numero_documento [get]
func (c *ContratosContratistaController) GetContratosContratista() {
	numero_documento := c.GetString(":numero_documento")
	
	if contratos_disponibilidad_rp, err:= contratos_contratista(numero_documento);err!=nil || len(contratos_disponibilidad_rp)==0{
		logs.Error(err)
		c.Data["mesaage"] = "Error service Get ContratosContratista: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	}else{
		c.Data["json"] = contratos_disponibilidad_rp
	}
	c.ServeJSON()

}

func contratos_contratista(numero_documento string) (contratos_disponibilidad_rp []models.ContratoDisponibilidadRp, err error) {
	var contratos_disponibilidad []models.ContratoDisponibilidad
	var novedades_postcontractuales []models.NovedadPostcontractual
	var novedades_novedad []models.NovedadPostcontractual
	var informacion_proveedores []models.InformacionProveedor
	contratos_persona := GetContratosPersona(numero_documento)
	if contratos_persona.ContratosPersonas.ContratoPersona == nil { // Si no tiene contrato
		if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+numero_documento, &informacion_proveedores); err == nil {
			for _, persona := range informacion_proveedores {
				if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/novedad_postcontractual/?query=Contratista:"+strconv.Itoa(persona.Id)+"&sortby=FechaInicio&order=desc&limit=1", &novedades_postcontractuales); err == nil {
					for _, novedad := range novedades_postcontractuales {
						var contrato models.InformacionContrato
						contrato = GetContrato(novedad.NumeroContrato, strconv.Itoa(novedad.Vigencia))
						var informacion_contrato_contratista models.InformacionContratoContratista
						informacion_contrato_contratista = GetInformacionContratoContratista(novedad.NumeroContrato, strconv.Itoa(novedad.Vigencia))

						if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/novedad_postcontractual/?query=NumeroContrato:"+novedad.NumeroContrato+",Vigencia:"+strconv.Itoa(novedad.Vigencia)+"&sortby=FechaInicio&order=desc&limit=1", &novedades_novedad); err == nil {

							for _, novedad_novedad := range novedades_novedad {

								if novedad_novedad != novedad {

									if (novedad_novedad.FechaInicio.Year() == time.Now().Year() && int(novedad_novedad.FechaFin.Month()) >= int(time.Now().Month()) && novedad_novedad.FechaFin.Year() == time.Now().Year()) ||
										(novedad_novedad.FechaInicio.Year() <= time.Now().Year() && int(novedad_novedad.FechaFin.Month()) <= int(time.Now().Month()) && novedad_novedad.FechaFin.Year() >= time.Now().Year() && novedad_novedad.FechaFin.Year() > novedad_novedad.FechaInicio.Year()) {

										if novedad_novedad.TipoNovedad == 219 { // si es una cesión

										} else {

											var cdprp models.InformacionCdpRp
											cdprp = GetRP(strconv.Itoa(novedad_novedad.NumeroCdp), strconv.Itoa(novedad_novedad.VigenciaCdp))
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

												contratos_disponibilidad_rp = append(contratos_disponibilidad_rp, contrato_disponibilidad_rp)
											}

										}
									}

								} else {

									if (novedad.FechaInicio.Year() == time.Now().Year() && int(novedad.FechaFin.Month()) >= int(time.Now().Month()) && novedad.FechaFin.Year() == time.Now().Year()) ||
										(novedad.FechaInicio.Year() <= time.Now().Year() && int(novedad.FechaFin.Month()) <= int(time.Now().Month()) && novedad.FechaFin.Year() >= time.Now().Year() && novedad.FechaFin.Year() > novedad.FechaInicio.Year()) {

										if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+contrato.Contrato.NumeroContrato+",Vigencia:"+contrato.Contrato.Vigencia, &contratos_disponibilidad); err == nil {

											for _, contrato_disponibilidad := range contratos_disponibilidad {

												var cdprp models.InformacionCdpRp
												cdprp = GetRP(strconv.Itoa(contrato_disponibilidad.NumeroCdp), strconv.Itoa(contrato_disponibilidad.VigenciaCdp))

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

													contratos_disponibilidad_rp = append(contratos_disponibilidad_rp, contrato_disponibilidad_rp)
												}

											}

										} else { // If contrato_disponibilidad get
											fmt.Println("Mirenme, me morí en If contrato_disponibilidad get, solucioname!!! ", err)

										}
									}

								}

							} //fin for novedad novedad

						} else { // If novedad_postcontractual get
							fmt.Println("Mirenme, me morí en If novedad_postcontractual de la novedad get, solucioname!!! ", err.Error())
						}
					}

				} else { // If novedad_postcontractual get
					fmt.Println("Mirenme, me morí en If novedad_postcontractual get, solucioname!!! ", err.Error())
				}
			}

		} else { // If informacion_proveedor get
			fmt.Println("Mirenme, me morí en If informacion_proveedor get, solucioname!!! ", err.Error())
		}

	} else { // si tiene contrato
		for _, contrato_persona := range contratos_persona.ContratosPersonas.ContratoPersona {
			var contrato models.InformacionContrato
			contrato = GetContrato(contrato_persona.NumeroContrato, contrato_persona.Vigencia)
			//var novedad_postcontractual models.NovedadPostcontractual
			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/novedad_postcontractual/?query=NumeroContrato:"+contrato_persona.NumeroContrato+",Vigencia:"+contrato_persona.Vigencia+"&sortby=FechaInicio&order=desc&limit=1", &novedades_postcontractuales); err == nil {
				//var	prueba []models.NovedadPostcontractual
				//	json.NewDecoder(r.Body).Decode(prueba)
				var informacion_contrato_contratista models.InformacionContratoContratista
				informacion_contrato_contratista = GetInformacionContratoContratista(contrato_persona.NumeroContrato, contrato_persona.Vigencia)
				if novedades_postcontractuales != nil { // Si tiene novedades
					for _, novedad := range novedades_postcontractuales {
						if novedad.TipoNovedad == 219 { // si es una cesión

							if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+numero_documento, &informacion_proveedores); err == nil {
								for _, persona := range informacion_proveedores {
									if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/novedad_postcontractual/?query=Contratista:"+strconv.Itoa(persona.Id)+"&sortby=FechaInicio&order=desc&limit=1", &novedades_postcontractuales); err == nil {
										for _, novedad := range novedades_postcontractuales {
											var contrato models.InformacionContrato
											contrato = GetContrato(novedad.NumeroContrato, strconv.Itoa(novedad.Vigencia))

											var informacion_contrato_contratista models.InformacionContratoContratista
											informacion_contrato_contratista = GetInformacionContratoContratista(novedad.NumeroContrato, strconv.Itoa(novedad.Vigencia))
											if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+contrato.Contrato.NumeroContrato+",Vigencia:"+contrato.Contrato.Vigencia, &contratos_disponibilidad); err == nil {

												for _, contrato_disponibilidad := range contratos_disponibilidad {

													var cdprp models.InformacionCdpRp
													cdprp = GetRP(strconv.Itoa(contrato_disponibilidad.NumeroCdp), strconv.Itoa(contrato_disponibilidad.VigenciaCdp))

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

														contratos_disponibilidad_rp = append(contratos_disponibilidad_rp, contrato_disponibilidad_rp)
													}

												}

											} else { // If contrato_disponibilidad get
												fmt.Println("Mirenme, me morí en If contrato_disponibilidad get, solucioname!!! ", err)

											}
										}

									} else { // If novedad_postcontractual get
										fmt.Println("Mirenme, me morí en If novedad_postcontractual get, solucioname!!! ", err.Error())
									}
								}

							} else { // If informacion_proveedor get
								fmt.Println("Mirenme, me morí en If informacion_proveedor get, solucioname!!! ", err.Error())
							}

						} else { // si no es una cesión
							var cdprp models.InformacionCdpRp
							cdprp = GetRP(strconv.Itoa(novedad.NumeroCdp), strconv.Itoa(novedad.VigenciaCdp))

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

								contratos_disponibilidad_rp = append(contratos_disponibilidad_rp, contrato_disponibilidad_rp)
							}

						}

					}
				} else { // si no tiene novedades

					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+contrato.Contrato.NumeroContrato+",Vigencia:"+contrato.Contrato.Vigencia, &contratos_disponibilidad); err == nil {
						for _, contrato_disponibilidad := range contratos_disponibilidad {
							var cdprp models.InformacionCdpRp
							cdprp = GetRP(strconv.Itoa(contrato_disponibilidad.NumeroCdp), strconv.Itoa(contrato_disponibilidad.VigenciaCdp))
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

								contratos_disponibilidad_rp = append(contratos_disponibilidad_rp, contrato_disponibilidad_rp)
							}

						}

					} else { // If contrato_disponibilidad get
						fmt.Println("Mirenme, me morí en If contrato_disponibilidad get, solucioname!!! ", err)
						return nil,err
					}

				}
			} else { // If novedad_postcontractual get
				fmt.Println("Mirenme, me morí en If novedad_postcontractual get, solucioname!!! ", err.Error())
				return nil, err
			}

		}

	}
	return
}

func GetRP(numero_cdp string, vigencia_cdp string) (rp models.InformacionCdpRp) {

	var temp map[string]interface{}
	var temp_cdp_rp models.InformacionCdpRp

	if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudFinanciera")+"/"+"cdprp/"+numero_cdp+"/"+vigencia_cdp+"/01", &temp); err == nil {
		json_cdp_rp, error_json := json.Marshal(temp)

		if error_json == nil {
			if err := json.Unmarshal(json_cdp_rp, &temp_cdp_rp); err == nil {
				rp = temp_cdp_rp
				return rp
			} else {
				fmt.Println(err)
			}

		} else {
			fmt.Println(error_json.Error())
		}

	} else {

		fmt.Println(err)
	}
	return rp
}


func GetContratosPersona(num_documento string) (contratos_persona models.InformacionContratosPersona) {
	var temp map[string]interface{}
	var contratos models.InformacionContratosPersona
	if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAdministrativa")+"/"+"contratos_persona/"+num_documento, &temp); err == nil {
		json_contratos, error_json := json.Marshal(temp)
		if error_json == nil {
			err := json.Unmarshal(json_contratos, &contratos)
			if err == nil {
				contratos_persona = contratos
				return contratos_persona
			} else {
				fmt.Println(err)
			}

		} else {
			fmt.Println(error_json.Error())
		}

	} else {

		fmt.Println(err)
	}

	return contratos_persona

}

func GetContrato(num_contrato_suscrito string, vigencia string) (informacion_contrato models.InformacionContrato) {

	var temp map[string]interface{}

	if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAdministrativa")+"/"+"contrato/"+num_contrato_suscrito+"/"+vigencia, &temp); err == nil {
		json_contrato, error_json := json.Marshal(temp)

		if error_json == nil {
			var contrato models.InformacionContrato
			if err := json.Unmarshal(json_contrato, &contrato); err == nil {
				informacion_contrato = contrato
				return informacion_contrato
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(error_json.Error())
		}

	} else {

		fmt.Println(err)
	}

	return informacion_contrato
}

func GetInformacionContratoContratista(num_contrato_suscrito string, vigencia string) (informacion_contrato_contratista models.InformacionContratoContratista) {

	var temp map[string]interface{}

	if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAdministrativa")+"/"+"informacion_contrato_contratista/"+num_contrato_suscrito+"/"+vigencia, &temp); err == nil {
		json_contrato, error_json := json.Marshal(temp)

		if error_json == nil {
			var contrato_contratista models.InformacionContratoContratista
			if err := json.Unmarshal(json_contrato, &contrato_contratista); err == nil {
				informacion_contrato_contratista = contrato_contratista
				return informacion_contrato_contratista
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(error_json.Error())
		}

	} else {

		fmt.Println(err)
	}

	return informacion_contrato_contratista
}
