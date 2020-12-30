package solicitud_coordinadorHelper

import (
	"flag"
	"os"
	"testing"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/helpers"
)

var parameters struct {
	ProtocolCrudCumplidos string
	UrlCrudCumplidos      string
	NsCrudCumplidos       string
	UrlcrudAgora          string
	NscrudAgora           string
	ProtocolAdmin         string
	UrlcrudAdmin          string
	NscrudAdmin           string
	UrlcrudOikos          string
	NscrudOikos           string
}

func TestMain(m *testing.M) {
	parameters.ProtocolCrudCumplidos = os.Getenv("ProtocolCrudCumplidos")
	parameters.UrlCrudCumplidos = os.Getenv("UrlCrudCumplidos")
	parameters.NsCrudCumplidos = os.Getenv("NsCrudCumplidos")
	parameters.UrlcrudAgora = os.Getenv("UrlcrudAgora")
	parameters.NscrudAgora = os.Getenv("NscrudAgora")
	parameters.ProtocolAdmin = os.Getenv("ProtocolAdmin")
	parameters.UrlcrudAdmin = os.Getenv("UrlcrudAdmin")
	parameters.NscrudAdmin = os.Getenv("NscrudAdmin")
	parameters.UrlcrudOikos = os.Getenv("UrlcrudOikos")
	parameters.NscrudOikos = os.Getenv("NscrudOikos")

	beego.AppConfig.Set("ProtocolCrudCumplidos", parameters.ProtocolCrudCumplidos)
	beego.AppConfig.Set("UrlCrudCumplidos", parameters.UrlCrudCumplidos)
	beego.AppConfig.Set("NsCrudCumplidos", parameters.NsCrudCumplidos)
	beego.AppConfig.Set("UrlCrudAgora", parameters.UrlcrudAgora)
	beego.AppConfig.Set("NsCrudAgora", parameters.NscrudAgora)
	beego.AppConfig.Set("ProtocolAdmin", parameters.ProtocolAdmin)
	beego.AppConfig.Set("UrlcrudAdmin", parameters.UrlcrudAdmin)
	beego.AppConfig.Set("NscrudAdmin", parameters.NscrudAdmin)
	beego.AppConfig.Set("UrlcrudOikos", parameters.UrlcrudOikos)
	beego.AppConfig.Set("NscrudOikos", parameters.NscrudOikos)

	flag.Parse()
	os.Exit(m.Run())
}

// CertificacionDocumentosAprobados ...
func TestSolicitudCoordinador(t *testing.T) {
	valor, err := helpers.SolicitudCoordinador("19346572")
	if err != nil {
		t.Error("No se pudieron obtener las solicitudes para este ordenador", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestSolicitudCoordinador Finalizado Correctamente (OK)")
	}
}

func TestEndPointCertificacion(t *testing.T) {

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
	t.Log("Testing EndPoint UrlcrudAdmin")
	t.Log(parameters.UrlcrudAdmin)
	t.Log("Testing EndPoint NscrudAdmin")
	t.Log(parameters.NscrudAdmin)
	t.Log("Testing EndPoint UrlcrudOikos")
	t.Log(parameters.UrlcrudOikos)
	t.Log("Testing EndPoint NscrudOikos")
	t.Log(parameters.NscrudOikos)
}
