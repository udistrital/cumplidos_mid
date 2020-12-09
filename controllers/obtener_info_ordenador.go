package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	_ "time"

	//"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/models"
	_ "github.com/udistrital/cumplidos_mid/helpers"
)

// ObtenerInfoOrdenadorController operations for obtener_info_ordenador
type ObtenerInfoOrdenadorController struct {
	beego.Controller
}

// URLMapping ...
func (c *ObtenerInfoOrdenadorController) URLMapping() {
	c.Mapping("ObtenerInfoOrdenador", c.ObtenerInfoOrdenador)
}

// AprobacionPagoController ...
// @Title ObtenerInfoOrdenador
// @Description create ObtenerInfoOrdenador
// @Param numero_contrato path string true "Numero de contrato en la tabla contrato general"
// @Param vigencia path int true "Vigencia del contrato en la tabla contrato general"
// @Success 201 {int} models.InformacionOrdenador
// @Failure 403 :numero_contrato is empty
// @Failure 403 :vigencia is empty
// @router /informacion_ordenador/:numero_contrato/:vigencia [get]
func (c *ObtenerInfoOrdenadorController) ObtenerInfoOrdenador() {
	numero_contrato := c.GetString(":numero_contrato")
	vigencia := c.GetString(":vigencia")

	
	if informacion_ordenador, err:= traer_info_ordenador(numero_contrato, vigencia,);err!=nil{
		logs.Error(err)
		c.Data["mesaage"] = "Error service Get SolicitudesOrdenadorContratistaDependencia: The request contains an incorrect parameter or no record exists"
		c.Abort("404")
	}else{
		c.Data["json"] = informacion_ordenador
	}
	c.ServeJSON()
}

func traer_info_ordenador(numero_contrato string, vigencia string)(informacion_ordenador models.InformacionOrdenador, err error){
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
