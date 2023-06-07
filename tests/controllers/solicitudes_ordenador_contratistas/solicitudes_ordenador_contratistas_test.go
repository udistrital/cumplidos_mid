package certificacionesHelper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

func TestGetSolicitudesOrdenadorContratistas(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_ordenador_contratistas/solicitudes/52204982"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndPoint(TestGetSolicitudesOrdenadorContratistas): Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetSolicitudesOrdenadorContratistas Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetSolicitudesOrdenadorContratistas:", err.Error())
		t.Fail()
	}

}

func TestGetSolicitudesOrdenadorContratistasError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_ordenador_contratistas/solicitudes/1"); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestEndPoint(TestGetSolicitudesOrdenadorContratistas): Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetSolicitudesOrdenadorContratistas Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetSolicitudesOrdenadorContratistas:", err.Error())
		t.Fail()
	}

}

func TestAprobarMultiplesPagosContratistas(t *testing.T) {

	arregloPrueba := []models.PagoContratistaCdpRp{
		{
			PagoMensual: &models.PagoMensual{
				Id:                 173298,
				NumeroContrato:     "1406",
				VigenciaContrato:   2020,
				Mes:                3,
				DocumentoPersonaId: "1032490151",
				EstadoPagoMensualId: &models.EstadoPagoMensual{
					Id: 8,
				},
				DocumentoResponsableId: "52204982",
				CargoResponsable:       "SUPERVISOR OFICINA ASESORA DE SISTEMAS",
				Ano:                    2020,
				Activo:                 false,
			},
			NombrePersona:     "DIEGO ALEJANDRO GUTIERREZ ROJAS",
			NumeroCdp:         "2400",
			VigenciaCdp:       "2020",
			NumeroRp:          "23478",
			VigenciaRp:        "2020",
			NombreDependencia: "OFICINA ASESORA DE SISTEMAS",
			Rubro:             "Inversión",
		},
		{
			PagoMensual: &models.PagoMensual{
				Id:                 10435,
				NumeroContrato:     "250",
				VigenciaContrato:   2018,
				Mes:                9,
				DocumentoPersonaId: "1032459747",
				EstadoPagoMensualId: &models.EstadoPagoMensual{
					Id: 8,
				},
				DocumentoResponsableId: "52204982",
				CargoResponsable:       "SUPERVISOR OFICINA ASESORA DE SISTEMAS",
				Ano:                    2018,
				Activo:                 true,
			},
			NombrePersona:     "CARLOS  ANDRES CONTRERAS BERNAL",
			NumeroCdp:         "765",
			VigenciaCdp:       "2018",
			NumeroRp:          "400",
			VigenciaRp:        "2018",
			NombreDependencia: "OFICINA ASESORA DE SISTEMAS",
			Rubro:             "Funcionamiento",
		},
	}
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(arregloPrueba); err != nil {
		beego.Error(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8090/v1/solicitudes_ordenador_contratistas/aprobar_pagos", b)
	r, err := client.Do(req)
	if err == nil {
		t.Log("TestAprobarMultiplesPagosContratistas Finalizado Correctamente (OK)")
	} else {
		t.Error("Error TestEndPoint(TestAprobarMultiplesPagosContratistas): Se esperaba 201 y se obtuvo", r.StatusCode)
		t.Fail()
	}
}

func TestCertificacionCumplidosContratistas(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_ordenador_contratistas/certificaciones/DEP14/10/2019"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndPoint(TestCertificacionCumplidosContratistas): Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestCertificacionCumplidosContratistas Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestCertificacionCumplidosContratistas:", err.Error())
		t.Fail()
	}

}

func TestCertificacionCumplidosContratistasError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_ordenador_contratistas/certificaciones/DEP14/10/201"); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestEndPoint(TestCertificacionCumplidosContratistasError): Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestCertificacionCumplidosContratistasError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestCertificacionCumplidosContratistasError:", err.Error())
		t.Fail()
	}

}

func TestGetSolicitudesOrdenadorContratistasDependencia(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_ordenador_contratistas/solicitudes_dependencia/19483708/DEP12/?limit=10&offset=0"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndPoint(TestGetSolicitudesOrdenadorContratistasDependencia): Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetSolicitudesOrdenadorContratistasDependencia Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetSolicitudesOrdenadorContratistasDependencia:", err.Error())
		t.Fail()
	}

}

func TestGetSolicitudesOrdenadorContratistasDependenciaError(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_ordenador_contratistas/solicitudes_dependencia/1/DEP12/?limit=10&offset=0"); err == nil {
		if response.StatusCode != 400 {
			t.Error("Error TestEndPoint(TestGetSolicitudesOrdenadorContratistasDependenciaError): Se esperaba 400 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetSolicitudesOrdenadorContratistasDependenciaError Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetSolicitudesOrdenadorContratistasDependenciaError:", err.Error())
		t.Fail()
	}

}

func TestGetCumplidosRevertiblesPorOrdenador(t *testing.T) {

	if response, err := http.Get("http://localhost:8090/v1/solicitudes_ordenador_contratistas/cumplidos_revertibles/19483708"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestEndPoint(TestGetCumplidosRevertiblesPorOrdenador): Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestGetCumplidosRevertiblesPorOrdenador Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error EndPoint TestGetCumplidosRevertiblesPorOrdenador:", err.Error())
		t.Fail()
	}

}
