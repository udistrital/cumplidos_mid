package models

import "time"

type ContratoDisponibilidadRp struct {
	NumeroContratoSuscrito string
	Vigencia               string
	NumeroCdp              string
	VigenciaCdp            string
	NumeroRp               string
	VigenciaRp             string
	NombreDependencia      string
	NumDocumentoSupervisor string
	FechaInicio            time.Time
	FechaFin               time.Time
}
