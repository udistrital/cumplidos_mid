// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/cumplidos_mid/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/contratos_contratista",
			beego.NSInclude(
				&controllers.ContratosContratistaController{},
			),
		),

		beego.NSNamespace("/solicitudes_supervisor_contratistas",
			beego.NSInclude(
				&controllers.SolicitudesSupervisorContratistasController{},
			),
		),

		beego.NSNamespace("/solicitudes_ordenador_contratistas",
			beego.NSInclude(
				&controllers.SolicitudesOrdenadorContratistasController{},
			),
		),

		beego.NSNamespace("/informacion_informe",
			beego.NSInclude(
				&controllers.InformacionInformeController{},
			),
		),
		beego.NSNamespace("/informe",
			beego.NSInclude(
				&controllers.InformeController{},
			),
		),
		beego.NSNamespace("/validacion_periodo_carga_cumplido",
			beego.NSInclude(
				&controllers.ValidacionFechaCargaCumplidoController{},
			),
		),
		beego.NSNamespace("/solicitudes_pagos",
			beego.NSInclude(
				&controllers.SolicitudesPagoMensualController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
