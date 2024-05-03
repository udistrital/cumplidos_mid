package models

import "time"

type CambioEstadoPago struct {
	Id                     int
	EstadoPagoMensualId    int
	DocumentoResponsableId string
	CargoResponsable       string
	PagoMensualId          string
	FechaCreacion          time.Time
	NombreEstado           string
	DescripcionEstado      string
}
