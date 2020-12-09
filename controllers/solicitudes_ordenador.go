package controllers

import (
	_ "encoding/json"
	"fmt"
	"strconv"
	_ "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
	"github.com/udistrital/cumplidos_mid/models"
)

// SolicitudesOrdenadorContratistasController operations for SolicitudesOrdenadorContratistas
type SolicitudesOrdenadorController struct {
	beego.Controller
}

//URLMapping ...
func (c *SolicitudesOrdenadorController) URLMapping() {
	c.Mapping("GetSolicitudesOrdenador", c.GetSolicitudesOrdenador)
	c.Mapping("GetSolicitudesOrdenadorContratistas", c.ObtenerDependenciaOrdenador)
	//c.Mapping("AprobarMultiplesPagosContratistas", c.AprobarMultiplesPagosContratistas)
}

// AprobacionPagoController ...
// @Title GetSolicitudesOrdenador
// @Description create GetSolicitudesOrdenador
// @Param docordenador path string true "Número del documento del ordenador"
// @Success 200 {object} []models.PagoPersonaProyecto
// @Failure 403 :docordenador is empty
// @router /:docordenador [get]
func (c *SolicitudesOrdenadorController) GetSolicitudesOrdenador() {

	doc_ordenador := c.GetString(":docordenador")
	//query := c.GetString("query")
	limit, _ := c.GetInt("limit")
	offset, _ := c.GetInt("offset")

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			c.Data["mesaage"] = "Error service Get solicitudes_coordinador: The request contains an incorrect parameter or no record exists"
			c.Abort("404")
		}
	}()

	//var v []models.PagoContratistaCdpRp
	if pagos_personas_proyecto, err := SolicitudesOrdenador(doc_ordenador, limit, offset); err == nil {
		c.Data["json"] = pagos_personas_proyecto
	} else {
		panic(err)
	}

	c.ServeJSON()

}

func SolicitudesOrdenador(doc_ordenador string, limit int, offset int) (pagos_personas_proyecto []models.PagoPersonaProyecto, err error) {

	var pagos_mensuales []models.PagoMensual
	var contratistas []models.InformacionProveedor
	//var pagos_personas_proyecto []models.PagoPersonaProyecto
	var pago_personas_proyecto models.PagoPersonaProyecto
	var vinculaciones_docente []models.VinculacionDocente
	var respuesta_peticion map[string]interface{}
	r := httplib.Get(beego.AppConfig.String("ProtocolCrudCumplidos") + "://" + beego.AppConfig.String("UrlCrudCumplidos") + "/" + beego.AppConfig.String("NsCrudCumplidos") + "/pago_mensual/")
	r.Param("offset", strconv.Itoa(offset))
	r.Param("limit", strconv.Itoa(limit))
	r.Param("query", "EstadoPagoMensualId.CodigoAbreviacion:AD,DocumentoResponsableId:"+doc_ordenador)

	if err := r.ToJSON(&respuesta_peticion); err == nil {
		//if err := r.ToJSON(&pagos_mensuales); err == nil {
		pagos_mensuales = []models.PagoMensual{}
		helpers.LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
		for x, pago_mensual := range pagos_mensuales {

			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAgora")+"/"+beego.AppConfig.String("NscrudAgora")+"/informacion_proveedor/?query=NumDocumento:"+pago_mensual.DocumentoPersonaId, &contratistas); err == nil {

				for _, contratista := range contratistas {
					// se podria armar un metodo desde aca y usar gorutines para reducir el tiempo de carga de las peticiones?
					if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudAdmin")+"/"+beego.AppConfig.String("NscrudAdmin")+"/vinculacion_docente/?limit=-1&query=NumeroContrato:"+pago_mensual.NumeroContrato+",Vigencia:"+strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64), &vinculaciones_docente); err == nil {

						for _, vinculacion := range vinculaciones_docente {
							var dep models.Dependencia

							if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudOikos")+"/"+beego.AppConfig.String("NscrudOikos")+"/dependencia/"+strconv.Itoa(vinculacion.IdProyectoCurricular), &dep); err == nil {

								pago_personas_proyecto.PagoMensual = &pagos_mensuales[x]
								pago_personas_proyecto.NombrePersona = contratista.NomProveedor
								pago_personas_proyecto.Dependencia = &dep
								pagos_personas_proyecto = append(pagos_personas_proyecto, pago_personas_proyecto)

							} else { //If dependencia get

								fmt.Println("Mirenme, me morí en If dependencia get, solucioname!!! ", err)
								return nil, err

							}

						}

					} else { // If vinculacion_docente_get

						fmt.Println("Mirenme, me morí en If vinculacion_docente get, solucioname!!! ", err)
						return nil, err
					}
				}
			} else { //If informacion_proveedor get

				fmt.Println("Mirenme, me morí en If informacion_proveedor get, solucioname!!! ", err)
				return nil, err
			}
		}
	} else { //If pago_mensual get

		fmt.Println("Mirenme, me morí en If pago_mensual get, solucioname!!! ", err)
		return nil, err
	}

	return pagos_personas_proyecto, nil
}

// AprobacionPagoController ...
// @Title ObtenerDependenciaOrdenador
// @Description create ObtenerDependenciaOrdenador
// @Param docordenador path string true "Número del documento del ordenador"
// @Success 200  int
// @Failure 403 :docordenador is empty
// @router /dependencia_ordenador/:docordenador [get]
func (c *SolicitudesOrdenadorController) ObtenerDependenciaOrdenador() {

	doc_ordenador := c.GetString(":docordenador")

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			c.Data["mesaage"] = "Error service Get solicitudes_coordinador: The request contains an incorrect parameter or no record exists"
			c.Abort("404")
		}
	}()

	if dependenciaId, err := DependenciaOrdenador(doc_ordenador); err == nil {
		c.Data["json"] = dependenciaId
	} else {
		panic(err)
	}

	c.ServeJSON()

}

func DependenciaOrdenador(doc_ordenador string) (dependenciaId int, err error) {

	var ordenadores_gasto []models.OrdenadorGasto
	var jefes_dependencia []models.JefeDependencia

	if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/jefe_dependencia/?query=TerceroId:"+doc_ordenador+"&sortby=FechaInicio&order=desc&limit=1", &jefes_dependencia); err == nil {
		for _, jefe := range jefes_dependencia {

			if err := getJson(beego.AppConfig.String("ProtocolAdmin")+"://"+beego.AppConfig.String("UrlcrudCore")+"/"+beego.AppConfig.String("NscrudCore")+"/ordenador_gasto/?query=DependenciaId:"+strconv.Itoa(jefe.DependenciaId), &ordenadores_gasto); err == nil {

				for _, ordenador := range ordenadores_gasto {

					return ordenador.DependenciaId, nil

				}

			} else { // If ordenador_gasto get
				fmt.Println("Mirenme, me morí en If ordenador_gasto get, solucioname!!! ", err)
				return dependenciaId, err
			}

		}

	} else { // If jefe_dependencia get
		fmt.Println("Mirenme, me morí en If jefe_dependencia get, solucioname!!! ", err)
		return dependenciaId, err
	}
	return
}
