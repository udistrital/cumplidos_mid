package contratosHelper

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
}

func TestMain(m *testing.M) {
	parameters.UrlcrudWSO2 = os.Getenv("UrlcrudWSO2")
	parameters.NscrudAdministrativa = os.Getenv("NscrudAdministrativa")
	beego.AppConfig.Set("UrlcrudWSO2", parameters.UrlcrudWSO2)
	beego.AppConfig.Set("NscrudAdministrativa", parameters.NscrudAdministrativa)
	flag.Parse()
	os.Exit(m.Run())
}

// GetContratosDependencia ...
func TestGetContratosDependencia(t *testing.T) {
	valor, err := helpers.GetContratosDependencia("DEP12", "2020-03")
	if err != nil {
		t.Error("No se pudo consultar las actas de recibido", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestCertificacionDocumentosAprobados Finalizado Correctamente (OK)")
	}
}

// GetContratosDependenciaFiltro ...
func TestGetContratosDependenciaFiltro(t *testing.T) {
	valor, err := helpers.GetContratosDependenciaFiltro("DEP12", "2020-12", "2020-11")
	if err != nil {
		t.Error("No se pudo consultar las actas de recibido por tipo", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestGetActasRecibidoTipo Finalizado Correctamente (OK)")
	}
}

// GetContratosOrdenadorDependencia ...
func TestGetContratosOrdenadorDependencia(t *testing.T) {
	valor, err := helpers.GetContratosOrdenadorDependencia("17", "4", "2019")
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
}