package controllers

import (
	_ "encoding/json"
	"fmt"
	"strconv"
	_ "time"

	//"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/models"
	"github.com/udistrital/cumplidos_mid/helpers"
)

// SolicitudesSupervisorContratistasController operations for solicitudes_supervisor_contratistas
type SolicitudesSupervisorContratistasController struct {
	beego.Controller
}

// URLMapping ...
func (c *SolicitudesSupervisorContratistasController) URLMapping() {
	c.Mapping("GetSolicitudesSupervisorContratistas", c.GetSolicitudesSupervisorContratistas)
}

// AprobacionPagoController ...
// @Title GetSolicitudesSupervisorContratistas
// @Description create GetSolicitudesSupervisorContratistas
// @Param docsupervisor path string true "Número del documento del supervisor"
// @Success 201
// @Failure 403 :docsupervisor is empty
// @router /solicitudes_supervisor_contratistas/:docsupervisor [get]
func (c *SolicitudesSupervisorContratistasController) GetSolicitudesSupervisorContratistas() {

	doc_supervisor := c.GetString(":docsupervisor")
	if pagos_contratista_cdp_rp, err := contratos_contratista_supervisor(doc_supervisor);err!=nil || len(pagos_contratista_cdp_rp)==0{
		logs.Error(err)
		c.Data["mesaage"] = "Error service Get SolicitudesSupervisorContratistas: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	}else{
		c.Data["json"] = pagos_contratista_cdp_rp
	}
	c.ServeJSON()

}

func contratos_contratista_supervisor(doc_supervisor string)(pagos_contratista_cdp_rp []models.PagoContratistaCdpRp, err error){
	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var contratos_disponibilidad []models.ContratoDisponibilidad
	var respuesta_peticion map[string]interface{}

	fmt.Println(beego.AppConfig.String("ProtocolCrudCumplidos")+"://"+beego.AppConfig.String("UrlCrudCumplidos")+"/"+beego.AppConfig.String("NsCrudCumplidos")+"/pago_mensual/?limit=-1&query=EstadoPagoMensualId.CodigoAbreviacion:PRS,DocumentoResponsableId:"+doc_supervisor )
	if err := getJson(beego.AppConfig.String("ProtocolCrudCumplidos")+"://"+beego.AppConfig.String("UrlCrudCumplidos")+"/"+beego.AppConfig.String("NsCrudCumplidos")+"/pago_mensual/?limit=-1&query=EstadoPagoMensualId.CodigoAbreviacion:PRS,DocumentoResponsableId:"+doc_supervisor, &respuesta_peticion); err == nil {
		helpers.LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
		for v, _ := range pagos_mensuales {

			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pagos_mensuales[v].DocumentoPersonaId, &contratistas); err == nil {
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
					}

				}
			} else { //If informacion_proveedor get

				fmt.Println("Mirenme, me morí en If informacion_proveedor get, solucioname!!! ", err)
				return nil, err
			}

		}
	} else { //If pago_mensual get

		fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
		return nil, err
	}
	return
}