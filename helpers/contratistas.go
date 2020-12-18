package helpers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/udistrital/cumplidos_mid/models"
)

/*func CertificacionCumplidosContratistas2(dependencia string, mes string, anio string) (personas []models.Persona, err error) {

	var contrato_dependencia models.ContratoDependencia
	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var persona models.Persona

	var nmes, _ = strconv.Atoi(mes)
	var respuesta_peticion map[string]interface{}

	contrato_dependencia = GetContratosDependencia2(dependencia, anio+"-"+mes)

	for _, cd := range contrato_dependencia.Contratos.Contrato {

		if err := getJson(beego.AppConfig.String("ProtocolCrudCumplidos")+"://"+beego.AppConfig.String("UrlcrudCumplidos")+"/"+beego.AppConfig.String("NscrudCumplidos")+"/pago_mensual/?query=EstadoPagoMensualId.CodigoAbreviacion.in:AS|AP,NumeroContrato:"+cd.NumeroContrato+",VigenciaContrato:"+cd.Vigencia+",Mes:"+strconv.Itoa(nmes)+",Ano:"+anio, &respuesta_peticion); err == nil {

			// se hace para limpiar la variable
			pagos_mensuales = []models.PagoMensual{}
			LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
			//fmt.Println("Pagos: ", respuesta_peticion)
			//fmt.Println("Pagos: ", len(pagos_mensuales))
			//fmt.Println(beego.AppConfig.String("ProtocolCrudCumplidos") + "://" + beego.AppConfig.String("UrlcrudCumplidos") + "/" + beego.AppConfig.String("NscrudCumplidos") + "/pago_mensual/?query=EstadoPagoMensualId.CodigoAbreviacion.in:AS|AP,NumeroContrato:" + cd.NumeroContrato + ",VigenciaContrato:" + cd.Vigencia + ",Mes:" + strconv.Itoa(nmes) + ",Ano:" + anio)
			for v := range pagos_mensuales {
				if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pagos_mensuales[v].DocumentoPersonaId, &contratistas); err == nil {
					//fmt.Println("agora URL: ", beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pagos_mensuales[v].DocumentoPersonaId)
					var contrato models.InformacionContrato
					contrato = GetContrato(pagos_mensuales[v].NumeroContrato, strconv.FormatFloat(pagos_mensuales[v].VigenciaContrato, 'f', 0, 64))

					for _, contratista := range contratistas {
						persona.NumDocumento = contratista.NumDocumento
						persona.Nombre = contratista.NomProveedor
						persona.NumeroContrato = pagos_mensuales[v].NumeroContrato
						persona.Vigencia, _ = strconv.Atoi(cd.Vigencia)
						persona.Rubro = contrato.Contrato.Rubro

						personas = append(personas, persona)
					}

				} else { //If informacion_proveedor get

					fmt.Println("Mirenme, me morí en If informacion_proveedor get, solucioname!!! ", err)
					return nil, err

				}
			}
		} else { //If pago_mensual get

			fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
			return nil, err

		}

	}
	return
}*/

func CertificacionCumplidosContratistas(dependencia string, mes string, anio string) (personas []models.Persona, err error) {

	//var contrato_dependencia models.ContratoDependencia
	var contrato_dependencia map[string]string
	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var persona models.Persona

	var nmes, _ = strconv.Atoi(mes)
	var respuesta_peticion map[string]interface{}
	//traemos los contratos activos para un mes en una dependencia
	contrato_dependencia = GetContratosDependencia(dependencia, anio+"-"+mes)
	if err := getJson(beego.AppConfig.String("ProtocolCrudCumplidos")+"://"+beego.AppConfig.String("UrlcrudCumplidos")+"/"+beego.AppConfig.String("NscrudCumplidos")+"/pago_mensual/?query=EstadoPagoMensualId.CodigoAbreviacion.in:AS|AP,Mes:"+strconv.Itoa(nmes)+",Ano:"+anio+"&limit=-1", &respuesta_peticion); err == nil {

		pagos_mensuales = []models.PagoMensual{}
		LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
		fmt.Println("total pagos: ", len(pagos_mensuales))
		for _, pago_mensual := range pagos_mensuales {
			if vigencia, ok := contrato_dependencia[pago_mensual.NumeroContrato]; ok == true && vigencia == strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64) {
				//fmt.Println("num contrato: ", contrato_dependencia[pago_mensual.NumeroContrato], " vig: ", vigencia)

				if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pago_mensual.DocumentoPersonaId, &contratistas); err == nil {
					//fmt.Println("agora URL: ", beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pagos_mensuales[v].DocumentoPersonaId)
					var contrato models.InformacionContrato

					contrato = GetContrato(pago_mensual.NumeroContrato, strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64))

					for _, contratista := range contratistas {
						persona.NumDocumento = contratista.NumDocumento
						persona.Nombre = contratista.NomProveedor
						persona.NumeroContrato = pago_mensual.NumeroContrato
						persona.Vigencia = int(pago_mensual.VigenciaContrato) //strconv.Atoi(cd.Vigencia)
						persona.Rubro = contrato.Contrato.Rubro
						personas = append(personas, persona)
					}

				} else { //If informacion_proveedor get

					fmt.Println("Mirenme, me morí en If informacion_proveedor get, solucioname!!! ", err)
					return nil, err

				}

			} else {
				//fmt.Println("se murio el if aspero")
			}

		}
	} else { //If pago_mensual get

		fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
		return nil, err

	}
	return
}

func AprobacionPagosContratistas(v []models.PagoContratistaCdpRp) (err error) {
	//var v []models.PagoContratistaCdpRp
	var response interface{}

	var pagos_mensuales []*models.PagoMensual

	var pago_mensual *models.PagoMensual
	for _, pm := range v {

		pago_mensual = pm.PagoMensual
		pagos_mensuales = append(pagos_mensuales, pago_mensual)
	}
	if err := sendJson(beego.AppConfig.String("ProtocolCrudCumplidos")+"://"+beego.AppConfig.String("UrlCrudCumplidos")+"/"+beego.AppConfig.String("NsCrudCumplidos")+"/tr_aprobacion_masiva_pagos", "POST", &response, pagos_mensuales); err != nil {
		fmt.Println(err)
		return err

	}
	return nil
}

func SolicitudesOrdenadorContratistas(doc_ordenador string, limit int, offset int) (pagos_contratista_cdp_rp []models.PagoContratistaCdpRp, err error) {
	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor

	var contratos_disponibilidad []models.ContratoDisponibilidad
	var respuesta_peticion map[string]interface{}

	r := httplib.Get(beego.AppConfig.String("ProtocolCrudCumplidos") + "://" + beego.AppConfig.String("UrlCrudCumplidos") + "/" + beego.AppConfig.String("NsCrudCumplidos") + "/pago_mensual/")
	r.Param("offset", strconv.Itoa(offset))
	r.Param("limit", strconv.Itoa(limit))
	r.Param("query", "EstadoPagoMensualId.CodigoAbreviacion:AS,DocumentoResponsableId:"+doc_ordenador)

	if err := r.ToJSON(&respuesta_peticion); err == nil {
		pagos_mensuales = []models.PagoMensual{}
		LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
		//fmt.Println(r, " respuesta: ", respuesta_peticion)
		//fmt.Println("pagos: ", pagos_mensuales)
		for v, _ := range pagos_mensuales {

			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pagos_mensuales[v].DocumentoPersonaId, &contratistas); err == nil {

				for _, contratista := range contratistas {

					var informacion_contrato_contratista models.InformacionContratoContratista
					informacion_contrato_contratista = GetInformacionContratoContratista(pagos_mensuales[v].NumeroContrato, strconv.FormatFloat(pagos_mensuales[v].VigenciaContrato, 'f', 0, 64))
					var contrato models.InformacionContrato
					contrato = GetContrato(pagos_mensuales[v].NumeroContrato, strconv.FormatFloat(pagos_mensuales[v].VigenciaContrato, 'f', 0, 64))

					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/contrato_disponibilidad/?query=NumeroContrato:"+contrato.Contrato.NumeroContrato+",Vigencia:"+contrato.Contrato.Vigencia, &contratos_disponibilidad); err == nil {

						for _, contrato_disponibilidad := range contratos_disponibilidad {

							var cdprp models.InformacionCdpRp
							cdprp = GetRP(strconv.Itoa(contrato_disponibilidad.NumeroCdp), strconv.Itoa(contrato_disponibilidad.VigenciaCdp))
							for _, rp := range cdprp.CdpXRp.CdpRp {
								var pago_contratista_cdp_rp models.PagoContratistaCdpRp

								pago_contratista_cdp_rp.PagoMensual = &pagos_mensuales[v]
								pago_contratista_cdp_rp.NombreDependencia = informacion_contrato_contratista.InformacionContratista.Dependencia
								pago_contratista_cdp_rp.NombrePersona = contratista.NomProveedor
								pago_contratista_cdp_rp.NumeroCdp = strconv.Itoa(contrato_disponibilidad.NumeroCdp)
								pago_contratista_cdp_rp.VigenciaCdp = strconv.Itoa(contrato_disponibilidad.VigenciaCdp)
								pago_contratista_cdp_rp.NumeroRp = rp.RpNumeroRegistro
								pago_contratista_cdp_rp.VigenciaRp = rp.RpVigencia
								pago_contratista_cdp_rp.Rubro = contrato.Contrato.Rubro

								pagos_contratista_cdp_rp = append(pagos_contratista_cdp_rp, pago_contratista_cdp_rp)

							}

						}

					} else { // If contrato_disponibilidad get
						fmt.Println("Mirenme, me morí en If contrato_disponibilidad get, solucioname!!! ", err)
						return nil, err

					}

				}
			} else { //If informacion_proveedor get
				fmt.Println("Mirenme, me morí en If informacion_proveedor get, solucioname!!! ", err)
				return nil, err
			}

		}
	} else { //If pago_mensual get
		fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
		return nil, err
	}

	return pagos_contratista_cdp_rp, nil
}
