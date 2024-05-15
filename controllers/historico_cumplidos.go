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
	c.Mapping("GetDependencias", c.GetDependencias)
}

func (c *HistoricoCumplidos) GetDependencias() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["message"] = beego.AppConfig.String("appname") + "/historicos/cambio-estado/" + "/" + localError["funcion"].(string)
			c.Data["data"] = localError["err"]
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	documento := c.GetString(":documento")

	dependencias, err := helpers.ObtenerDependencias(documento)

	if err != nil {
		panic(c.Data)
	} else if dependencias == nil {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 204, "Message": "No hay datos", "Data": dependencias}
	} else {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 200, "Message": "Consulta completa", "Data": dependencias}
	}
	c.ServeJSON()

}
