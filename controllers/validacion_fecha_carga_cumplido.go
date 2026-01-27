package controllers

import (

	//"net/http"

	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
)

// ValidacionFechaCargaCumplidoController operations for validacion_fecha_carga_cumplido
type ValidacionFechaCargaCumplidoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ValidacionFechaCargaCumplidoController) URLMapping() {
	c.Mapping("GetValidacionPeriodo", c.GetValidacionPeriodo)
}

// GetValidacionPeriodo ...
// @Title GetValidacionPeriodo
// @Description valida si se esta dentro del periodo de carga de cumplidos
// @Param dependencia_supervisor path string true "Dependencia supervisor"
// @Param anio path string true "anio"
// @Param mes path string true "mes"
// @Success 200 {object} []models.ValidacionFechaCargaCumplido
// @Failure 404 not found resource
// @router /:dependencia_supervisor/:anio/:mes [get]
func (c *ValidacionFechaCargaCumplidoController) GetValidacionPeriodo() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ValidacionFechaCargaCumplidoController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	dependencia_supervisor := c.GetString(":dependencia_supervisor")
	anio := c.GetString(":anio")
	mes := c.GetString(":mes")

	_, err_anio := strconv.Atoi(anio)
	_, err_mes := strconv.Atoi(mes)

	if err_anio != nil || err_mes != nil || len(anio) != 4 || len(mes) > 2 {
		panic(map[string]interface{}{"funcion": "GetValidacionPeriodo", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if validacion, err := helpers.ValidarPeriodoCargaCumplido(dependencia_supervisor, anio, mes); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": validacion}
	} else {
		panic(err)
	}
	c.ServeJSON()
}
