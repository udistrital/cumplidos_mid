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

	type BodyParams struct {
		Dependencias        string `json:"dependencias"`
		Vigencias           string `json:"vigencias"`
		DocumentosPersonaId string `json:"documentos_persona_id"`
		NumerosContratos    string `json:"numeros_contratos"`
		Meses               string `json:"meses"`
		Anios               string `json:"anios"`
		EstadosPagos        string `json:"estados_pagos"`
	}
	var v BodyParams

	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	dependencias := helpers.StringToSlice(v.Dependencias)
	vigencias := helpers.StringToSlice(v.Vigencias)
	helpers.ConvertInt(vigencias)
	documentos_persona_id := helpers.StringToSlice(v.DocumentosPersonaId)
	helpers.ConvertInt(documentos_persona_id)
	numeros_contratos := helpers.StringToSlice(v.NumerosContratos)
	helpers.ConvertInt(numeros_contratos)
	meses := helpers.StringToSlice(v.Meses)
	helpers.ConvertInt(numeros_contratos)
	anios := helpers.StringToSlice(v.Anios)
	helpers.ConvertInt(anios)
	estados_pagos := helpers.StringToSlice(v.EstadosPagos)

	filtros_pago, err := helpers.SolicitudesPagoMensual(dependencias, vigencias, documentos_persona_id, numeros_contratos, meses, anios, estados_pagos)

	if err != nil {
		c.Ctx.Output.SetStatus(204)
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 204, "Message": "No hay datos que coincidan con los filtros", "Data": nil}
	} else {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 200, "Message": "Consulta exitosa", "Data": filtros_pago}
	}

	c.ServeJSON()
}
