package helpers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/udistrital/cumplidos_mid/models"
)

func CertificacionDocumentosAprobados(dependencia string, anio string, mes string) (personas []models.Persona, err error) {

	var contrato_ordenador_dependencia models.ContratoOrdenadorDependencia

	var pagos_mensuales []models.PagoMensual
	//var personas []models.Persona
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

func GetContratosOrdenadorDependencia(dependencia string, fechaInicio string, fechaFin string) (contratos_ordenador_dependencia models.ContratoOrdenadorDependencia) {

	r := httplib.Get("http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudAdministrativa") + "/" + "contratos_ordenador_dependencia/" + dependencia + "/" + fechaInicio + "/" + fechaFin)
	r.Header("Accept", "application/json")
	if err := r.ToJSON(&contratos_ordenador_dependencia); err == nil {
	} else {

		fmt.Println(err)
	}

	return contratos_ordenador_dependencia
}
