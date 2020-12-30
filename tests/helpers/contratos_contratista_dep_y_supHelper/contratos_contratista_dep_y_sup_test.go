package contratos_contratista_dep_y_supHelper

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
	ProtocolCrudCumplidos string
	UrlCrudCumplidos string
	NsCrudCumplidos string
	NscrudFinanciera string
}

func TestMain(m *testing.M) {
	parameters.UrlcrudWSO2 = os.Getenv("UrlcrudWSO2")
	parameters.NscrudAdministrativa = os.Getenv("NscrudAdministrativa")
	parameters.ProtocolAdmin = os.Getenv("ProtocolAdmin")
	parameters.UrlcrudAgora = os.Getenv("UrlcrudAgora")
	parameters.NscrudAgora = os.Getenv("NscrudAgora")
	parameters.ProtocolCrudCumplidos = os.Getenv("ProtocolCrudCumplidos")
	parameters.UrlCrudCumplidos = os.Getenv("UrlCrudCumplidos")
	parameters.NsCrudCumplidos = os.Getenv("NsCrudCumplidos")
	parameters.NscrudFinanciera = os.Getenv("NscrudFinanciera")
	beego.AppConfig.Set("UrlcrudWSO2", parameters.UrlcrudWSO2)
	beego.AppConfig.Set("NscrudAdministrativa", parameters.NscrudAdministrativa)
	beego.AppConfig.Set("ProtocolAdmin", parameters.ProtocolAdmin)
	beego.AppConfig.Set("UrlcrudAgora", parameters.UrlcrudAgora)
	beego.AppConfig.Set("NscrudAgora", parameters.NscrudAgora)
	beego.AppConfig.Set("ProtocolCrudCumplidos", parameters.ProtocolCrudCumplidos)
	beego.AppConfig.Set("UrlCrudCumplidos", parameters.UrlCrudCumplidos)
	beego.AppConfig.Set("NsCrudCumplidos", parameters.NsCrudCumplidos)
	beego.AppConfig.Set("NscrudFinanciera", parameters.NscrudFinanciera)
	flag.Parse()
	os.Exit(m.Run())
}

// CertificacionDocumentosAprobados ...
func TestContratosContratistaDependencia(t *testing.T) {
	valor, err := helpers.ContratosContratistaDependencia("19483708", "DEP12", 10, 0)
	if err != nil {
		t.Error("No se pudo consultar las actas de recibido", err)
		t.Fail()	
	} else {
		t.Log(valor)
		t.Log("TestCertificacionDocumentosAprobados Finalizado Correctamente (OK)")
	}
}

// CertificadoVistoBueno ...
func TestContratosContratistaSupervisor(t *testing.T) {
	valor, err := helpers.ContratosContratistaSupervisor("52204982")
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
	t.Log("Testing EndPoint ProtocolCrudCumplidos")
	t.Log(parameters.ProtocolCrudCumplidos)
	t.Log("Testing EndPoint UrlCrudCumplidos")
	t.Log(parameters.UrlCrudCumplidos)
	t.Log("Testing EndPoint NsCrudCumplidos")
	t.Log(parameters.NsCrudCumplidos)
	t.Log("Testing EndPoint NscrudFinanciera")
	t.Log(parameters.NscrudFinanciera)
}