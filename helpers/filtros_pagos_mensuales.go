package helpers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

//Funcion para construir el query dinamico

func build_query(slices []string, columna string) string {

	query := ""

	if len(slices) == 1 {
		query += fmt.Sprintf("%s.in:%v,", columna, slices[0])
	}
	if len(slices) > 1 {
		for i, dato := range slices {
			if i == 0 {
				query += fmt.Sprintf("%s.in:%v|", columna, dato)
			} else if i < len(slices)-1 {
				query += fmt.Sprintf("%s|", dato)
			} else {
				query += fmt.Sprintf("%s,", dato)
			}
		}
		return query
	}
	return query
}

func GetPagosFiltrados(numeros_contratos []string, numeros_documentos []string, anios []string, meses []string, estados_pagos []string) (PagoMensual []models.PagoMensual, outputError interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError := map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al filtrar los pagos",
				"Funcion": "GetPagosFiltrados",
				"Error":   err,
			}
			panic(outputError)
		}
	}()

	var respuesta_peticion map[string]interface{}

	//Se contruye dinamicamente el query

	query := strings.TrimSuffix(("?query=" + build_query(numeros_contratos, "NumeroContrato") + build_query(numeros_documentos, "DocumentoPersonaId") +
		build_query(anios, "Ano") + build_query(meses, "Mes") + build_query(estados_pagos, "EstadoPagoMensualId__Id")), ",")
	order := "&order=desc"
	sortby := "&sortby=Ano"
	limit := "&limit=0"

	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/pago_mensual/"+query+sortby+order+limit, &respuesta_peticion); (err == nil) && (response == 200) {

		if respuesta_peticion != nil {
			LimpiezaRespuestaRefactor(respuesta_peticion, &PagoMensual)
		} else {
			return nil, outputError
		}

	}
	return PagoMensual, nil

}
