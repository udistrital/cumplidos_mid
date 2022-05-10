package models

import "time"

type ValidacionFechaCargaCumplido struct {
	CargaHabilitada bool
	FechaActual     time.Time
	Periodo         struct {
		Inicio time.Time
		Fin    time.Time
	}
}

type FechasCargaCumplidos struct {
	Id                  int
	DocumentoSupervisor float64
	Activo              bool
	FechaCreacion       time.Time
	FechaModificacion   time.Time
	FechaInicio         time.Time
	FechaFin            time.Time
	Anio                float64
	Mes                 float64
	Dependencia         string
}
