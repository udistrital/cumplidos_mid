package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
)

type DescargarDocumentosSolicitudesPagosController struct {
	beego.Controller
}

func (c *DescargarDocumentosSolicitudesPagosController) URLMapping() {
	c.Mapping("GetDocumentosPagoZip", c.GetDocumentosPagoZip)
}

// GetDocumentosPagoMensual ...
// @Title GetDocumentosPagoMensual
// @Description Download in a .zip file all the documents uploaded in the payment request form
// @Param id_pago_mensual path string true "Id Pago Mensual"
// @Success 200
// @Failure 404 not found resource
// @router /:pago_mensual_id [get]
func (c *DescargarDocumentosSolicitudesPagosController) GetDocumentosPagoZip() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "DescargarDocumentosSolicitudesPagosController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	pago_mensual_id := c.GetString(":pago_mensual_id")
	_, err := strconv.Atoi(pago_mensual_id)

	if err != nil {
		panic(map[string]interface{}{"funcion": "GetDocumentosPagoMensual", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if data, err := helpers.DescargarDocumentosSolicitudesPagos(pago_mensual_id); err == nil {
		if data.Nombre != "___0_0" {
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": data}
		} else {
			c.Ctx.Output.SetStatus(204)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": "Ningun Pago coincide con el parametro de busqueda"}
		}

	} else {
		panic(err)
	}
	c.ServeJSON()

}
