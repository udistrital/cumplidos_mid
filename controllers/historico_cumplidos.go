package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
)

type HistoricoCumplidos struct {
	beego.Controller
}

func (c *HistoricoCumplidos) URLMapping() {
	c.Mapping("GetCambioEstado", c.GetCambioEstado)
	c.Mapping("GetDependencias", c.GetDependencias)
}

// @Title GetCambioEstado
// @Description get the state change history for a given payment ID
// @Param idPagoMensual path string true "ID of the monthly payment"
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 404 {object} map[string]interface{} "Error"
// @router /cambio_estado_pago/:idPagoMensual [get]
func (c *HistoricoCumplidos) GetCambioEstado() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = beego.AppConfig.String("appname") + "/cambio_estado_pago/" + "/" + localError["funcion"].(string)
			c.Data["data"] = localError["err"]
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	idPagoMensual := c.GetString(":idPagoMensual")

	estadospago, err := helpers.GetEstadosPago(idPagoMensual)

	if err != nil {
		panic(c.Data)
	} else if len(estadospago) < 1 {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 201, "Message": "No hay datos", "Data": estadospago}
	} else {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 201, "Message": "Consulta completa", "Data": estadospago}
	}
	c.ServeJSON()
}

// @Title GetDependencias
// @Description get the dependencies for a given document
// @Param documento path string true "Document number"
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 404 {object} map[string]interface{} "Error"
// @router /dependencias/:documento [get]
func (c *HistoricoCumplidos) GetDependencias() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = beego.AppConfig.String("appname") + "/dependencias/" + "/" + localError["funcion"].(string)
			c.Data["data"] = localError["err"]
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()


	documentoUsuario := c.GetString(":documento")
	dependencias, err := helpers.ObtenerDependencias(documentoUsuario)

	if err != nil {
		panic(c.Data)
	} else if dependencias == nil {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 204, "Message": "No hay datos", "Data": dependencias}
	} else {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 200, "Message": "Consulta completa", "Data": dependencias}
	}
	c.ServeJSON()

}
// @Title GetDependenciasGeneral
// @Description get all  dependencies 
// @Success 200 {object} map[string]interface{} "Success"
// @Failure 404 {object} map[string]interface{} "Error"
// @router /dependenciasgen/
func (c *HistoricoCumplidos) GetDependenciasGeneral() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = beego.AppConfig.String("appname") + "/dependencias/" + "/" + localError["funcion"].(string)
			c.Data["data"] = localError["err"]
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	dependencias, err := helpers.GetDependenciasRolGeneral()

	if err != nil {
		panic(c.Data)
	} else if dependencias == nil {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 204, "Message": "No hay datos en general", "Data": dependencias}
	} else {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 200, "Message": "Consulta completa", "Data": dependencias}
	}
	c.ServeJSON()

}