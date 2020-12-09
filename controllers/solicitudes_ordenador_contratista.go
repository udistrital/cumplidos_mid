package controllers

import (
	"fmt"
	"strconv"
	"github.com/astaxie/beego/httplib"

	//"net/http"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

// ContratosContratistaController operations for contratos_contratista
type GetSolicitudesOrdenadorContratistasController struct {
	beego.Controller
}

//URLMapping ...
func (c *GetSolicitudesOrdenadorContratistasController) URLMapping() {	

}


// AprobacionPagoController ...
// @Title GetSolicitudesOrdenadorContratistas
// @Description create GetSolicitudesOrdenadorContratistas
// @Param docordenador path string true "Número del documento del supervisor"
// @Success 201
// @Failure 403 :docordenador is empty
// @router /solicitudes_ordenador_contratistas/:docordenador [get]
func (c *GetSolicitudesOrdenadorContratistasController) GetSolicitudesOrdenadorContratistas() {

	var alertErr models.Alert
	// alertas := append([]interface{}{"Response:"})

	doc_ordenador := c.GetString(":docordenador")
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var pagos_contratista_cdp_rp []models.PagoContratistaCdpRp
	var contratos_disponibilidad []models.ContratoDisponibilidad
	r := httplib.Get(beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudAdmin") + "/" + beego.AppConfig.String("NscrudAdmin") + "/pago_mensual/")
	r.Param("offset", strconv.Itoa(offset))
	r.Param("limit", strconv.Itoa(limit))
	r.Param("query", "EstadoPagoMensual.CodigoAbreviacion:AS,Responsable:"+doc_ordenador)
	if err := r.ToJSON(&pagos_mensuales); err == nil {

		for v, _ := range pagos_mensuales {

			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pagos_mensuales[v].Persona, &contratistas); err == nil {

				for _, contratista := range contratistas {

					var informacion_contrato_contratista models.InformacionContratoContratista
					informacion_contrato_contratista = GetInformacionContratoContratista(pagos_mensuales[v].NumeroContrato, strconv.FormatFloat(pagos_mensuales[v].VigenciaContrato, 'f', 0, 64))
					var contrato models.InformacionContrato
					contrato = GetContrato(pagos_mensuales[v].NumeroContrato, strconv.FormatFloat(pagos_mensuales[v].VigenciaContrato, 'f', 0, 64))

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
						alertErr.Type = "error"
						alertErr.Code = "404"
						alertErr.Body = "" + beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String("UrlcrudAgora") + "/" + beego.AppConfig.String("NscrudAgora") + "/contrato_disponibilidad/?query=NumeroContrato:" + contrato.Contrato.NumeroContrato + ",Vigencia:" + contrato.Contrato.Vigencia + " numero del contrato : " + pagos_mensuales[v].NumeroContrato + " vigencia: " + strconv.FormatFloat(pagos_mensuales[v].VigenciaContrato, 'f', 0, 64)
						c.Data["json"] = alertErr
						c.ServeJSON()
					}

				}
			} else { //If informacion_proveedor get

				fmt.Println("Mirenme, me morí en If informacion_proveedor get, solucioname!!! ", err)
				return
			}

		}
	} else { //If pago_mensual get

		fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
		return
	}
	c.Data["json"] = pagos_contratista_cdp_rp

	c.ServeJSON()

}