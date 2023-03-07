package InformacionInforme_test

import (
	"flag"
	"os"
	"testing"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/helpers"
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

//InformacionInforme ...
func TestInformacionInforme(t *testing.T) {
	InformacionInforme, err := helpers.InformacionInforme("94158")
	if err != nil {
		t.Error("No se pudo consultar la informacion del informe", err)
		t.Fail()
	} else {
		t.Log(InformacionInforme)
		t.Log("TestInformacionInforme Finalizado Correctamente (OK)")
	}
}
