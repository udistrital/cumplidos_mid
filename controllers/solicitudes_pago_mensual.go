package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

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

	dependencias := stringToSlice(v.Dependencias)
	vigencias := stringToSlice(v.Vigencias)
	convertInt(vigencias)
	documentos_persona_id := stringToSlice(v.DocumentosPersonaId)
	convertInt(documentos_persona_id)
	numeros_contratos := stringToSlice(v.NumerosContratos)
	convertInt(numeros_contratos)
	meses := stringToSlice(v.Meses)
	convertInt(numeros_contratos)
	anios := stringToSlice(v.Anios)
	convertInt(anios)
	estados_pagos := stringToSlice(v.EstadosPagos)

	filtros_pago, err := helpers.SolicitudesPagoMensual(dependencias, vigencias, documentos_persona_id, numeros_contratos, meses, anios, estados_pagos)

	if err != nil {
		panic(c.Data)
	} else if filtros_pago == nil {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 200, "Message": "No hay datos que coincidan con los filtros", "Data": filtros_pago}
	} else {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 200, "Message": "Consulta exitosa", "Data": filtros_pago}
	}

	c.ServeJSON()
}

// Funcion para agregar los datos a un slice
func stringToSlice(cadena string) (slice []string) {
	parts := strings.Split(cadena, ",")

	if cadena != "" {
		for _, part := range parts {
			slice = append(slice, part)
		}
	}
	return slice
}

//Funcion para Verificar que se ingresen datos correctos cuando el parametro sean números

func convertInt(data []string) {
	for _, str := range data {
		_, err := strconv.Atoi(str)
		if err != nil && len(data) > 0 {
			panic(map[string]interface{}{"funcion: ": "convertInt", "err": "El valor " + str + "no es un número"})
		}
	}
}
