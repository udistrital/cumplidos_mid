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

    beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:GetSolicitudesOrdenadorContratistasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cumplidos_mid/controllers:GetSolicitudesOrdenadorContratistasController"],
        beego.ControllerComments{
            Method: "GetSolicitudesOrdenadorContratistas",
            Router: `/solicitudes_ordenador_contratistas/:docordenador`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
