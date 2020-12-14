package helpers

//esta clase deberia guardar todos los get de contratos porque se pueden reutilizar
import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

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


func GetContratosOrdenadorDependencia(dependencia string, fechaInicio string, fechaFin string) (contratos_ordenador_dependencia models.ContratoOrdenadorDependencia) {

	r := httplib.Get("http://" + beego.AppConfig.String("UrlcrudWSO2") + "/" + beego.AppConfig.String("NscrudAdministrativa") + "/" + "contratos_ordenador_dependencia/" + dependencia + "/" + fechaInicio + "/" + fechaFin)
	r.Header("Accept", "application/json")
	if err := r.ToJSON(&contratos_ordenador_dependencia); err == nil {
	} else {

		fmt.Println(err)
	}

	return contratos_ordenador_dependencia
}
