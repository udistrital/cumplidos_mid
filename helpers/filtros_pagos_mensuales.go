package helpers

import (
	"fmt"
	"strconv"
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
			outputError = map[string]interface{}{
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

//Funcion para filtrar por una lista de dependencias

func FiltrosDependencia(dependencias []string, vigencias []string) (contratos []models.ContratoSuscritoDependencia, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al filtrar los pagos, Ningun parametro coincide con los filtros",
				"Funcion": "FiltrosDependencia",
			}
			panic(outputError)
		}
	}()

	type BodyParams struct {
		Dependencias string `json:"dependencias"`
		Vigencias    string `json:"vigencias"`
	}

	//Se agregan los años actuales de vigencias de la OATI, en un futuro cuando hayan más solo se debe agregar el año al slice

	var vigencias_oficina = []string{"2017", "2018", "2019", "2020", "2021", "2022", "2023"}

	var parametro_dependencia string
	var parametro_vigencia string

	if len(dependencias) == 0 {
		outputError := map[string]interface{}{
			"Succes":  false,
			"Status":  502,
			"Message": "Debe procporcionar como minimo una dependencia",
			"Funcion": "FiltrosDependencia",
		}
		return nil, outputError
	} else {
		for j, dependencia := range dependencias {
			if j == (len(dependencias) - 1) {
				parametro_dependencia += dependencia
			} else {
				parametro_dependencia += dependencia + ","
			}

		}
	}
	if len(vigencias) == 0 {
		for v, vigencia_oficina := range vigencias_oficina {
			if v == (len(vigencias_oficina) - 1) {
				parametro_vigencia += vigencia_oficina
			} else {
				parametro_vigencia += vigencia_oficina + ","
			}

		}
	} else {
		for i, vigencia := range vigencias {
			if i == (len(vigencias) - 1) {
				parametro_vigencia += vigencia
			} else {
				parametro_vigencia += vigencia + ","
			}

		}
	}

	parametro := BodyParams{
		Dependencias: parametro_dependencia,
		Vigencias:    parametro_vigencia,
	}

	var respuesta_peticion map[string]interface{}

	err := sendJson3(beego.AppConfig.String("UrlPruebasAdministrativaJBPM")+"/contratos_dependencias", "POST", &respuesta_peticion, parametro)

	if err == nil {

		contratosMap := respuesta_peticion["contratos"].(map[string]interface{})["contrato"].([]interface{})
		for _, contrato := range contratosMap {
			contratoMap := contrato.(map[string]interface{})
			vigencia, _ := strconv.Atoi(contratoMap["vigencia"].(string))
			numeroContratoSuscrito, _ := strconv.Atoi(contratoMap["numero_contrato_suscrito"].(string))
			contratoModel := models.ContratoSuscritoDependencia{
				Vigencia:               vigencia,
				NumeroContratoSuscrito: numeroContratoSuscrito,
			}
			contratos = append(contratos, contratoModel)
		}
		fmt.Println(contratos)
		return contratos, nil
	} else {
		outputError = map[string]interface{}{"funcion": "/FiltrosDependencia", "err": err.Error(), "status": "404"}
		fmt.Println(err)
		return nil, outputError
	}

	return
}

// Funcion para filtrar pagos por lista de codigos dependencias, listas de vigencias, lista de numeros documentos contratistas,
//lista de numeros de contratos, lista de meses, lista de años o listas de id de estados

/*
func GetFiltros(codigos_dependencias []string, vigencias []string, documentos_contratistas []string, numeros_contratos []string, meses []string, anios []string, estados []string) (pagos []models.SolicitudPago, outputError interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetFiltros", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var contratistas []models.InformacionProveedor

	pagos_mensuales, outputError := GetPagosFiltrados(numeros_contratos, documentos_contratistas, anios, meses, estados)
	if outputError != nil {
		return nil, outputError
	}

	pagos_dependencias, outputError := FiltrosDependencia(codigos_dependencias, vigencias)
	if outputError != nil {
		return nil, outputError
	}

	for _, pago_dependencia := range pagos_dependencias {
		for _, pago_mensual := range pagos_mensuales {
			var pago models.SolicitudPago
			if contratoExists(strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64), pago_mensual.NumeroContrato, pago_dependencia.Contratos.Contrato) {
				if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pago_mensual.DocumentoPersonaId, &contratistas); (err == nil) && (response == 200) {
					var contrato models.InformacionContrato
					contrato, outputError = GetContrato(pago_mensual.NumeroContrato, strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64))
					if outputError == nil {
						for _, contratista := range contratistas {
							pago.NombreDependencia = " "
							pago.Rubro = contrato.Contrato.Rubro
							pago.DocumentoContratista = contratista.NumDocumento
							pago.NombreContratista = contratista.NomProveedor
							pago.Vigencia = strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64)
							pago.Ano = strconv.FormatFloat(pago_mensual.Ano, 'f', 0, 64)
							pago.Mes = strconv.FormatFloat(pago_mensual.Mes, 'f', 0, 64)
							pago.Estado = pago_mensual.EstadoPagoMensualId.Nombre
							pagos = append(pagos, pago)
						}
					} else {
						return nil, outputError
					}
				} else {
					logs.Error(err)
					outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas", "err": err.Error(), "status": "404"}
					return nil, outputError
				}
			} else {
				outputError = map[string]interface{}{"funcion": "/CertificacionCumplidosContratistas", "Message": "El contrato buscado no existe", "status": "201"}
			}
		}
	}

	return
}
*/
