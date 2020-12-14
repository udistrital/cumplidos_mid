package controllers

import (
	"encoding/json"
	_ "encoding/json"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
	_ "github.com/udistrital/cumplidos_mid/helpers"
	"github.com/udistrital/cumplidos_mid/models"
)

// SolicitudesOrdenadorContratistasController operations for SolicitudesOrdenadorContratistas
type SolicitudesOrdenadorContratistasController struct {
	beego.Controller
}

//URLMapping ...
func (c *SolicitudesOrdenadorContratistasController) URLMapping() {
	c.Mapping("GetSolicitudesOrdenadorContratistas", c.GetSolicitudesOrdenadorContratistas)
	c.Mapping("AprobarMultiplesPagosContratistas", c.AprobarMultiplesPagosContratistas)
	c.Mapping("CertificacionCumplidosContratistas", c.CertificacionCumplidosContratistas)
}

// AprobacionPagoController ...
// @Title GetSolicitudesOrdenadorContratistas
// @Description create GetSolicitudesOrdenadorContratistas
// @Param docordenador path string true "Número del documento del supervisor"
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.PagoContratistaCdpRp
// @Failure 403 :docordenador is empty
// @router /solicitudes/:docordenador [get]
func (c *SolicitudesOrdenadorContratistasController) GetSolicitudesOrdenadorContratistas() {

	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")
	if pagos_contratista_cdp_rp, err := helpers.SolicitudesOrdenadorContratistas(c.GetString(":docordenador"), limit, offset); err != nil {
		logs.Error(err)
		c.Data["mesaage"] = "Error service Get solicitudes_coordinador: The request contains an incorrect parameter or no record exists"
		c.Abort("404")

	} else {
		c.Data["json"] = pagos_contratista_cdp_rp
	}

	c.ServeJSON()

}

// AprobacionPagoController ...
// @Title AprobarMultiplesPagosContratistas
// @Description create AprobarMultiplesPagosContratistas
// @Param	body		body 	[]models.PagoContratistaCdpRp	true		"body for SoportePagoMensual content"
// @Success 201
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *SolicitudesOrdenadorContratistasController) AprobarMultiplesPagosContratistas() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			c.Data["mesaage"] = "Error service Get solicitudes_coordinador: The request contains an incorrect parameter or no record exists"
			c.Abort("404")
		}
	}()

	var v []models.PagoContratistaCdpRp
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := helpers.AprobacionPagosContratistas(v); err == nil {
			c.Data["json"] = "OK"
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}

	c.ServeJSON()
}

// AprobacionPagoController ...
// @Title certificacion_cumplidos_contratistas
// @Description get certificacion_cumplidos_contratistas
// @Param dependencia path string true "Dependencia supervisor"
// @Param mes path string true "Mes del certificado"
// @Param anio path string true "Año del certificado "
// @Success 201 {object} []models.Persona
// @Failure 403 :dependencia is empty
// @Failure 403 :mes is empty
// @Failure 403 :anio is empty
// @router /certificaciones/:dependencia/:mes/:anio [get]
func (c *SolicitudesOrdenadorContratistasController) CertificacionCumplidosContratistas() {

	dependencia := c.GetString(":dependencia")
	mes := c.GetString(":mes")
	anio := c.GetString(":anio")

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			c.Data["mesaage"] = "Error service Get solicitudes_coordinador: The request contains an incorrect parameter or no record exists"
			c.Abort("404")
		}
	}()

	//var v []models.PagoContratistaCdpRp
	if personas, err := helpers.CertificacionCumplidosContratistas(dependencia, mes, anio); err == nil {
		c.Data["json"] = personas
	} else {
		panic(err)
	}

	c.ServeJSON()
}

// AprobacionPagoController ...
// @Title GetSolicitudesOrdenadorContratistasDependencia
// @Description create GetSolicitudesOrdenadorContratistasDependencia
// @Param docordenador path string true "Número del documento del supervisor"
// @Param cod_dependencia path string true "cod_dependencia"
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 201
// @Failure 403 :docordenador is empty
// @router /solicitudes_ordenador_contratistas_dependencia/:docordenador/:cod_dependencia [get]
func (c *SolicitudesOrdenadorContratistasController) GetSolicitudesOrdenadorContratistasDependencia() {

	doc_ordenador := c.GetString(":docordenador")
	cod_dependencia := c.GetString(":cod_dependencia")
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")
	var alerta models.Alert

	if pagos_contratista_cdp_rp, err := helpers.ContratosContratistaDependencia(doc_ordenador, cod_dependencia, limit, offset); err != nil || len(pagos_contratista_cdp_rp) == 0 {
		logs.Error(err)
		c.Data["json"] = alerta
		c.Data["mesaage"] = "Error service Get SolicitudesOrdenadorContratistaDependencia: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	} else {
		c.Data["json"] = pagos_contratista_cdp_rp
	}
	c.ServeJSON()

}
