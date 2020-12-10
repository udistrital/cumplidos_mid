package controllers

import (

	//"net/http"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
)

// ContratosContratistaController operations for contratos_contratista
type ContratosContratistaController struct {
	beego.Controller
}

//URLMapping ...
func (c *ContratosContratistaController) URLMapping() {
	// c.Mapping("ObtenerInfoCoordinador", c.ObtenerInfoCoordinador)
	c.Mapping("GetContratosContratista", c.GetContratosContratista)
	// c.Mapping("ObtenerInfoOrdenador", c.ObtenerInfoOrdenador)
	// c.Mapping("PagoAprobado", c.PagoAprobado)
	// c.Mapping("CertificacionVistoBueno", c.CertificacionVistoBueno)
	// c.Mapping("CertificacionDocumentosAprobados", c.CertificacionDocumentosAprobados)
	// c.Mapping("ObtenerDependenciaOrdenador", c.ObtenerDependenciaOrdenador)

}

// GetContratosContratista ...
// @Title GetContratosContratista
// @Description create ContratosContratista
// @Param numero_documento path string true "Número documento"
// @Success 200 {object} []models.ContratoDisponibilidadRp
// @Failure 404 not found resource
// @router /contratos_contratista/:numero_documento [get]
func (c *ContratosContratistaController) GetContratosContratista() {
	numero_documento := c.GetString(":numero_documento")

	if contratos_disponibilidad_rp, err := helpers.ContratosContratista(numero_documento); err != nil || len(contratos_disponibilidad_rp) == 0 {
		logs.Error(err)
		c.Data["mesaage"] = map[string]interface{}{"Function": "FuncionalidadMidController:getUserAgora", "Error": err}
		c.Abort("404")
	} else {
		c.Data["json"] = contratos_disponibilidad_rp
	}
	c.ServeJSON()

}
