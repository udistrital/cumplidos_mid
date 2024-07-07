package helpers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
	"github.com/udistrital/utils_oas/request"
)

func GetNovedadesPostcontractuales(tipo_novedad models.TipoNovedad, query string, sortby string, order string, limit string, offset string, fields string, target interface{}) (status int, err_nov error) {
	url_base := beego.AppConfig.String("UrlcrudAgora") + "/novedad_postcontractual/?"
	var peticion string

	peticion = url_base
	if tipo_novedad != models.TipoNovedadTodas {
		peticion += "query=TipoNovedad:" + tipo_novedad.String()
	}
	if query != "" {
		if tipo_novedad != models.TipoNovedadTodas {
			peticion += "," + query
		} else {
			peticion += "query=" + query
		}
	}

	if sortby != "" {
		peticion += "&sortby=" + sortby
	}

	if order != "" {
		peticion += "&order=" + order
	}

	if limit != "" {
		peticion += "&limit=" + limit
	}

	if offset != "" {
		peticion += "&offset=" + offset
	}

	if fields != "" {
		peticion += "&fields" + fields
	}

	fmt.Println(peticion)
	if response, err := request.GetJsonTest2(peticion, target); (err == nil) && (response == 200) {
		return 200, nil
	} else {
		err_nov = err
	}
	return 400, err_nov
}
