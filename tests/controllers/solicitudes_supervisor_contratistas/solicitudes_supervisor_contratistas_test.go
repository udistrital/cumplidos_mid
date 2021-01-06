package solicitudes_supervisor_contratistas

import (
	"net/http"
	"testing"
)

func TestGetSolicitudesSupervisorContratistas(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_supervisor_contratistas/52204982"); err == nil {
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

func TestGetSolicitudesSupervisorContratistasError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_supervisor_contratistas/52204982s"); err == nil {
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
