package controllers

import (
	_ "time"

	//"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
	_ "github.com/udistrital/cumplidos_mid/helpers"
)

// ObtenerInfoOrdenadorController operations for obtener_info_ordenador
type ObtenerInfoOrdenadorController struct {
	beego.Controller
}

// URLMapping ...
func (c *ObtenerInfoOrdenadorController) URLMapping() {
	c.Mapping("ObtenerInfoOrdenador", c.ObtenerInfoOrdenador)
}

// AprobacionPagoController ...
// @Title ObtenerInfoOrdenador
// @Description create ObtenerInfoOrdenador
// @Param numero_contrato path string true "Numero de contrato en la tabla contrato general"
// @Param vigencia path int true "Vigencia del contrato en la tabla contrato general"
// @Success 201 {int} models.InformacionOrdenador
// @Failure 403 :numero_contrato is empty
// @Failure 403 :vigencia is empty
// @router /informacion_ordenador/:numero_contrato/:vigencia [get]
func (c *ObtenerInfoOrdenadorController) ObtenerInfoOrdenador() {
	numero_contrato := c.GetString(":numero_contrato")
	vigencia := c.GetString(":vigencia")

	if informacion_ordenador, err := helpers.TraerInfoOrdenador(numero_contrato, vigencia); err != nil {
		logs.Error(err)
		c.Data["mesaage"] = "Error service Get SolicitudesOrdenadorContratistaDependencia: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	} else {
		c.Data["json"] = informacion_ordenador
	}
	c.ServeJSON()
}
