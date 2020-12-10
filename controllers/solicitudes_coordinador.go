package controllers

import (
	_ "encoding/json"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
)

// AprobacionPagoController operations for AprobacionPago
type SolicitudesCoordinadorController struct {
	beego.Controller
}

//URLMapping ...
func (c *SolicitudesCoordinadorController) URLMapping() {
	c.Mapping("GetSolicitudesCoordinador", c.GetSolicitudesCoordinador)
}

// AprobacionPagoController ...
// @Title GetSolicitudesCoordinador
// @Description create GetSolicitudesCoordinador
// @Param doccoordinador path string true "Número del documento del coordinador"
// @Success 200 {object} []models.PagoPersonaProyecto
// @Failure 404 not found resource
// @router /:doccoordinador [get]
func (c *SolicitudesCoordinadorController) GetSolicitudesCoordinador() {
	doc_coordinador := c.GetString(":doccoordinador")
	//fmt.Println("salida2: ", pagos_personas_proyecto,"   ", len(pagos_personas_proyecto))
	if pagos_personas_proyecto, err := helpers.SolicitudCoordinador(doc_coordinador); err != nil || len(pagos_personas_proyecto) == 0 {
		logs.Error(err)
		c.Data["mesaage"] = "Error service Get solicitudes_coordinador: The request contains an incorrect parameter or no record exists"
		c.Abort("404")

	} else {
		c.Data["json"] = pagos_personas_proyecto
	}

	c.ServeJSON()

}
