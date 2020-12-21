package controllers

import (
	_ "encoding/json"
	"strconv"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
)

// AprobacionPagoController operations for AprobacionPago
type CertificacionController struct {
	beego.Controller
}

//URLMapping ...
func (c *CertificacionController) URLMapping() {
	c.Mapping("GetCertificacionDocumentosAprobados", c.GetCertificacionDocumentosAprobados)
	c.Mapping("CertificacionVistoBueno", c.CertificacionVistoBueno)
}

// AprobacionPagoController ...
// @Title CertificacionDocumentosAprobados
// @Description create CertificacionDocumentosAprobados  trae
// @Param dependencia path int true "Dependencia del contrato en la tabla ordenador_gasto"
// @Param mes path int true "Mes del pago mensual"
// @Param anio path int true "Año del pago mensual"
// @Success 201
// @Failure 403 :dependencia is empty
// @Failure 403 :mes is empty
// @Failure 403 :ano is empty
// @router /documentos_aprobados/:dependencia/:mes/:ano [get]
func (c *CertificacionController) GetCertificacionDocumentosAprobados() {

	dependencia := c.GetString(":dependencia")
	mes := c.GetString(":mes")
	ano := c.GetString(":ano")

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			respuesta := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "CertificacionController" + "/" + (respuesta["funcion"]).(string))
			c.Data["data"] = (respuesta["err"])
			c.Abort("404")
		}
	}()

	_, err1 := strconv.Atoi(dependencia)
	mess, err2 := strconv.Atoi(mes)
	_, err3 := strconv.Atoi(ano)
	if (mess == 0) || (len(ano) != 4) || (mess > 12) || (err1 != nil) || (err2 != nil) || (err3 != nil) {
		panic(map[string]interface{}{"funcion": "GetCertificacionDocumentosAprobados", "err": "Error en los parametros de ingreso"})
	}

	if personas, err := helpers.CertificacionDocumentosAprobados(dependencia, ano, mes); err == nil {
		c.Data["json"] = personas
	} else {
		panic(err)
	}

	c.ServeJSON()

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
func (c *CertificacionController) CertificacionVistoBueno() {
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
