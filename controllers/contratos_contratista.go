package controllers

import (

	//"net/http"

	_ "fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
)

// ContratosContratistaController operations for contratos_contratista
type ContratosContratistaController struct {
	beego.Controller
}

//URLMapping ...
func (c *ContratosContratistaController) URLMapping() {
	c.Mapping("GetContratosContratista", c.GetContratosContratista)
}

// GetContratosContratista ...
// @Title GetContratosContratista
// @Description create ContratosContratista
// @Param numero_documento path string true "Número documento"
// @Success 200 {object} []models.ContratoDisponibilidadRp
// @Failure 404 not found resource
// @router /:numero_documento [get]
func (c *ContratosContratistaController) GetContratosContratista() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ContratosContratistaController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	numero_documento := c.GetString(":numero_documento")
	_, err := strconv.Atoi(numero_documento)

	if err != nil || len(numero_documento) < 2 {
		panic(map[string]interface{}{"funcion": "GetContratosContratista", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if contratos_disponibilidad_rp, err := helpers.ContratosContratista(numero_documento); (err == nil) || (len(contratos_disponibilidad_rp) != 0) {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": contratos_disponibilidad_rp}
	} else {
		panic(err)
	}
	c.ServeJSON()

}
