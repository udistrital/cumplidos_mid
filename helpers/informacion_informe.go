package helpers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/models"
)

func InformacionInforme(pago_mensual_id string) (informacion_informe models.InformacionInforme, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": outputError["funcion"],
				"err":     err,
				"status":  "502",
			}
			panic(outputError)
		}
	}()

	var totalGiradoBase int
	var totalGiradoOtrosi int

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
		outputError = map[string]interface{}{
			"funcion": "/InformacionInforme",
			"err":     err,
			"status":  "502",
		}
		panic(outputError)
	}

	var resInfo map[string]interface{}
	var informacion_contrato models.InformacionContrato

	if response, err := getJsonWSO2Test(
		beego.AppConfig.String("UrlAdministrativaJBPM")+
			"/informacion_contrato/"+contrato+"/"+vigencia,
		&resInfo,
	); (err == nil) && (response == 200) {

		b, _ := json.Marshal(resInfo)
		json.Unmarshal(b, &informacion_contrato)

		val, _ := strconv.Atoi(strings.Split(informacion_contrato.Contrato.ValorContrato, ".")[0])
		informacion_informe.ValorContrato = val
		informacion_informe.ValorTotalContrato = val
		informacion_informe.Objeto = informacion_contrato.Contrato.ObjetoContrato
		informacion_informe.ActividadesEspecificas = informacion_contrato.Contrato.Actividades
		informacion_informe.FechaCPS = informacion_contrato.Contrato.FechaSuscripcion
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{
			"funcion": "/InformacionInforme/informacion_contrato",
			"err":     err,
			"status":  "502",
		}
		panic(outputError)
	}

	var informacion_persona []models.InformacionPersonaNatural
	if response, err := getJsonTest(
		beego.AppConfig.String("UrlcrudAgora")+
			"/informacion_persona_natural?query=Id:"+num_documento,
		&informacion_persona,
	); (err == nil) && (response == 200) {

		inf := informacion_persona[0]
		informacion_informe.InformacionContratista.Nombre =
			inf.PrimerNombre + " " + inf.SegundoNombre + " " + inf.PrimerApellido + " " + inf.SegundoApellido

		informacion_informe.InformacionContratista.TipoIdentificacion =
			inf.TipoDocumento.ValorParametro

		var ciudad models.Ciudad
		if response, err := getJsonTest(
			beego.AppConfig.String("UrlcrudCore")+"/ciudad/"+strconv.Itoa(inf.IdCiudadExpedicionDocumento),
			&ciudad,
		); (err == nil) && (response == 200) {
			informacion_informe.InformacionContratista.CiudadExpedicion = ciudad.Nombre
		} else {
			logs.Error(err)
			outputError = map[string]interface{}{
				"funcion": "/InformacionInforme/ciudad",
				"err":     err,
				"status":  "502",
			}
			panic(outputError)
		}

	} else {
		logs.Error(err)
		outputError = map[string]interface{}{
			"funcion": "/InformacionInforme/informacion_persona_natural",
			"err":     err,
			"status":  "502",
		}
		panic(outputError)
	}

	var infoContratoContr models.InformacionContratoContratista
	if response, err := getJsonWSO2Test(
		beego.AppConfig.String("UrlAdministrativaJBPM")+
			"/informacion_contrato_contratista/"+contrato+"/"+vigencia,
		&resInfo,
	); (err == nil) && (response == 200) {
		b, _ := json.Marshal(resInfo)
		json.Unmarshal(b, &infoContratoContr)
		informacion_informe.Dependencia =
			infoContratoContr.InformacionContratista.Dependencia
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{
			"funcion": "/InformacionInforme/informacion_contrato_contratista",
			"err":     err,
			"status":  "502",
		}
		panic(outputError)
	}

	var contrato_general []models.ContratoGeneral
	var sede []models.Sede
	var supervisor []models.SupervisorContrato

	if response, err := getJsonTest(
		beego.AppConfig.String("UrlcrudAgora")+
			"/contrato_general/?query=ContratoSuscrito.NumeroContratoSuscrito:"+contrato+
			",VigenciaContrato:"+vigencia,
		&contrato_general,
	); (err == nil) && (response == 200) {

		getJsonTest(
			beego.AppConfig.String("UrlcrudAgora")+
				"/sedes_SIC/?query=ESFIDSEDE:"+contrato_general[0].LugarEjecucion.Sede,
			&sede,
		)
		informacion_informe.Sede = sede[0].ESFSEDE

		getJsonTest(
			beego.AppConfig.String("UrlcrudAgora")+
				"/supervisor_contrato?query=DependenciaSupervisor:"+contrato_general[0].Supervisor.DependenciaSupervisor+
				"&sortby=FechaInicio&order=desc",
			&supervisor,
		)

		informacion_informe.Supervisor.Cargo = supervisor[0].Cargo
		informacion_informe.Supervisor.Nombre = supervisor[0].Nombre

	} else {
		logs.Error(err)
		outputError = map[string]interface{}{
			"funcion": "/InformacionInforme/contrato_general",
			"err":     err,
			"status":  "502",
		}
		panic(outputError)
	}

	if acta, err := GetActaDeInicio(contrato_general[0].Id, contrato_general[0].VigenciaContrato); err == nil {
		informacion_informe.FechaInicio = acta.FechaInicio
		informacion_informe.FechaFin = acta.FechaFin
	} else {
		outputError = map[string]interface{}{
			"funcion": "/InformacionInforme/Acta_inicio",
			"err":     err,
			"status":  "502",
		}
		panic(outputError)
	}

	var temp map[string]interface{}
	var cdpRp models.InformacionCdpRp

	if response, err := getJsonWSO2Test(
		beego.AppConfig.String("UrlFinancieraJBPM")+
			"/cdprp/"+cdp+"/"+vigencia_cdp+"/01",
		&temp,
	); (err == nil) && (response == 200) {
		b, _ := json.Marshal(temp)
		json.Unmarshal(b, &cdpRp)

		informacion_informe.CDP.Consecutivo = cdpRp.CdpXRp.CdpRp[0].CdpNumeroDisponibilidad
		informacion_informe.CDP.Fecha = cdpRp.CdpXRp.CdpRp[0].CdpFechaExpedicion
		informacion_informe.RP.Consecutivo = cdpRp.CdpXRp.CdpRp[0].RpNumeroRegistro
		informacion_informe.RP.Fecha = cdpRp.CdpXRp.CdpRp[0].RpFechaRegistro
	}

	var novedades []models.NovedadPoscontractual
	var novedadesStruct []models.Noveda
	query := contrato + "/" + vigencia

	if response, err := GetNovedadesPostcontractuales(query, &novedades); (err == nil) && (response == 200) {

		for _, nov := range novedades {
			switch nov.TipoNovedad {
			case 1:
				if v, e := ConstruirNovedadSuspension(nov); e == nil {
					novedadesStruct = append(novedadesStruct, v)
				}
			case 2:
				if v, e := ConstruirNovedadCesion(nov); e == nil {
					novedadesStruct = append(novedadesStruct, v)
				}
			case 5:
				if v, e := ConstruirNovedadTerminacion(nov); e == nil {
					novedadesStruct = append(novedadesStruct, v)
				}
			case 6, 7, 8:
				otrosi, err := ConstruirNovedadOtroSi(nov)
				if err != nil {
					logs.Error(err)
					continue
				}

				valPagado, errVal := getValorGiradoPorCdp(
					strconv.Itoa(otrosi.NumeroCdp),
					strconv.Itoa(otrosi.VigenciaCdp),
					strconv.Itoa(contrato_general[0].UnidadEjecutora),
				)
				if errVal == nil {
					otrosi.ValorNovedadPagado = valPagado
					totalGiradoOtrosi += valPagado
					informacion_informe.ValorTotalContrato += otrosi.ValorAdicion
				}

				novedadesStruct = append(novedadesStruct, otrosi)
			}
		}
	}
	informacion_informe.Novedades = novedadesStruct

	var contratos_disponibilidad []models.ContratoDisponibilidad
	numContrato := contrato_general[0].Id

	if response, err := getJsonTest(
		beego.AppConfig.String("UrlcrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+numContrato,
		&contratos_disponibilidad,
	); (err == nil) && (response == 200) {

		if len(contratos_disponibilidad) == 0 {
			err := errors.New("No se encontro CDP asociado")
			outputError = map[string]interface{}{
				"funcion": "/InformacionInforme",
				"err":     err,
				"status":  "502",
			}
			return informacion_informe, outputError
		}

		disp := contratos_disponibilidad[0]

		valBase, err := getValorGiradoPorCdp(
			strconv.Itoa(disp.NumeroCdp),
			strconv.Itoa(disp.VigenciaCdp),
			strconv.Itoa(contrato_general[0].UnidadEjecutora),
		)
		if err == nil {
			totalGiradoBase = valBase
		}
	}

	totalGirado := totalGiradoBase + totalGiradoOtrosi

	informacion_informe.EjecutadoDinero.Pagado = totalGirado
	informacion_informe.EjecutadoDinero.Faltante =
		informacion_informe.ValorTotalContrato - totalGirado

	if fechasConNov, err := FechasContratoConNovedades(contrato, vigencia, cdp, num_documento); err == nil {
		informacion_informe.FechasConNovedades = fechasConNov
	}

	return
}

func getPagoMensual(id string) (pago models.PagoMensual, err error) {
	var pagos []models.PagoMensual
	var resp map[string]interface{}

	if response, err := getJsonTest(
		beego.AppConfig.String("UrlcrudCumplidos")+
			"/pago_mensual/?query=Id:"+id,
		&resp,
	); (err == nil) && (response == 200) {

		LimpiezaRespuestaRefactor(resp, &pagos)
		if len(pagos) == 0 {
			return pago, errors.New("No se encontró pago mensual")
		}
		return pagos[0], nil
	}

	return pago, errors.New("Error en petición PagoMensual")
}

func GetPreliquidacion(id string) ([]models.PreliquidacionTitan, map[string]interface{}) {

	var outputError map[string]interface{}

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "/InformacionInforme",
				"err":     err,
				"status":  "502",
			}
			panic(outputError)
		}
	}()

	var pagos []models.PreliquidacionTitan
	var resp map[string]interface{}

	pago, err := getPagoMensual(id)
	if err != nil {
		return pagos, map[string]interface{}{
			"funcion": "/InformacionInforme/Preliquidacion",
			"err":     err,
			"status":  "502",
		}
	}

	anio := strconv.Itoa(int(pago.Ano))
	mes := strconv.Itoa(int(pago.Mes))
	contrato := pago.NumeroContrato
	vigencia := strconv.Itoa(int(pago.VigenciaContrato))
	doc := pago.DocumentoPersonaId
	cdp, _ := strconv.Atoi(pago.NumeroCDP)

	if response, err := getJsonTest(
		beego.AppConfig.String("UrlTitanMid")+
			"/detalle_preliquidacion/obtener_detalle_CT/"+anio+"/"+mes+"/"+contrato+"/"+vigencia+"/"+doc,
		&resp,
	); (err == nil) && (response == 201) {

		LimpiezaRespuestaRefactor(resp, &pagos)
		if preliq, err := seleccionarPreliquidacion(pagos, cdp); err == nil {
			return darFormatoPreliquidacion(preliq), nil
		} else {
			return pagos, map[string]interface{}{
				"funcion": "/InformacionInforme/SeleccionPreliquidacion",
				"err":     err,
				"status":  "400",
			}
		}
	}

	return pagos, map[string]interface{}{
		"funcion": "/InformacionInforme/Preliquidacion",
		"err":     errors.New("Error obteniendo preliquidacion"),
		"status":  "502",
	}
}

func seleccionarPreliquidacion(prel []models.PreliquidacionTitan, cdp int) ([]models.PreliquidacionTitan, error) {

	if len(prel) == 0 {
		return prel, errors.New("No hay preliquidaciones")
	}

	if len(prel) == 1 {
		return prel, nil
	}

	var result []models.PreliquidacionTitan
	for _, p := range prel {
		if p.Detalle[0].ContratoPreliquidacionId.ContratoId.Cdp == cdp {
			result = append(result, p)
		}
	}
	return result, nil
}

func darFormatoPreliquidacion(prel []models.PreliquidacionTitan) []models.PreliquidacionTitan {

	var out []models.PreliquidacionTitan

	for i, p := range prel {
		out = append(out, p)

		out[i].TotalDevengadoConFormato =
			FormatMoneyString(formatNumberString(strconv.Itoa(int(p.TotalDevengado)), 0, ",", "."), 0)
		out[i].TotalDescuentosConFormato =
			FormatMoneyString(formatNumberString(strconv.Itoa(int(p.TotalDescuentos)), 0, ",", "."), 0)
		out[i].TotalPagoConFormato =
			FormatMoneyString(formatNumberString(strconv.Itoa(int(p.TotalPago)), 0, ",", "."), 0)

		for j := range p.Detalle {
			out[i].Detalle[j].ValorCalculadoConFormato =
				FormatMoneyString(formatNumberString(strconv.Itoa(int(p.Detalle[j].ValorCalculado)), 0, ",", "."), 0)
		}
	}

	return out
}

func getValorGiradoPorCdp(cdp string, vigencia string, unidad string) (int, error) {

	var temp map[string]interface{}
	var giros models.GirosTercero
	total := 0

	url := beego.AppConfig.String("UrlFinancieraJBPM") +
		"/giros_tercero/" + cdp + "/" + vigencia + "/" + unidad

	response, err := getJsonWSO2Test(url, &temp)
	if err != nil || response != 200 {
		return total, err
	}

	b, err := json.Marshal(temp)
	if err != nil {
		return total, err
	}

	json.Unmarshal(b, &giros)

	for _, g := range giros.Giros.Tercero {
		val, err := strconv.Atoi(g.ValorBrutoGirado)
		if err == nil {
			total += val
		}
	}

	return total, nil
}
