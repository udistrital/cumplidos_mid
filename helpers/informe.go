package helpers

import (
	"fmt"

	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/models"
)

func Informe(contrato string, vigencia string, mes string, anio string) (informe []models.Informe, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": "/Informe", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	// var aux_informe models.Informe
	var query string
	var respuesta_peticion map[string]interface{}
	query = "contrato:" + contrato + ",vigencia:" + vigencia + ",mes:" + mes + ",anio:" + anio
	fmt.Println(beego.AppConfig.String("UrlCrudCumplidos") + "/informe/?query=" + query)
	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/informe/?query="+query, &respuesta_peticion); (err == nil) && (response == 200) {
		fmt.Println("informe:", respuesta_peticion)
		if len(respuesta_peticion["Data"].([]interface{})[0].(map[string]interface{})) != 0 {

			LimpiezaRespuestaRefactor(respuesta_peticion, &informe)
			for i, inf := range informe {
				idInforme := strconv.Itoa(inf.Id)
				actividadesEsp, err := getActividadesEspecificas(idInforme)
				fmt.Println(actividadesEsp)
				if err == nil {
					informe[i].ActividadesEspecificas = &actividadesEsp
				}
				fmt.Println(inf.ActividadesEspecificas)
			}
		}

		// aux_informe.Id= respuesta_peticion.Id
		// informe = append(informe, aux_informe)
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/Informe", "err": err, "status": "502"}
		return nil, outputError
	}

	return
}

func getActividadesEspecificas(idInforme string) (actividades_especificas []models.ActividadEspecifica, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": "/Informe", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	// var aux_informe models.Informe
	var query string
	var respuesta_peticion map[string]interface{}
	query = "informeid:" + idInforme
	fmt.Println(beego.AppConfig.String("UrlCrudCumplidos") + "/actividad_especifica/?query=" + query)
	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/actividad_especifica/?query="+query, &respuesta_peticion); (err == nil) && (response == 200) {
		fmt.Println("Actividades especificas:", respuesta_peticion)
		if len(respuesta_peticion["Data"].([]interface{})[0].(map[string]interface{})) != 0 {

			LimpiezaRespuestaRefactor(respuesta_peticion, &actividades_especificas)
			for i, actEsp := range actividades_especificas {
				idactEsp := strconv.Itoa(actEsp.Id)
				actividadesRea, err := getActividadesRealizadas(idactEsp)
				if err == nil {
					actividades_especificas[i].ActividadesRealizadas = &actividadesRea
				}
			}
		}

		// aux_informe.Id= respuesta_peticion.Id
		// informe = append(informe, aux_informe)
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/Informe", "err": err, "status": "502"}
		return nil, outputError
	}

	return
}

func getActividadesRealizadas(idActividadEspecifica string) (actividades_realizadas []models.ActividadRealizada, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": "/Informe", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	// var aux_informe models.Informe
	var query string
	var respuesta_peticion map[string]interface{}
	query = "actividadespecificaid:" + idActividadEspecifica
	fmt.Println(beego.AppConfig.String("UrlCrudCumplidos") + "/actividad_realizada/?query=" + query)
	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/actividad_realizada/?query="+query, &respuesta_peticion); (err == nil) && (response == 200) {
		fmt.Println("informe:", respuesta_peticion)
		if len(respuesta_peticion["Data"].([]interface{})[0].(map[string]interface{})) != 0 {

			LimpiezaRespuestaRefactor(respuesta_peticion, &actividades_realizadas)
		}

		// aux_informe.Id= respuesta_peticion.Id
		// informe = append(informe, aux_informe)
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/Informe", "err": err, "status": "502"}
		return nil, outputError
	}

	return
}
