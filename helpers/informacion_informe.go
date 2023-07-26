package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

func InformacionInforme(pago_mensual_id string) (informacion_informe models.InformacionInforme, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": outputError["funcion"], "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var contrato string
	var vigencia string
	var cdp string
	var vigencia_cdp string
	var num_documento string
	if pago_mensual, err := getPagoMensual(pago_mensual_id); err == nil {

		contrato = pago_mensual.NumeroContrato
		vigencia = strconv.Itoa(int(pago_mensual.VigenciaContrato))
		num_documento = pago_mensual.DocumentoPersonaId
		cdp = pago_mensual.NumeroCDP
		vigencia_cdp = strconv.Itoa(int(pago_mensual.VigenciaCDP))
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme", "err": err, "status": "502"}
		panic(outputError)
	}

	var res_informacion_proveedor map[string]interface{}
	var informacion_contrato models.InformacionContrato
	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlAdministrativaJBPM")+"/"+"informacion_contrato/"+contrato+"/"+vigencia, &res_informacion_proveedor); (err == nil) && (response == 200) {
		//fmt.Println("informacion_contrato:", res_informacion_proveedor)
		b, err := json.Marshal(res_informacion_proveedor)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(b, &informacion_contrato)
		informacion_informe.ValorContrato, err = strconv.Atoi(strings.Split(informacion_contrato.Contrato.ValorContrato, ".")[0])
		if err != nil {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/InformacionInforme", "err": err, "status": "502"}
			panic(outputError)
		}
		informacion_informe.ValorTotalContrato = informacion_informe.ValorContrato
		informacion_informe.Objeto = informacion_contrato.Contrato.ObjetoContrato
		informacion_informe.ActividadesEspecificas = informacion_contrato.Contrato.Actividades
		informacion_informe.FechaCPS = informacion_contrato.Contrato.FechaSuscripcion
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/informacion_contrato", "err": err, "status": "502"}
		panic(outputError)
	}
	//fmt.Println(informacion_informe)

	var informacion_persona_natural []models.InformacionPersonaNatural
	fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/informacion_persona_natural?query=Id:" + num_documento)
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_persona_natural?query=Id:"+num_documento, &informacion_persona_natural); (err == nil) && (response == 200) {
		//fmt.Println("informacion_persona natural:", informacion_persona_natural)
		//informacion_informe.InformacionContratista.CiudadExpedicion = informacion_contrato_contratista.InformacionContratista.Documento.Ciudad
		informacion_informe.InformacionContratista.Nombre = informacion_persona_natural[0].PrimerNombre + " " + informacion_persona_natural[0].SegundoNombre + " " + informacion_persona_natural[0].PrimerApellido + " " + informacion_persona_natural[0].SegundoApellido
		informacion_informe.InformacionContratista.TipoIdentificacion = informacion_persona_natural[0].TipoDocumento.ValorParametro

		fmt.Println(beego.AppConfig.String("UrlcrudCore") + "/ciudad/" + strconv.Itoa(informacion_persona_natural[0].IdCiudadExpedicionDocumento))
		var ciudad models.Ciudad
		if response, err := getJsonTest(beego.AppConfig.String("UrlcrudCore")+"/ciudad/"+strconv.Itoa(informacion_persona_natural[0].IdCiudadExpedicionDocumento), &ciudad); (err == nil) && (response == 200) {
			//fmt.Println("ciudad:", ciudad)
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
		//fmt.Println("informacion_contrato_contratista:", res_informacion_proveedor)
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
		//fmt.Println("contrato_general:", contrato_general)
		fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/sedes_SIC/?query=ESFIDSEDE:" + contrato_general[0].LugarEjecucion.Sede)
		if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/sedes_SIC/?query=ESFIDSEDE:"+contrato_general[0].LugarEjecucion.Sede, &sede); (err == nil) && (response == 200) {
			//fmt.Println("sede:", sede)
			informacion_informe.Sede = sede[0].ESFSEDE
		} else {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/InformacionInforme/contrato_general/sedes_SIC", "err": err, "status": "502"}
			panic(outputError)
		}

		if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/supervisor_contrato?query=DependenciaSupervisor:"+contrato_general[0].Supervisor.DependenciaSupervisor+"&sortby=FechaInicio&order=desc", &supervisor_contrato); (err == nil) && (response == 200) {
			//fmt.Println("supervisor_contrato:", supervisor_contrato)
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

	//Fechas acta de inicio
	if acta_inicio, err := GetActaDeInicio(contrato_general[0].Id, contrato_general[0].VigenciaContrato); err == nil {
		informacion_informe.FechaInicio = acta_inicio.FechaInicio
		informacion_informe.FechaFin = acta_inicio.FechaFin
	} else {
		outputError = map[string]interface{}{"funcion": "/Informacion_informe/Acta_inicio", "err": err, "status": "502"}
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
		//fmt.Println("cdp_rp:", cdp_rp)
		informacion_informe.CDP.Consecutivo = cdp_rp.CdpXRp.CdpRp[0].CdpNumeroDisponibilidad
		informacion_informe.CDP.Fecha = cdp_rp.CdpXRp.CdpRp[0].CdpFechaExpedicion
		informacion_informe.RP.Consecutivo = cdp_rp.CdpXRp.CdpRp[0].RpNumeroRegistro
		informacion_informe.RP.Fecha = cdp_rp.CdpXRp.CdpRp[0].RpFechaRegistro
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/cdprp", "err": err, "status": "502"}
		panic(outputError)
	}

	// Consulta novedades OtroSi
	var otrosi []models.Otrosi
	fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/novedad_postcontractual?query=TipoNovedad:220,NumeroContrato:" + contrato + ",Vigencia:" + vigencia + "&sortby=FechaInicio&order=desc")
	if response, err := GetNovedadesPostcontractuales(models.TipoNovedadOtrosi, "NumeroContrato:"+contrato+",Vigencia:"+vigencia, "FechaInicio", "desc", "-1", "", "", &otrosi); (err == nil) && (response == 200) {
		for i, os := range otrosi {
			if valor_girado_otrosi, err := getValorGiradoPorCdp(strconv.Itoa(os.NumeroCdp), strconv.Itoa(os.VigenciaCdp), strconv.Itoa(contrato_general[0].UnidadEjecutora)); err == nil {
				otrosi[i].ValorNovedadPagado = valor_girado_otrosi
				informacion_informe.EjecutadoDinero.Faltante = informacion_informe.EjecutadoDinero.Faltante + (otrosi[i].ValorNovedad - otrosi[i].ValorNovedadPagado)
				informacion_informe.ValorTotalContrato = informacion_informe.ValorTotalContrato + otrosi[i].ValorNovedad
			} else {
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/InformacionInforme/Novedades/OtroSi", "err": err, "status": "502"}
				panic(outputError)
			}
		}
		informacion_informe.Novedades.Otrosi = otrosi
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/Novedades/OtroSi", "err": err, "status": "502"}
		panic(outputError)
	}

	var cesion []models.Cesion
	fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/novedad_postcontractual?query=TipoNovedad:219,NumeroContrato:" + contrato + ",Vigencia:" + vigencia + "&sortby=FechaInicio&order=desc")
	if response, err := GetNovedadesPostcontractuales(models.TipoNovedadCesion, "NumeroContrato:"+contrato+",Vigencia:"+vigencia, "FechaInicio", "desc", "-1", "", "", &cesion); (err == nil) && (response == 200) {
		//fmt.Println("Cesion:", cesion)
		informacion_informe.Novedades.Cesion = cesion
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/Novedades/Cesion", "err": err, "status": "502"}
		panic(outputError)
	}

	var terminacion []models.Terminacion
	fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/novedad_postcontractual?query=TipoNovedad:218,NumeroContrato:" + contrato + ",Vigencia:" + vigencia + "&sortby=FechaInicio&order=desc")
	if response, err := GetNovedadesPostcontractuales(models.TipoNovedadTerminacion, "NumeroContrato:"+contrato+",Vigencia:"+vigencia, "FechaInicio", "desc", "-1", "", "", &terminacion); (err == nil) && (response == 200) {
		//fmt.Println("terminacion:", terminacion)
		informacion_informe.Novedades.Terminacion = terminacion
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/Novedades/Terminacion", "err": err, "status": "502"}
		panic(outputError)
	}

	var suspencion []models.Suspencion
	fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/novedad_postcontractual?query=TipoNovedad:216,NumeroContrato:" + contrato + ",Vigencia:" + vigencia + "&sortby=FechaInicio&order=desc")
	if response, err := GetNovedadesPostcontractuales(models.TipoNovedadSuspension, "NumeroContrato:"+contrato+",Vigencia:"+vigencia, "FechaInicio", "desc", "-1", "", "", &suspencion); (err == nil) && (response == 200) {
		//fmt.Println("Suspencion:", suspencion)
		informacion_informe.Novedades.Suspencion = suspencion
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/Novedades/Suspencion", "err": err, "status": "502"}
		panic(outputError)
	}

	//Girado contrato original
	var contratos_disponibilidad []models.ContratoDisponibilidad
	numero_contrato := contrato_general[0].Id

	fmt.Println(beego.AppConfig.String("UrlcrudAgora") + "/contrato_disponibilidad/?query=NumeroContrato:" + numero_contrato)
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+numero_contrato, &contratos_disponibilidad); (err == nil) && (response == 200) {

		if len(contratos_disponibilidad) == 0 {
			err = errors.New("No se encontro cdp asociado al contrato")
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/Informacion_informe", "err": err, "status": "502"}
			return informacion_informe, outputError
		}
		contrato_disponibilidad := contratos_disponibilidad[0]
		if valor_girado, err := getValorGiradoPorCdp(strconv.Itoa(contrato_disponibilidad.NumeroCdp), strconv.Itoa(contrato_disponibilidad.VigenciaCdp), strconv.Itoa(contrato_general[0].UnidadEjecutora)); err == nil {
			//fmt.Println(informacion_informe.EjecutadoDinero.Pagado)
			informacion_informe.EjecutadoDinero.Pagado = informacion_informe.EjecutadoDinero.Pagado + valor_girado
			informacion_informe.EjecutadoDinero.Faltante = informacion_informe.EjecutadoDinero.Faltante + (informacion_informe.ValorContrato - valor_girado)
		} else {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/Informacion_informe", "err": err.Error(), "status": "502"}
			return informacion_informe, outputError
		}

	} else { // If contrato_disponibilidad get
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/Informacion_informe", "err": err.Error(), "status": "502"}
		return informacion_informe, outputError
	}

	//fechasConNovedades
	if vig, err := strconv.Atoi(vigencia); err == nil {
		if fechasConNov, err := FechasContratoConNovedades(contrato, vig, contrato_general[0].Id); err == nil {
			informacion_informe.FechasConNovedades = fechasConNov
		}
	}
	return
}

func getPagoMensual(pago_mensual_id string) (pago_mensual models.PagoMensual, err error) {

	var pagos_mensuales []models.PagoMensual
	var respuesta_peticion map[string]interface{}

	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudCumplidos")+"/pago_mensual/?query=Id:"+pago_mensual_id, &respuesta_peticion); (err == nil) && (response == 200) {

		pagos_mensuales = []models.PagoMensual{}
		LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
		if len(pagos_mensuales) == 0 {
			err = errors.New("No se encontro pago mensual asociado al id")
			return pago_mensual, err
		}
		return pagos_mensuales[0], err

	} else {
		logs.Error(err)
		err = errors.New("Error en la peticion")
		return pago_mensual, err
	}
	return
}

func GetPreliquidacion(pago_mensual_id string) (preliquidacion models.PreliquidacionTitan, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": "/InformacionInforme", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var anio string
	var mes string
	var contrato string
	var vigencia_contrato string
	var documento_contratista string
	var numero_cdp int

	if pago_mensual, err := getPagoMensual(pago_mensual_id); err == nil {
		anio = strconv.Itoa(int(pago_mensual.Ano))
		mes = strconv.Itoa(int(pago_mensual.Mes))
		contrato = pago_mensual.NumeroContrato
		vigencia_contrato = strconv.Itoa(int(pago_mensual.VigenciaContrato))
		documento_contratista = pago_mensual.DocumentoPersonaId
		numero_cdp, err = strconv.Atoi(pago_mensual.NumeroCDP)
		if err != nil {
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/InformacionInforme/Preliquidacion/", "err": err, "status": "502"}
			panic(outputError)
		}

	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/Preliquidacion", "err": err, "status": "502"}
		panic(outputError)
	}

	var preliquidaciones []models.PreliquidacionTitan
	var respuesta_peticion_prel map[string]interface{}
	fmt.Println(beego.AppConfig.String("UrlTitanMid") + "/detalle_preliquidacion/obtener_detalle_CT/" + anio + "/" + mes + "/" + contrato + "/" + vigencia_contrato + "/" + documento_contratista)
	if response, err := getJsonTest(beego.AppConfig.String("UrlTitanMid")+"/detalle_preliquidacion/obtener_detalle_CT/"+anio+"/"+mes+"/"+contrato+"/"+vigencia_contrato+"/"+documento_contratista, &respuesta_peticion_prel); (err == nil) && (response == 201) {
		LimpiezaRespuestaRefactor(respuesta_peticion_prel, &preliquidaciones)
		if preliquidacion, err := seleccionarPreliquidacion(preliquidaciones, numero_cdp); err == nil {
			//fmt.Println(preliquidacion)
			return darFormatoPreliquidacion(preliquidacion), nil
		} else {
			//fmt.Println(err)
			outputError = map[string]interface{}{"funcion": "/InformacionInforme/Preliquidacion/SeleccionPreliquidacion:" + err.Error(), "err": err, "status": "400"}
			return preliquidacion, outputError
		}

	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/InformacionInforme/Preliquidacion", "err": err, "status": "502"}
		panic(outputError)
	}

	return
}

func seleccionarPreliquidacion(preliquidaciones []models.PreliquidacionTitan, cdp int) (preliquidacion models.PreliquidacionTitan, err error) {
	if len(preliquidaciones) == 0 {
		err = errors.New("No se encontraron preliquidaciones")
		return preliquidacion, err
	}
	if len(preliquidaciones) == 1 {
		return preliquidaciones[0], nil
	}

	for _, prel := range preliquidaciones {
		if prel.Detalle[0].ContratoPreliquidacionId.ContratoId.Cdp == cdp {
			return prel, nil
		}
	}

	err = errors.New("No se encuentra preliquidacion asociada al CDP")
	return preliquidacion, err
}

func darFormatoPreliquidacion(preliquidacion models.PreliquidacionTitan) (preliquidacionConFormato models.PreliquidacionTitan) {
	preliquidacionConFormato = preliquidacion
	preliquidacionConFormato.TotalDevengadoConFormato = FormatMoneyString(formatNumberString(strconv.Itoa(int(preliquidacion.TotalDevengado)), 0, ",", "."), 0)
	preliquidacionConFormato.TotalDescuentosConFormato = FormatMoneyString(formatNumberString(strconv.Itoa(int(preliquidacion.TotalDescuentos)), 0, ",", "."), 0)
	preliquidacionConFormato.TotalPagoConFormato = FormatMoneyString(formatNumberString(strconv.Itoa(int(preliquidacion.TotalPago)), 0, ",", "."), 0)
	for i, _ := range preliquidacionConFormato.Detalle {
		preliquidacionConFormato.Detalle[i].ValorCalculadoConFormato = FormatMoneyString(formatNumberString(strconv.Itoa(int(preliquidacion.Detalle[i].ValorCalculado)), 0, ",", "."), 0)
	}
	return
}

func getValorGiradoPorCdp(cdp string, vigencia_cdp string, unidad_ejecucion string) (valor_girado int, err error) {
	var temp_giros_tercero map[string]interface{}
	var giros_tercero models.GirosTercero
	valor_girado = 0
	fmt.Println(beego.AppConfig.String("UrlFinancieraJBPM") + "/" + "giros_tercero/" + cdp + "/" + vigencia_cdp + "/" + unidad_ejecucion)
	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlFinancieraJBPM")+"/"+"giros_tercero/"+cdp+"/"+vigencia_cdp+"/"+unidad_ejecucion, &temp_giros_tercero); (err == nil) && (response == 200) {
		json_giros_tercero, error_json := json.Marshal(temp_giros_tercero)
		if error_json == nil {
			if err := json.Unmarshal(json_giros_tercero, &giros_tercero); err == nil {
				//fmt.Println("giros "+cdp, giros_tercero)
				for _, giro := range giros_tercero.Giros.Tercero {
					total_girado, err := strconv.Atoi(giro.ValorBrutoGirado)
					//fmt.Println(total_girado)
					if err == nil {
						valor_girado = valor_girado + total_girado
					}
				}
				//fmt.Println(valor_girado)
				return valor_girado, nil

			} else {
				err = errors.New("Error Unmarshal giros_tercero")
				return valor_girado, err
			}

		} else {
			err = errors.New("Error Marshal giros_tercero")
			return valor_girado, err
		}

	} else {
		return valor_girado, err
	}
	return
}
