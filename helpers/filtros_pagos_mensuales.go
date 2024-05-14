package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
	sortby := "&sortby=Ano,Mes,DocumentoPersonaId"
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

	//Vigencias de la oficina, esta lista se crea para retornar por defecto todas las
	//vigencias de la oficina en caso de que no se le asigne un valor al filtro

	vigencias_oficina := []string{}

	for vigencia := 2017; vigencia <= time.Now().Year(); vigencia++ {
		vigencias_oficina = append(vigencias_oficina, strconv.Itoa(vigencia))
	}

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

	err := sendJson3(beego.AppConfig.String("UrlAdministrativaJBPMContratosDependencia")+"/contratos_dependencias", "POST", &respuesta_peticion, parametro)

	if err == nil {

		contratosMap := respuesta_peticion["contratos"].(map[string]interface{})["contrato"].([]interface{})
		for _, contrato := range contratosMap {
			contratoMap := contrato.(map[string]interface{})
			vigencia, _ := contratoMap["vigencia"].(string)
			numeroContratoSuscrito, _ := contratoMap["numero_contrato_suscrito"].(string)
			contratoModel := models.ContratoSuscritoDependencia{
				Vigencia:               vigencia,
				NumeroContratoSuscrito: numeroContratoSuscrito,
			}
			contratos = append(contratos, contratoModel)

		}
	} else {
		outputError = map[string]interface{}{"funcion": "/FiltrosDependencia", "err": err, "status": "404"}
		return nil, outputError
	}
	return contratos, nil
}

// Funcion para filtrar pagos por lista de codigos dependencias, listas de vigencias, lista de numeros documentos contratistas,
//lista de numeros de contratos, lista de meses, lista de años o listas de id de estados

func SolicitudesPagoMensual(codigos_dependencias []string, vigencias []string, documentos_contratistas []string, numeros_contratos []string, meses []string, anios []string, estados []string) (pagos []models.SolicitudPago, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al filtrar los pagos, ṕor dependencia",
				"Funcion": "GetPagosFiltradosDependencia",
			}
			panic(outputError)
		}
	}()

	var pagos_filtrados []models.PagoMensual
	var filtros_dependencias []models.ContratoSuscritoDependencia

	pagos_filtrados, err := GetPagosFiltrados(numeros_contratos, documentos_contratistas, anios, meses, estados)
	if err != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  502,
			"Message": "Error al filtrar los pagos, Ningun parametro coincide con los filtros",
		}
		return nil, outputError
	}
	filtros_dependencias, error := FiltrosDependencia(codigos_dependencias, vigencias)
	if error != nil {
		outputError := map[string]interface{}{
			"Succes":  false,
			"Status":  502,
			"Message": "Error al filtrar por dependencias, Ningun parametro coincide con los filtros",
			"Error":   error,
		}
		return nil, outputError
	}

	for _, filtro_dependencia := range filtros_dependencias {
		for _, pago_filtrado := range pagos_filtrados {
			var filtro_pago_dependencia models.SolicitudPago
			if filtro_dependencia.NumeroContratoSuscrito == pago_filtrado.NumeroContrato && filtro_dependencia.Vigencia == strconv.Itoa(int(pago_filtrado.VigenciaContrato)) {
				var contrato models.InformacionContrato
				contrato, outputError = GetInformacionContrato(pago_filtrado.NumeroContrato, strconv.FormatFloat(pago_filtrado.VigenciaContrato, 'f', 0, 64))
				var informacion_contrato_contratista models.InformacionContratoContratista
				informacion_contrato_contratista, outputError = GetInformacionContratoContratista(pago_filtrado.NumeroContrato, strconv.FormatFloat(pago_filtrado.VigenciaContrato, 'f', 0, 64))
				if outputError == nil {
					filtro_pago_dependencia.NombreDependencia = informacion_contrato_contratista.InformacionContratista.Dependencia
					filtro_pago_dependencia.Rubro = contrato.Contrato.Rubro
					filtro_pago_dependencia.Vigencia = strconv.FormatFloat(pago_filtrado.VigenciaContrato, 'f', 0, 64)
					filtro_pago_dependencia.Ano = strconv.FormatFloat(pago_filtrado.Ano, 'f', 0, 64)
					filtro_pago_dependencia.Mes = strconv.FormatFloat(pago_filtrado.Mes, 'f', 0, 64)
					filtro_pago_dependencia.Estado = pago_filtrado.EstadoPagoMensualId.Nombre
					filtro_pago_dependencia.DocumentoContratista = informacion_contrato_contratista.InformacionContratista.Documento.Numero
					filtro_pago_dependencia.NombreContratista = informacion_contrato_contratista.InformacionContratista.NombreCompleto
					pagos = append(pagos, filtro_pago_dependencia)
				} else {
					return nil, outputError
				}
			}
		}
	}

	return pagos, nil
}

func GetInformacionContrato(num_contrato_suscrito string, vigencia string) (informacion_contrato models.InformacionContrato, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetContrato", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var temp map[string]interface{}
	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlAdministrativaJBPM")+"/"+"informacion_contrato/"+num_contrato_suscrito+"/"+vigencia, &temp); (err == nil) && (response == 200) {
		json_contrato, error_json := json.Marshal(temp)
		if error_json == nil {
			var contrato models.InformacionContrato
			if err := json.Unmarshal(json_contrato, &contrato); err == nil {
				informacion_contrato = contrato
				//Se valida si esta vacio el objeto
				if informacion_contrato == (models.InformacionContrato{}) {
					logs.Error(err)
					outputError = map[string]interface{}{"funcion": "/GetContrato/EmptyResponse", "err": err, "status": "502"}
					return informacion_contrato, outputError
				}
				return informacion_contrato, nil
			} else {
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/GetContrato", "err": err, "status": "502"}
				return informacion_contrato, outputError
			}
		} else {
			logs.Error(error_json.Error())
			outputError = map[string]interface{}{"funcion": "/GetContrato", "err": error_json.Error(), "status": "502"}
			return informacion_contrato, outputError
		}

	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetContrato", "err": err, "status": "502"}
		return informacion_contrato, outputError
	}

	return informacion_contrato, nil
}
