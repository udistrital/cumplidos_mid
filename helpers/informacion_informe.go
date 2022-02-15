package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

func InformacionInforme(num_documento string, contrato string, vigencia string, cdp string, vigencia_cdp string) (informacion_informe models.InformacionInforme, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": "/InformacionInforme", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var res_informacion_proveedor map[string]interface{}
	var informacion_contrato models.InformacionContrato
	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlAdministrativaJBPM")+"/"+"informacion_contrato/"+contrato+"/"+vigencia, &res_informacion_proveedor); (err == nil) && (response == 200) {
		fmt.Println("informacion_contrato:", res_informacion_proveedor)
		b, err := json.Marshal(res_informacion_proveedor)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(b, &informacion_contrato)
		informacion_informe.ValorContrato = informacion_contrato.Contrato.ValorContrato
		informacion_informe.Objeto = informacion_contrato.Contrato.ObjetoContrato
		informacion_informe.ActividadesEspecificas = informacion_contrato.Contrato.Actividades
		informacion_informe.FechaCPS = informacion_contrato.Contrato.FechaSuscripcion
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/informacion_contrato", "err": err, "status": "502"}
		panic(outputError)
	}

	var informacion_contrato_contratista models.InformacionContratoContratista
	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlAdministrativaJBPM")+"/"+"informacion_contrato_contratista/"+contrato+"/"+vigencia, &res_informacion_proveedor); (err == nil) && (response == 200) {
		fmt.Println("informacion_contrato_contratista:", res_informacion_proveedor)
		b, err := json.Marshal(res_informacion_proveedor)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(b, &informacion_contrato_contratista)
		informacion_informe.InformacionContratista.CiudadExpedicion = informacion_contrato_contratista.InformacionContratista.Documento.Ciudad
		informacion_informe.InformacionContratista.Nombre = informacion_contrato_contratista.InformacionContratista.NombreCompleto
		informacion_informe.InformacionContratista.TipoIdentificacion = informacion_contrato_contratista.InformacionContratista.Documento.Tipo
		informacion_informe.Dependencia = informacion_contrato_contratista.InformacionContratista.Dependencia

	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/informacion_contrato_contratista", "err": err, "status": "502"}
		panic(outputError)
	}

	var contrato_general []models.ContratoGeneral
	var sede []models.Sede
	var supervisor_contrato []models.SupervisorContrato
	fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/contrato_general/?query=ContratoSuscrito.NumeroContratoSuscrito:" + contrato + ",VigenciaContrato:" + vigencia)
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/contrato_general/?query=ContratoSuscrito.NumeroContratoSuscrito:"+contrato+",VigenciaContrato:"+vigencia, &contrato_general); (err == nil) && (response == 200) {
		fmt.Println("contrato_general:", contrato_general)
		fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/sedes_SIC/?query=ESFIDSEDE:" + contrato_general[0].LugarEjecucion.Sede)
		if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/sedes_SIC/?query=ESFIDSEDE:"+contrato_general[0].LugarEjecucion.Sede, &sede); (err == nil) && (response == 200) {
			fmt.Println("sede:", sede)
			informacion_informe.Sede = sede[0].ESFSEDE
		} else {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/InformacionInforme/contrato_general/sedes_SIC", "err": err, "status": "502"}
			panic(outputError)
		}

		if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/supervisor_contrato?query=DependenciaSupervisor:"+contrato_general[0].Supervisor.DependenciaSupervisor+"&sortby=FechaInicio&order=desc", &supervisor_contrato); (err == nil) && (response == 200) {
			fmt.Println("supervisor_contrato:", supervisor_contrato)
			informacion_informe.Supervisor.Cargo = supervisor_contrato[0].Cargo
			informacion_informe.Supervisor.Nombre = supervisor_contrato[0].Nombre
		} else {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/InformacionInforme/contrato_general/sedes_SIC", "err": err, "status": "502"}
			panic(outputError)
		}

	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/contrato_general", "err": err, "status": "502"}
		panic(outputError)
	}

	var temp map[string]interface{}
	var cdp_rp models.InformacionCdpRp
	fmt.Println(beego.AppConfig.String("UrlFinancieraJBPM") + "/" + "cdprp/" + cdp + "/" + vigencia_cdp + "/01")
	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlFinancieraJBPM")+"/"+"cdprp/"+cdp+"/"+vigencia_cdp+"/01", &temp); (err == nil) && (response == 200) {
		json_cdp_rp, error_json := json.Marshal(temp)

		if error_json == nil {
			if err := json.Unmarshal(json_cdp_rp, &cdp_rp); err == nil {

			} else {
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/UnmarshalCDP", "err": err, "status": "502"}
				panic(outputError)
			}
		} else {
			logs.Error(error_json)
			outputError = map[string]interface{}{"funcion": "/marshalCDP", "err": error_json, "status": "502"}
			panic(outputError)
		}
		fmt.Println("cdp_rp:", cdp_rp)
		informacion_informe.CDP.Consecutivo = cdp_rp.CdpXRp.CdpRp[0].CdpNumeroDisponibilidad
		informacion_informe.CDP.Fecha = cdp_rp.CdpXRp.CdpRp[0].CdpFechaExpedicion
		informacion_informe.RP.Consecutivo = cdp_rp.CdpXRp.CdpRp[0].RpNumeroRegistro
		informacion_informe.RP.Fecha = cdp_rp.CdpXRp.CdpRp[0].RpFechaRegistro
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/cdprp", "err": err, "status": "502"}
		panic(outputError)
	}

	var temp_novedades models.RespNov
	fmt.Println(beego.AppConfig.String("UrlNovedadesMid") + "/novedad/" + contrato + "/" + vigencia)
	if response, err := getJsonTest(beego.AppConfig.String("UrlNovedadesMid")+"/novedad/"+contrato+"/"+vigencia, &temp_novedades); (err == nil) && (response == 200) {

		fmt.Println("temp_novedades", temp_novedades)
		novedades := temp_novedades.Body
		if novedades_cesion, err := getNovedadCesion(novedades); err == nil {
			informacion_informe.Novedades.Cesion = novedades_cesion

		}
		if novedades_otrosi, err := getNovedadOtroSi(novedades); err == nil {
			informacion_informe.Novedades.Otrosi = novedades_otrosi
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/novedades", "err": err, "status": "502"}
		panic(outputError)
	}

	return
}

func getNovedadOtroSi(novedades []models.Novedad) (novedad_otrosi []models.Otrosi, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/getNovedadOtroSi", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	for _, novedad := range novedades {
		var otrosi models.Otrosi
		if novedad.TipoNovedad == 7 || novedad.TipoNovedad == 8 {
			otrosi.FechaAdiccion = novedad.FechaAdiccion
			otrosi.FechaProrroga = novedad.FechaProrroga
			otrosi.TiempoProrroga = novedad.TiempoProrroga
			fmt.Println("otrosi: ", otrosi)
			novedad_otrosi = append(novedad_otrosi, otrosi)
			fmt.Println("lista otrosi: ", novedad_otrosi)
		}
	}

	return
}

func getNovedadCesion(novedades []models.Novedad) (novedad_cesion []models.Cesion, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/getNovedadCesion", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	for _, novedad := range novedades {
		var cesion models.Cesion
		if novedad.TipoNovedad == 2 {
			cesion.FechaCesion = novedad.FechaCesion
			fmt.Println("cesion:", cesion)
			novedad_cesion = append(novedad_cesion, cesion)
			fmt.Println("lista cesion: ", novedad_cesion)
		}
	}

	return
}
