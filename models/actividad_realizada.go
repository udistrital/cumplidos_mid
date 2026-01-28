package models

import "time"

type ActividadRealizada struct {
	Id                int
	Activo            bool
	FechaCreacion     time.Time
	FechaModificacion time.Time
	Actividad         string
	ProductoAsociado  string
	Evidencia         string
}
