package controllers

import (

	//"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
	"github.com/udistrital/cumplidos_mid/models"
)

// SolicitudesOrdenadorContratistasDependenciaController operations for solicitudes_ordenador_contratistas_dependencia
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
func (c *SolicitudesOrdenadorContratistasDependenciaController) GetSolicitudesOrdenadorContratistasDependencia() {

	doc_ordenador := c.GetString(":docordenador")
	cod_dependencia := c.GetString(":cod_dependencia")
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")
	var alerta models.Alert

	if pagos_contratista_cdp_rp, err := helpers.ContratosContratistaDependencia(doc_ordenador, cod_dependencia, limit, offset); err != nil || len(pagos_contratista_cdp_rp) == 0 {
		logs.Error(err)
		c.Data["json"] = alerta
		c.Data["mesaage"] = "Error service Get SolicitudesOrdenadorContratistaDependencia: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	} else {
		c.Data["json"] = pagos_contratista_cdp_rp
	}
	c.ServeJSON()

}
