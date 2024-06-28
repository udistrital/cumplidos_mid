package helpers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

func GetNovedadesPostcontractuales(tipo_novedad models.TipoNovedad, query string, sortby string, order string, limit string, offset string, fields string, target interface{}) (status int, err_nov error) {
	url_base := beego.AppConfig.String("UrlNovedadesCrud") + "/novedades_poscontractuales/?"
	var peticion string

	peticion = url_base
	if tipo_novedad != models.TipoNovedadTodas {
		peticion += "query=" + tipo_novedad.String()
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
	if response, err := getJsonTest(peticion, target); (err == nil) && (response == 200) {
		return 200, nil
	} else {
		err_nov = err
	}
	return 400, err_nov
}

func ConstruirNovedadOtroSi(id string, cdp string, novedadRes *models.Novedades) {
	var fechaTemp []models.Fecha
	novedaux := models.Noveda{}
	url := beego.AppConfig.String("UrlNovedadesCrud") + "/fechas/?query=IdNovedadesPoscontractuales.Id:" + id
	if status, err := getJsonTest(url, &fechaTemp); err == nil && status == 200 {
		for _, f := range fechaTemp {
			tipoInterface := f.IdTipoFecha.(map[string]interface{})
			idTipo := tipoInterface["Id"].(float64)
			switch idTipo {
			case 5:
				novedaux.FechaCreacion = f.Fecha[:10]
			case 1:
				novedaux.FechaInicio = f.Fecha[:10]
			case 12:
				novedaux.FechaFin = f.Fecha[:10]
			}
		}
	} else {
		fmt.Println(err)
	}
	novedaux.NumeroCdp = cdp
	novedadRes.Novedades = append(novedadRes.Novedades, novedaux)
}

func ConstruirNovedadCesion(id string, novedadRes *models.Novedades) {
	var fechaTemp []models.Fecha
	var propiedadTemp []models.Propiedad
	novedaux := models.Noveda{}
	url := beego.AppConfig.String("UrlNovedadesCrud") + "/fechas/?query=IdTipoFecha.Id:2,IdNovedadesPoscontractuales.Id:" + id
	if status, err := getJsonTest(url, &fechaTemp); err == nil && status == 200 {
		novedaux.FechaInicio = fechaTemp[0].Fecha[:10]
	} else {
		fmt.Println(err)
	}
	url = beego.AppConfig.String("UrlNovedadesCrud") + "/propiedad/?query=IdNovedadesPoscontractuales.Id:" + id
	if status, err := getJsonTest(url, &propiedadTemp); err == nil && status == 200 {
		for _, p := range propiedadTemp {
			tipoInterface := p.IdTipoPropiedad.(map[string]interface{})
			idTipo := tipoInterface["Id"].(float64)
			if idTipo == 1 {
				novedaux.Cedente = strconv.Itoa(p.Propiedad)
			}
			if idTipo == 2 {
				novedaux.Cesionario = strconv.Itoa(p.Propiedad)
			}
		}
	} else {
		fmt.Println(err)
	}
	novedadRes.Novedades = append(novedadRes.Novedades, novedaux)
}

func ConstruirNovedadSuspension(id string, novedadRes *models.Novedades) {
	var fechaTemp []models.Fecha
	novedaux := models.Noveda{}
	url := beego.AppConfig.String("UrlNovedadesCrud") + "/fechas/?query=IdNovedadesPoscontractuales.Id:" + id
	if status, err := getJsonTest(url, &fechaTemp); err == nil && status == 200 {
		for _, f := range fechaTemp {
			tipoInterface := f.IdTipoFecha.(map[string]interface{})
			idTipo := tipoInterface["Id"].(float64)
			switch idTipo {
			case 5:
				novedaux.FechaCreacion = f.Fecha[:10]
			case 8:
				novedaux.FechaInicio = f.Fecha[:10]
			case 11:
				novedaux.FechaFinSus = f.Fecha[:10]
			case 12:
				novedaux.FechaFin = f.Fecha[:10]
			}
		}
	} else {
		fmt.Println(err)
	}
	novedadRes.Novedades = append(novedadRes.Novedades, novedaux)
}

func ConstruirNovedadTerminacion(id string, novedadRes *models.Novedades) {

}
