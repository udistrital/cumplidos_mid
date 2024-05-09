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

func GetPagosFiltrados(numeros_contratos []string, numeros_documentos []string, anios []string, meses []string, estados_pagos []string, vigencias []string) (PagoMensual []models.PagoMensual, outputError interface{}) {
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
		build_query(anios, "Ano") + build_query(meses, "Mes") + build_query(estados_pagos, "EstadoPagoMensualId__Id") + build_query(vigencias, "VigenciaContrato")), ",")
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

//Funcion para filtrar por una lista de dependencias

func GetDependencias(dependencias []string) (contratos []models.ContratoDependencia, err){

	
	
	return contratos, nil
}


// Funcion para filtrar pagos por lista, de codigos dependencias, listas de vigencias, lista de numeros documentos contratistas,
//lista de numeros de contratos, lista de meses, lista de años o listas de id de estados

func GetFiltros(codigos_dependencias []string, vigencias []string, documentos_contratistas []string, numeros_contratos []string, meses []string, anios []string, estados []string ) (pagos []models.PagosFiltrados, outputError map[string]interface{}){

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetFiltros", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var contratistas []models.InformacionProveedor
	var persona models.Persona
	var respuesta_peticion map[string]interface{}

	if pagos_mensuales, outputError := GetPagosFiltrados(numeros_contratos, documentos_contratistas, anios, meses, estados, vigencias); outputError != nil {
		return nil, outputError
	}

	if pagos_dependencias, outputError := GetDependencias(codigos_dependencias); outputError != nil {
		return nil, outputError
	}

	for _, pago_dependencia := range pagos_dependencias{
		for _, pago_mensual := range pagos_mensuales{

			if contratoExists(strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64), pago_mensual.NumeroContrato, pago_dependencia.Contratos.Contrato ){
				if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pago_mensual.DocumentoPersonaId, &contratistas); (err == nil) && (response == 200) {
					var contrato models.InformacionContrato
					contrato, outputError = GetContrato(pago_mensual.NumeroContrato, strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64))
					if outputError == nil{
						for _, contratista := range contratistas {
							persona.NombreDependencia = " "
							persona.Rubro = contrato.Contrato.Rubro
							persona.DocumentoContratista = contratista.NumDocumento
							persona.NombreContratista = contratista.NomProveedor
							persona.Vigencia = int(pago_mensual.VigenciaContrato)
							persona.Ano = int(pago_mensual.Ano)
							persona.Mes = int(pago_mensual.Mes)
							persona.Estado = pago_mensual.EstadoPagoMensual
							pagos = append(pagos, persona)
						}
					}else {
						return nil, outputError
					}
				}else {
					logs.Error(err)
					outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas", "err": err.Error(), "status": "404"}
					return nil, outputError
				}
			}else {
				outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas", "Message": "El contrato buscado no existe", "status": "201"}
			}
		}
	}else{
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas", "err": err.Error(), "status": "404"}
		return nil, outputError
	}

	return
}