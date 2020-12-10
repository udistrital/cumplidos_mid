package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/udistrital/cumplidos_mid/models"
)

func TraerInfoOrdenador(numero_contrato string, vigencia string) (informacion_ordenador models.InformacionOrdenador, err error) {
	var temp map[string]interface{}
	var contrato_elaborado models.ContratoElaborado
	var ordenadores_gasto []models.OrdenadorGasto
	var jefes_dependencia []models.JefeDependencia
	var informacion_proveedores []models.InformacionProveedor
	//var informacion_ordenador models.InformacionOrdenador
	var ordenadores []models.Ordenador

	if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAdministrativa")+"/"+"contrato_elaborado/"+numero_contrato+"/"+vigencia, &temp); err == nil && temp != nil {
		json_contrato_elaborado, error_json := json.Marshal(temp)
		if error_json == nil {
			if err := json.Unmarshal(json_contrato_elaborado, &contrato_elaborado); err == nil {
				if contrato_elaborado.Contrato.TipoContrato == "2" || contrato_elaborado.Contrato.TipoContrato == "3" || contrato_elaborado.Contrato.TipoContrato == "18" {
					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/ordenador_gasto/?query=Id:"+contrato_elaborado.Contrato.OrdenadorGasto, &ordenadores_gasto); err == nil {

						for _, ordenador_gasto := range ordenadores_gasto {

							if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/jefe_dependencia/?query=DependenciaId:"+strconv.Itoa(ordenador_gasto.DependenciaId)+"&sortby=FechaInicio&order=desc&limit=1", &jefes_dependencia); err == nil {

								for _, jefe_dependencia := range jefes_dependencia {

									if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+strconv.Itoa(jefe_dependencia.TerceroId), &informacion_proveedores); err == nil {

										for _, informacion_proveedor := range informacion_proveedores {

											informacion_ordenador.NumeroDocumento = jefe_dependencia.TerceroId
											informacion_ordenador.Cargo = ordenador_gasto.Cargo
											informacion_ordenador.Nombre = informacion_proveedor.NomProveedor
											informacion_ordenador.IdDependencia = jefe_dependencia.DependenciaId
											//c.Data["json"] = informacion_ordenador

										}

									} else {

										fmt.Println(err)
									}

								}

							} else {
								fmt.Println(err)
							}

						}

					} else {
						fmt.Println(err)
					}

				} else { //si no son docentes
					fmt.Println(contrato_elaborado.Contrato.OrdenadorGasto)
					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/ordenadores/?query=IdOrdenador:"+contrato_elaborado.Contrato.OrdenadorGasto+"&sortby=FechaInicio&order=desc&limit=1", &ordenadores); err == nil {
						for _, ordenador := range ordenadores {
							informacion_ordenador.NumeroDocumento = ordenador.Documento
							informacion_ordenador.Cargo = ordenador.RolOrdenador
							informacion_ordenador.Nombre = ordenador.NombreOrdenador
							//c.Data["json"] = informacion_ordenador

						}

					} else {

						fmt.Println(err)

					}

				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(error_json.Error())
		}
	} else {
		fmt.Println(err)

	}
	return
}

func SolicitudesOrdenador(doc_ordenador string, limit int, offset int) (pagos_personas_proyecto []models.PagoPersonaProyecto, err error) {

	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	//var pagos_personas_proyecto []models.PagoPersonaProyecto
	var pago_personas_proyecto models.PagoPersonaProyecto
	var vinculaciones_docente []models.VinculacionDocente
	var respuesta_peticion map[string]interface{}
	r := httplib.Get(beego.AppConfig.String("ProtocolCrudCumplidos") + "://" + beego.AppConfig.String("UrlCrudCumplidos") + "/" + beego.AppConfig.String("NsCrudCumplidos") + "/pago_mensual/")
	r.Param("offset", strconv.Itoa(offset))
	r.Param("limit", strconv.Itoa(limit))
	r.Param("query", "EstadoPagoMensualId.CodigoAbreviacion:AD,DocumentoResponsableId:"+doc_ordenador)

	if err := r.ToJSON(&respuesta_peticion); err == nil {
		//if err := r.ToJSON(&pagos_mensuales); err == nil {
		pagos_mensuales = []models.PagoMensual{}
		LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
		for x, pago_mensual := range pagos_mensuales {

			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pago_mensual.DocumentoPersonaId, &contratistas); err == nil {

				for _, contratista := range contratistas {
					// se podria armar un metodo desde aca y usar gorutines para reducir el tiempo de carga de las peticiones?
					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?limit=-1&query=NumeroContrato:"+pago_mensual.NumeroContrato+",Vigencia:"+strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64), &vinculaciones_docente); err == nil {

						for _, vinculacion := range vinculaciones_docente {
							var dep models.Dependencia

							if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia/"+strconv.Itoa(vinculacion.IdProyectoCurricular), &dep); err == nil {

								pago_personas_proyecto.PagoMensual = &pagos_mensuales[x]
								pago_personas_proyecto.NombrePersona = contratista.NomProveedor
								pago_personas_proyecto.Dependencia = &dep
								pagos_personas_proyecto = append(pagos_personas_proyecto, pago_personas_proyecto)

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

	return pagos_personas_proyecto, nil
}

func DependenciaOrdenador(doc_ordenador string) (dependenciaId int, err error) {

	var ordenadores_gasto []models.OrdenadorGasto
	var jefes_dependencia []models.JefeDependencia

	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/jefe_dependencia/?query=TerceroId:"+doc_ordenador+"&sortby=FechaInicio&order=desc&limit=1", &jefes_dependencia); err == nil {
		for _, jefe := range jefes_dependencia {

			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/ordenador_gasto/?query=DependenciaId:"+strconv.Itoa(jefe.DependenciaId), &ordenadores_gasto); err == nil {

				for _, ordenador := range ordenadores_gasto {

					return ordenador.DependenciaId, nil

				}

			} else { // If ordenador_gasto get
				fmt.Println("Mirenme, me morí en If ordenador_gasto get, solucioname!!! ", err)
				return dependenciaId, err
			}

		}

	} else { // If jefe_dependencia get
		fmt.Println("Mirenme, me morí en If jefe_dependencia get, solucioname!!! ", err)
		return dependenciaId, err
	}
	return
}
