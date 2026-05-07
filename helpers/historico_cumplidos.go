package helpers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

func GetEstadosPago(idPagoMensual string) (cambiosEstado []models.CambioEstadoPago, outputError interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError := map[string]interface{}{
				"funcion": "/GetEstadosPago",
				"err":     err,
				"status":  "502"}
			panic(outputError)
		}
	}()
	//Query de solicitud
	query := "PagoMensualId.Id:" + idPagoMensual
	var respuesta_peticion map[string]interface{}

	if response, err := GetJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/cambio_estado_pago/?query="+query+"&sortby=FechaCreacion&order=asc&limit=-1", &respuesta_peticion); (err == nil) && (response == 200) {
		//Ejecuta si no hay error y estado = 200
		data, ok := respuesta_peticion["Data"].([]interface{})
		if ok && len(data) > 0 {
			LimpiezaRespuestaRefactor(respuesta_peticion, &cambiosEstado)

			for i, cambioEstado := range cambiosEstado {
				nombreResponsable, _ := GetNombreResponable(cambioEstado.DocumentoResponsableId)
				estado, err := GetEstado(strconv.Itoa(cambioEstado.EstadoPagoMensualId.Id))
				if err != nil || estado == nil {
					outputError = map[string]interface{}{"funcion": "/GetEstadosPago/GetEstado", "err": err, "status": "502"}
					panic(outputError)
				} else {
					cambiosEstado[i].CargoResponsable = capitalizarPrimeraLetra(cambiosEstado[i].CargoResponsable)
					cambiosEstado[i].NombreEstado = estado.Nombre
					cambiosEstado[i].DescripcionEstado = estado.Descripcion
					if idPagoMensualInt, err := strconv.Atoi(idPagoMensual); err == nil {
						cambiosEstado[i].PagoMensualId.Id = idPagoMensualInt
					}
					cambiosEstado[i].NombreResponsable = nombreResponsable
					cambiosEstado[i].NombreEstado = capitalizarPrimeraLetra(cambiosEstado[i].NombreEstado)

				}

			}

		}
	} else {
		//Ejecutar si hay un error o status !=200
		outputError = map[string]interface{}{"funcion": "/GetEstadosPago", "err": err, "status": "502"}
		return nil, outputError
	}
	return cambiosEstado, nil
}

func ObtenerDependencias(documento string) (dependencias map[string]interface{}, errorOutput interface{}) {

	defer func() {

		if err := recover(); err != nil {
			errorOutput = map[string]interface{}{
				"Success": true,
				"Status":  502,
				"Message": "Error al consultar las dependencias: " + documento,
				"Error":   err,
			}
			panic(errorOutput)
		}
	}()
	dependencias = make(map[string]interface{})
	dependencias["Dependencias Supervisor"], errorOutput = GetDependenciasSupervisor(documento)
	dependencias["Dependencias Ordenador"], errorOutput = GetDependenciasOrdenador(documento)
	if dependencias != nil {
		return dependencias, nil
	}

	return nil, errorOutput
}
