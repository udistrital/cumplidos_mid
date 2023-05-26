package InformacionInforme_test

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/helpers"
	"github.com/udistrital/cumplidos_mid/models"
)

var parameters struct {
	UrlcrudWSO2          string
	NscrudAdministrativa string
	ProtocolAdmin        string
	UrlcrudAgora         string
	NscrudAgora          string
	NscrudFinanciera     string
}

func TestMain(m *testing.M) {
	parameters.UrlcrudWSO2 = os.Getenv("UrlcrudWSO2")
	parameters.NscrudAdministrativa = os.Getenv("NscrudAdministrativa")
	parameters.ProtocolAdmin = os.Getenv("ProtocolAdmin")
	parameters.UrlcrudAgora = os.Getenv("UrlcrudAgora")
	parameters.NscrudAgora = os.Getenv("NscrudAgora")
	parameters.NscrudFinanciera = os.Getenv("NscrudFinanciera")
	beego.AppConfig.Set("UrlcrudWSO2", parameters.UrlcrudWSO2)
	beego.AppConfig.Set("NscrudAdministrativa", parameters.NscrudAdministrativa)
	beego.AppConfig.Set("ProtocolAdmin", parameters.ProtocolAdmin)
	beego.AppConfig.Set("UrlcrudAgora", parameters.UrlcrudAgora)
	beego.AppConfig.Set("NscrudAgora", parameters.NscrudAgora)
	beego.AppConfig.Set("NscrudFinanciera", parameters.NscrudFinanciera)
	flag.Parse()
	os.Exit(m.Run())
}

//Informe ...
func TestInforme(t *testing.T) {
	Informe, err := helpers.Informe("94162")
	if err != nil {
		t.Error("No se pudo consultar la informacion del informe", err)
		t.Fail()
	} else {
		t.Log(Informe)
		t.Log("TestInforme Finalizado Correctamente (OK)")
	}
}

//GetActividadesEspecificas ...
func TestGetActividadesEspecificas(t *testing.T) {
	ActividadesEspecificas, err := helpers.GetActividadesEspecificas("96")
	if err != nil {
		t.Error("No se pudo consultar la informacion de las actividades especificas", err)
		t.Fail()
	} else {
		t.Log(ActividadesEspecificas)
		t.Log("TestGetActividadesEspecificas Finalizado Correctamente (OK)")
	}
}

//GetActividadesRealizadas ...
func TestGetActividadesRealizadas(t *testing.T) {
	ActividadesRealizadas, err := helpers.GetActividadesRealizadas("94")
	if err != nil {
		t.Error("No se pudo consultar la informacion de las actividades realizadas", err)
		t.Fail()
	} else {
		t.Log(ActividadesRealizadas)
		t.Log("TestGetActividadesRealizadas Finalizado Correctamente (OK)")
	}
}

//AddInforme ...
func TestAddInforme(t *testing.T) {
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
	informeCreado, err := helpers.AddInforme(informePrueba)
	if err != nil {
		t.Error("No se pudo crear el informe", err)
		t.Fail()
	} else {
		t.Log(informeCreado)
		t.Log("AddInforme Finalizado Correctamente (OK)")
	}
}

//AddActividadEspecifica ...
func TestAddActividadEspecifica(t *testing.T) {
	var actividad_esp = map[string]interface{}{"ActividadEspecifica": "prueba", "Avance": 20, "InformeId": map[string]interface{}{"Id": 46}}
	actEsp, err := helpers.AddActividadEspecifica(actividad_esp)
	if err != nil {
		t.Error("No se pudo crear las actividades especificas", err)
		t.Fail()
	} else {
		t.Log(actEsp)
		t.Log("TestAddActividadEspecifica Finalizado Correctamente (OK)")
	}
}

//AddActividadRealizada ...
func TestAddActividadRealizada(t *testing.T) {
	var actividad_rea = map[string]interface{}{"Actividad": "prueba", "ProductoAsociado": "prueba", "Evidencia": "prueba", "ActividadEspecificaId": map[string]interface{}{"Id": 96}}
	actRea, err := helpers.AddActividadRealizada(actividad_rea)
	if err != nil {
		t.Error("No se pudo crear las actividades realizadas", err)
		t.Fail()
	} else {
		t.Log(actRea)
		t.Log("TestAddActividadRealizada Finalizado Correctamente (OK)")
	}
}

//UpdateInformeById ...
func TestUpdateInformeById(t *testing.T) {
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
	err := helpers.UpdateInformeById(informePrueba)
	if err != nil {
		t.Error("No se pudo actualizar el informe", err)
		t.Fail()
	} else {
		t.Log("TestUpdateInformeById Finalizado Correctamente (OK)")
	}
}

//UltimoInformeContratista ...
func TestUltimoInformeContratista(t *testing.T) {
	informe, err := helpers.UltimoInformeContratista("94162")
	if err != nil {
		t.Error("No se pudo consultar la informacion del ultimo informe", err)
		t.Fail()
	} else {
		t.Log(informe)
		t.Log("TestUltimoInformeContratista Finalizado Correctamente (OK)")
	}
}
