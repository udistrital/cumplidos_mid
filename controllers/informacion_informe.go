package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
)

// InformacionInformeController operations for InformacionInforme
type InformacionInformeController struct {
	beego.Controller
}

// URLMapping ...
func (c *InformacionInformeController) URLMapping() {
	c.Mapping("GetOne", c.GetInformacionInforme)

}

// GetInformacionInforme ...
// @Title GetInformacionInforme
// @Description get InformacionInforme by num_documento
// @Param	num_documento	path 	string	true		"numero documento contratista"
// @Param	contrato	path 	string	true		"numero de contrato"
// @Param	vigencia	path 	string	true		"vigencia del contrato"
// @Param	cdp	path 	string	true		"cdp del contrato"
// @Param	vigencia_cdp	path 	string	true		"vigencia del cdp del contrato"
// @Success 200 {object} models.InformacionInforme
// @Failure 403 :num_documento is empty
// @Failure 403 :contrato is empty
// @Failure 403 :vigencia is empty
// @Failure 403 :cdp is empty
// @Failure 403 :vigencia_cdp is empty
// @router /:num_documento/:contrato/:vigencia/:cdp/:vigencia_cdp [get]
func (c *InformacionInformeController) GetInformacionInforme() {

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

	fmt.Println("Numero documento:", c.GetString(":num_documento"))
	fmt.Println("Contrato:", c.GetString(":contrato"))
	fmt.Println("Vigencia:", c.GetString(":vigencia"))
	fmt.Println("Cdp:", c.GetString(":cdp"))

	num_documento := c.GetString(":num_documento")
	contrato := c.GetString(":contrato")
	vigencia := c.GetString(":vigencia")
	cdp := c.GetString(":cdp")
	vigencia_cdp := c.GetString(":vigencia_cdp")

	_, err := strconv.Atoi(num_documento)
	_, err2 := strconv.Atoi(contrato)
	_, err3 := strconv.Atoi(vigencia)
	_, err4 := strconv.Atoi(cdp)
	_, err5 := strconv.Atoi(vigencia_cdp)

	if err != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || len(num_documento) < 2 {
		panic(map[string]interface{}{"funcion": "GetOne", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if informacion_informe, err := helpers.InformacionInforme(num_documento, contrato, vigencia, cdp, vigencia_cdp); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": informacion_informe}
	} else {
		panic(err)
	}
	c.ServeJSON()

}
