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
	TipoNovedad    string `json:"tipoNovedad"`
	FechaCreacion  string `json:"fechaCreacion"`
	FechaInicio    string `json:"fechaInicio"`
	FechaFin       string `json:"fechaFin"`
	FechaFinSus    string `json:"fechaFinSus"`
	PlazoEjecucion string `json:"plazoEjecucion"`
	NumeroCdp      string `json:"numeroCdp"`
	Cedente        string `json:"cedente"`
	Cesionario     string `json:"cesionario"`
}
