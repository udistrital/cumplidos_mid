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
}

func (c *HistoricoCumplidos) GetCambioEstado() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = beego.AppConfig.String("appname") + "/historico/cambio-estado/" + "/" + localError["funcion"].(string)
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
