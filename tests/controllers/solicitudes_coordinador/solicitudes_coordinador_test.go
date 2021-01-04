package certificacionesHelper

import (
	"net/http"
	"testing"
)

func TestGetSolicitudesCoordinador(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_coordinador/19346572"); err == nil {
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

func TesGetSolicitudesCoordinadorError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_coordinador/1"); err == nil {
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
