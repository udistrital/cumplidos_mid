package certificaciones

import (
	"net/http"
	"testing"
)

func TestCertificadoVistoBueno(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/certificacion/certificacion_visto_bueno/38/10/2018"); err == nil {
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

func TestCertificadoVistoBuenoError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/certificacion/certificacion_visto_bueno/38/10/18"); err == nil {
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

func TestCertificacionDocumentosAprobados(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/certificacion/documentos_aprobados/35/10/2019"); err == nil {
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

func TestCertificacionDocumentosAprobadosError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/certificacion/documentos_aprobados/35/10/19"); err == nil {
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
