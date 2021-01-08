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
// @Param ano path int true "Año del pago mensual"
// @Success 200 {object} []models.Persona
// @Failure 404 not found resource
// @router /documentos_aprobados/:dependencia/:mes/:ano [get]
func (c *CertificacionController) GetCertificacionDocumentosAprobados() {

	dependencia := c.GetString(":dependencia")
	mes := c.GetString(":mes")
	ano := c.GetString(":ano")

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "CertificacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	_, err1 := strconv.Atoi(dependencia)
	mess, err2 := strconv.Atoi(mes)
	_, err3 := strconv.Atoi(ano)
	if (mess == 0) || (len(ano) != 4) || (mess > 12) || (err1 != nil) || (err2 != nil) || (err3 != nil) {
		panic(map[string]interface{}{"funcion": "GetCertificacionDocumentosAprobados", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if personas, err := helpers.CertificacionDocumentosAprobados(dependencia, ano, mes); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": personas}
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
// @Param ano path int true "Año del pago mensual"
// @Success 200 {object} []models.Persona
// @Failure 404 not found source
// @router /certificacion_visto_bueno/:dependencia/:mes/:ano [get]
func (c *CertificacionController) CertificacionVistoBueno() {
	dependencia := c.GetString(":dependencia")
	mes := c.GetString(":mes")
	ano := c.GetString(":ano")

	//función que maneja el error
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "CertificacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	//Validación de parametros de entrada
	_, err1 := strconv.Atoi(dependencia)
	mess, err2 := strconv.Atoi(mes)
	_, err3 := strconv.Atoi(ano)
	if (mess == 0) || (len(ano) != 4) || (mess > 12) || (err1 != nil) || (err2 != nil) || (err3 != nil) {
		panic(map[string]interface{}{"funcion": "CertificacionVistoBueno", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if personas, err := helpers.CertificadoVistoBueno(dependencia, mes, ano); err != nil || len(personas) == 0 {
		if err == nil {
			panic(map[string]interface{}{"funcion": "CertificacionVistoBueno", "err": "No se encontraron registros"})
		} else {
			panic(err)
		}

	} else {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": personas}
	}
	c.ServeJSON()
}
