package InformacionInforme_test

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

//InformacionInforme ...
func TestInformacionInforme(t *testing.T) {
	InformacionInforme, err := helpers.InformacionInforme("94158")
	if err != nil {
		t.Error("Error helper func InformacionInforme:", err)
		t.Fail()
	} else {
		t.Log(InformacionInforme)
		t.Log("TestInformacionInforme Finalizado Correctamente (OK)")
	}
}

//GetPreliquidacion ...
func TestGetPreliquidacion(t *testing.T) {
	InformacionInforme, err := helpers.GetPreliquidacion("94158")
	if err != nil {
		t.Error("Error helper func TestGetPreliquidacion:", err)
		t.Fail()
	} else {
		t.Log(InformacionInforme)
		t.Log("TestGetPreliquidacion Finalizado Correctamente (OK)")
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
