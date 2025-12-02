package helpers

//esta clase deberia guardar todos los get de contratos porque se pueden reutilizar
import (
	"encoding/json"
	_ "fmt"

	_ "github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

//RFC 47758 Se elimina función debido a que sólo contemplaba la última vigencia para un número de contrato, se usará la función GetContratosDependenciaFiltro
/*func GetContratosDependencia(dependencia string, fecha string) (salida map[string]string, outputError map[string]interface{}) {
	salida = make(map[string]string)
	var temp map[string]interface{}
	var contratos_dependencia models.ContratoDependencia
	//var salida map[string]string

	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlAdministrativaJBPM")+"/"+"contratos_dependencia/"+dependencia+"/"+fecha+"/"+fecha, &temp); (err == nil) && (response == 200) {
		if json_contrato, error_json := json.Marshal(temp); error_json == nil {
			if err := json.Unmarshal(json_contrato, &contratos_dependencia); err == nil {
				for _, cd := range contratos_dependencia.Contratos.Contrato {
					salida[cd.NumeroContrato] = cd.Vigencia
				}
				return salida, nil
			} else {
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/GetContratosDependencia", "err": err, "status": "502"}
				return salida, outputError
			}
		} else {
			logs.Error(error_json)
			outputError = map[string]interface{}{"funcion": "/GetContratosDependencia", "err": error_json.Error(), "status": "502"}
			return salida, outputError
		}

	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetContratosDependencia", "err": err.Error(), "status": "502"}
		return salida, outputError
	}
	//return contratos_dependencia
	return salida, outputError
}*/

// practicamente es el mismo metodo anterior
func GetContratosDependenciaFiltro(dependencia string, fecha_inicio string, fecha_fin string) (contratos_dependencia models.ContratoDependencia, outputError map[string]interface{}) {
	var temp map[string]interface{}
	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlHomologacionDepsJBPM")+"/oikos_argo/"+dependencia, &temp); (err == nil) && (response == 200) {
		json_dep_oikos, error_json := json.Marshal(temp)
		if error_json == nil {
			var depOikos models.HomologacionDepOikos
			if err := json.Unmarshal(json_dep_oikos, &depOikos); err == nil {

				if len(depOikos.Dependencias.Dependencia) != 0 {
					if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlAdministrativaJBPM")+"/contratos_dependencia_oikos/"+depOikos.Dependencias.Dependencia[0].IDMaster+"/"+fecha_inicio+"/"+fecha_fin, &temp); (err == nil) && (response == 200) {
						json_contrato, error_json := json.Marshal(temp)
						if error_json == nil {
							if err := json.Unmarshal(json_contrato, &contratos_dependencia); err == nil {
								return contratos_dependencia, nil
							} else {
								logs.Error(err)
								outputError = map[string]interface{}{"funcion": "/GetContratosDependenciaFiltro/contratos_dependencia_oikos", "err": err.Error(), "status": "502"}
								return contratos_dependencia, outputError

							}
						} else {
							logs.Error(error_json)
							outputError = map[string]interface{}{"funcion": "/GetContratosDependenciaFiltro/contratos_dependencia_oikos", "err": error_json.Error(), "status": "502"}
							return contratos_dependencia, outputError
						}

					} else {
						logs.Error(err)
						outputError = map[string]interface{}{"funcion": "/GetContratosDependenciaFiltro/contratos_dependencia_oikos", "err": err.Error(), "status": "502"}
						return contratos_dependencia, outputError
					}
				} else {
					outputError = map[string]interface{}{"funcion": "/GetContratosDependenciaFiltro/oikos_argo", "err": "no hay dependencia homologada en oikos", "status": "502"}
					return contratos_dependencia, outputError

				}

			} else {
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/GetContratosDependenciaFiltro/oikos_argo", "err": err.Error(), "status": "502"}
				return contratos_dependencia, outputError

			}
		} else {
			logs.Error(error_json)
			outputError = map[string]interface{}{"funcion": "/GetContratosDependenciaFiltro/oikos_argo", "err": error_json.Error(), "status": "502"}
			return contratos_dependencia, outputError
		}
	}
	return
}

func GetContratosOrdenadorDependencia(dependencia string, fechaInicio string, fechaFin string) (contratos_ordenador_dependencia models.ContratoOrdenadorDependencia, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "/GetContratosOrdenadorDependencia", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var temp map[string]interface{}
	if response, err := getJsonWSO2Test(beego.AppConfig.String("UrlAdministrativaJBPM")+"/"+"contratos_ordenador_dependencia/"+dependencia+"/"+fechaInicio+"/"+fechaFin, &temp); (err == nil) && (response == 200) {
		json_contrato_dependencia, error_json := json.Marshal(temp)
		if error_json == nil {
			if err := json.Unmarshal(json_contrato_dependencia, &contratos_ordenador_dependencia); err == nil {
				return contratos_ordenador_dependencia, nil
			} else {
				logs.Error(error_json.Error())
				outputError = map[string]interface{}{"funcion": "/GetInformacionContratoContratista", "err": error_json.Error(), "status": "502"}
				return contratos_ordenador_dependencia, outputError
			}
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetContratosOrdenadorDependencia", "err": err, "status": "502"}
		return contratos_ordenador_dependencia, outputError
	}

	return
}
