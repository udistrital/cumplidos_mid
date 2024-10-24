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

type NovedadPoscontractual struct {
	Id                         int
	Aclaracion                 string
	Contrato                   int
	Cedente                    int
	Cesionario                 int
	CodAbreviacionTipo         string
	Vigencia                   int
	Enlace                     string
	EntidadAseguradora         int
	Estado                     string
	FechaAdicion               string
	FechaCesion                string
	FechaExpedicion            string
	FechaFinefectiva           string
	FechaOficio                string
	FechaLiquidacion           string
	FechaProrroga              string
	FechaRegistro              string
	FechaReinicio              string
	FechaSolicitud             string
	FechaSuspension            string
	FechaFinSuspension         string
	FechaTerminacionanticipada string
	Motivo                     string
	PeriodoSuspension          int
	PlazoActual                int
	TiempoProrroga             int
	ValorAdicion               int
	ValorFinalContrato         int
	NombreEstado               string
	NombreTipoNovedad          string
	NumeroActaEntrega          string
	NumeroCdp                  int
	VigenciaCdp                int
	NumeroOficioordenador      string
	NumeroOficiosupervisor     string
	NumeroSolicitud            string
	Observacion                string
	Poliza                     string
	TipoNovedad                int
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
	NumeroContrato     int    `json:"NumeroContrato"`
	Vigencia           int    `json:"Vigencia"`
	TipoNovedad        string `json:"TipoNovedad"`
	FechaCreacion      string `json:"FechaCreacion"`
	FechaInicio        string `json:"FechaInicio"`
	FechaFin           string `json:"FechaFin"`
	FechaFinSus        string `json:"FechaFinSus"`
	PlazoEjecucion     int    `json:"PlazoEjecucion"`
	NumeroCdp          int    `json:"NumeroCdp"`
	VigenciaCdp        int    `json:"VigenciaCdp"`
	Cedente            string `json:"Cedente"`
	Cesionario         string `json:"Cesionario"`
	ValorNovedadPagado int    `json:"ValorNovedadPagado"`
}
