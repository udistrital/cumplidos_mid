package helpers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

func InformacionInforme(num_documento string, contrato string, vigencia string) (informacion_informe models.InformacionInforme, outputError map[string]interface{}) {

	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("error", err)
			outputError = map[string]interface{}{"funcion": "/ContratosContratista", "err": err, "status": "502"}
			panic(outputError)
		}
	}()

	var informacion_proveedores []models.InformacionProveedor
	//var informacion_informe []models.InformacionInforme
	if response, err := getJsonTest(beego.AppConfig.String("UrlcrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+num_documento, &informacion_proveedores); (err == nil) && (response == 200) {
		fmt.Println("informacion_proveedor:", informacion_proveedores)
		for x, persona := range informacion_proveedores {
			fmt.Println("Persona "+strconv.Itoa(x)+":", persona)
			informacion_informe.NombreContratista = persona.NomProveedor
			informacion_informe.Sede = persona.NomProveedor
			informacion_informe.Dependencia = persona.NomProveedor
			informacion_informe.TipoIdentificacion = persona.NomProveedor
			informacion_informe.NumeroDocumento = persona.NomProveedor
		}
	} else {
		logs.Error(err)
		outputError = map[string]interface{}{"funcion": "/GetContrato", "err": err, "status": "502"}
		return nil, outputError
	}

	return
}
