package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:ContratosContratistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:ContratosContratistaController"],
        beego.ControllerComments{
            Method: "GetContratosContratista",
            Router: `/contratos_contratista/:numero_documento`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:CertificacionVistoBuenoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:CertificacionVistoBuenoController"],
        beego.ControllerComments{
            Method: "CertificacionVistoBueno",
            Router: `/certificacion_visto_bueno/:dependencia/:mes/:anio`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesOrdenadorContratistasDependenciaController"],
        beego.ControllerComments{
            Method: "GetSolicitudesOrdenadorContratistasDependencia",
            Router: `/solicitudes_ordenador_contratistas_dependencia/:docordenador/:cod_dependencia`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesSupervisorContratistasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesSupervisorContratistasController"],
        beego.ControllerComments{
            Method: "GetSolicitudesSupervisorContratistas",
            Router: `/solicitudes_supervisor_contratistas/:docsupervisor`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:ObtenerInfoOrdenadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:ObtenerInfoOrdenadorController"],
        beego.ControllerComments{
            Method: "ObtenerInfoOrdenador",
            Router: `/informacion_ordenador/:numero_contrato/:vigencia`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
