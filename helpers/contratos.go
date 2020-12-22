package helpers

//esta clase deberia guardar todos los get de contratos porque se pueden reutilizar
import (
	"encoding/json"
	"fmt"

	_ "github.com/astaxie/beego/httplib"
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

func GetContratosDependencia(dependencia string, fecha string) (salida map[string]string) {
	salida = make(map[string]string)
	var temp map[string]interface{}
	var contratos_dependencia models.ContratoDependencia
	//var salida map[string]string
	if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAdministrativa")+"/"+"contratos_dependencia/"+dependencia+"/"+fecha+"/"+fecha, &temp); err == nil {

		json_contrato, error_json := json.Marshal(temp)
		if error_json == nil {
			if err := json.Unmarshal(json_contrato, &contratos_dependencia); err == nil {

				for _, cd := range contratos_dependencia.Contratos.Contrato {
					//fmt.Println("dependencia: ", cd)
					//fmt.Println(cd.NumeroContrato, cd.Vigencia)
					//salida["prueba"] = "cd.Vigencia"
					//fmt.Println(cd.NumeroContrato)
					salida[cd.NumeroContrato] = cd.Vigencia
					//fmt.Println("salida: ", salida)
				}
				return salida
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(error_json.Error())
		}

	} else {

		fmt.Println(err)
	}

	//return contratos_dependencia
	return salida
}

func GetContratosDependenciaFiltro(dependencia string, fecha_inicio string, fecha_fin string) (contratos_dependencia models.ContratoDependencia) {

	var temp map[string]interface{}

	if err := getJsonWSO2("http://"+beego.AppConfig.String("UrlcrudWSO2")+"/"+beego.AppConfig.String("NscrudAdministrativa")+"/"+"contratos_dependencia/"+dependencia+"/"+fecha_fin+"/"+fecha_inicio, &temp); err == nil {
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

func GetContratosOrdenadorDependencia(dependencia string, fechaInicio string, fechaFin string) (contratos_ordenador_dependencia models.ContratoOrdenadorDependencia, outputError map[string]interface{}) {

	var temp map[string]interface{}
	if response, err := getJsonWSO2Test("http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudAdministrativa") + "/" + "contratos_ordenador_dependencia/" + dependencia + "/" + fechaInicio + "/" + fechaFin, &temp); (err == nil)  && (response == 200) {
		json_contrato_dependencia, error_json := json.Marshal(temp)
		if error_json == nil{
			if err := json.Unmarshal(json_contrato_dependencia, &contratos_ordenador_dependencia); err == nil {
				return contratos_ordenador_dependencia, nil
			}else{
				fmt.Println(error_json.Error())
			}
		}
	}else{
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/CertificacionDocumentosAprobados/GetContratosOrdenadorDependencia", "err": err}
		return contratos_ordenador_dependencia, outputError
	}

	//r := httplib.Get("http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudAdministrativa") + "/" + "contratos_ordenador_dependencia/" + dependencia + "/" + fechaInicio + "/" + fechaFin)
	//r.Header("Accept", "application/json")

	//if err := r.ToJSON(&contratos_ordenador_dependencia); err == nil {
	//	return contratos_ordenador_dependencia, nil
	//} else {
		
	//}

	return
}
