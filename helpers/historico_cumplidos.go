package helpers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/models"
)

func GetEstadosPago(idPagoMensual string) (cambios_estado []models.CambioEstadoPago, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError := map[string]interface{}{"funcion": "/getEstadosPago", "err": err, "status": "502"}
			panic(outputError)
		}
	}()
	//Query de solicitud
	query := "PagoMensualId.Id:" + idPagoMensual
	var respuesta_peticion map[string]interface{}
	println(beego.AppConfig.String("UrlCrudCumplidos") + "/cambio_estado_pago/?query=" + query)

	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/cambio_estado_pago/?query="+query, &respuesta_peticion); (err == nil) && (response == 200) {
		//Ejecuta si no hay error y estado = 200
		if len(respuesta_peticion["Data"].([]interface{})[0].(map[string]interface{})) != 0 {
			LimpiezaRespuestaRefactor(respuesta_peticion, &cambios_estado)
		}
	} else {
		//Ejecutar si hay un error o status !=200
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/getEstadosPago", "err": err, "status": "502"}
		return nil, outputError
	}
	return cambios_estado, nil
}
