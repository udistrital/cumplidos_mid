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
	ValorContrato string
	FechaCPS      string
	Dependencia   string
	Sede          string
	Objeto        string
	CDP           struct {
		Consecutivo string
		Fecha       time.Time
	}
	RP struct {
		Consecutivo string
		Fecha       time.Time
	}
	Novedades struct {
		Otrosi []Otrosi
		Cesion []Cesion
	}
	InformacionContratista struct {
		Nombre             string
		TipoIdentificacion string
		CiudadExpedicion   string
	}
	ActividadesEspecificas string
}

type Otrosi struct {
	FechaAdiccion  string
	FechaProrroga  string
	TiempoProrroga interface{}
}

type Cesion struct {
	FechaCesion string
}
