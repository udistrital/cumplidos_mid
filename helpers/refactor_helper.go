package helpers
import (
	_"fmt"
	"encoding/json"
)



func LimpiezaRespuestaRefactor(respuesta map[string]interface{} , v interface{}){
	//fmt.Println("salida: ", respuesta)
	//b, ok := prueba["Data"].(*[]byte)
	b, err := json.Marshal(respuesta["Data"])
	if err!= nil{
		panic(err)
	}
	json.Unmarshal(b, &v)
	//fmt.Println("salida2: ", v)
}