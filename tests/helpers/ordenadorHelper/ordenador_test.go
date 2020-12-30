package ordenadorHelper

import (
	"flag"
	"os"
	"testing"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/helpers"
)

var parameters struct {
	UrlcrudWSO2           string
	NscrudAdministrativa  string
	ProtocolCrudCumplidos string
	UrlCrudCumplidos      string
	NsCrudCumplidos       string
	UrlcrudAgora          string
	NscrudAgora           string
	ProtocolAdmin         string
	NscrudFinanciera      string
	UrlcrudCore           string
	NscrudCore            string
	UrlcrudOikos          string
	NscrudOikos           string
	UrlcrudAdmin          string
	NscrudAdmin           string
}

func TestMain(m *testing.M) {

	parameters.UrlcrudWSO2 = os.Getenv("UrlcrudWSO2")
	parameters.NscrudAdministrativa = os.Getenv("NscrudAdministrativa")
	parameters.ProtocolCrudCumplidos = os.Getenv("ProtocolCrudCumplidos")
	parameters.UrlCrudCumplidos = os.Getenv("UrlCrudCumplidos")
	parameters.NsCrudCumplidos = os.Getenv("NsCrudCumplidos")
	parameters.UrlcrudAgora = os.Getenv("UrlcrudAgora")
	parameters.NscrudAgora = os.Getenv("NscrudAgora")
	parameters.ProtocolAdmin = os.Getenv("ProtocolAdmin")
	//parameters.NscrudFinanciera = os.Getenv("NscrudFinanciera")
	parameters.UrlcrudCore = os.Getenv("UrlcrudCore")
	parameters.NscrudCore = os.Getenv("NscrudCore")
	parameters.UrlcrudOikos = os.Getenv("UrlcrudOikos")
	parameters.NscrudOikos = os.Getenv("NscrudOikos")
	parameters.UrlcrudAdmin = os.Getenv("UrlcrudAdmin")
	parameters.NscrudAdmin = os.Getenv("NscrudAdmin")

	beego.AppConfig.Set("UrlcrudWSO2", parameters.UrlcrudWSO2)
	beego.AppConfig.Set("NscrudAdministrativa", parameters.NscrudAdministrativa)
	beego.AppConfig.Set("ProtocolAdmin", parameters.ProtocolAdmin)
	beego.AppConfig.Set("UrlcrudCore", parameters.UrlcrudCore)
	beego.AppConfig.Set("NscrudCore", parameters.NscrudCore)
	beego.AppConfig.Set("UrlCrudAgora", parameters.UrlcrudAgora)
	beego.AppConfig.Set("NsCrudAgora", parameters.NscrudAgora)
	beego.AppConfig.Set("ProtocolCrudCumplidos", parameters.ProtocolCrudCumplidos)
	beego.AppConfig.Set("UrlCrudCumplidos", parameters.UrlCrudCumplidos)
	beego.AppConfig.Set("NsCrudCumplidos", parameters.NsCrudCumplidos)
	beego.AppConfig.Set("UrlcrudOikos", parameters.UrlcrudOikos)
	beego.AppConfig.Set("NscrudOikos", parameters.NscrudOikos)
	beego.AppConfig.Set("UrlcrudAdmin", parameters.UrlcrudAdmin)
	beego.AppConfig.Set("NscrudAdmin", parameters.NscrudAdmin)

	flag.Parse()
	os.Exit(m.Run())
}

// TraerInfoOrdenador ...
func TestTraerInfoOrdenador(t *testing.T) {
	t.Log("-----------------------------------------------------")
	valor, err := helpers.TraerInfoOrdenador("DVE1569", "2020")
	if err != nil {
		t.Error("No se pudo consultar la información del ordenadoor del gasto", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestTraerInfoOrdenador Finalizado Correctamente (OK)")
	}
	t.Log("-----------------------------------------------------")
}

// SolicitudesOrdenador ...
func TestSolicitudesOrdenador(t *testing.T) {
	t.Log("-----------------------------------------------------")
	valor, err := helpers.SolicitudesOrdenador("19128837", 5, 0)
	if err != nil {
		t.Error("No se pudo consultar las solicitudes correspondientes al ordenador del gasto", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestSolicitudesOrdenador Finalizado Correctamente (OK)")
	}
	t.Log("-----------------------------------------------------")
}

// DependenciaOrdenador ...
func TestDependenciaOrdenador(t *testing.T) {
	t.Log("-----------------------------------------------------")
	valor, err := helpers.DependenciaOrdenador("19400342")
	if err != nil {
		t.Error("No se pudo consultar la dependencia del ordenador", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestDependenciaOrdenador Finalizado Correctamente (OK)")
	}
	t.Log("-----------------------------------------------------")
}

func TestEndPointCertificacion(t *testing.T) {
	t.Log("-----------------------------------------------------")
	t.Log("Testing EndPoint UrlcrudWSO2 ")
	t.Log(parameters.UrlcrudWSO2)
	t.Log("Testing EndPoint NscrudAdministrativa")
	t.Log(parameters.NscrudAdministrativa)
	t.Log("Testing EndPoint ProtocolCrudCumplidos")
	t.Log(parameters.ProtocolCrudCumplidos)
	t.Log("Testing EndPoint UrlCrudCumplidos")
	t.Log(parameters.UrlCrudCumplidos)
	t.Log("Testing EndPoint NsCrudCumplidos")
	t.Log(parameters.NsCrudCumplidos)
	t.Log("Testing EndPoint UrlcrudAgora")
	t.Log(parameters.UrlcrudAgora)
	t.Log("Testing EndPoint NscrudAgora")
	t.Log(parameters.NscrudAgora)
	t.Log("Testing EndPoint ProtocolAdmin")
	t.Log(parameters.ProtocolAdmin)
	t.Log("-----------------------------------------------------")
}
