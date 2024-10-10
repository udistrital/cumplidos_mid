package helpers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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

	fmt.Println("PETICION ", peticion)
	if response, err := getJsonTest(peticion, target); (err == nil) && (response == 200) {
		return 200, nil
	} else {
		err_nov = err
	}
	return 400, err_nov
}

func ConstruirNovedadOtroSi(id string, cdp string, vigenciaCdp string, novedadRes []models.Noveda) (novedaux models.Noveda, outputError map[string]interface{}) {
	var fechaTemp []models.Fecha
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
		novedaux.NumeroCdp = cdp
		novedaux.VigenciaCdp = vigenciaCdp
	} else {
		fmt.Println(err)
	}
	novedaux.TipoNovedad = "NP_ADPRO"
	return novedaux, outputError
}

func ConstruirNovedadCesion(id string, novedadRes []models.Noveda) (novedaux models.Noveda, outputError map[string]interface{}) {
	var fechaTemp []models.Fecha
	var propiedadTemp []models.Propiedad
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
				var informacion_proveedor []models.InformacionProveedor
				if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor?query=Id:"+strconv.Itoa(p.Propiedad), &informacion_proveedor); (err == nil) && (response == 200) {
					novedaux.Cedente = informacion_proveedor[0].NomProveedor
				} else {
					logs.Error(err)
					outputError = map[string]interface{}{"funcion": "/ConstruirNovedadCesion/informacion_proveedor", "err": err, "status": "502"}
					panic(outputError)
				}
			}
			if idTipo == 2 {
				var informacion_proveedor []models.InformacionProveedor
				if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor?query=Id:"+strconv.Itoa(p.Propiedad), &informacion_proveedor); (err == nil) && (response == 200) {
					novedaux.Cesionario = informacion_proveedor[0].NomProveedor
				} else {
					logs.Error(err)
					outputError = map[string]interface{}{"funcion": "/ConstruirNovedadCesion/informacion_proveedor", "err": err, "status": "502"}
					panic(outputError)
				}
			}
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/ConstruirNovedadCesion/propiedad", "err": err, "status": "502"}
		panic(outputError)
	}
	novedaux.TipoNovedad = "NP_CES"
	return novedaux, outputError
}

func ConstruirNovedadSuspension(id string, novedadRes []models.Noveda) (novedaux models.Noveda, outputError map[string]interface{}) {
	var fechaTemp []models.Fecha
	var propiedadTemp []models.Propiedad
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
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/ConstruirNovedadCesion/propiedad", "err": err, "status": "502"}
		panic(outputError)
	}
	url = beego.AppConfig.String("UrlNovedadesCrud") + "/propiedad/?query=IdTipoPropiedad.Id:3,IdNovedadesPoscontractuales.Id:" + id
	if status, err := getJsonTest(url, &propiedadTemp); err == nil && status == 200 {
		novedaux.PlazoEjecucion = strconv.Itoa(propiedadTemp[0].Propiedad)
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/ConstruirNovedadCesion/propiedad", "err": err, "status": "502"}
		panic(outputError)
	}
	novedaux.TipoNovedad = "NP_SUS"
	return novedaux, outputError
}

func ConstruirNovedadTerminacion(id string, novedadRes []models.Noveda) (novedaux models.Noveda, outputError map[string]interface{}) {
	var fechaTemp []models.Fecha
	url := beego.AppConfig.String("UrlNovedadesCrud") + "/fechas/?query=IdNovedadesPoscontractuales.Id:" + id
	if status, err := getJsonTest(url, &fechaTemp); err == nil && status == 200 {
		for _, f := range fechaTemp {
			tipoInterface := f.IdTipoFecha.(map[string]interface{})
			idTipo := tipoInterface["Id"].(float64)
			switch idTipo {
			case 12:
				novedaux.FechaFin = f.Fecha[:10]
			}
		}
	} else {
		fmt.Println(err)
	}
	novedaux.TipoNovedad = "NP_TER"
	return novedaux, outputError
}
