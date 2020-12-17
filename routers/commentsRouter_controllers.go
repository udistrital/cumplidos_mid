package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:ContratosContratistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:ContratosContratistaController"],
		beego.ControllerComments{
			Method:           "GetContratosContratista",
			Router:           `/:numero_documento`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:CertificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:CertificacionController"],
		beego.ControllerComments{
			Method:           "CertificacionVistoBueno",
			Router:           `/certificacion_visto_bueno/:dependencia/:mes/:anio`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:CertificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:CertificacionController"],
		beego.ControllerComments{
			Method:           "GetCertificacionDocumentosAprobados",
			Router:           "/documentos_aprobados/:dependencia/:mes/:anio",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesSupervisorContratistasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesSupervisorContratistasController"],
		beego.ControllerComments{
			Method:           "GetSolicitudesSupervisorContratistas",
			Router:           `/:docsupervisor`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:AprobacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:AprobacionPagoController"],
		beego.ControllerComments{
			Method:           "GetContratosContratista",
			Router:           "/contratos_contratista/:numero_documento",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesCoordinadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesCoordinadorController"],
		beego.ControllerComments{
			Method:           "GetSolicitudesCoordinador",
			Router:           "/:doccoordinador",
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
			Method:           "AprobarMultiplesPagosContratistas",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"],
		beego.ControllerComments{
			Method:           "CertificacionCumplidosContratistas",
			Router:           "/certificaciones/:dependencia/:mes/:anio",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasController"],
		beego.ControllerComments{
			Method:           "GetSolicitudesOrdenadorContratistasDependencia",
			Router:           `/solicitudes_ordenador_contratistas_dependencia/:docordenador/:cod_dependencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorController"],
		beego.ControllerComments{
			Method:           "GetSolicitudesOrdenador",
			Router:           "/solicitudes/:docordenador",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorController"],
		beego.ControllerComments{
			Method:           "ObtenerDependenciaOrdenador",
			Router:           "/dependencia_ordenador/:docordenador",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorController"],
		beego.ControllerComments{
			Method:           "ObtenerInfoOrdenador",
			Router:           "/informacion_ordenador/:numero_contrato/:vigencia",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
