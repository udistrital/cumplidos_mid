package filtros_pagos_mensuales

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

// TestGetPagosFiltrados ...
func TestGetPagosFiltrados(t *testing.T) {
	t.Log("---------------------------------------------")

	numerosContratos := []string{"619", "48"}
	numerosDocumentos := []string{"80145797", "1022926896"}
	anios := []string{"2018"}
	meses := []string{"5", "4"}
	estadosPagos := []string{"12"}

	pagos, outputError := helpers.GetPagosFiltrados(
		numerosContratos,
		numerosDocumentos,
		anios,
		meses,
		estadosPagos,
	)

	if outputError != nil {
		t.Error("Error en la función GetPagosFiltrados:", outputError)
		t.Fail()
	} else {
		t.Log(pagos)
		t.Log("TestGetPagosFiltrados Finalizado Correctamente (OK)")
	}

	t.Log("---------------------------------------------")
}

func TestFiltrosDependencia(t *testing.T) {
	t.Log("-----------------------------------------------------")

	dependencias := []string{"'DEP12'", "'DEP626'"}
	vigencias := []string{"2017"}
	contratos, outputError := helpers.FiltrosDependencia(dependencias, vigencias)

	if outputError != nil {
		t.Error("Error en la función FiltrosDependencia:", outputError)
		t.Fail()
	} else {
		t.Log(contratos)
		t.Log("TestFiltrosDependencia Finalizado Correctamente (OK)")
	}

	t.Log("-----------------------------------------------------")
}

func TestSolicitudesPagoMensual(t *testing.T) {

	t.Log("-----------------------------------------------------")
	codigos_dependencias := []string{"'DEP12'", "'DEP626'"}
	vigencias := []string{"2017", "2018"}
	documentos_contratistas := []string{"1109843316", "1109843316"}
	numeros_contratos := []string{}
	meses := []string{"5", "6"}
	anios := []string{"2018"}
	estados := []string{}

	pagos, outputError := helpers.SolicitudesPagoMensual(codigos_dependencias, vigencias, documentos_contratistas, numeros_contratos, meses, anios, estados)

	if outputError != nil {
		t.Error("Error en la función SolicitudesPagoMensual:", outputError)
		t.Fail()
	} else {
		t.Log(pagos)
		t.Log("TestSolicitudesPagoMensual Finalizado Correctamente (OK)")
	}

	t.Log("-----------------------------------------------------")
}

func TestGetInformacionContrato(t *testing.T) {
	t.Log("-----------------------------------------------------")
	informacionContrato, outputError := helpers.GetInformacionContrato("1406", "2020")
	if outputError != nil {
		t.Error("Error en la función GetInformacionContrato")
		t.Fail()
	} else {
		t.Log(informacionContrato)
		t.Log("TestGetInformacionContrato Finalizado Correctamente (OK)")
	}

	t.Log("-----------------------------------------------------")
}
