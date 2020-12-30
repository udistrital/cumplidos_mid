package controllers

import (
	_ "encoding/json"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
)

// AprobacionPagoController operations for AprobacionPago
type SolicitudesCoordinadorController struct {
	beego.Controller
}

//URLMapping ...
func (c *SolicitudesCoordinadorController) URLMapping() {
	c.Mapping("GetSolicitudesCoordinador", c.GetSolicitudesCoordinador)
}

// AprobacionPagoController ...
// @Title GetSolicitudesCoordinador
// @Description create GetSolicitudesCoordinador
// @Param doccoordinador path string true "Número del documento del coordinador"
// @Success 200 {object} []models.PagoPersonaProyecto
// @Failure 404 not found resource
// @router /:doccoordinador [get]
func (c *SolicitudesCoordinadorController) GetSolicitudesCoordinador() {
	doc_coordinador := c.GetString(":doccoordinador")
	//función que maneja el error
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "SolicitudesCoordinadorController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	//Validación de parametros de entrada

	//docc_coordinador, err1 := strconv.Atoi(doc_coordinador)
	if len(doc_coordinador) < 2 {
		panic(map[string]interface{}{"funcion": "GetSolicitudesCoordinador", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if pagos_personas_proyecto, err := helpers.SolicitudCoordinador(doc_coordinador); err != nil || len(pagos_personas_proyecto) == 0 {
		if err == nil {
			panic(map[string]interface{}{"funcion": "GetSolicitudesCoordinador", "err": "No se encontraron registros"})
		} else {
			panic(err)
		}

	} else {
		c.Data["json"] = pagos_personas_proyecto
	}

	c.ServeJSON()

}
