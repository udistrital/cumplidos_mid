package models

import "time"

type ActividadEspecifica struct {
	Id                    int
	Activo                bool
	FechaCreacion         time.Time
	FechaModificacion     time.Time
	ActividadEspecifica   string
	Avance                int
	ActividadesRealizadas []ActividadRealizada
}
