package controllers

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
)

type FiltrosPagosMensualesController struct {
	beego.Controller
}

//URLMapping ...
func (c *FiltrosPagosMensualesController) URLMapping() {
	c.Mapping("GetPagos", c.GetPagos)
}

// GetPagos ...
// @Title GetPagos
// @Description get PagoMensual
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.PagoMensual
// @Failure 400 not found resource
// @router / [get]
func (c *FiltrosPagosMensualesController) GetPagos() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ValidacionFechaCargaCumplidoController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	//Capturar de la URL los datos filtros

	DocumentoPersonaId := c.GetString("DocumentoPersonaId")
	Ano := c.GetString("Ano")
	Mes := c.GetString("Mes")
	EstadoPagoMensualId := c.GetString("EstadoPagoMensualId")
	NumeroContrato := c.GetString("NumeroContrato")

	documentos := stringToSlice(DocumentoPersonaId)
	anios := stringToSlice(Ano)
	convertInt(anios)
	meses := stringToSlice(Mes)
	convertInt(meses)
	estados_pagos := stringToSlice(EstadoPagoMensualId)
	convertInt(estados_pagos)
	numeros_contratos := stringToSlice(NumeroContrato)
	convertInt(numeros_contratos)

	filtrospago, err := helpers.GetPagosFiltrados(numeros_contratos, documentos, anios, meses, estados_pagos)

	if err != nil {
		panic(c.Data)
	} else if filtrospago == nil {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 201, "Message": "No hay datos que coincidan con los filtros", "Data": filtrospago}
	} else {
		c.Data["json"] = map[string]interface{}{"Succes": true, "Status:": 201, "Message": "Consulta exitosa", "Data": filtrospago}
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

//Funcion para Verificar que se ingresen datos correctos cuando se necesita un

func convertInt(data []string) {
	for _, str := range data {
		_, err := strconv.Atoi(str)
		if err != nil && len(data) > 0 {
			panic(map[string]interface{}{"funcion: ": "convertInt", "err": "El valor " + str + "no es un número"})
		}
	}
}
