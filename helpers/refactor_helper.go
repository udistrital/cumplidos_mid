package helpers

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"
)

func LimpiezaRespuestaRefactor(respuesta map[string]interface{}, v interface{}) {
	if respuesta == nil {
		return
	}

	b, err := json.Marshal(respuesta["Data"])
	if err != nil {
		logs.Error("[LIMPIEZA] Error al marshalar respuesta[\"Data\"]: %v", err)
		panic(err)
	}

	err = json.Unmarshal(b, v)
	if err != nil {
		logs.Error("[LIMPIEZA] Error al unmarshal: %v", err)
	}
}
