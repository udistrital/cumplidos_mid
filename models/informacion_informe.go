package models

import "time"

//import (
//	"time"
//)

type InformacionInforme struct {
	Supervisor struct {
		Cargo  string
		Nombre string
	}
	ValorContrato      int
	ValorTotalContrato int
	FechaCPS           string
	FechaInicio        time.Time
	FechaFin           time.Time
	Dependencia        string
	Sede               string
	Objeto             string
	CDP                struct {
		Consecutivo string
		Fecha       time.Time
	}
	RP struct {
		Consecutivo string
		Fecha       time.Time
	}
	Novedades struct {
		Otrosi      []Otrosi
		Cesion      []Cesion
		Terminacion []Terminacion
		Suspencion  []Suspencion
	}
	ArrayNovedades         []Noveda
	InformacionContratista struct {
		Nombre             string
		TipoIdentificacion string
		CiudadExpedicion   string
	}
	ActividadesEspecificas string
	EjecutadoDinero        struct {
		Pagado   int
		Faltante int
	}
	FechasConNovedades FechasConNovedades
}

type Otrosi struct {
	FechaCreacion      string
	ValorNovedad       int
	ValorNovedadPagado int
	NumeroCdpId        int
	VigenciaCdp        int
	TipoNovedad        string
}

type Cesion struct {
	FechaInicio string
}

type Terminacion struct {
	FechaFin string
}

type Suspencion struct {
	FechaInicio    string
	FechaFin       string
	PlazoEjecucion int
}

type FechasConNovedades struct {
	FechaInicio time.Time
	FechaFin    time.Time
}
