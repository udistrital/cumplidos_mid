package contratos_contratistaHelper

import (
	"flag"
	"os"
	"testing"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/helpers"
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

// CertificacionDocumentosAprobados ...
func TestContratosContratista(t *testing.T) {
	valor, err := helpers.ContratosContratista("1032490151")
	if err != nil {
		t.Error("Error helper func ContratosContratista", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestContratosContratista Finalizado Correctamente (OK)")
	}
}

// GetRP
func TestGetRP(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("GetRP panic: %v", r)
		}
	}()

	numeroCDP := "2400"
	vigenciaCDP := "2020"
	unidadEjecucion := "1"

	valor, outputError := helpers.GetRP(
		numeroCDP,
		vigenciaCDP,
		unidadEjecucion,
	)

	if outputError != nil {
		t.Fatalf("Error helper func GetRP: %v", outputError)
	}

	t.Log(valor)
	t.Log("TestGetRP Finalizado Correctamente (OK)")
}

// GetContratosPersona ...
func TestGetContratosPersona(t *testing.T) {
	valor, err := helpers.GetContratosPersona("1032490151")
	if err != nil {
		t.Error("Error helper func GetContratosPersona", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestGetContratosPersona Finalizado Correctamente (OK)")
	}
}

// GetContrato ...
func TestGetContrato(t *testing.T) {
	valor, err := helpers.GetContrato("1406", "2020")
	if err != nil {
		t.Error("Error helper func GetContrato", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestGetContrato Finalizado Correctamente (OK)")
	}
}

// GetInformacionContratoContratista ...
func TestGetInformacionContratoContratista(t *testing.T) {
	valor, err := helpers.GetInformacionContratoContratista("1406", "2020")
	if err != nil {
		t.Error("Error helper func GetInformacionContratoContratista", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestGetInformacionContratoContratista Finalizado Correctamente (OK)")
	}
}

// GetActaDeInicio ...
func TestGetActaDeInicio(t *testing.T) {
	valor, err := helpers.GetActaDeInicio("1406", 2020)
	if err != nil {
		t.Error("Error helper func GetActaDeInicio", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestGetActaDeInicio Finalizado Correctamente (OK)")
	}
}

func TestEndPointCumplidosCpsMid(t *testing.T) {
	t.Log("-----------------------------------------------------")
	t.Log("Testing EndPoint UrlCrudCumplidos ")
	t.Log("URLcrud: ", parameters.UrlCrudCumplidos)
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
