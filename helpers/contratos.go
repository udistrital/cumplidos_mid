package helpers

//esta clase deberia guardar todos los get de contratos porque se pueden reutilizar
import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

/*func GetContratosDependencia2(dependencia string, fecha string) (contratos_dependencia models.ContratoDependencia) {

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
}*/

func GetContratosDependencia(dependencia string, fecha string) (salida map[string]string, outputError map[string]interface{}) {
	salida = make(map[string]string)
	var temp map[string]interface{}
	var contratos_dependencia models.ContratoDependencia
	//var salida map[string]string

	r := httplib.Get("http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudAdministrativa") + "/" + "contratos_dependencia/" + dependencia + "/" + fecha + "/" + fecha)
	r.Header("Accept", "application/json")

	if err := r.ToJSON(&temp); err == nil {
		if json_contrato, error_json := json.Marshal(temp); error_json == nil {
			if err := json.Unmarshal(json_contrato, &contratos_dependencia); err == nil {
				for _, cd := range contratos_dependencia.Contratos.Contrato {
					salida[cd.NumeroContrato] = cd.Vigencia
				}
				return salida, nil
			} else {
				//fmt.Println(err)
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/GetContratosDependencia", "err": err, "status": "502"}
				return salida, outputError
			}
		} else {
			fmt.Println(error_json.Error())
			logs.Error(err)
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
}

//practicamente es el mismo metodo anterior
func GetContratosDependenciaFiltro(dependencia string, fecha_inicio string, fecha_fin string) (contratos_dependencia models.ContratoDependencia, outputError map[string]interface{}) {

	var temp map[string]interface{}
	r := httplib.Get("http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudAdministrativa") + "/" + "contratos_dependencia/" + dependencia + "/" + fecha_fin + "/" + fecha_inicio)
	r.Header("Accept", "application/json")
	if err := r.ToJSON(&temp); err == nil {
		//if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAdministrativa")+"/"+"contratos_dependencia/"+dependencia+"/"+fecha_fin+"/"+fecha_inicio, &temp); err == nil {
		json_contrato, error_json := json.Marshal(temp)
		if error_json == nil {
			if err := json.Unmarshal(json_contrato, &contratos_dependencia); err == nil {
				return contratos_dependencia, nil
			} else {
				fmt.Println(err)
				logs.Error(err)
				outputError = map[string]interface{}{"funcion": "/GetContratosDependenciaFiltro", "err": err.Error(), "status": "502"}
				return contratos_dependencia, outputError

			}
		} else {
			fmt.Println(error_json.Error())
			logs.Error(err)
			outputError = map[string]interface{}{"funcion": "/GetContratosDependenciaFiltro", "err": err.Error(), "status": "502"}
			return contratos_dependencia, outputError
		}

	} else {
		fmt.Println(err)
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetContratosDependenciaFiltro", "err": err.Error(), "status": "502"}
		return contratos_dependencia, outputError
	}

	return
}

func GetContratosOrdenadorDependencia(dependencia string, fechaInicio string, fechaFin string) (contratos_ordenador_dependencia models.ContratoOrdenadorDependencia, outputError map[string]interface{}) {

	r := httplib.Get("http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudAdministrativa") + "/" + "contratos_ordenador_dependencia/" + dependencia + "/" + fechaInicio + "/" + fechaFin)
	r.Header("Accept", "application/json")

	if err := r.ToJSON(&contratos_ordenador_dependencia); err == nil {
		return contratos_ordenador_dependencia, nil
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetContratosOrdenadorDependencia", "err": err, "status": "502"}
		return contratos_ordenador_dependencia, outputError
	}

	return
}
