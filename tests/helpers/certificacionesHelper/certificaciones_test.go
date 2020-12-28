package certificacionesHelper

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
	UrlcrudAdmin string
	NscrudAdmin string
	ProtocolCrudCumplidos string
	UrlCrudCumplidos string
	NsCrudCumplidos string
}

func TestMain(m *testing.M) {
	parameters.UrlcrudWSO2 = os.Getenv("UrlcrudWSO2")
	parameters.NscrudAdministrativa = os.Getenv("NscrudAdministrativa")
	parameters.ProtocolAdmin = os.Getenv("ProtocolAdmin")
	parameters.UrlcrudAdmin = os.Getenv("UrlcrudAdmin")
	parameters.NscrudAdmin = os.Getenv("NscrudAdmin")
	parameters.ProtocolCrudCumplidos = os.Getenv("ProtocolCrudCumplidos")
	parameters.UrlCrudCumplidos = os.Getenv("UrlCrudCumplidos")
	parameters.NsCrudCumplidos = os.Getenv("NsCrudCumplidos")
	beego.AppConfig.Set("UrlcrudWSO2", parameters.UrlcrudWSO2)
	beego.AppConfig.Set("NscrudAdministrativa", parameters.NscrudAdministrativa)
	beego.AppConfig.Set("ProtocolAdmin", parameters.ProtocolAdmin)
	beego.AppConfig.Set("UrlcrudAdmin", parameters.UrlcrudAdmin)
	beego.AppConfig.Set("NscrudAdmin", parameters.NscrudAdmin)
	beego.AppConfig.Set("ProtocolCrudCumplidos", parameters.ProtocolCrudCumplidos)
	beego.AppConfig.Set("UrlCrudCumplidos", parameters.UrlCrudCumplidos)
	beego.AppConfig.Set("NsCrudCumplidos", parameters.NsCrudCumplidos)
	flag.Parse()
	os.Exit(m.Run())
}

// CertificacionDocumentosAprobados ...
func TestCertificacionDocumentosAprobados(t *testing.T) {
	valor, err := helpers.CertificacionDocumentosAprobados("17", "2020", "6")
	if err != nil {
		t.Error("No se pudo consultar las actas de recibido", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestCertificacionDocumentosAprobados Finalizado Correctamente (OK)")
	}
}

// CertificadoVistoBueno ...
func TestCertificadoVistoBueno(t *testing.T) {
	valor, err := helpers.CertificadoVistoBueno("38", "10", "2018")
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
	t.Log("Testing EndPoint UrlcrudAdmin")
	t.Log(parameters.UrlcrudAdmin)
	t.Log("Testing EndPoint NscrudAdmin")
	t.Log(parameters.NscrudAdmin)
	t.Log("Testing EndPoint ProtocolCrudCumplidos")
	t.Log(parameters.ProtocolCrudCumplidos)
	t.Log("Testing EndPoint UrlCrudCumplidos")
	t.Log(parameters.UrlCrudCumplidos)
	t.Log("Testing EndPoint NsCrudCumplidos")
	t.Log(parameters.NsCrudCumplidos)
}