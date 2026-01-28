package helpers

import (
	"strconv"

	"github.com/astaxie/beego"
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

func ConstruirNovedadOtroSi(nov models.NovedadPoscontractual) (otrosi models.NovedadCesion, outputError map[string]interface{}) {
	if docProv, nomProv, err := ConsultarProveedorNovedad(nov.Cesionario); err == nil {

		otrosi.DocCesionario = docProv
		otrosi.NombreCesionario = nomProv
		otrosi.NumeroContrato = nov.Contrato
		otrosi.Vigencia = nov.Vigencia
		otrosi.TipoNovedad = nov.CodAbreviacionTipo
		otrosi.FechaCreacion = FormatoFechaNovedad(nov.FechaRegistro)
		otrosi.FechaInicio = FormatoFechaNovedad(nov.FechaAdicion)
		otrosi.FechaFin = FormatoFechaNovedad(nov.FechaFinefectiva)
		otrosi.PlazoEjecucion = nov.TiempoProrroga
		otrosi.ValorAdicion = nov.ValorAdicion
		otrosi.NumeroCdp = nov.NumeroCdp
		otrosi.VigenciaCdp = nov.VigenciaCdp

		return otrosi, nil
	} else {
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/ConstruirNovedadSuspension", "err": err, "status": "502"}
		return otrosi, outputError
	}
}

func ConstruirNovedadCesion(nov models.NovedadPoscontractual) (cesion models.NovedadCesion, outputError map[string]interface{}) {
	cesion.TipoNovedad = nov.CodAbreviacionTipo

	if docCedente, nomCedente, err := ConsultarProveedorNovedad(nov.Cedente); err == nil {
		if docCesionario, nomCesionario, err := ConsultarProveedorNovedad(nov.Cesionario); err == nil {
			cesion.DocCedente = docCedente
			cesion.NombreCedente = nomCedente
			cesion.DocCesionario = docCesionario
			cesion.NombreCesionario = nomCesionario
			cesion.NumeroContrato = nov.Contrato
			cesion.Vigencia = nov.Vigencia
			cesion.FechaCreacion = FormatoFechaNovedad(nov.FechaRegistro)
			cesion.FechaInicio = FormatoFechaNovedad(nov.FechaCesion)
			cesion.FechaFin = FormatoFechaNovedad(nov.FechaFinefectiva)
			cesion.TipoNovedad = nov.CodAbreviacionTipo

			return cesion, nil
		} else {
			outputError = map[string]interface{}{"funcion": "/InformacionInforme/ConstruirNovedadCesion", "err": err, "status": "502"}
			return cesion, outputError
		}
	} else {
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/ConstruirNovedadCesion", "err": err, "status": "502"}
		return cesion, outputError
	}
}

func ConstruirNovedadSuspension(nov models.NovedadPoscontractual) (suspension models.NovedadCesion, outputError map[string]interface{}) {

	if docProv, nomProv, err := ConsultarProveedorNovedad(nov.Cesionario); err == nil {

		suspension.DocCesionario = docProv
		suspension.NombreCesionario = nomProv
		suspension.NumeroContrato = nov.Contrato
		suspension.Vigencia = nov.Vigencia
		suspension.TipoNovedad = nov.CodAbreviacionTipo
		suspension.FechaCreacion = FormatoFechaNovedad(nov.FechaRegistro)
		suspension.FechaInicio = FormatoFechaNovedad(nov.FechaSuspension)
		suspension.FechaFin = FormatoFechaNovedad(nov.FechaFinefectiva)
		suspension.FechaFinSus = FormatoFechaNovedad(nov.FechaFinSuspension)
		suspension.PlazoEjecucion = nov.PeriodoSuspension
		suspension.NumeroCdp = nov.NumeroCdp

		return suspension, nil
	} else {
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/ConstruirNovedadSuspension", "err": err, "status": "502"}
		return suspension, outputError
	}
}

func ConstruirNovedadTerminacion(nov models.NovedadPoscontractual) (terminacion models.NovedadCesion, outputError map[string]interface{}) {

	if docProv, nomProv, err := ConsultarProveedorNovedad(nov.Cesionario); err == nil {

		terminacion.DocCesionario = docProv
		terminacion.NombreCesionario = nomProv
		terminacion.NumeroContrato = nov.Contrato
		terminacion.Vigencia = nov.Vigencia
		terminacion.TipoNovedad = nov.CodAbreviacionTipo
		terminacion.FechaCreacion = FormatoFechaNovedad(nov.FechaRegistro)
		terminacion.FechaInicio = FormatoFechaNovedad(nov.FechaTerminacionanticipada)
		terminacion.FechaFin = FormatoFechaNovedad(nov.FechaFinefectiva)

		return terminacion, nil
	} else {
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/ConstruirNovedadTerminacion", "err": err, "status": "502"}
		return terminacion, outputError
	}
}

func ConsultarProveedorNovedad(id int) (string, string, map[string]interface{}) {

	var proveedor models.InformacionProveedor
	var docProveedor string
	var nomProveedor string
	var outputError map[string]interface{}

	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor/"+strconv.Itoa(id), &proveedor); (err == nil) && (response == 200) {
		docProveedor = proveedor.NumDocumento
		nomProveedor = proveedor.NomProveedor
		outputError = nil
	} else {
		docProveedor = ""
		nomProveedor = ""
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/ConsultarProveedorNovedad", "err": err, "status": "502"}
	}
	return docProveedor, nomProveedor, outputError
}

func FormatoFechaNovedad(fecha string) string {
	if fecha != "" {
		return fecha[0:10]
	}
	return ""
}
