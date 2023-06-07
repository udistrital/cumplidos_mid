package certificacionesHelper

import (
	"net/http"
	"testing"
)

func TestGetContratosContratista(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/contratos_contratista/1032490151"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndPoint(TestGetContratosContratista): Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetContratosContratista Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetContratosContratista:", err.Error())
		t.Fail()
	}

}

func TestGetContratosContratistaError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/contratos_contratista/1"); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestEndPoint(TestGetContratosContratistaError): Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetContratosContratistaError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetContratosContratistaError:", err.Error())
		t.Fail()
	}

}

func TestGetDocumentosPagoMensual(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/contratos_contratista/documentos_pago_mensual/8207"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndPoint(TestGetDocumentosPagoMensual): Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetDocumentosPagoMensual Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetDocumentosPagoMensual:", err.Error())
		t.Fail()
	}

}

func TestGetDocumentosPagoMensualError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/contratos_contratista/documentos_pago_mensual/8207"); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestEndPoint(TestGetDocumentosPagoMensualError): Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetDocumentosPagoMensualError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetDocumentosPagoMensualError:", err.Error())
		t.Fail()
	}

}
