package filtros_pagos_mensuales

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

// TestGetPagosFiltrados ...
func TestGetPagosFiltrados(t *testing.T) {
	t.Log("-----------------------------------------------------")

	// Definir los parámetros de prueba
	numerosContratos := []string{"619", "48"}
	numerosDocumentos := []string{"80145797", "1022926896"}
	anios := []string{"2018"}
	meses := []string{"5", "4"}
	estadosPagos := []string{"12"}

	// Llamar a la función GetPagosFiltrados
	pagosMensuales, outputError := helpers.GetPagosFiltrados(numerosContratos, numerosDocumentos, anios, meses, estadosPagos)

	// Verificar si hubo un error
	if outputError != nil {
		t.Error("Error en la función GetPagosFiltrados:", outputError)
		t.Fail()
	} else {
		// Imprimir los pagos mensuales obtenidos
		for _, pago := range pagosMensuales {
			t.Logf("PagoMensual: %+v\n", pago)
		}
		t.Log("TestGetPagosFiltrados Finalizado Correctamente (OK)")
	}

	t.Log("-----------------------------------------------------")
}
