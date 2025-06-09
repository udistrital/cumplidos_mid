package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

func ContratosContratista(numero_documento string) (contratos_disponibilidad_rp []models.ContratoDisponibilidadRp, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/ContratosContratista", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	if contratos_persona, outputError := GetContratosPersona(numero_documento); outputError == nil {
		for _, contrato_persona := range contratos_persona.ContratosPersonas.ContratoPersona {
			contrato_persona.FechaInicio = time.Date(contrato_persona.FechaInicio.Year(), contrato_persona.FechaInicio.Month(), contrato_persona.FechaInicio.Day(), 0, 0, 0, 0, contrato_persona.FechaInicio.Location())
			contrato_persona.FechaFin = time.Date(contrato_persona.FechaFin.Year(), contrato_persona.FechaFin.Month(), contrato_persona.FechaFin.Day(), 0, 0, 0, 0, contrato_persona.FechaFin.Location())

			if time.Now().AddDate(-1, 0, 0).Before(contrato_persona.FechaFin) {
				var contrato models.InformacionContrato
				contrato, outputError = GetContrato(contrato_persona.NumeroContrato, contrato_persona.Vigencia)

				if (contrato == models.InformacionContrato{} || outputError != nil) {
					continue
				}

				var informacion_contrato_contratista models.InformacionContratoContratista
				informacion_contrato_contratista, outputError = GetInformacionContratoContratista(contrato_persona.NumeroContrato, contrato_persona.Vigencia)

				// Obtener la unidad ejecutora desde el contrato
				unidadEjecucion := contrato.Contrato.UnidadEjecutora

				if cdprp, outputError := GetRP(contrato_persona.NumeroCDP, contrato_persona.Vigencia, unidadEjecucion); outputError == nil {
					for _, rp := range cdprp.CdpXRp.CdpRp {
						var contrato_disponibilidad_rp models.ContratoDisponibilidadRp
						contrato_disponibilidad_rp.NumeroContratoSuscrito = contrato_persona.NumeroContrato
						contrato_disponibilidad_rp.Vigencia = contrato_persona.Vigencia
						contrato_disponibilidad_rp.NumeroCdp = contrato_persona.NumeroCDP
						contrato_disponibilidad_rp.VigenciaCdp = contrato_persona.Vigencia
						contrato_disponibilidad_rp.NumeroRp = rp.RpNumeroRegistro
						contrato_disponibilidad_rp.VigenciaRp = rp.RpVigencia
						contrato_disponibilidad_rp.NombreDependencia = informacion_contrato_contratista.InformacionContratista.Dependencia
						contrato_disponibilidad_rp.NumDocumentoSupervisor = contrato.Contrato.Supervisor.DocumentoIdentificacion
						contrato_disponibilidad_rp.FechaInicio = contrato_persona.FechaInicio
						contrato_disponibilidad_rp.FechaFin = contrato_persona.FechaFin
						contratos_disponibilidad_rp = append(contratos_disponibilidad_rp, contrato_disponibilidad_rp)
					}
				} else {
					logs.Error(outputError)
					continue
				}
			}
		}
	} else {
		logs.Error(outputError)
		outputError = map[string]interface{}{"funcion": "/contratosContratista/GetContratosPersona", "err": outputError, "status": "502"}
		return nil, outputError
	}
	return contratos_disponibilidad_rp, nil
}

func GetRP(numero_cdp string, vigencia_cdp string, unidad_ejecucion string) (rp models.InformacionCdpRp, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetRP0", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var temp map[string]interface{}
	var temp_cdp_rp models.InformacionCdpRp

	url := fmt.Sprintf("%s/cdprp/%s/%s/%s", beego.AppConfig.String("UrlFinancieraJBPM"), numero_cdp, vigencia_cdp, unidad_ejecucion)
	fmt.Println(url)

	if response, err := getJsonWSO2Test(url, &temp); (err == nil) && (response == 200) {
		json_cdp_rp, error_json := json.Marshal(temp)

		if error_json == nil {
			if err := json.Unmarshal(json_cdp_rp, &temp_cdp_rp); err == nil {
				rp = temp_cdp_rp
				return rp, nil
			} else {
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/GetRP1", "err": err, "status": "502"}
				return rp, outputError
			}
		} else {
			logs.Error(error_json)
			outputError = map[string]interface{}{"funcion": "/GetRP2", "err": error_json, "status": "502"}
			return rp, outputError
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetRP3", "err": err, "status": "502"}
		return rp, outputError
	}
}

func GetContratosPersona(num_documento string) (contratos_persona models.InformacionContratosPersona, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetContratosPersona", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var temp map[string]interface{}
	var contratos models.InformacionContratosPersona
	fmt.Println(beego.AppConfig.String("UrlAdministrativaJBPM") + "/contratos_contratista/" + num_documento)
	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlAdministrativaJBPM")+"/contratos_contratista/"+num_documento, &temp); (err == nil) && (response == 200) {
		json_contratos, error_json := json.Marshal(temp)
		if error_json == nil {
			err := json.Unmarshal(json_contratos, &contratos)
			if err == nil {
				contratos_persona = contratos
				//fmt.Println("Contratos personas", contratos_persona)
				return contratos_persona, nil
			} else {
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/GetContratosPersona", "err": err, "status": "502"}
				return contratos_persona, outputError
			}

		} else {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/GetContratosPersona", "err": error_json.Error(), "status": "502"}
			return contratos_persona, outputError
		}

	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetContratosPersona", "err": err, "status": "502"}
		return contratos_persona, outputError
	}

	return contratos_persona, nil

}

func GetContrato(num_contrato_suscrito string, vigencia string) (informacion_contrato models.InformacionContrato, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetContrato", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var temp map[string]interface{}
	//fmt.Println(beego.AppConfig.String("UrlAdministrativaJBPM") + "/" + "contrato/" + num_contrato_suscrito + "/" + vigencia)
	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlAdministrativaJBPM")+"/"+"contrato/"+num_contrato_suscrito+"/"+vigencia, &temp); (err == nil) && (response == 200) {
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

func GetInformacionContratoContratista(num_contrato_suscrito string, vigencia string) (informacion_contrato_contratista models.InformacionContratoContratista, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetInformacionContratoContratista", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var temp map[string]interface{}

	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlAdministrativaJBPM")+"/"+"informacion_contrato_contratista/"+num_contrato_suscrito+"/"+vigencia, &temp); (err == nil) && (response == 200) {
		json_contrato, error_json := json.Marshal(temp)
		if error_json == nil {
			var contrato_contratista models.InformacionContratoContratista
			if err := json.Unmarshal(json_contrato, &contrato_contratista); err == nil {
				informacion_contrato_contratista = contrato_contratista
				return informacion_contrato_contratista, nil
			} else {
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/GetInformacionContratoContratista", "err": err, "status": "502"}
				return informacion_contrato_contratista, outputError
			}
		} else {
			logs.Error(error_json.Error())
			outputError = map[string]interface{}{"funcion": "/GetInformacionContratoContratista", "err": error_json.Error(), "status": "502"}
			return informacion_contrato_contratista, outputError
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/getInformacionContratosContratista", "err": err, "status": "502"}
		return informacion_contrato_contratista, outputError
	}
}

func GetActaDeInicio(numero_contrato string, vigencia_contrato int) (acta_inicio models.ActaInicio, outputError map[string]interface{}) {

	var actasInicio []models.ActaInicio
	fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/acta_inicio/?query=NumeroContrato:" + numero_contrato + ",Vigencia:" + strconv.Itoa(vigencia_contrato))
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/acta_inicio/?query=NumeroContrato:"+numero_contrato+",Vigencia:"+strconv.Itoa(vigencia_contrato), &actasInicio); (err == nil) && (response == 200) {
		if len(actasInicio) == 0 {
			outputError = map[string]interface{}{"funcion": "/getInformacionContratosContratista", "err": errors.New("No se encontro acta de inicio"), "status": "502"}
			return acta_inicio, outputError
		}
		acta_inicio = actasInicio[0]
		return acta_inicio, nil

	} else {
		outputError = map[string]interface{}{"funcion": "/getInformacionContratosContratista", "err": err, "status": "502"}
		return acta_inicio, outputError
	}
}

func FechasContratoConNovedades(numero_contrato string, vigencia_contrato string, numero_cdp string, num_doc string) (fechas models.FechasConNovedades, outputError map[string]interface{}) {

	if contratos_persona, err := GetContratosPersona(num_doc); err == nil {
		for _, contrato := range contratos_persona.ContratosPersonas.ContratoPersona {
			if contrato.NumeroContrato == numero_contrato && contrato.Vigencia == vigencia_contrato && contrato.NumeroCDP == numero_cdp {
				fechas.FechaInicio = contrato.FechaInicio
				fechas.FechaFin = contrato.FechaFin
				return fechas, nil
			}
		}
		outputError = map[string]interface{}{"funcion": "/FechasContratoConNovedades", "err": "No se encontro el contrato", "status": "502"}
	} else {
		outputError = map[string]interface{}{"funcion": "/FechasContratoConNovedades/GetContratoPersona", "err": err, "status": "502"}
	}

	return fechas, outputError
}
