package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"

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

	var informacion_persona_natural []models.InformacionPersonaNatural
	fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/informacion_persona_natural?query=Id:" + num_documento)
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_persona_natural?query=Id:"+num_documento, &informacion_persona_natural); (err == nil) && (response == 200) {
		fmt.Println("informacion_persona natural:", informacion_persona_natural)
		//informacion_informe.InformacionContratista.CiudadExpedicion = informacion_contrato_contratista.InformacionContratista.Documento.Ciudad
		informacion_informe.InformacionContratista.Nombre = informacion_persona_natural[0].PrimerNombre + " " + informacion_persona_natural[0].SegundoNombre + " " + informacion_persona_natural[0].PrimerApellido + " " + informacion_persona_natural[0].SegundoApellido
		informacion_informe.InformacionContratista.TipoIdentificacion = informacion_persona_natural[0].TipoDocumento.ValorParametro

		fmt.Println(beego.AppConfig.String("UrlcrudCore") + "/ciudad/" + strconv.Itoa(informacion_persona_natural[0].IdCiudadExpedicionDocumento))
		var ciudad models.Ciudad
		if response, err := getJsonTest(beego.AppConfig.String("UrlcrudCore")+"/ciudad/"+strconv.Itoa(informacion_persona_natural[0].IdCiudadExpedicionDocumento), &ciudad); (err == nil) && (response == 200) {
			fmt.Println("ciudad:", ciudad)
			informacion_informe.InformacionContratista.CiudadExpedicion = ciudad.Nombre
		} else {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/InformacionInforme/ciudad", "err": err, "status": "502"}
			panic(outputError)
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/Informacion_persona_natural", "err": err, "status": "502"}
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

	// var temp_novedades models.RespNov
	// fmt.Println(beego.AppConfig.String("UrlNovedadesMid") + "/novedad/" + contrato + "/" + vigencia)
	// if response, err := getJsonTest(beego.AppConfig.String("UrlNovedadesMid")+"/novedad/"+contrato+"/"+vigencia, &temp_novedades); (err == nil) && (response == 200) {

	// 	fmt.Println("temp_novedades", temp_novedades)
	// 	novedades := temp_novedades.Body
	// 	if novedades_cesion, err := getNovedadCesion(novedades); err == nil {
	// 		informacion_informe.Novedades.Cesion = novedades_cesion

	// 	}
	// 	if novedades_otrosi, err := getNovedadOtroSi(novedades); err == nil {
	// 		informacion_informe.Novedades.Otrosi = novedades_otrosi
	// 	}
	// } else {
	// 	logs.Error(err)
	// 	outputError = map[string]interface{}{"funcion": "/InformacionInforme/novedades", "err": err, "status": "502"}
	// 	panic(outputError)
	// }

	// Consulta novedades OtroSi
	var otrosi []models.Otrosi
	fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/novedad_postcontractual?query=TipoNovedad:220,NumeroContrato:" + contrato + ",Vigencia:" + vigencia + "&sortby=FechaInicio&order=desc")
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/novedad_postcontractual?query=TipoNovedad:220,NumeroContrato:"+contrato+",Vigencia:"+vigencia+"&sortby=FechaInicio&order=desc", &otrosi); (err == nil) && (response == 200) {
		fmt.Println("Otro si:", otrosi)
		informacion_informe.Novedades.Otrosi = otrosi
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/Novedades/OtroSi", "err": err, "status": "502"}
		panic(outputError)
	}

	var cesion []models.Cesion
	fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/novedad_postcontractual?query=TipoNovedad:219,NumeroContrato:" + contrato + ",Vigencia:" + vigencia + "&sortby=FechaInicio&order=desc")
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/novedad_postcontractual?query=TipoNovedad:219,NumeroContrato:"+contrato+",Vigencia:"+vigencia+"&sortby=FechaInicio&order=desc", &cesion); (err == nil) && (response == 200) {
		fmt.Println("Cesion:", cesion)
		informacion_informe.Novedades.Cesion = cesion
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/Novedades/OtroSi", "err": err, "status": "502"}
		panic(outputError)
	}

	var terminacion models.Terminacion
	fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/novedad_postcontractual?query=TipoNovedad:218,NumeroContrato:" + contrato + ",Vigencia:" + vigencia + "&sortby=FechaInicio&order=desc")
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/novedad_postcontractual?query=TipoNovedad:218,NumeroContrato:"+contrato+",Vigencia:"+vigencia+"&sortby=FechaInicio&order=desc", &terminacion); (err == nil) && (response == 200) {
		fmt.Println("terminacion:", cesion)
		informacion_informe.Novedades.Terminacion = terminacion
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/Novedades/Terminacion", "err": err, "status": "502"}
		panic(outputError)
	}

	var suspencion []models.Suspencion
	fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/novedad_postcontractual?query=TipoNovedad:216,NumeroContrato:" + contrato + ",Vigencia:" + vigencia + "&sortby=FechaInicio&order=desc")
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/novedad_postcontractual?query=TipoNovedad:216,NumeroContrato:"+contrato+",Vigencia:"+vigencia+"&sortby=FechaInicio&order=desc", &suspencion); (err == nil) && (response == 200) {
		fmt.Println("Suspencion:", cesion)
		informacion_informe.Novedades.Suspencion = suspencion
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/Novedades/Suspencion", "err": err, "status": "502"}
		panic(outputError)
	}
	return
}
