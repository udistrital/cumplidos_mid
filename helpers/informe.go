package helpers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/models"
)

func Informe(pago_mensual_id string) (informe []models.Informe, outputError map[string]interface{}) {
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
	query = "PagoMensualId.Id:" + pago_mensual_id
	fmt.Println(beego.AppConfig.String("UrlCrudCumplidos") + "/informe/?query=" + query)
	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/informe/?query="+query, &respuesta_peticion); (err == nil) && (response == 200) {
		if len(respuesta_peticion["Data"].([]interface{})[0].(map[string]interface{})) != 0 {

			LimpiezaRespuestaRefactor(respuesta_peticion, &informe)
			var pagos_mensuales []models.PagoMensual
			if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/pago_mensual/?query=Id:"+pago_mensual_id, &respuesta_peticion); (err == nil) && (response == 200) {
				LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
				informe[0].PagoMensualId.NumeroContrato = pagos_mensuales[0].NumeroContrato
				informe[0].PagoMensualId.VigenciaContrato = pagos_mensuales[0].VigenciaContrato
				informe[0].PagoMensualId.Mes = pagos_mensuales[0].Mes
				informe[0].PagoMensualId.Ano = pagos_mensuales[0].Ano
				informe[0].PagoMensualId.DocumentoPersonaId = pagos_mensuales[0].DocumentoPersonaId
				informe[0].PagoMensualId.DocumentoResponsableId = pagos_mensuales[0].DocumentoResponsableId
				informe[0].PagoMensualId.EstadoPagoMensualId = pagos_mensuales[0].EstadoPagoMensualId
				informe[0].PagoMensualId.CargoResponsable = pagos_mensuales[0].CargoResponsable
				informe[0].PagoMensualId.Activo = pagos_mensuales[0].Activo
				informe[0].PagoMensualId.FechaCreacion = pagos_mensuales[0].FechaCreacion
				informe[0].PagoMensualId.FechaModificacion = pagos_mensuales[0].FechaModificacion
			}
			for i, inf := range informe {
				idInforme := strconv.Itoa(inf.Id)
				actividadesEsp, err := GetActividadesEspecificas(idInforme)
				if err == nil {
					informe[i].ActividadesEspecificas = actividadesEsp
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

func GetActividadesEspecificas(idInforme string) (actividades_especificas []models.ActividadEspecifica, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": "/GetActividadesEspecificas", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	// var aux_informe models.Informe
	var query string
	var respuesta_peticion map[string]interface{}
	query = "informeid:" + idInforme + ",Activo:true"
	fmt.Println(beego.AppConfig.String("UrlCrudCumplidos") + "/actividad_especifica/?query=" + query + "&limit=-1&sortby=FechaCreacion&order=asc")
	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/actividad_especifica/?query="+query+"&limit=-1&sortby=FechaCreacion&order=asc", &respuesta_peticion); (err == nil) && (response == 200) {
		if len(respuesta_peticion["Data"].([]interface{})[0].(map[string]interface{})) != 0 {

			LimpiezaRespuestaRefactor(respuesta_peticion, &actividades_especificas)
			for i, actEsp := range actividades_especificas {
				idactEsp := strconv.Itoa(actEsp.Id)
				actividadesRea, err := GetActividadesRealizadas(idactEsp)
				if err == nil {
					actividades_especificas[i].ActividadesRealizadas = actividadesRea
				}
			}
		}

		// aux_informe.Id= respuesta_peticion.Id
		// informe = append(informe, aux_informe)
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetActividadesEspecificas", "err": err, "status": "502"}
		return nil, outputError
	}

	return
}

func GetActividadesRealizadas(idActividadEspecifica string) (actividades_realizadas []models.ActividadRealizada, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": "/GetActividadesRealizadas", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	// var aux_informe models.Informe
	var query string
	var respuesta_peticion map[string]interface{}
	query = "actividadespecificaid:" + idActividadEspecifica + ",Activo:true"
	fmt.Println(beego.AppConfig.String("UrlCrudCumplidos") + "/actividad_realizada/?query=" + query + "&limit=-1&sortby=FechaCreacion&order=asc")
	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/actividad_realizada/?query="+query+"&limit=-1&sortby=FechaCreacion&order=asc", &respuesta_peticion); (err == nil) && (response == 200) {
		if len(respuesta_peticion["Data"].([]interface{})[0].(map[string]interface{})) != 0 {

			LimpiezaRespuestaRefactor(respuesta_peticion, &actividades_realizadas)
		}

		// aux_informe.Id= respuesta_peticion.Id
		// informe = append(informe, aux_informe)
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetActividadesRealizadas", "err": err, "status": "502"}
		return nil, outputError
	}

	return
}

func AddInforme(informe models.Informe) (response map[string]interface{}, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": "/AddInforme", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var informe_creado models.Informe
	informe_creado.PeriodoInformeFin = informe.PeriodoInformeFin
	informe_creado.PeriodoInformeInicio = informe.PeriodoInformeInicio
	informe_creado.Proceso = informe.Proceso
	informe_creado.PagoMensualId = informe.PagoMensualId
	actividades_especificas := informe.ActividadesEspecificas
	var res map[string]interface{}
	if err := sendJson(beego.AppConfig.String("UrlCrudCumplidos")+"/informe", "POST", &res, informe_creado); err == nil {
		LimpiezaRespuestaRefactor(res, &informe_creado)
		for _, actEsp := range actividades_especificas {
			var actividad_esp = map[string]interface{}{"ActividadEspecifica": actEsp.ActividadEspecifica, "Avance": actEsp.Avance, "InformeId": map[string]interface{}{"Id": informe_creado.Id}}
			if res, err := AddActividadEspecifica(actividad_esp); err == nil {
				for _, actRea := range actEsp.ActividadesRealizadas {
					var actividad_rea = map[string]interface{}{"Actividad": actRea.Actividad, "ProductoAsociado": actRea.ProductoAsociado, "Evidencia": actRea.Evidencia, "ActividadEspecificaId": map[string]interface{}{"Id": res.Id}}
					if err == nil {
						if _, err := AddActividadRealizada(actividad_rea); err == nil {
							response = map[string]interface{}{"result": "succesfully created"}
						}
					}
				}
			}
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/Informe/AddInforme", "err": err, "status": "502"}
		panic(outputError)
	}

	return
}

func AddActividadEspecifica(actividad_especifica map[string]interface{}) (actividad_especifica_creada models.ActividadEspecifica, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": "/AddActividadEspecifica", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var response map[string]interface{}
	if err := sendJson(beego.AppConfig.String("UrlCrudCumplidos")+"/actividad_especifica", "POST", &response, actividad_especifica); err == nil {
		LimpiezaRespuestaRefactor(response, &actividad_especifica_creada)
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/Informe/AddActividadEspecifica", "err": err, "status": "502"}
		panic(outputError)

	}

	return actividad_especifica_creada, outputError
}

func AddActividadRealizada(actividad_realizada map[string]interface{}) (actividad_realizada_creada models.ActividadRealizada, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": "/AddActividadRealizada", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	var response map[string]interface{}
	if err := sendJson(beego.AppConfig.String("UrlCrudCumplidos")+"/actividad_realizada", "POST", &response, actividad_realizada); err == nil {
		LimpiezaRespuestaRefactor(response, &actividad_realizada_creada)
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/Informe/AddActividadRealizada", "err": err, "status": "502"}
		panic(outputError)
	}

	return
}

func UpdateInformeById(informe models.Informe) (outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": "/UpdateInforme", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var informe_Actualizado models.Informe
	informe_Actualizado.Id = informe.Id
	informe_Actualizado.FechaCreacion = informe.FechaCreacion
	informe_Actualizado.PeriodoInformeFin = informe.PeriodoInformeFin
	informe_Actualizado.PeriodoInformeInicio = informe.PeriodoInformeInicio
	informe_Actualizado.Proceso = informe.Proceso
	informe_Actualizado.PagoMensualId = informe.PagoMensualId
	informe_Actualizado.Activo = informe.Activo
	actividades_especificas_update := informe.ActividadesEspecificas
	id := strconv.Itoa(informe.Id)
	var res map[string]interface{}
	//Actualiza el informe
	if err := sendJson(beego.AppConfig.String("UrlCrudCumplidos")+"/informe/"+id, "PUT", &res, informe_Actualizado); err == nil {
		query := "informeid:" + id
		var actividades_especificas []models.ActividadEspecifica
		var respuesta_peticion map[string]interface{}
		//consulta todas las antiguas actividades especificas
		if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/actividad_especifica/?query="+query+"&limit=-1&sortby=FechaCreacion&order=asc", &respuesta_peticion); (err == nil) && (response == 200) {
			if len(respuesta_peticion["Data"].([]interface{})[0].(map[string]interface{})) != 0 {
				LimpiezaRespuestaRefactor(respuesta_peticion, &actividades_especificas)
				for _, actEsp := range actividades_especificas {
					idactEsp := strconv.Itoa(actEsp.Id)
					if err := sendJson(beego.AppConfig.String("UrlCrudCumplidos")+"/actividad_especifica/"+idactEsp, "DELETE", &res, informe_Actualizado); err == nil {
					}
				}
			}
		}
		//Crea nuevamente las actividades especificas y realizadas
		for _, actEsp := range actividades_especificas_update {
			var actividad_esp = map[string]interface{}{"ActividadEspecifica": actEsp.ActividadEspecifica, "FechaCreacion": actEsp.FechaCreacion, "Avance": actEsp.Avance, "InformeId": map[string]interface{}{"Id": informe.Id}}
			if res, err := AddActividadEspecifica(actividad_esp); err == nil {
				for _, actRea := range actEsp.ActividadesRealizadas {
					var actividad_rea = map[string]interface{}{"Actividad": actRea.Actividad, "ProductoAsociado": actRea.ProductoAsociado, "FechaCreacion": actRea.FechaCreacion, "Evidencia": actRea.Evidencia, "ActividadEspecificaId": map[string]interface{}{"Id": res.Id}}
					if err == nil {
						if _, err := AddActividadRealizada(actividad_rea); err == nil {
						}
					}
				}
			}
		}
	}

	return
}

func UltimoInformeContratista(contrato string, vigencia string, documento string) (informe []models.Informe, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": "/UltimoInformeContratista", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	// var aux_informe models.Informe
	var respuesta_peticion map[string]interface{}
	query := "NumeroContrato:" + contrato + ",VigenciaContrato:" + vigencia + ",DocumentoPersonaId:" + documento + ",EstadoPagoMensualId.CodigoAbreviacion.in:AP|PRC|PAD|AD|AS|RD|RP|RC|PRS|RS|RO"
	order := "&order=desc,desc"
	sortby := "&sortby=Ano,Mes"
	limit := "&limit=1"
	var pagos_mensuales []models.PagoMensual
	fmt.Println(beego.AppConfig.String("UrlCrudCumplidos") + "/pago_mensual/?query=" + query + sortby + order + limit)
	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/pago_mensual/?query="+query+sortby+order+limit, &respuesta_peticion); (err == nil) && (response == 200) {
		if len(respuesta_peticion["Data"].([]interface{})[0].(map[string]interface{})) != 0 {

			LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
			pago_mensual_id := strconv.Itoa(pagos_mensuales[0].Id)
			return Informe(pago_mensual_id)
		}

		// aux_informe.Id= respuesta_peticion.Id
		// informe = append(informe, aux_informe)
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/UltimoInformeContratista", "err": err, "status": "502"}
		return nil, outputError
	}

	return
}
