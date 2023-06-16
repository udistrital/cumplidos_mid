package InformacionInformetest

import (
	"net/http"
	"testing"
)

func TestGetInformacionInforme(t *testing.T) {
	if response, err := http.Get("http://localhost:8090/v1/informacion_informe/94158"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndPoint(TestGetInformacionInforme): Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetInformacionInforme Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetInformacionInforme:", err.Error())
		t.Fail()
	}
}

func TestGetPreliquidacion(t *testing.T) {
	if response, err := http.Get("http://localhost:8090/v1/informacion_informe/preliquidacion/94181"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndPoint(TestGetPreliquidacion): Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetPreliquidacion Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetPreliquidacion:", err.Error())
		t.Fail()
	}
}
