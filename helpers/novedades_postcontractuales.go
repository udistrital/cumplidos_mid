package helpers

import (
	// "fmt"
	// "strconv"

	"strconv"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/models"
)

func GetNovedadesPostcontractuales(query string, target *[]models.NovedadPoscontractual) (status int, err_nov error) {
	var responseWrapper models.RespNov
	url := beego.AppConfig.String("UrlNovedadesMid") + "/novedad/" + query
	if response, err := getJsonTest(url, &responseWrapper); (err == nil) && (response == 200) {
		*target = responseWrapper.Body
		return 200, nil
	} else {
		err_nov = err
	}
	return 400, err_nov
}

func ConstruirNovedadOtroSi(nov models.NovedadPoscontractual) (novedaux models.Noveda, outputError map[string]interface{}) {
	novedaux.NumeroCdp = nov.NumeroCdp
	novedaux.VigenciaCdp = nov.VigenciaCdp
	novedaux.TipoNovedad = "NP_ADPRO"
	return novedaux, outputError
}

func ConstruirNovedadCesion(nov models.NovedadPoscontractual) (novedaux models.Noveda, outputError map[string]interface{}) {
	novedaux.TipoNovedad = "NP_CES"
	return novedaux, outputError
}

func ConstruirNovedadSuspension(nov models.NovedadPoscontractual) (suspension models.Noveda, outputError map[string]interface{}) {

	if prov, err := ConsultarProveedorNovedad(nov.Cesionario); err == nil {

		suspension.Cesionario = prov

		suspension.NumeroContrato = nov.Contrato
		suspension.Vigencia = nov.Vigencia
		suspension.TipoNovedad = nov.CodAbreviacionTipo
		suspension.FechaCreacion = nov.FechaRegistro
		suspension.FechaInicio = nov.FechaSuspension
		suspension.FechaFin = nov.FechaTerminacionanticipada
		suspension.FechaFinSus = nov.FechaFinSuspension
		suspension.PlazoEjecucion = nov.PeriodoSuspension
		suspension.NumeroCdp = nov.NumeroCdp

		return suspension, nil
	} else {
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/ConstruirNovedadSuspension", "err": err, "status": "502"}
		return suspension, outputError
	}
}

func ConstruirNovedadTerminacion(nov models.NovedadPoscontractual) (novedaux models.Noveda, outputError map[string]interface{}) {

	novedaux.TipoNovedad = "NP_TER"
	return novedaux, outputError
}

func ConsultarProveedorNovedad(id int) (docProveedor string, outputError map[string]interface{}) {

	var proveedor models.InformacionProveedor
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor/"+strconv.Itoa(id), &proveedor); (err == nil) && (response == 200) {

		return docProveedor, nil
	} else {
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/ConsultarProveedorNovedad", "err": err, "status": "502"}
		return "", outputError
	}
}
