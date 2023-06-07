package Informe_test

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
	UrlCrudCumplidos      string
	UrlcrudAgora          string
	UrlFinancieraJBPM     string
	UrlAdministrativaJBPM string
	UrlcrudCore           string
	UrlcrudOikos          string
	UrlNovedadesMid       string
	UrlTitanMid           string
	UrlDocumentosCrud     string
	UrlGestorDocumental   string
}

func TestMain(m *testing.M) {

	parameters.UrlCrudCumplidos = os.Getenv("UrlCrudCumplidos")
	parameters.UrlcrudAgora = os.Getenv("UrlcrudAgora")
	parameters.UrlFinancieraJBPM = os.Getenv("UrlFinancieraJBPM")
	parameters.UrlAdministrativaJBPM = os.Getenv("UrlAdministrativaJBPM")
	parameters.UrlcrudCore = os.Getenv("UrlcrudCore")
	parameters.UrlcrudOikos = os.Getenv("UrlcrudOikos")
	parameters.UrlNovedadesMid = os.Getenv("UrlNovedadesMid")
	parameters.UrlTitanMid = os.Getenv("UrlTitanMid")
	parameters.UrlDocumentosCrud = os.Getenv("UrlDocumentosCrud")
	parameters.UrlGestorDocumental = os.Getenv("UrlGestorDocumental")

	beego.AppConfig.Set("UrlCrudCumplidos", parameters.UrlCrudCumplidos)
	beego.AppConfig.Set("UrlcrudAgora", parameters.UrlcrudAgora)
	beego.AppConfig.Set("UrlFinancieraJBPM", parameters.UrlFinancieraJBPM)
	beego.AppConfig.Set("UrlAdministrativaJBPM", parameters.UrlAdministrativaJBPM)
	beego.AppConfig.Set("UrlcrudCore", parameters.UrlcrudCore)
	beego.AppConfig.Set("UrlcrudOikos", parameters.UrlcrudOikos)
	beego.AppConfig.Set("UrlNovedadesMid", parameters.UrlNovedadesMid)
	beego.AppConfig.Set("UrlTitanMid", parameters.UrlTitanMid)
	beego.AppConfig.Set("UrlDocumentosCrud", parameters.UrlDocumentosCrud)
	beego.AppConfig.Set("UrlGestorDocumental", parameters.UrlGestorDocumental)

	flag.Parse()
	os.Exit(m.Run())
}

//Informe ...
func TestInforme(t *testing.T) {
	Informe, err := helpers.Informe("94162")
	if err != nil {
		t.Error("Error helper func TestInforme", err)
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
		t.Error("Error helper func GetActividadesEspecificas", err)
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
		t.Error("Error helper func GetActividadesRealizadas", err)
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
		t.Error("Error helper func AddInforme", err)
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
		t.Error("Error helper func AddActividadEspecifica", err)
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
		t.Error("Error helper func AddActividadRealizada", err)
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
		t.Error("Error helper func UpdateInformeById", err)
		t.Fail()
	} else {
		t.Log("TestUpdateInformeById Finalizado Correctamente (OK)")
	}
}

//UltimoInformeContratista ...
func TestUltimoInformeContratista(t *testing.T) {
	informe, err := helpers.UltimoInformeContratista("94162")
	if err != nil {
		t.Error("Error helper func UltimoInformeContratista", err)
		t.Fail()
	} else {
		t.Log(informe)
		t.Log("TestUltimoInformeContratista Finalizado Correctamente (OK)")
	}
}

func TestEndPointCumplidosCpsMid(t *testing.T) {
	t.Log("-----------------------------------------------------")
	t.Log("Testing EndPoint UrlCrudCumplidos ")
	t.Log(parameters.UrlCrudCumplidos)
	t.Log("Testing EndPoint UrlcrudAgora")
	t.Log(parameters.UrlcrudAgora)
	t.Log("Testing EndPoint UrlFinancieraJBPM")
	t.Log(parameters.UrlFinancieraJBPM)
	t.Log("Testing EndPoint UrlAdministrativaJBPM")
	t.Log(parameters.UrlAdministrativaJBPM)
	t.Log("Testing EndPoint UrlcrudCore")
	t.Log(parameters.UrlcrudCore)
	t.Log("Testing EndPoint UrlcrudOikos")
	t.Log(parameters.UrlcrudOikos)
	t.Log("Testing EndPoint UrlNovedadesMid")
	t.Log(parameters.UrlNovedadesMid)
	t.Log("Testing EndPoint UrlTitanMid")
	t.Log(parameters.UrlTitanMid)
	t.Log("Testing EndPoint UrlDocumentosCrud")
	t.Log(parameters.UrlDocumentosCrud)
	t.Log("Testing EndPoint UrlGestorDocumental")
	t.Log(parameters.UrlGestorDocumental)
	t.Log("-----------------------------------------------------")
}
