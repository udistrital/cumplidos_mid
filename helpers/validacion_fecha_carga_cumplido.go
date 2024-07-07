package helpers

import (
	_ "fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
	"github.com/udistrital/utils_oas/request"
)

func ValidarPeriodoCargaCumplido(dependencia_supervisor string, anio string, mes string) (Validacion models.ValidacionFechaCargaCumplido, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": "/ValidarPeriodoCargaCumplido", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var respuesta_peticion map[string]interface{}
	var fechasCargaCumplido []models.FechasCargaCumplidos
	query := "?query=Dependencia:" + dependencia_supervisor + ",Mes.in:" + mes + "|0,Anio.in:" + anio + "|0"
	order := "&order=desc"
	sortby := "&sortby=FechaModificacion"
	limit := "&limit=0"
	//fmt.Println(beego.AppConfig.String("UrlCrudCumplidos") + "/fechas_carga_cumplidos/" + query + sortby + order + limit)
	if response, err := request.GetJsonTest2(beego.AppConfig.String("UrlCrudCumplidos")+"/fechas_carga_cumplidos/"+query+sortby+order+limit, &respuesta_peticion); (err == nil) && (response == 200) {
		//Se encontro alguna parametrizacion

		//fmt.Println("fechas carga cumplido:", respuesta_peticion)
		//fecha actual
		fecha_actual := time.Now().Add(-time.Hour * 5)
		Validacion.FechaActual = fecha_actual
		if len(respuesta_peticion["Data"].([]interface{})[0].(map[string]interface{})) != 0 {

			LimpiezaRespuestaRefactor(respuesta_peticion, &fechasCargaCumplido)
			//fmt.Println(fechasCargaCumplido)
			fechas_periodo_parametrizado := fechasCargaCumplido[0]

			//fmt.Println("año inicio", fechas_periodo_parametrizado.FechaInicio.Year())
			//fmt.Println("año fin", fechas_periodo_parametrizado.FechaFin.Year())
			if fechas_periodo_parametrizado.FechaInicio.Year() == 1 && fechas_periodo_parametrizado.FechaFin.Year() == 1 { //Sin periodo parametrizado
				Validacion.CargaHabilitada = true
			} else { //Periodo parametrizado
				fechas_periodo_parametrizado.FechaInicio = fechas_periodo_parametrizado.FechaInicio.Add(-time.Hour*time.Duration(fechas_periodo_parametrizado.FechaInicio.Hour()) - time.Minute*time.Duration(fechas_periodo_parametrizado.FechaInicio.Minute()) - time.Second*time.Duration(fechas_periodo_parametrizado.FechaInicio.Second()))
				fechas_periodo_parametrizado.FechaFin = fechas_periodo_parametrizado.FechaFin.Add(-time.Hour*time.Duration(fechas_periodo_parametrizado.FechaFin.Hour()) - time.Minute*time.Duration(fechas_periodo_parametrizado.FechaFin.Minute()) - time.Second*time.Duration(fechas_periodo_parametrizado.FechaFin.Second()))
				//fmt.Println("fechara fin a cero", fechas_periodo_parametrizado.FechaFin)
				fechas_periodo_parametrizado.FechaFin = fechas_periodo_parametrizado.FechaFin.Add(time.Hour*23 + time.Minute*59 + time.Second*59)
				Validacion.Periodo.Inicio = fechas_periodo_parametrizado.FechaInicio
				Validacion.Periodo.Fin = fechas_periodo_parametrizado.FechaFin
				//fmt.Println("fecha actual", fecha_actual)
				//fmt.Println("periodo", fechas_periodo_parametrizado)
				if fecha_actual.Before(fechas_periodo_parametrizado.FechaFin) && fecha_actual.After(fechas_periodo_parametrizado.FechaInicio) { //dentro del periodo patrametrizado
					Validacion.CargaHabilitada = true
				} else { //fuera del periodo parametrizado
					Validacion.CargaHabilitada = false
				}
			}

		} else { //no hay ninguna parametrizacion para la carga
			Validacion.CargaHabilitada = true
		}
	}

	return
}
