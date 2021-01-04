package certificacionesHelper

import (
	"net/http"
	"testing"
)

func TestGetContratosContratista(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/contratos_contratista/1032490151"); err == nil {
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

func TestGetContratosContratistaError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/contratos_contratista/1"); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestEndPoint: Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestEndPoint Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint:", err.Error())
		t.Fail()
	}

}
