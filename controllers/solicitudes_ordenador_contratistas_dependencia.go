package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	//"net/http"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

// CertificacionVistoBuenoController operations for certificacion_visto_bueno
type SolicitudesOrdenadorContratistasDependenciaController struct {
	beego.Controller
}

// URLMapping ...
func (c *SolicitudesOrdenadorContratistasDependenciaController) URLMapping() {
	c.Mapping("GetSolicitudesOrdenadorContratistasDependencia", c.GetSolicitudesOrdenadorContratistasDependencia)
}

// AprobacionPagoController ...
// @Title GetSolicitudesOrdenadorContratistasDependencia
// @Description create GetSolicitudesOrdenadorContratistasDependencia
// @Param docordenador path string true "Número del documento del supervisor"
// @Param cod_dependencia path string true "cod_dependencia"
// @Success 201
// @Failure 403 :docordenador is empty
// @router /solicitudes_ordenador_contratistas_dependencia/:docordenador/:cod_dependencia [get]
func (c *SolicitudesOrdenadorContratistasDependenciaController	) GetSolicitudesOrdenadorContratistasDependencia() {

	var contrato_dependencia models.ContratoDependencia

	var contratistas []models.InformacionProveedor
	var pagos_contratista_cdp_rp []models.PagoContratistaCdpRp

	var conteo_offset, conteo_limit = 0, 0
	var contratos_disponibilidad []models.ContratoDisponibilidad

	var alerta models.Alert

	doc_ordenador := c.GetString(":docordenador")
	cod_dependencia := c.GetString(":cod_dependencia")
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	t := time.Now()
	fecha_inicio := fmt.Sprintf("%d-%02d",
		t.Year(), t.Month())

	t2 := t.AddDate(0, -1, 0)
	fecha_final := fmt.Sprintf("%d-%02d",
		t2.Year(), t2.Month())

	contrato_dependencia = GetContratosDependenciaFiltro(cod_dependencia, fecha_inicio, fecha_final)
	for _, cd := range contrato_dependencia.Contratos.Contrato {

		var pagos_mensuales []models.PagoMensual

		if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/pago_mensual/?query=NumeroContrato:"+cd.NumeroContrato+",VigenciaContrato:"+cd.Vigencia+",EstadoPagoMensual.CodigoAbreviacion:AS,Responsable:"+doc_ordenador, &pagos_mensuales); err == nil {
			fmt.Println(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/pago_mensual/?query=NumeroContrato:"+cd.NumeroContrato+",VigenciaContrato:"+cd.Vigencia+",EstadoPagoMensual.CodigoAbreviacion:AS,Responsable:"+doc_ordenador)
			for v, pm := range pagos_mensuales {

				if conteo_limit == limit && conteo_offset == offset {
					break
				}

				var pago_mensual models.PagoMensual

				if pm.NumeroContrato != "" {

					if conteo_offset == offset && conteo_limit < limit {
						pago_mensual.Persona = pm.Persona
						pago_mensual.VigenciaContrato = pm.VigenciaContrato
						pago_mensual.NumeroContrato = pm.NumeroContrato

						if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pago_mensual.Persona, &contratistas); err == nil {

							for _, contratista := range contratistas {

								var informacion_contrato_contratista models.InformacionContratoContratista
								informacion_contrato_contratista = GetInformacionContratoContratista(pago_mensual.NumeroContrato, strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64))
								var contrato models.InformacionContrato
								contrato = GetContrato(pago_mensual.NumeroContrato, strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64))

								if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+contrato.Contrato.NumeroContrato+",Vigencia:"+contrato.Contrato.Vigencia, &contratos_disponibilidad); err == nil {

									for _, contrato_disponibilidad := range contratos_disponibilidad {

										var cdprp models.InformacionCdpRp
										cdprp = GetRP(strconv.Itoa(contrato_disponibilidad.NumeroCdp), strconv.Itoa(contrato_disponibilidad.VigenciaCdp))

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

									}

								} else { // If contrato_disponibilidad get
									fmt.Println("Mirenme, me morí en If contrato_disponibilidad get, solucioname!!! ", err)
								}

							}
							conteo_limit = conteo_limit + 1
						} else { //If informacion_proveedor get

							fmt.Println("Mirenme, me morí en If informacion_proveedor get, solucioname!!! ", err)
							return
						}

					} else {
						conteo_offset = conteo_offset + 1
					}
				}
			}

		} else { //If pago_mensual get

			fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
			return
		}

	}

	if pagos_contratista_cdp_rp == nil {
		c.Data["json"] = alerta
		c.ServeJSON()

	} else {
		c.Data["json"] = pagos_contratista_cdp_rp
		c.ServeJSON()

	}

}

func GetContratosDependenciaFiltro(dependencia string, fecha_inicio string, fecha_fin string) (contratos_dependencia models.ContratoDependencia) {

	var temp map[string]interface{}

	if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAdministrativa")+"/"+"contratos_dependencia/"+dependencia+"/"+fecha_fin+"/"+fecha_inicio, &temp); err == nil {
		fmt.Println("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAdministrativa")+"/"+"contratos_dependencia/"+dependencia+"/"+fecha_fin+"/"+fecha_inicio)
		json_contrato, error_json := json.Marshal(temp)
		if error_json == nil {
			if err := json.Unmarshal(json_contrato, &contratos_dependencia); err == nil {
				return contratos_dependencia
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(error_json.Error())
		}

	} else {

		fmt.Println(err)
	}

	return contratos_dependencia
}