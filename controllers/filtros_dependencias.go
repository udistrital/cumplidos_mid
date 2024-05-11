package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/helpers"
)

type FiltrosDependenciasController struct {
	beego.Controller
}

//URLMapping ...
func (c *FiltrosDependenciasController) URLMapping() {
	c.Mapping("FiltrarPorDependencias", c.FiltrarPorDependencias)
}

// Post ...
// @Title FiltrarPorDependencias
// @Description filter contracts and periods of validity by units and periods of validity
// @Param   body     body    BodyParams  true   "body for units content"
// @Success 201 {int} models.ContratoDependencia
// @Failure 403 body is empty
// @router / [post]
func (c *FiltrosDependenciasController) FiltrarPorDependencias() {

	type BodyParams struct {
		Dependencias string `json:"dependencias"`
		Vigencias    string `json:"vigencias"`
	}
	var v BodyParams

	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	dependencias := stringToSlice(v.Dependencias)
	vigencias := stringToSlice(v.Vigencias)

	if response, err := helpers.FiltrosDependencia(dependencias, vigencias); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Successful", "Data": response}

	} else {
		c.Data["json"] = map[string]interface{}{"Success": false, "Status": 404, "Function": "FiltrosDependencia", "Error": err, "Data": response}
		panic(err)
	}

	c.ServeJSON()

}
