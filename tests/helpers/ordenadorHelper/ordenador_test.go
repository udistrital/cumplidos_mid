package ordenadorHelper

import (
	"flag"
	"os"
	"testing"

	_ "github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/helpers"
)

var parameters struct {
	Certificacion string
}

func TestMain(m *testing.M) {
	parameters.Certificacion = os.Getenv("Certificacion")
	//beego.AppConfig.Set("ActaRecibidoService", os.Getenv("ACTA_RECIBIDO_CRUD"))
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
	valor, err := helpers.CertificadoVistoBueno("17", "2020", "6")
	if err != nil {
		t.Error("No se pudo consultar las actas de recibido por tipo", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestGetActasRecibidoTipo Finalizado Correctamente (OK)")
	}
}


func TestEndPointCertificacion(t *testing.T) {
	t.Log("Testing EndPoint Certificacion")
	t.Log(parameters.Certificacion)
}