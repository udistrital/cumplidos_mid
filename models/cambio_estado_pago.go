package models

import "time"

type CambioEstadoPago struct {
	Id                     int
	EstadoPagoMensualId    int
	DocumentoResponsableId string
	CargoResponsable       string
	PagoMensualId          *PagoMensual
	Activo                 bool
	FechaCreacion          time.Time
	FechaModificacion      time.Time
}
