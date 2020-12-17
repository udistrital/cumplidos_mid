package controllers

import (
	_ "encoding/json"
	_ "time"

	//"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
// @router /:docsupervisor [get]
func (c *SolicitudesSupervisorContratistasController) GetSolicitudesSupervisorContratistas() {

	doc_supervisor := c.GetString(":docsupervisor")
	if pagos_contratista_cdp_rp, err := helpers.ContratosContratistaSupervisor(doc_supervisor); err != nil || len(pagos_contratista_cdp_rp) == 0 {
		logs.Error(err)
		c.Data["mesaage"] = "Error service Get SolicitudesSupervisorContratistas: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	} else {
		c.Data["json"] = pagos_contratista_cdp_rp
	}
	c.ServeJSON()

}
