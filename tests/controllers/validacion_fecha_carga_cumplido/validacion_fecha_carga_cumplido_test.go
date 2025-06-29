package validacionfechacargacumplidotest

import (
	"net/http"
	"testing"
)

func TestGetValidacionPeriodo(t *testing.T) {
	if response, err := http.Get("http://localhost:8090/v1/validacion_periodo_carga_cumplido/DEP12/2021/7"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndPoint(TestGetValidacionPeriodo): Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetValidacionPeriodo Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetValidacionPeriodo:", err.Error())
		t.Fail()
	}
}
