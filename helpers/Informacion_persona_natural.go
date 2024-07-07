package helpers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/models"
	"github.com/udistrital/utils_oas/request"

)

func GetNombreResponable(id string) (nombreCompleto string, outputError interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError := map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al consultar info persona:" + id,
				"Error":   err}
			panic(outputError)
		}
	}()

	var respuesta_peticion []*models.InformacionPersonaNatural
	//println(beego.AppConfig.String("UrlcrudAgora") + "/informacion_persona_natural/" + id)
	query := "Id:" + id

	if response, err := request.GetJsonTest2(beego.AppConfig.String("UrlcrudAgora")+"/informacion_persona_natural?fields=PrimerNombre,SegundoNombre,PrimerApellido,SegundoApellido&limit=0&query="+query, &respuesta_peticion); (err == nil) && (response == 200) {

		if respuesta_peticion != nil {
			nombreCompleto = capitalizarPrimeraLetra(respuesta_peticion[0].PrimerNombre) +
				" " + capitalizarPrimeraLetra(respuesta_peticion[0].SegundoNombre) +
				" " + capitalizarPrimeraLetra(respuesta_peticion[0].PrimerApellido) +
				" " + capitalizarPrimeraLetra(respuesta_peticion[0].SegundoApellido)
				println(capitalizarPrimeraLetra(respuesta_peticion[0].SegundoApellido))
		}
	} else {

		return "", outputError
	}

	return nombreCompleto, nil
}
