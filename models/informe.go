package models

import "time"

type Informe struct {
	Id                     int
	Activo                 bool
	FechaCreacion          time.Time
	FechaModificacion      time.Time
	PeriodoInformeInicio   time.Time
	PeriodoInformeFin      time.Time
	Proceso                string
	PagoMensualId          *PagoMensual
	ActividadesEspecificas []ActividadEspecifica
}
