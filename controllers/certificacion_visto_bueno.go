package controllers

import (

	//"net/http"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
)

// CertificacionVistoBuenoController operations for certificacion_visto_bueno
type CertificacionVistoBuenoController struct {
	beego.Controller
}

// URLMapping ...
func (c *CertificacionVistoBuenoController) URLMapping() {
	c.Mapping("CertificacionVistoBueno", c.CertificacionVistoBueno)
}

// CertificacionVistoBuenoController ...
// @Title CertificacionVistoBueno
// @Description create CertificacionVistoBueno
// @Param dependencia path int true "Dependencia del contrato en la tabla vinculacion"
// @Param mes path int true "Mes del pago mensual"
// @Param anio path int true "Año del pago mensual"
// @Success 200 {object} []models.Persona
// @Failure 404 not found source
// @router /certificacion_visto_bueno/:dependencia/:mes/:anio [get]
func (c *CertificacionVistoBuenoController) CertificacionVistoBueno() {
	dependencia := c.GetString(":dependencia")
	mes := c.GetString(":mes")
	anio := c.GetString(":anio")

	if personas, err := helpers.CertificadoVistoBueno(dependencia, mes, anio); err != nil || len(personas) == 0 {
		logs.Error(err)
		c.Data["mesaage"] = "Error service Get CertificacionVistoBueno: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	} else {
		c.Data["json"] = personas
	}
	c.ServeJSON()
}
