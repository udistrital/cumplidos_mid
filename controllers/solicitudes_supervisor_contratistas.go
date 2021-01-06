package controllers

import (
	_ "encoding/json"
	_ "time"

	//"net/http"
	"strconv"

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
// @Success 200
// @Failure 403 :docsupervisor is empty
// @router /:docsupervisor [get]
func (c *SolicitudesSupervisorContratistasController) GetSolicitudesSupervisorContratistas() {

	doc_supervisor := c.GetString(":docsupervisor")

	//defer helpers.GestionError(c)
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "SolicitudesSupervisorContratistasController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	_, err := strconv.Atoi(doc_supervisor)

	if err != nil {
		panic(map[string]interface{}{"funcion": "GetSolicitudesSupervisorContratistas", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if pagos_contratista_cdp_rp, err := helpers.ContratosContratistaSupervisor(doc_supervisor); err == nil || len(pagos_contratista_cdp_rp) != 0 {
		//c.Data["json"] = pagos_contratista_cdp_rp
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": pagos_contratista_cdp_rp}
	} else {
		panic(err)
	}
	c.ServeJSON()

}
