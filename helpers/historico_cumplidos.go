package helpers

//PagoMensual models.PagoMensual

func DevolverDatosFiltrados(contratos []string, anos []string, meses []string, documentos []string, estados []string) (outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/DevolverDatosFiltrados", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var respuesta_peticion map[string]interface{}

	

		
	//Se contruye dinamicamente el query

	query := "pago_mensual?query" 
	order :=
	sortby :=
	limit := 

}