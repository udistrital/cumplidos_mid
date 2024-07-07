package helpers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
	"github.com/udistrital/utils_oas/request"
)

func GetEstado(idEstado string) (estado *models.EstadoPagoMensual, outputError interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError := map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al consultar Estado   pago :" + idEstado,
				"Error":   err}
			panic(outputError)
		}
	}()

	var respuesta_peticion map[string]interface{}
	///println(beego.AppConfig.String("UrlCrudCumplidos") + "/estado_pago_mensual/" + idEstado)
	if response, err := request.GetJsonTest2(beego.AppConfig.String("UrlCrudCumplidos")+"/estado_pago_mensual/"+idEstado, &respuesta_peticion); (err == nil) && (response == 200) {
		//Ejecuta si no hay error y estado = 200
		if respuesta_peticion != nil {
			LimpiezaRespuestaRefactor(respuesta_peticion, &estado)

		}
	} else {
		//Ejecuta si hay un error o status !=200
		return nil, outputError
	}

	return estado, nil
}
