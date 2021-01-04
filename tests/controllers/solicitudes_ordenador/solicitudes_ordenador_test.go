package certificacionesHelper

import (
	"net/http"
	"testing"
)

func TestGetSolicitudesOrdenador(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_ordenador/solicitudes/19128837/?limit=5&offset=0"); err == nil {
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

func TestGetSolicitudesOrdenadorError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_ordenador/solicitudes/1912883a/?limit=5&offset=0"); err == nil {
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

func TestObtenerDependenciaOrdenador(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_ordenador/dependencia_ordenador/19400342"); err == nil {
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

func TestObtenerDependenciaOrdenadorError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_ordenador/dependencia_ordenador/1912883a"); err == nil {
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

func TestInformacion_ordenador(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_ordenador/informacion_ordenador/DVE1569/2020"); err == nil {
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

func TestInformacion_ordenadorError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_ordenador/informacion_ordenador/DVE1569/202a"); err == nil {
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
