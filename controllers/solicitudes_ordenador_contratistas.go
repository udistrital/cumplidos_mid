package controllers

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"strconv"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
	_ "github.com/udistrital/cumplidos_mid/helpers"
	"github.com/udistrital/cumplidos_mid/models"
)

// SolicitudesOrdenadorContratistasController operations for SolicitudesOrdenadorContratistas
type SolicitudesOrdenadorContratistasController struct {
	beego.Controller
}

//URLMapping ...
func (c *SolicitudesOrdenadorContratistasController) URLMapping() {
	c.Mapping("GetSolicitudesOrdenadorContratistas", c.GetSolicitudesOrdenadorContratistas)
	c.Mapping("AprobarMultiplesPagosContratistas", c.AprobarMultiplesPagosContratistas)
	c.Mapping("CertificacionCumplidosContratistas", c.CertificacionCumplidosContratistas)
}

// AprobacionPagoController ...
// @Title GetSolicitudesOrdenadorContratistas
// @Description create GetSolicitudesOrdenadorContratistas
// @Param docordenador path string true "Número del documento del supervisor"
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.PagoContratistaCdpRp
// @Failure 403 :docordenador is empty
// @router /solicitudes:docordenador [get]
func (c *SolicitudesOrdenadorContratistasController) GetSolicitudesOrdenadorContratistas() {

	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")
	if pagos_contratista_cdp_rp, err := SolicitudesOrdenadorContratistas(c.GetString(":docordenador"), limit, offset); err != nil {
		logs.Error(err)
		c.Data["mesaage"] = "Error service Get solicitudes_coordinador: The request contains an incorrect parameter or no record exists"
		c.Abort("404")

	} else {
		c.Data["json"] = pagos_contratista_cdp_rp
	}

	c.ServeJSON()

}

// AprobacionPagoController ...
// @Title AprobarMultiplesPagosContratistas
// @Description create AprobarMultiplesPagosContratistas
// @Param	body		body 	[]models.PagoContratistaCdpRp	true		"body for SoportePagoMensual content"
// @Success 201
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *SolicitudesOrdenadorContratistasController) AprobarMultiplesPagosContratistas() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			c.Data["mesaage"] = "Error service Get solicitudes_coordinador: The request contains an incorrect parameter or no record exists"
			c.Abort("404")
		}
	}()

	var v []models.PagoContratistaCdpRp
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := AprobacionPagosContratistas(v); err == nil {
			c.Data["json"] = "OK"
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}

	c.ServeJSON()
}

// AprobacionPagoController ...
// @Title certificacion_cumplidos_contratistas
// @Description get certificacion_cumplidos_contratistas
// @Param dependencia path string true "Dependencia supervisor"
// @Param mes path string true "Mes del certificado"
// @Param anio path string true "Año del certificado "
// @Success 201 {object} []models.Persona
// @Failure 403 :dependencia is empty
// @Failure 403 :mes is empty
// @Failure 403 :anio is empty
// @router /certificaciones/:dependencia/:mes/:anio [get]
func (c *SolicitudesOrdenadorContratistasController) CertificacionCumplidosContratistas() {

	dependencia := c.GetString(":dependencia")
	mes := c.GetString(":mes")
	anio := c.GetString(":anio")

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			c.Data["mesaage"] = "Error service Get solicitudes_coordinador: The request contains an incorrect parameter or no record exists"
			c.Abort("404")
		}
	}()

	//var v []models.PagoContratistaCdpRp
	if personas, err := CertificacionCumplidosContratistas(dependencia, mes, anio); err == nil {
		c.Data["json"] = personas
	} else {
		panic(err)
	}

	c.ServeJSON()
}

func CertificacionCumplidosContratistas(dependencia string, mes string, anio string) (personas []models.Persona, err error) {

	var contrato_dependencia models.ContratoDependencia
	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var persona models.Persona

	var nmes, _ = strconv.Atoi(mes)
	var respuesta_peticion map[string]interface{}

	contrato_dependencia = GetContratosDependencia(dependencia, anio+"-"+mes)

	for _, cd := range contrato_dependencia.Contratos.Contrato {

		if err := getJson(beego.AppConfig.String("ProtocolCrudCumplidos")+"://"+beego.AppConfig.String("UrlcrudCumplidos")+"/"+beego.AppConfig.String("NscrudCumplidos")+"/pago_mensual/?query=EstadoPagoMensualId.CodigoAbreviacion.in:AS|AP,NumeroContrato:"+cd.NumeroContrato+",VigenciaContrato:"+cd.Vigencia+",Mes:"+strconv.Itoa(nmes)+",Ano:"+anio, &respuesta_peticion); err == nil {

			// se hace para limpiar la variable
			pagos_mensuales = []models.PagoMensual{}
			helpers.LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)

			for v := range pagos_mensuales {

				if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pagos_mensuales[v].DocumentoPersonaId, &contratistas); err == nil {

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
		helpers.LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
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

func GetContratosDependencia(dependencia string, fecha string) (contratos_dependencia models.ContratoDependencia) {

	var temp map[string]interface{}

	if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAdministrativa")+"/"+"contratos_dependencia/"+dependencia+"/"+fecha+"/"+fecha, &temp); err == nil {
		json_contrato, error_json := json.Marshal(temp)
		if error_json == nil {
			if err := json.Unmarshal(json_contrato, &contratos_dependencia); err == nil {
				return contratos_dependencia
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(error_json.Error())
		}

	} else {

		fmt.Println(err)
	}

	return contratos_dependencia
}
