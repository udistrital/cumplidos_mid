package descarga_pagos_mensuales

import (
	"flag"
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

func TestDescargarDocumentosSolicitudesPagos(t *testing.T) {
	t.Log("-----------------------------------------------------")
	documentos, outputError := helpers.DescargarDocumentosSolicitudesPagos("8207")
	if outputError != nil {
		t.Error("Error en la función DescargarDocumentosSolicitudesPagos")
		t.Fail()
	} else {
		t.Log(documentos)
		t.Log("TestDescargarDocumentosSolicitudesPagos Finalizado Correctamente (OK)")
	}
	t.Log("-----------------------------------------------------")
}
