package contratistasHelper

import (
	"flag"
	"os"
	"testing"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/helpers"
	"github.com/udistrital/cumplidos_mid/models"
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
	//UrlcrudAdmin          string
	//NscrudAdmin           string
	//UrlcrudOikos          string
	//NscrudOikos           string
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
	parameters.NscrudFinanciera = os.Getenv("NscrudFinanciera")
	//parameters.UrlcrudAdmin = os.Getenv("UrlcrudAdmin")
	//parameters.NscrudAdmin = os.Getenv("NscrudAdmin")
	//parameters.UrlcrudOikos = os.Getenv("UrlcrudOikos")
	//parameters.NscrudOikos = os.Getenv("NscrudOikos")

	beego.AppConfig.Set("UrlcrudWSO2", parameters.UrlcrudWSO2)
	beego.AppConfig.Set("NscrudAdministrativa", parameters.NscrudAdministrativa)
	beego.AppConfig.Set("ProtocolCrudCumplidos", parameters.ProtocolCrudCumplidos)
	beego.AppConfig.Set("UrlCrudCumplidos", parameters.UrlCrudCumplidos)
	beego.AppConfig.Set("NsCrudCumplidos", parameters.NsCrudCumplidos)
	beego.AppConfig.Set("UrlCrudAgora", parameters.UrlcrudAgora)
	beego.AppConfig.Set("NsCrudAgora", parameters.NscrudAgora)
	beego.AppConfig.Set("ProtocolAdmin", parameters.ProtocolAdmin)
	beego.AppConfig.Set("NscrudFinanciera", parameters.NscrudFinanciera)
	//beego.AppConfig.Set("NscrudAdmin", parameters.NscrudAdmin)
	//beego.AppConfig.Set("UrlcrudOikos", parameters.UrlcrudOikos)
	//beego.AppConfig.Set("NscrudOikos", parameters.NscrudOikos)

	flag.Parse()
	os.Exit(m.Run())
}

// CertificacionCumplidosContratistas ...
func TestCertificacionCumplidosContratistas(t *testing.T) {
	valor, err := helpers.CertificacionCumplidosContratistas("DEP14", "10", "2019")
	t.Log("-----------------------------------------------------")
	if err != nil {
		t.Error("No se pudo consultar los certifiados de cumplidos de los contratistas por: ", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestCertificacionCumplidosContratistas Finalizado Correctamente (OK)")
	}
	t.Log("-----------------------------------------------------")
}

// AprobacionPagosContratistas ...
func TestAprobacionPagosContratistas(t *testing.T) {
	arregloPrueba := []models.PagoContratistaCdpRp{
		{
			PagoMensual: &models.PagoMensual{
				Id:                 173298,
				NumeroContrato:     "1406",
				VigenciaContrato:   2020,
				Mes:                3,
				DocumentoPersonaId: "1032490151",
				EstadoPagoMensualId: &models.EstadoPagoMensual{
					Id: 8,
				},
				DocumentoResponsableId: "52204982",
				CargoResponsable:       "SUPERVISOR OFICINA ASESORA DE SISTEMAS",
				Ano:                    2020,
				Activo:                 false,
			},
			NombrePersona:     "DIEGO ALEJANDRO GUTIERREZ ROJAS",
			NumeroCdp:         "2400",
			VigenciaCdp:       "2020",
			NumeroRp:          "23478",
			VigenciaRp:        "2020",
			NombreDependencia: "OFICINA ASESORA DE SISTEMAS",
			Rubro:             "Inversión",
		},
		{
			PagoMensual: &models.PagoMensual{
				Id:                 10435,
				NumeroContrato:     "250",
				VigenciaContrato:   2018,
				Mes:                9,
				DocumentoPersonaId: "1032459747",
				EstadoPagoMensualId: &models.EstadoPagoMensual{
					Id: 8,
				},
				DocumentoResponsableId: "52204982",
				CargoResponsable:       "SUPERVISOR OFICINA ASESORA DE SISTEMAS",
				Ano:                    2018,
				Activo:                 true,
			},
			NombrePersona:     "CARLOS  ANDRES CONTRERAS BERNAL",
			NumeroCdp:         "765",
			VigenciaCdp:       "2018",
			NumeroRp:          "400",
			VigenciaRp:        "2018",
			NombreDependencia: "OFICINA ASESORA DE SISTEMAS",
			Rubro:             "Funcionamiento",
		},
	}
	err := helpers.AprobacionPagosContratistas(arregloPrueba)
	if err != nil {
		t.Error("No se pudo generar la aprobación de pagos", err)
		t.Fail()
	} else {
		t.Log("TestAprobacionPagosContratistas Finalizado Correctamente (OK)")
	}
}

// SolicitudesOrdenadorContratistas ...
func TestSolicitudesOrdenadorContratistas(t *testing.T) {
	t.Log("-----------------------------------------------------")
	valor, err := helpers.SolicitudesOrdenadorContratistas("52204982", 5, 0)
	if err != nil {
		t.Error("No se pudo consultar las solicitudes de contratistas correspondientes a un ordenador por: ", err)
		t.Fail()
	} else {
		t.Log(valor)
		t.Log("TestGetActasRecibidoTipo Finalizado Correctamente (OK)")
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
