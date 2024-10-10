package models

type NovedadPostcontractual struct {
	Id                int
	NumeroSolicitud   string
	ContratoId        int
	NumeroCdpId       int
	Motivo            string
	Aclaracion        string
	Observacion       string
	Vigencia          int
	VigenciaCdp       int
	FechaCreacion     string
	FechaModificacion string
	Activo            bool
	TipoNovedad       int
	OficioSupervisor  string
	OficioOrdenador   string
	Estado            string
	EnlaceDocumento   string
}

type Fecha struct {
	Fecha                       string
	Activo                      bool
	FechaCreacion               string
	FechaModificacion           string
	IdTipoFecha                 interface{}
	IdNovedadesPoscontractuales interface{}
}

type Propiedad struct {
	Propiedad                   int
	Activo                      bool
	FechaCreacion               string
	FechaModificacion           string
	IdTipoPropiedad             interface{}
	IdNovedadesPoscontractuales interface{}
}

type Novedades struct {
	Novedades []Noveda `JSON:"novedades"`
}

type Noveda struct {
	TipoNovedad    string `json:"TipoNovedad"`
	FechaCreacion  string `json:"FechaCreacion"`
	FechaInicio    string `json:"FechaInicio"`
	FechaFin       string `json:"FechaFin"`
	FechaFinSus    string `json:"FechaFinSus"`
	PlazoEjecucion string `json:"PlazoEjecucion"`
	NumeroCdp      string `json:"NumeroCdp"`
	VigenciaCdp    string `json:"VigenciaCdp"`
	Cedente        string `json:"Cedente"`
	Cesionario     string `json:"Cesionario"`
}
