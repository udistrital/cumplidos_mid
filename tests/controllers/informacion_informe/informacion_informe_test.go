package InformacionInformetest

import (
	"net/http"
	"testing"
)

func TestGetInformacionInforme(t *testing.T) {
	if response, err := http.Get("http://localhost:8090/v1/informacion_informe/1014294957/1265/2021/1668/2021"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndPoint: Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestEndPoint Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint:", err.Error())
		t.Fail()
	}
}
