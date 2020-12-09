package controllers

import (
	"fmt"
	"strconv"

	//"net/http"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

// CertificacionVistoBuenoController operations for certificacion_visto_bueno
type CertificacionVistoBuenoController struct {
	beego.Controller
}

// URLMapping ...
func (c *CertificacionVistoBuenoController) URLMapping() {
	c.Mapping("CertificacionVistoBueno", c.CertificacionVistoBueno)
}

// CertificacionVistoBuenoController ...
// @Title CertificacionVistoBueno
// @Description create CertificacionVistoBueno
// @Param dependencia path int true "Dependencia del contrato en la tabla vinculacion"
// @Param mes path int true "Mes del pago mensual"
// @Param anio path int true "Año del pago mensual"
// @Success 201
// @Failure 403 :dependencia is empty
// @Failure 403 :mes is empty
// @Failure 403 :anio is empty
// @router /certificacion_visto_bueno/:dependencia/:mes/:anio [get]
func (c *CertificacionVistoBuenoController) CertificacionVistoBueno() {
	dependencia := c.GetString(":dependencia")
	mes := c.GetString(":mes")
	anio := c.GetString(":anio")
	var vinculaciones_docente []models.VinculacionDocente
	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var personas []models.Persona
	var persona models.Persona
	var actasInicio []models.ActaInicio
	var mes_cer, _ = strconv.Atoi(mes)
	var anio_cer, _ = strconv.Atoi(anio)
	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?limit=-1&query=IdProyectoCurricular:"+dependencia, &vinculaciones_docente); err == nil {
		for _, vinculacion_docente := range vinculaciones_docente {
			if vinculacion_docente.NumeroContrato.Valid == true {

				if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/?query=NumeroContrato:"+vinculacion_docente.NumeroContrato.String+",Vigencia:"+strconv.FormatInt(vinculacion_docente.Vigencia.Int64, 10), &actasInicio); err == nil {

					for _, actaInicio := range actasInicio {
						//If Estado = 4
						if int(actaInicio.FechaInicio.Month()) <= mes_cer && actaInicio.FechaInicio.Year() <= anio_cer && int(actaInicio.FechaFin.Month()) >= mes_cer && actaInicio.FechaFin.Year() >= anio_cer {

							if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/pago_mensual/?query=EstadoPagoMensual.CodigoAbreviacion.in:PAD|AD|AP,NumeroContrato:"+vinculacion_docente.NumeroContrato.String+",VigenciaContrato:"+strconv.FormatInt(vinculacion_docente.Vigencia.Int64, 10)+",Mes:"+mes+",Ano:"+anio, &pagos_mensuales); err == nil {

								if pagos_mensuales == nil {

									if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+vinculacion_docente.IdPersona, &contratistas); err == nil {

										for _, contratista := range contratistas {

											persona.NumDocumento = contratista.NumDocumento
											persona.Nombre = contratista.NomProveedor
											persona.NumeroContrato = actaInicio.NumeroContrato
											persona.Vigencia = actaInicio.Vigencia

											personas = append(personas, persona)

										}

									} else { //If informacion_proveedor get

										fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
									}

								}

							} else { //If pago_mensual get
								fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
								return
							}
						}
					}
				} else { //If contrato_estado get
					fmt.Println("Mirenme, me morí en If contrato_estado get, solucioname!!! ", err)
					return
				}
			}
		}

	} else { //If vinculacion_docente get

		fmt.Println("Mirenme, me morí en If vinculacion_docente get, solucioname!!! ", err)
	}
	c.Data["json"] = personas

	c.ServeJSON()

}