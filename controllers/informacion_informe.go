package controllers

import (
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
	c.Mapping("GetOne", c.GetPreliquidacion)
}

// GetInformacionInforme ...
// @Title GetInformacionInforme
// @Description get InformacionInforme by pago_mensual_id
// @Param	pago_mensual_id	path 	string	true		"id del pago mensual"
// @Success 200 {object} models.InformacionInforme
// @Failure 403 :pago_mensual_id is empty
// @router /:num_documento/:pago_mensual_id[get]
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

	pago_mensual_id := c.GetString(":pago_mensual_id")

	_, err := strconv.Atoi(pago_mensual_id)

	if err != nil {
		panic(map[string]interface{}{"funcion": "GetOne", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if informacion_informe, err := helpers.InformacionInforme(pago_mensual_id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": informacion_informe}
	} else {
		panic(err)
	}
	c.ServeJSON()

}

// GetPreliquidacion ...
// @Title GetPreliquidacion
// @Description get preliquidacion correspondiente a una solicitud de pago mensual
// @Param	pago_mensual_id	path 	string	true		"id del pago mensual"
// @Success 200 {object} models.PreliquidacionTitan
// @Failure 403 :pago_mensual_id is empty
// @router /preliquidacion/:pago_mensual_id [get]
func (c *InformacionInformeController) GetPreliquidacion() {

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

	pago_mensual_id := c.GetString(":pago_mensual_id")
	//fmt.Println("Pago mensual id:", pago_mensual_id)

	_, err := strconv.Atoi(pago_mensual_id)

	if err != nil {
		panic(map[string]interface{}{"funcion": "GetOne", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if preliquidacion, err := helpers.GetPreliquidacion(pago_mensual_id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": preliquidacion}
	} else {
		panic(err)
	}
	c.ServeJSON()

}
