package solicitudes_pago_mensual

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/astaxie/beego"
)

func TestGetSolicitudesPagoMensual(t *testing.T) {

	var body map[string]interface{}

	body = map[string]interface{}{
		"dependencias":          "'DEP12'",
		"vigencias":             "2017",
		"documentos_persona_id": "",
		"numeros_contratos":     "",
		"meses":                 "",
		"anios":                 "",
		"estados_pagos":         "",
	}

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(body); err != nil {
		beego.Error(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8090/v1/solicitudes_pagos", b)
	r, err := client.Do(req)
	if err != nil {
		t.Error("Error TestEndPoint(TestPostInforme): Se esperaba 200 y se obtuvo", r.StatusCode)
		t.Fail()
	} else {
		t.Log("TestPostInforme Finalizado Correctamente (OK)")
	}

}
