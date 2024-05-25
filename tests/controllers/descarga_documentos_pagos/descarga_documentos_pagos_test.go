package descarga_documentos_pagos

import (
	"net/http"
	"testing"
)

func TestGetDocumentosPagoMensual(t *testing.T) {

	if response, error := http.Get("http://localhost:8090/v1/download_documents/8207"); error == nil {
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

func TestGetDocumentosPagoMensualEmpty(t *testing.T) {

	if response, error := http.Get("http://localhost:8090/v1/download_documents/827"); error == nil {
		if response.StatusCode != 204 {
			t.Error("Error TestEndPoint(TestGetDocumentosPagoMensual): Se esperaba 204 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetDocumentosPagoMensualError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetDocumentosPagoMensualError:", error.Error())
		t.Fail()
	}

}

func TestGetDocumentosPagoMensualError(t *testing.T) {

	if response, error := http.Get("http://localhost:8090/v1/download_documents/"); error == nil {
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
