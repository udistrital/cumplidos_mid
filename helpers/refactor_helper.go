package helpers

import (
	"encoding/json"
	"fmt"
	_ "fmt"
)

func LimpiezaRespuestaRefactor(respuesta map[string]interface{}, v interface{}) {
	//fmt.Println("salida: ", respuesta)
	//b, ok := prueba["Data"].(*[]byte)
	//v = nil
	b, err := json.Marshal(respuesta["Data"])
	if err != nil {
		fmt.Println("fallo")
		panic(err)
	}
	json.Unmarshal(b, v)

	//fmt.Println("salida2: ", v)
}
