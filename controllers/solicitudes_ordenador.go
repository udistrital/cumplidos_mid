package controllers

import (
	_ "encoding/json"
	_ "time"

	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
)

// SolicitudesOrdenadorContratistasController operations for SolicitudesOrdenadorContratistas
type SolicitudesOrdenadorController struct {
	beego.Controller
}

//URLMapping ...
func (c *SolicitudesOrdenadorController) URLMapping() {
	c.Mapping("GetSolicitudesOrdenador", c.GetSolicitudesOrdenador)
	c.Mapping("GetSolicitudesOrdenadorContratistas", c.ObtenerDependenciaOrdenador)
	c.Mapping("ObtenerInfoOrdenador", c.ObtenerInfoOrdenador)
	//c.Mapping("AprobarMultiplesPagosContratistas", c.AprobarMultiplesPagosContratistas)
}

// AprobacionPagoController ...
// @Title GetSolicitudesOrdenador
// @Description Trae todas las solicitudes AD aprobadas por el decano que son responsabilidad de un ordenador (Aparentemente solo para docentes)
// @Param docordenador path string true "Número del documento del ordenador"
// @Success 200 {object} []models.PagoPersonaProyecto
// @Failure 403 :docordenador is empty
// @router /solicitudes/:docordenador [get]
func (c *SolicitudesOrdenadorController) GetSolicitudesOrdenador() {

	doc_ordenador := c.GetString(":docordenador")
	//query := c.GetString("query")
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	//defer helpers.GestionError(c)
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "SolicitudesOrdenadorController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	_, err := strconv.Atoi(doc_ordenador)

	if err != nil {
		panic(map[string]interface{}{"funcion": "GetSolicitudesOrdenador", "err": "Error en los parametros de ingreso", "status": "400"})
	}
	//var v []models.PagoContratistaCdpRp
	if pagos_personas_proyecto, err := helpers.SolicitudesOrdenador(doc_ordenador, limit, offset); err == nil {
		c.Data["json"] = pagos_personas_proyecto
	} else {
		panic(err)
	}
	c.ServeJSON()

}

// AprobacionPagoController ...
// @Title ObtenerDependenciaOrdenador
// @Description create ObtenerDependenciaOrdenador
// @Param docordenador path string true "Número del documento del ordenador"
// @Success 200  int
// @Failure 403 :docordenador is empty
// @router /dependencia_ordenador/:docordenador [get]
func (c *SolicitudesOrdenadorController) ObtenerDependenciaOrdenador() {

	doc_ordenador := c.GetString(":docordenador")

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "SolicitudesOrdenadorController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	_, err := strconv.Atoi(doc_ordenador)

	if err != nil {
		panic(map[string]interface{}{"funcion": "ObtenerDependenciaOrdenador", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if dependenciaId, err := helpers.DependenciaOrdenador(doc_ordenador); err == nil {
		c.Data["json"] = dependenciaId
	} else {
		panic(err)
	}

	c.ServeJSON()

}

// AprobacionPagoController ...
// @Title ObtenerInfoOrdenador
// @Description ObtenerInfoOrdenador trae la informacion de un ordenador del gasto a partir de su numero de contrato y ano de vigencia
// @Param numero_contrato path string true "Numero de contrato en la tabla contrato general"
// @Param vigencia path int true "Vigencia del contrato en la tabla contrato general"
// @Success 201 {int} models.InformacionOrdenador
// @Failure 403 :numero_contrato is empty
// @Failure 403 :vigencia is empty
// @router /informacion_ordenador/:numero_contrato/:vigencia [get]
func (c *SolicitudesOrdenadorController) ObtenerInfoOrdenador() {
	numero_contrato := c.GetString(":numero_contrato")
	vigencia := c.GetString(":vigencia")

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "SolicitudesOrdenadorController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	_, err1 := strconv.Atoi(vigencia)

	if (err1 != nil) || (len(vigencia) != 4) {
		panic(map[string]interface{}{"funcion": "ObtenerDependenciaOrdenador", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if informacion_ordenador, err := helpers.TraerInfoOrdenador(numero_contrato, vigencia); err == nil {
		c.Data["json"] = informacion_ordenador
	} else {
		panic(err)
	}
	c.ServeJSON()
}
