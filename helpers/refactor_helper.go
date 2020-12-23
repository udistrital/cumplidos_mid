package helpers

import (
	"encoding/json"
	_ "fmt"
)

func LimpiezaRespuestaRefactor(respuesta map[string]interface{}, v interface{}) {
	b, err := json.Marshal(respuesta["Data"])
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, v)
}

/*func GestionError(c map[string]interface{}) {
	c.(beego.Controller)
	if err := recover(); err != nil {
		logs.Error(err)
		fmt.Println("aca es")
		respuesta := err.(map[string]interface{})
		c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "CertificacionController" + "/" + (respuesta["funcion"]).(string))
		c.Data["data"] = (respuesta["err"])
		if status, ok := respuesta["status"]; ok {
			c.Abort(status.(string))
		} else {
			c.Abort("404")
		}
	}
}*/
