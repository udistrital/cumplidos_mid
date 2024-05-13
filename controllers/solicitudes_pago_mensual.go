package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/helpers"
)

type SolicitudesPagoMensualController struct {
	beego.Controller
}

// URLMapping maps the SolicitudesPagoMensualController methods to GET requests.
func (c *SolicitudesPagoMensualController) URLMapping() {
	c.Mapping("GetSolicitudesPagoMensual", c.GetSolicitudesPagoMensual)
}

// GetSolicitudesPagoMensual handles POST requests to filter payment requests.
// @Title GetSolicitudesPagoMensual
// @Description filter payment requests by various parameters
// @Param   codigos_dependencias	query	[]string	true	"List of dependency codes"
// @Param   vigencias	query	[]string	true	"List of years"
// @Param   documentos_contratistas	query	[]string	true	"List of contractor IDs"
// @Param   numeros_contratos	query	[]string	true	"List of contract numbers"
// @Param   meses	query	[]string	true	"List of months"
// @Param   anios	query	[]string	true	"List of years"
// @Param   estados	query	[]string	true	"List of payment states"
// @Success 200 {object} []models.SolicitudPago
// @Failure 400 Bad request
// @router / [post]
func (c *SolicitudesPagoMensualController) GetSolicitudesPagoMensual() {

	//Capturar de la URL los datos de los filtros

	DocumentoPersonaId := c.GetString("DocumentoPersonaId")
	NumeroContrato := c.GetString("NumeroContrato")
	Ano := c.GetString("Ano")
	Mes := c.GetString("Mes")
	EstadoPagoMensualId := c.GetString("EstadoPagoMensualId")
	/*
		dependencia := c.GetString("dependencias")
		vigencia := c.GetString("vigencias")
	*/

	documentos := stringToSlice(DocumentoPersonaId)
	anios := stringToSlice(Ano)
	convertInt(anios)
	meses := stringToSlice(Mes)
	convertInt(meses)
	estados_pagos := stringToSlice(EstadoPagoMensualId)
	convertInt(estados_pagos)
	numeros_contratos := stringToSlice(NumeroContrato)
	convertInt(numeros_contratos)
	/*
		dependencias := stringToSlice(dependencia)
		vigencias := stringToSlice(vigencia)
	*/

	//Obtener dependencias y vigencias
	type BodyParams struct {
		Dependencias string `json:"dependencias"`
		Vigencias    string `json:"vigencias"`
	}
	var v BodyParams

	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	dependencias := stringToSlice(v.Dependencias)
	vigencias := stringToSlice(v.Vigencias)

	filtros_pago, err := helpers.SolicitudesPagoMensual(dependencias, vigencias, documentos, numeros_contratos, meses, anios, estados_pagos)

	if err != nil {
		panic(c.Data)
	} else if filtros_pago == nil {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 200, "Message": "No hay datos que coincidan con los filtros", "Data": filtros_pago}
	} else {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 200, "Message": "Consulta exitosa", "Data": filtros_pago}
	}

	c.ServeJSON()
}
