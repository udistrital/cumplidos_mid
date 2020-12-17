package helpers

import (
	_ "encoding/json"
	"fmt"
	"strconv"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

func SolicitudCoordinador(doc_coordinador string) (pagos_personas_proyecto []models.PagoPersonaProyecto, err error) {

	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	var pago_personas_proyecto models.PagoPersonaProyecto
	var vinculaciones_docente []models.VinculacionDocente
	var respuesta_peticion map[string]interface{}
	if response, err := getJsonTest(beego.AppConfig.String("ProtocolCrudCumplidos")+"://"+beego.AppConfig.String("UrlCrudCumplidos")+"/"+beego.AppConfig.String("NsCrudCumplidos")+"/pago_mensual/?limit=-1&query=EstadoPagoMensualId.CodigoAbreviacion:PRC,DocumentoResponsableId:"+doc_coordinador, &respuesta_peticion); err == nil {
		pagos_mensuales = []models.PagoMensual{}
		LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
		for x, _ := range pagos_mensuales {

			if response, err := getJsonTest(beego.AppConfig.String("ProtocolCrudCumplidos")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pagos_mensuales[x].DocumentoPersonaId, &contratistas); err == nil {
				for _, contratista := range contratistas {

					if response, err := getJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?limit=-1&query=NumeroContrato:"+pagos_mensuales[x].NumeroContrato+",Vigencia:"+strconv.FormatFloat(pagos_mensuales[x].VigenciaContrato, 'f', 0, 64), &vinculaciones_docente); err == nil {
						for y, _ := range vinculaciones_docente {
							var dep []models.Dependencia

							if response, err := getJsonTest(beego.AppConfig.String("ProtocolCrudCumplidos")+"://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia/?query=Id:"+strconv.Itoa(vinculaciones_docente[y].IdProyectoCurricular), &dep); err == nil {
								for z, _ := range dep {
									pago_personas_proyecto.PagoMensual = &pagos_mensuales[x]
									pago_personas_proyecto.NombrePersona = contratista.NomProveedor
									pago_personas_proyecto.Dependencia = &dep[z]
									pagos_personas_proyecto = append(pagos_personas_proyecto, pago_personas_proyecto)
									return pagos_personas_proyecto, nil
								}

							} else { //If dependencia get

								fmt.Println("Mirenme, me morí en If dependencia get, solucioname!!! ", err)
								return nil, err
							}
						}

					} else { // If vinculacion_docente_get
						fmt.Println("Mirenme, me morí en If vinculacion_docente get, solucioname!!! ", err)
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

	return
}
