package historico_cumplidos

import (
	"net/http"
	"testing"
)

func TestGetCambioEstado(t *testing.T) {

	if response, error := http.Get("http://localhost:8090/v1/historicos/cambio-estado/89499"); error == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndPoint(TestGetDocumentosPagoMensual): Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetDocumentosPagoMensualError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetDocumentosPagoMensualError:", error.Error())
		t.Fail()
	}
}

func TestGetCambioEstadoError(t *testing.T) {

	if response, error := http.Get("http://localhost:8090/v1/historicos/cambio-estado/"); error == nil {
		if response.StatusCode != 404 {
			t.Error("Error TestEndPoint(TestGetDocumentosPagoMensual): Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetDocumentosPagoMensualError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetDocumentosPagoMensualError:", error.Error())
		t.Fail()
	}
}

func TestGetDependencias(t *testing.T) {

	if response, error := http.Get("http://localhost:8090/v1/historicos/dependencias/19483708"); error == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndPoint(TestGetDocumentosPagoMensual): Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetDocumentosPagoMensualError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetDocumentosPagoMensualError:", error.Error())
		t.Fail()
	}
}

func TestGetDependenciasError(t *testing.T) {

	if response, error := http.Get("http://localhost:8090/v1/historicos/dependencias/"); error == nil {
		if response.StatusCode != 404 {
			t.Error("Error TestEndPoint(TestGetDocumentosPagoMensual): Se esperaba 404 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetDocumentosPagoMensualError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetDocumentosPagoMensualError:", error.Error())
		t.Fail()
	}
}
