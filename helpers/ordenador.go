package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/models"
)

func TraerInfoOrdenador(numero_contrato string, vigencia string) (informacion_ordenador models.InformacionOrdenador, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/TraerInfoOrdenador", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var temp map[string]interface{}
	var contrato_elaborado models.ContratoElaborado
	var ordenadores_gasto []models.OrdenadorGasto
	var jefes_dependencia []models.JefeDependencia
	var informacion_proveedores []models.InformacionProveedor
	//var informacion_ordenador models.InformacionOrdenador
	var ordenadores []models.Ordenador

	if response, err := getJsonWSO2Test("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAdministrativa")+"/"+"contrato_elaborado/"+numero_contrato+"/"+vigencia, &temp); err == nil && temp != nil && response == 200 {
		json_contrato_elaborado, error_json := json.Marshal(temp)
		if error_json == nil {
			if err := json.Unmarshal(json_contrato_elaborado, &contrato_elaborado); err == nil {
				if contrato_elaborado.Contrato.TipoContrato == "2" || contrato_elaborado.Contrato.TipoContrato == "3" || contrato_elaborado.Contrato.TipoContrato == "18" {
					if response, err := getJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/ordenador_gasto/?query=Id:"+contrato_elaborado.Contrato.OrdenadorGasto, &ordenadores_gasto); (err == nil) && (response == 200) {

						for _, ordenador_gasto := range ordenadores_gasto {

							if response, err := getJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/jefe_dependencia/?query=DependenciaId:"+strconv.Itoa(ordenador_gasto.DependenciaId)+"&sortby=FechaInicio&order=desc&limit=1", &jefes_dependencia); (err == nil) && (response == 200) {

								for _, jefe_dependencia := range jefes_dependencia {

									if response, err := getJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+strconv.Itoa(jefe_dependencia.TerceroId), &informacion_proveedores); (err == nil) && (response == 200) {

										for _, informacion_proveedor := range informacion_proveedores {

											informacion_ordenador.NumeroDocumento = jefe_dependencia.TerceroId
											informacion_ordenador.Cargo = ordenador_gasto.Cargo
											informacion_ordenador.Nombre = informacion_proveedor.NomProveedor
											informacion_ordenador.IdDependencia = jefe_dependencia.DependenciaId
											//c.Data["json"] = informacion_ordenador
										}
									} else {
										logs.Error(err)
										outputError = map[string]interface{}{"funcion": "/TraerInfoOrdenador", "err": err.Error(), "status": "502"}
										return informacion_ordenador, outputError
									}

								}

							} else {
								logs.Error(err)
								outputError = map[string]interface{}{"funcion": "/TraerInfoOrdenador", "err": err.Error(), "status": "502"}
								return informacion_ordenador, outputError
							}

						}

					} else {
						logs.Error(err)
						outputError = map[string]interface{}{"funcion": "/TraerInfoOrdenador", "err": err.Error(), "status": "502"}
						return informacion_ordenador, outputError
					}

				} else { //si no son docentes
					if response, err := getJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/ordenadores/?query=IdOrdenador:"+contrato_elaborado.Contrato.OrdenadorGasto+"&sortby=FechaInicio&order=desc&limit=1", &ordenadores); (err == nil) && (response == 200) {
						for _, ordenador := range ordenadores {
							informacion_ordenador.NumeroDocumento = ordenador.Documento
							informacion_ordenador.Cargo = ordenador.RolOrdenador
							informacion_ordenador.Nombre = ordenador.NombreOrdenador
							//c.Data["json"] = informacion_ordenador
						}

					} else {
						logs.Error(err)
						outputError = map[string]interface{}{"funcion": "/TraerInfoOrdenador", "err": err.Error(), "status": "502"}
						return informacion_ordenador, outputError
					}

				}
			} else {
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/TraerInfoOrdenador", "err": err.Error(), "status": "502"}
				return informacion_ordenador, outputError

			}
		} else {
			fmt.Println(error_json.Error())
			logs.Error(error_json.Error())
			outputError = map[string]interface{}{"funcion": "/TraerInfoOrdenador", "err": error_json.Error(), "status": "502"}
			return informacion_ordenador, outputError
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/TraerInfoOrdenador", "err": err.Error(), "status": "502"}
		return informacion_ordenador, outputError
	}
	return
}

func SolicitudesOrdenador(doc_ordenador string, limit int, offset int) (pagos_personas_proyecto []models.PagoPersonaProyecto, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/SolicitudesOrdenador", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	//var pagos_personas_proyecto []models.PagoPersonaProyecto
	var pago_personas_proyecto models.PagoPersonaProyecto
	var vinculaciones_docente []models.VinculacionDocente
	var respuesta_peticion map[string]interface{}

	if response, err := getJsonTest(beego.AppConfig.String("ProtocolCrudCumplidos")+"://"+beego.AppConfig.String("UrlCrudCumplidos")+"/"+beego.AppConfig.String("NsCrudCumplidos")+"/pago_mensual/?offset="+strconv.Itoa(offset)+"&limit="+strconv.Itoa(limit)+"&query=EstadoPagoMensualId.CodigoAbreviacion:AD,DocumentoResponsableId:"+doc_ordenador, &respuesta_peticion); (err == nil) && (response == 200) {
		pagos_mensuales = []models.PagoMensual{}
		LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
		for x, pago_mensual := range pagos_mensuales {
			if response, err := getJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pago_mensual.DocumentoPersonaId, &contratistas); (err == nil) && (response == 200) {
				for _, contratista := range contratistas {
					// se podria armar un metodo desde aca y usar gorutines para reducir el tiempo de carga de las peticiones?
					if response, err := getJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?limit=-1&query=NumeroContrato:"+pago_mensual.NumeroContrato+",Vigencia:"+strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64), &vinculaciones_docente); (err == nil) && (response == 200) {
						for _, vinculacion := range vinculaciones_docente {
							var dep models.Dependencia
							if response, err := getJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia/"+strconv.Itoa(vinculacion.IdProyectoCurricular), &dep); (err == nil) && (response == 200) {
								pago_personas_proyecto.PagoMensual = &pagos_mensuales[x]
								pago_personas_proyecto.NombrePersona = contratista.NomProveedor
								pago_personas_proyecto.Dependencia = &dep
								pagos_personas_proyecto = append(pagos_personas_proyecto, pago_personas_proyecto)

							} else { //If dependencia get
								logs.Error(err)
								outputError = map[string]interface{}{"funcion": "/SolicitudesOrdenador1", "err": err.Error(), "status": "502"}
								return nil, outputError
							}

						}

					} else { // If vinculacion_docente_get
						logs.Error(err)
						outputError = map[string]interface{}{"funcion": "/SolicitudesOrdenador2", "err": err.Error(), "status": "502"}
						return nil, outputError
					}
				}
			} else { //If informacion_proveedor get
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/SolicitudesOrdenador3", "err": err.Error(), "status": "502"}
				return nil, outputError
			}
		}

	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/SolicitudesOrdenador4", "err": err.Error(), "status": "502"}
		return nil, outputError
	}
	return pagos_personas_proyecto, nil
}

func DependenciaOrdenador(doc_ordenador string) (dependenciaId int, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/DependenciaOrdenador", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var ordenadores_gasto []models.OrdenadorGasto
	var jefes_dependencia []models.JefeDependencia
	// las consultas de la linea 164 y 187 se pueden unir y dejar una sola que use sql puro para relacionarlas
	if response, err := getJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/jefe_dependencia/?query=TerceroId:"+doc_ordenador+"&sortby=FechaInicio&order=desc&limit=1", &jefes_dependencia); (err == nil) && (response == 200) {
		for _, jefe := range jefes_dependencia {

			if response, err := getJsonTest(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/ordenador_gasto/?query=DependenciaId:"+strconv.Itoa(jefe.DependenciaId), &ordenadores_gasto); (err == nil) && (response == 200) {

				for _, ordenador := range ordenadores_gasto {

					return ordenador.DependenciaId, nil

				}

			} else { // If ordenador_gasto get
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/DependenciaOrdenador", "err": err.Error(), "status": "502"}
				return 0, outputError
			}

		}

	} else { // If jefe_dependencia get
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/DependenciaOrdenador", "err": err.Error(), "status": "502"}
		return 0, outputError
	}
	return
}
