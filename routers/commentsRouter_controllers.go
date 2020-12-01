package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:AprobacionPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:AprobacionPagoController"],
        beego.ControllerComments{
            Method: "GetContratosContratista",
            Router: "/contratos_contratista/:numero_documento",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})
    
    beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesCoordinadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:SolicitudesCoordinadorController"],
        beego.ControllerComments{
            Method: "GetSolicitudesCoordinador",
            Router: "/solicitudes_coordinador/:doccoordinador",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})
}
