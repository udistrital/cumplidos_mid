package controllers

import (
	"encoding/json"
	_ "encoding/json"
	_ "fmt"
	"strconv"
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
// @Failure 404 not found resource
// @router /solicitudes/:docordenador [get]
func (c *SolicitudesOrdenadorContratistasController) GetSolicitudesOrdenadorContratistas() {

	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	//defer helpers.GestionError(c)
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "SolicitudesOrdenadorContratistasController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	if len(c.GetString(":docordenador")) < 3 {
		panic(map[string]interface{}{"funcion": "GetSolicitudesOrdenadorContratistas", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if pagos_contratista_cdp_rp, err := helpers.SolicitudesOrdenadorContratistas(c.GetString(":docordenador"), limit, offset); err != nil {
		panic(err)

	} else {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": pagos_contratista_cdp_rp}
	}

	c.ServeJSON()

}

// AprobacionPagoController ...
// @Title AprobarMultiplesPagosContratistas
// @Description create AprobarMultiplesPagosContratistas
// @Param	body		body 	[]models.PagoContratistaCdpRp	true		"body for SoportePagoMensual content"
// @Success 201			Ok
// @Failure 400 the request contains incorrect syntax
// @router /aprobar_pagos [post]
func (c *SolicitudesOrdenadorContratistasController) AprobarMultiplesPagosContratistas() {
	//defer helpers.GestionError(c)
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "SolicitudesOrdenadorContratistasController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	var v []models.PagoContratistaCdpRp
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := helpers.AprobacionPagosContratistas(v); err == nil {
			//c.Data["json"] = "OK"
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Successful", "Data": "OK"}
		} else {
			panic(err)
		}
	} else {
		panic(map[string]interface{}{"funcion": "AprobarMultiplesPagosContratistas", "err": err.Error(), "status": "400"})
	}

	c.ServeJSON()
}

// AprobacionPagoController ...
// @Title certificacion_cumplidos_contratistas
// @Description get certificacion_cumplidos_contratistas
// @Param dependencia path string true "Dependencia supervisor"
// @Param mes path string true "Mes del certificado"
// @Param ano path string true "Año del certificado "
// @Success 200 {object} []models.Persona
// @Failure 404 not found resource
// @router /certificaciones/:dependencia/:mes/:ano [get]
func (c *SolicitudesOrdenadorContratistasController) CertificacionCumplidosContratistas() {

	dependencia := c.GetString(":dependencia")
	mes := c.GetString(":mes")
	ano := c.GetString(":ano")

	//función que maneja el error
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "SolicitudesOrdenadorContratistasController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	//Validación de parametros de entrada
	//_, err1 := strconv.Atoi(dependencia)
	mess, err1 := strconv.Atoi(mes)
	_, err2 := strconv.Atoi(ano)
	if (mess == 0) || (len(ano) != 4) || (mess > 12) || (err1 != nil) || (err2 != nil) {
		panic(map[string]interface{}{"funcion": "CertificacionCumplidosContratistas", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	//var v []models.PagoContratistaCdpRp
	if personas, err := helpers.CertificacionCumplidosContratistas(dependencia, mes, ano); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": personas}
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
// @Success 200 {object} []models.PagoContratistaCdpRp
// @Failure 404 not found resource
// @router /solicitudes_dependencia/:docordenador/:cod_dependencia [get]
func (c *SolicitudesOrdenadorContratistasController) GetSolicitudesOrdenadorContratistasDependencia() {

	doc_ordenador := c.GetString(":docordenador")
	cod_dependencia := c.GetString(":cod_dependencia")
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	//función que maneja el error
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "SolicitudesOrdenadorContratistasController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	//Validación de parametros de entrada

	if len(doc_ordenador) < 2 {
		panic(map[string]interface{}{"funcion": "GetSolicitudesOrdenadorContratistasDependencia", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if pagos_contratista_cdp_rp, err := helpers.ContratosContratistaDependencia(doc_ordenador, cod_dependencia, limit, offset); err != nil || len(pagos_contratista_cdp_rp) == 0 {
		if err == nil {
			panic(map[string]interface{}{"funcion": "GetSolicitudesOrdenadorContratistasDependencia", "err": "No se encontraron registros"})
		} else {
			panic(err)
		}
	} else {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": pagos_contratista_cdp_rp}
	}
	c.ServeJSON()

}
