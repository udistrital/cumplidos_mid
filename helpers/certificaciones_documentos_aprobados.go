package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

func CertificacionDocumentosAprobados(dependencia string, anio string, mes string) (personas []models.Persona, err error) {

	var contrato_ordenador_dependencia models.ContratoOrdenadorDependencia

	var pagos_mensuales []models.PagoMensual
	var persona models.Persona
	var vinculaciones_docente []models.VinculacionDocente

	var mes_cer, _ = strconv.Atoi(mes)

	if mes_cer < 10 {

		mes = "0" + mes

	}

	contrato_ordenador_dependencia = GetContratosOrdenadorDependencia(dependencia, anio+"-"+mes, anio+"-"+mes)

	for _, contrato := range contrato_ordenador_dependencia.ContratosOrdenadorDependencia.InformacionContratos {

		if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?limit=-1&query=NumeroContrato:"+contrato.NumeroContrato+",Vigencia:"+contrato.Vigencia, &vinculaciones_docente); err == nil {

			for _, vinculacion_docente := range vinculaciones_docente {
				if vinculacion_docente.NumeroContrato.Valid == true {

					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/pago_mensual/?query=EstadoPagoMensual.CodigoAbreviacion:AP,NumeroContrato:"+contrato.NumeroContrato+",VigenciaContrato:"+contrato.Vigencia+",Mes:"+strconv.Itoa(mes_cer)+",Ano:"+anio, &pagos_mensuales); err == nil {

						if pagos_mensuales == nil {

							persona.NumDocumento = contrato.Documento
							persona.Nombre = contrato.NombreContratista
							persona.NumeroContrato = contrato.NumeroContrato
							persona.Vigencia, _ = strconv.Atoi(contrato.Vigencia)

							personas = append(personas, persona)

						}

					} else { //If informacion_proveedor get

						fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
						return nil, err

					}

				}
			}

		} else { //If vinculacion_docente get

			fmt.Println("Mirenme, me morí en If vinculacion_docente get, solucioname!!! ", err)
			return nil, err

		}

	}

	return
}

func CertificadoVistoBueno(dependencia string, mes string, anio string) (personas []models.Persona, err error) {
	var vinculaciones_docente []models.VinculacionDocente
	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var persona models.Persona
	var actasInicio []models.ActaInicio
	var mes_cer, _ = strconv.Atoi(mes)
	var anio_cer, _ = strconv.Atoi(anio)
	var respuesta_peticion map[string]interface{}
	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?limit=-1&query=IdProyectoCurricular:"+dependencia, &vinculaciones_docente); err == nil {
		for _, vinculacion_docente := range vinculaciones_docente {
			if vinculacion_docente.NumeroContrato.Valid == true {

				if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/acta_inicio/?query=NumeroContrato:"+vinculacion_docente.NumeroContrato.String+",Vigencia:"+strconv.FormatInt(vinculacion_docente.Vigencia.Int64, 10), &actasInicio); err == nil {

					for _, actaInicio := range actasInicio {
						//If Estado = 4
						if int(actaInicio.FechaInicio.Month()) <= mes_cer && actaInicio.FechaInicio.Year() <= anio_cer && int(actaInicio.FechaFin.Month()) >= mes_cer && actaInicio.FechaFin.Year() >= anio_cer {

							if err := getJson(beego.AppConfig.String("ProtocolCrudCumplidos")+"://"+beego.AppConfig.String("UrlCrudCumplidos")+"/"+beego.AppConfig.String("NsCrudCumplidos")+"/pago_mensual/?query=EstadoPagoMensualId.CodigoAbreviacion.in:PAD|AD|AP,NumeroContrato:"+vinculacion_docente.NumeroContrato.String+",VigenciaContrato:"+strconv.FormatInt(vinculacion_docente.Vigencia.Int64, 10)+",Mes:"+mes+",Ano:"+anio, &respuesta_peticion); err == nil {
								fmt.Println(respuesta_peticion)
								b, _ := json.Marshal(respuesta_peticion["Data"])
								if len(string(b)) <= 4 {
									pagos_mensuales = nil
								} else {
									pagos_mensuales = []models.PagoMensual{}
									LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
								}
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

							}
						}
					}
				} else { //If contrato_estado get
					fmt.Println("Mirenme, me morí en If contrato_estado get, solucioname!!! ", err)
				}
			}
		}

	} else { //If vinculacion_docente get

		fmt.Println("Mirenme, me morí en If vinculacion_docente get, solucioname!!! ", err)
	}
	return
}

