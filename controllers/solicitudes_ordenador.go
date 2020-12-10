package controllers

import (
	_ "encoding/json"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
)

// SolicitudesOrdenadorContratistasController operations for SolicitudesOrdenadorContratistas
type SolicitudesOrdenadorController struct {
	beego.Controller
}

//URLMapping ...
func (c *SolicitudesOrdenadorController) URLMapping() {
	c.Mapping("GetSolicitudesOrdenador", c.GetSolicitudesOrdenador)
	c.Mapping("GetSolicitudesOrdenadorContratistas", c.ObtenerDependenciaOrdenador)
	//c.Mapping("AprobarMultiplesPagosContratistas", c.AprobarMultiplesPagosContratistas)
}

// AprobacionPagoController ...
// @Title GetSolicitudesOrdenador
// @Description create GetSolicitudesOrdenador
// @Param docordenador path string true "Número del documento del ordenador"
// @Success 200 {object} []models.PagoPersonaProyecto
// @Failure 403 :docordenador is empty
// @router /:docordenador [get]
func (c *SolicitudesOrdenadorController) GetSolicitudesOrdenador() {

	doc_ordenador := c.GetString(":docordenador")
	//query := c.GetString("query")
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			c.Data["mesaage"] = "Error service Get solicitudes_coordinador: The request contains an incorrect parameter or no record exists"
			c.Abort("404")
		}
	}()

	//var v []models.PagoContratistaCdpRp
	if pagos_personas_proyecto, err := helpers.SolicitudesOrdenador(doc_ordenador, limit, offset); err == nil {
		c.Data["json"] = pagos_personas_proyecto
	} else {
		panic(err)
	}

	c.ServeJSON()

}

// AprobacionPagoController ...
// @Title ObtenerDependenciaOrdenador
// @Description create ObtenerDependenciaOrdenador
// @Param docordenador path string true "Número del documento del ordenador"
// @Success 200  int
// @Failure 403 :docordenador is empty
// @router /dependencia_ordenador/:docordenador [get]
func (c *SolicitudesOrdenadorController) ObtenerDependenciaOrdenador() {

	doc_ordenador := c.GetString(":docordenador")

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			c.Data["mesaage"] = "Error service Get solicitudes_coordinador: The request contains an incorrect parameter or no record exists"
			c.Abort("404")
		}
	}()

	if dependenciaId, err := helpers.DependenciaOrdenador(doc_ordenador); err == nil {
		c.Data["json"] = dependenciaId
	} else {
		panic(err)
	}

	c.ServeJSON()

}
