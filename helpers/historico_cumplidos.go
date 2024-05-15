package helpers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

func GetEstadosPago(documento string) (depenendicas interface{}, outputError interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError := map[string]interface{}{
				"Succes":  true,
				"Status":  502,
				"Message": "Error al consultar cambios de estados del pago :" + documento,
				"Error":   err}
			panic(outputError)
		}
	}()

	var respuesta_peticion models.DependenciasXmln
	if response, err := getXMLTest(beego.AppConfig.String("UrlAdministrativaJBPM")+"/dependencias_sic/"+documento, &respuesta_peticion); err == nil && response == 200 {

		for _, dep := range respuesta_peticion.Dependencias {
			// Aquí puedes hacer lo que necesites con cada dependencia

			fmt.Println(dep.EsfCodigoDep)
		}

	} else {
		//Ejecutar si hay un error o status !=200
		return nil, outputError
	}
	return depenendicas, nil
}
