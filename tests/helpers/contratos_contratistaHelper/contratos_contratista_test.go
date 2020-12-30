package contratos_contratistaHelper

import (
	"flag"
	"os"
	"testing"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/helpers"
)

var parameters struct {
	UrlcrudWSO2 string
	NscrudAdministrativa string
	ProtocolAdmin string
	UrlcrudAgora string
	NscrudAgora string
	NscrudFinanciera string
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

// CertificacionDocumentosAprobados ...
func TestContratosContratista(t *testing.T) {
	valor, err := helpers.ContratosContratista("1032490151")
	if err != nil {
		t.Error("No se pudo consultar las actas de recibido", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestCertificacionDocumentosAprobados Finalizado Correctamente (OK)")
	}
}

//GetRP
func TestGetRP(t *testing.T) {
	valor, err := helpers.GetRP("2400", "2020")
	if err != nil {
		t.Error("No se pudo consultar las actas de recibido por tipo", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestGetActasRecibidoTipo Finalizado Correctamente (OK)")
	}
}

// GetContratosPersona ...
func TestGetContratosPersona(t *testing.T) {
	valor, err := helpers.GetContratosPersona("1032490151")
	if err != nil {
		t.Error("No se pudo consultar las actas de recibido por tipo", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestGetActasRecibidoTipo Finalizado Correctamente (OK)")
	}
}

// GetContrato ...
func TestGetContrato(t *testing.T) {
	valor, err := helpers.GetContrato("1406", "2020")
	if err != nil {
		t.Error("No se pudo consultar las actas de recibido por tipo", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestGetActasRecibidoTipo Finalizado Correctamente (OK)")
	}
}

// GetInformacionContratoContratista ...
func TestGetInformacionContratoContratista(t *testing.T) {
	valor, err := helpers.GetInformacionContratoContratista("1406", "2020")
	if err != nil {
		t.Error("No se pudo consultar las actas de recibido por tipo", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestGetActasRecibidoTipo Finalizado Correctamente (OK)")
	}
}


func TestEndPointCertificacion(t *testing.T) {
	t.Log("Testing EndPoint UrlcrudWSO2")
	t.Log(parameters.UrlcrudWSO2)
	t.Log("Testing EndPoint NscrudAdministrativa")
	t.Log(parameters.NscrudAdministrativa)
	t.Log("Testing EndPoint ProtocolAdmin")
	t.Log(parameters.ProtocolAdmin)
	t.Log("Testing EndPoint UrlcrudAgora")
	t.Log(parameters.UrlcrudAgora)
	t.Log("Testing EndPoint NscrudAgora")
	t.Log(parameters.NscrudAgora)
	t.Log("Testing EndPoint NscrudFinanciera")
	t.Log(parameters.NscrudFinanciera)
}