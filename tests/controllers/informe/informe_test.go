package Informetest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

func TestGetInforme(t *testing.T) {
	if response, err := http.Get("http://localhost:8090/v1/informe/94162"); err == nil {
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

func TestGetUltimoInformeContratista(t *testing.T) {
	if response, err := http.Get("http://localhost:8090/v1/informe/1265/2021/1014294957"); err == nil {
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

func TestPostInforme(t *testing.T) {

	informePrueba := models.Informe{
		PeriodoInformeInicio: time.Now(),
		PeriodoInformeFin:    time.Now(),
		Proceso:              "Gestión de los Sistemas de Información y las Telecomunicaciones",
		PagoMensualId: &models.PagoMensual{
			Id: 94162,
		},
		ActividadesEspecificas: []models.ActividadEspecifica{
			{
				ActividadEspecifica: "Diseñar y desarrollar los componentes de desarrollo (menús, funcionalidades, motor de reglas, servicios, servicios interoperables, etc, que apliquen)",
				Avance:              20,
				ActividadesRealizadas: []models.ActividadRealizada{
					{
						Actividad:        "Actividad 1",
						ProductoAsociado: "Producto 1",
						Evidencia:        "https://github.com/udistrital/cumplidos_cliente/blob/feature/AutomatizacionInforme/app/scripts/controllers/seguimientoycontrol/tecnico/informe_gestion_y_certificado_cumplimiento.js",
					},
					{
						Actividad:        "Actividad 2",
						ProductoAsociado: "Producto 2",
						Evidencia:        "https://github.com/udistrital/cumplidos_cliente/blob/feature/AutomatizacionInforme/app/scripts/controllers/seguimientoycontrol/tecnico",
					},
					{
						Actividad:        "Actividad 3",
						ProductoAsociado: "Producto 3",
						Evidencia:        "https://github.com/udistrital/cumplidos_cliente/blob/feature/AutomatizacionInforme/app/scripts/controllers/seguimientoycontrol",
					},
					{
						Actividad:        "Actividad 4",
						ProductoAsociado: "Producto 4",
						Evidencia:        "https://github.com/udistrital/cumplidos_cliente/blob/feature/AutomatizacionInforme/app/scripts/controllers",
					},
				},
			},
		},
	}

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(informePrueba); err != nil {
		beego.Error(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8090/v1/informe", b)
	r, err := client.Do(req)
	if err == nil {
		t.Log("TestEndPoint Finalizado Correctamente (OK)")
	} else {
		t.Error("Error TestEndPoint: Se esperaba 201 y se obtuvo", r.StatusCode)
		t.Fail()
	}
}

func TestPutInforme(t *testing.T) {
	informePrueba := models.Informe{
		PeriodoInformeInicio: time.Now(),
		PeriodoInformeFin:    time.Now(),
		Proceso:              "Gestión de los Sistemas de Información y las Telecomunicaciones",
		PagoMensualId: &models.PagoMensual{
			Id: 94162,
		},
		ActividadesEspecificas: []models.ActividadEspecifica{
			{
				ActividadEspecifica: "Diseñar y desarrollar los componentes de desarrollo (menús, funcionalidades, motor de reglas, servicios, servicios interoperables, etc, que apliquen)",
				Avance:              20,
				ActividadesRealizadas: []models.ActividadRealizada{
					{
						Actividad:        "Actividad 1",
						ProductoAsociado: "Producto 1",
						Evidencia:        "https://github.com/udistrital/cumplidos_cliente/blob/feature/AutomatizacionInforme/app/scripts/controllers/seguimientoycontrol/tecnico/informe_gestion_y_certificado_cumplimiento.js",
					},
					{
						Actividad:        "Actividad 2",
						ProductoAsociado: "Producto 2",
						Evidencia:        "https://github.com/udistrital/cumplidos_cliente/blob/feature/AutomatizacionInforme/app/scripts/controllers/seguimientoycontrol/tecnico",
					},
					{
						Actividad:        "Actividad 3",
						ProductoAsociado: "Producto 3",
						Evidencia:        "https://github.com/udistrital/cumplidos_cliente/blob/feature/AutomatizacionInforme/app/scripts/controllers/seguimientoycontrol",
					},
					{
						Actividad:        "Actividad 4",
						ProductoAsociado: "Producto 4",
						Evidencia:        "https://github.com/udistrital/cumplidos_cliente/blob/feature/AutomatizacionInforme/app/scripts/controllers",
					},
				},
			},
		},
	}

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(informePrueba); err != nil {
		beego.Error(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8090/v1/informe", b)
	r, err := client.Do(req)
	if err == nil {
		t.Log("TestEndPoint Finalizado Correctamente (OK)")
	} else {
		t.Error("Error TestEndPoint: Se esperaba 200 y se obtuvo", r.StatusCode)
		t.Fail()
	}
}
