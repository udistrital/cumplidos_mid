package dependencias

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/helpers"
)

var parameters struct {
	UrlCrudCumplidos                          string
	UrlcrudAgora                              string
	UrlFinancieraJBPM                         string
	UrlAdministrativaJBPM                     string
	UrlcrudCore                               string
	UrlcrudOikos                              string
	UrlNovedadesMid                           string
	UrlTitanMid                               string
	UrlDocumentosCrud                         string
	UrlGestorDocumental                       string
	UrlAdministrativaJBPMContratosDependencia string
}

func TestMain(m *testing.M) {

	err := os.Setenv("UrlCrudCumplidos", "http://pruebasapi.intranetoas.udistrital.edu.co:8511/v1")
	err = os.Setenv("UrlAdministrativaJBPMContratosDependencia", "http://busservicios.intranetoas.udistrital.edu.co:8282/administrativa")
	err = os.Setenv("UrlAdministrativaJBPM", "http://busservicios.intranetoas.udistrital.edu.co:8282/wso2eiserver/services/administrativa_pruebas")
	err = os.Setenv("UrlcrudAgora", "http://pruebasapi.intranetoas.udistrital.edu.co:8104/v1")
	if err != nil {
		fmt.Println("Error estableciendo la variable de entorno:", err)
		return
	}

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
	parameters.UrlAdministrativaJBPMContratosDependencia = os.Getenv("UrlAdministrativaJBPMContratosDependencia")

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
	beego.AppConfig.Set("UrlAdministrativaJBPMContratosDependencia", parameters.UrlAdministrativaJBPMContratosDependencia)

	flag.Parse()
	os.Exit(m.Run())
}

func TestGetDependenciasSupervisor(t *testing.T) {
	t.Log("-----------------------------------------------------")
	dependenciasSupervisor, outputError := helpers.GetDependenciasSupervisor("80093200")
	if outputError != nil {
		t.Error("Error en la función GetDependenciasSupervisor")
		t.Fail()
	} else {
		t.Log(dependenciasSupervisor)
		t.Log("TestGetDependenciasSupervisor Finalizado Correctamente (OK)")
	}
	t.Log("-----------------------------------------------------")
}

func TestGetDependenciasOrdenador(t *testing.T) {
	t.Log("-----------------------------------------------------")
	dependenciasOrdenador, outputError := helpers.GetDependenciasOrdenador("19483708")
	if outputError != nil {
		t.Error("Error en la función GetDependenciasOrdenador")
		t.Fail()
	} else {
		t.Log(dependenciasOrdenador)
		t.Log("TestGetDependenciasOrdenador Finalizado Correctamente (OK)")
	}
	t.Log("-----------------------------------------------------")
}
