package controllers

import (
	_ "encoding/json"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
)

// AprobacionPagoController operations for AprobacionPago
type CertificacionDocumentosAprobadosController struct {
	beego.Controller
}

//URLMapping ...
func (c *CertificacionDocumentosAprobadosController) URLMapping() {
	c.Mapping("GetCertificacionDocumentosAprobados", c.GetCertificacionDocumentosAprobados)
}

// AprobacionPagoController ...
// @Title CertificacionDocumentosAprobados
// @Description create CertificacionDocumentosAprobados
// @Param dependencia path int true "Dependencia del contrato en la tabla ordenador_gasto"
// @Param mes path int true "Mes del pago mensual"
// @Param anio path int true "Año del pago mensual"
// @Success 201
// @Failure 403 :dependencia is empty
// @Failure 403 :mes is empty
// @Failure 403 :anio is empty
// @router /:dependencia/:mes/:anio [get]
func (c *CertificacionDocumentosAprobadosController) GetCertificacionDocumentosAprobados() {

	dependencia := c.GetString(":dependencia")
	mes := c.GetString(":mes")
	anio := c.GetString(":anio")

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			c.Data["mesaage"] = "Error service Get solicitudes_coordinador: The request contains an incorrect parameter or no record exists"
			c.Abort("404")
		}
	}()

	//var v []models.PagoContratistaCdpRp
	if personas, err := helpers.CertificacionDocumentosAprobados(dependencia, anio, mes); err == nil {
		c.Data["json"] = personas
	} else {
		panic(err)
	}

	c.ServeJSON()

}
