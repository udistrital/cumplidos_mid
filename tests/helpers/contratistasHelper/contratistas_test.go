package contratistasHelper

import (
	"flag"
	"os"
	"testing"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
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

// CertificacionCumplidosContratistas ...
func TestCertificacionCumplidosContratistas(t *testing.T) {
	valor, err := helpers.CertificacionCumplidosContratistas("DEP14", "10", "2019")
	t.Log("-----------------------------------------------------")
	if err != nil {
		t.Error("Error helper func CertificacionCumplidosContratistas", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestCertificacionCumplidosContratistas Finalizado Correctamente (OK)")
	}
	t.Log("-----------------------------------------------------")
}

// AprobacionPagosContratistas ...
func TestAprobacionPagosContratistas(t *testing.T) {
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
	err := helpers.AprobacionPagosContratistas(arregloPrueba)
	if err != nil {
		t.Error("Error helper func AprobacionPagosContratistas", err)
		t.Fail()
	} else {
		t.Log("TestAprobacionPagosContratistas Finalizado Correctamente (OK)")
	}
}

// SolicitudesOrdenadorContratistas ...
func TestSolicitudesOrdenadorContratistas(t *testing.T) {
	t.Log("-----------------------------------------------------")
	valor, err := helpers.SolicitudesOrdenadorContratistas("52204982", 5, 0)
	if err != nil {
		t.Error("Error helper func SolicitudesOrdenadorContratistas", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestSolicitudesOrdenadorContratistas Finalizado Correctamente (OK)")
	}
	t.Log("-----------------------------------------------------")
}

// TraerInfoOrdenador ...
func TestTraerInfoOrdenador(t *testing.T) {
	t.Log("-----------------------------------------------------")
	valor, err := helpers.TraerInfoOrdenador("11824", "2023")
	if err != nil {
		t.Error("Error helper func TraerInfoOrdenador", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestTraerInfoOrdenador Finalizado Correctamente (OK)")
	}
	t.Log("-----------------------------------------------------")
}

// GetCumplidosRevertiblesPorOrdenador ...
func TestGetCumplidosRevertiblesPorOrdenador(t *testing.T) {
	t.Log("-----------------------------------------------------")
	valor, err := helpers.GetCumplidosRevertiblesPorOrdenador("19483708")
	if err != nil {
		t.Error("Error helper func GetCumplidosRevertiblesPorOrdenador", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestGetCumplidosRevertiblesPorOrdenador Finalizado Correctamente (OK)")
	}
	t.Log("-----------------------------------------------------")
}

// TraerEnlacesDocumentosAsociadosPagoMensual ...
func TestTraerEnlacesDocumentosAsociadosPagoMensual(t *testing.T) {
	t.Log("-----------------------------------------------------")
	valor, err := helpers.TraerEnlacesDocumentosAsociadosPagoMensual("94190")
	if err != nil {
		t.Error("Error helper func TraerEnlacesDocumentosAsociadosPagoMensual", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestTraerEnlacesDocumentosAsociadosPagoMensual Finalizado Correctamente (OK)")
	}
	t.Log("-----------------------------------------------------")
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
