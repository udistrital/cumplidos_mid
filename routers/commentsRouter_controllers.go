package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:ContratosContratistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:ContratosContratistaController"],
		beego.ControllerComments{
			Method:           "GetContratosContratista",
			Router:           "/:numero_documento",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:ContratosContratistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:ContratosContratistaController"],
		beego.ControllerComments{
			Method:           "GetDocumentosPagoMensual",
			Router:           "/documentos_pago_mensual/:pago_mensual_id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:DescargarDocumentosSolicitudesPagosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:DescargarDocumentosSolicitudesPagosController"],
		beego.ControllerComments{
			Method:           "GetDocumentosPagoMensual",
			Router:           "/:pago_mensual_id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:HistoricoCumplidos"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:HistoricoCumplidos"],
		beego.ControllerComments{
			Method:           "GetCambioEstado",
			Router:           "/cambio-estado/:idPagoMensual",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:HistoricoCumplidos"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:HistoricoCumplidos"],
		beego.ControllerComments{
			Method:           "GetDependencias",
			Router:           "/dependencias/:documento",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:InformacionInformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:InformacionInformeController"],
		beego.ControllerComments{
			Method:           "GetInformacionInforme",
			Router:           "/:pago_mensual_id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:InformacionInformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:InformacionInformeController"],
		beego.ControllerComments{
			Method:           "GetPreliquidacion",
			Router:           "/preliquidacion/:pago_mensual_id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:InformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:InformeController"],
		beego.ControllerComments{
			Method:           "PostInforme",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:InformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:InformeController"],
		beego.ControllerComments{
			Method:           "PutInforme",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:InformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:InformeController"],
		beego.ControllerComments{
			Method:           "GetInforme",
			Router:           "/:pago_mensual_id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:InformeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:InformeController"],
		beego.ControllerComments{
			Method:           "GetUltimoInformeContratista",
			Router:           "/ultimo_informe/:pago_mensual_id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"],
		beego.ControllerComments{
			Method:           "AprobarMultiplesPagosContratistas",
			Router:           "/aprobar_pagos",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"],
		beego.ControllerComments{
			Method:           "CertificacionCumplidosContratistas",
			Router:           "/certificaciones/:dependencia/:mes/:ano",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"],
		beego.ControllerComments{
			Method:           "GetCumplidosRevertiblesPorOrdenador",
			Router:           "/cumplidos_revertibles/:docordenador",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"],
		beego.ControllerComments{
			Method:           "GetInfoOrdanador",
			Router:           "/informacion_ordenador/:numero_contrato/:vigencia",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"],
		beego.ControllerComments{
			Method:           "GetSolicitudesOrdenadorContratistas",
			Router:           "/solicitudes/:docordenador",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"],
		beego.ControllerComments{
			Method:           "GetSolicitudesOrdenadorContratistasDependencia",
			Router:           "/solicitudes_dependencia/:docordenador/:cod_dependencia",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesPagoMensualController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesPagoMensualController"],
		beego.ControllerComments{
			Method:           "GetSolicitudesPagoMensual",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesSupervisorContratistasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesSupervisorContratistasController"],
		beego.ControllerComments{
			Method:           "GetSolicitudesSupervisorContratistas",
			Router:           "/:docsupervisor",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:ValidacionFechaCargaCumplidoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:ValidacionFechaCargaCumplidoController"],
		beego.ControllerComments{
			Method:           "GetValidacionPeriodo",
			Router:           "/:dependencia_supervisor/:anio/:mes",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
