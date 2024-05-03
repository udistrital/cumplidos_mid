package helpers

import (
	"encoding/json"
)

func LimpiezaRespuestaRefactor(respuesta map[string]interface{}, v interface{}) {
	b, err := json.Marshal(respuesta["Data"])
	if err != nil {
		panic(err)
	}

	json.Unmarshal(b, v)
}

func LimpiezaEstadoPago(respuesta map[string]interface{}, cambios_estado *interface{}) {

	data, ok := respuesta["Data"].([]interface{})

	///Si respuesta esa vacion conttinua
	if !ok {
		return
	}
	///recorre element y elimina PagoMensualId
	for _, element := range data {
		if data, ok := element.(map[string]interface{}); ok {
			if _, pagoId := data["PagoMensualId"]; pagoId {
				delete(data, "PagoMensualId")
			}
		}
	}
	*cambios_estado = respuesta
}
