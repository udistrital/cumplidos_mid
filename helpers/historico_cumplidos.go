package helpers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
	"strconv"
)

func GetEstadosPago(idPagoMensual string) (cambiosEstado []models.CambioEstadoPago, outputError interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError := map[string]interface{}{
				"Succes":  true,
				"Status":  502,
				"Message": "Error al consultar cambios de estados del pago :" + idPagoMensual,
				"Error":   err}
			panic(outputError)
		}
	}()
	//Query de solicitud
	query := "PagoMensualId.Id:" + idPagoMensual
	var respuesta_peticion map[string]interface{}
	//println(beego.AppConfig.String("UrlCrudCumplidos") + "/cambio_estado_pago/?query=" + query)

	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/cambio_estado_pago/?query="+query+"&sortby=FechaCreacion&order=asc&limit=-1", &respuesta_peticion); (err == nil) && (response == 200) {
		//Ejecuta si no hay error y estado = 200
		if len(respuesta_peticion["Data"].([]interface{})[0].(map[string]interface{})) != 0 {
			LimpiezaRespuestaRefactor(respuesta_peticion, &cambiosEstado)

			for i, cambioEstado := range cambiosEstado {
				persona, _ := GetInformacionPersona(cambioEstado.DocumentoResponsableId)
				estado, err := GetEstado(strconv.Itoa(cambioEstado.EstadoPagoMensualId))
				if err != nil {
					panic(outputError)
				} else {
					cambiosEstado[i].NombreEstado = estado.Nombre
					cambiosEstado[i].DescripcionEstado = estado.Descripcion
					cambiosEstado[i].PagoMensualId = idPagoMensual
					cambiosEstado[i].NombreResponsable = persona.PrimerNombre

				}

			}

		}
	} else {
		//Ejecutar si hay un error o status !=200
		return nil, outputError
	}
	return cambiosEstado, nil
}
