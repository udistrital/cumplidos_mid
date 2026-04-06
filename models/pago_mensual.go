package models

import (
	"time"
)

type PagoMensual struct {
	Id                     int                `json:"Id"`
	NumeroContrato         string             `json:"NumeroContrato"`
	VigenciaContrato       float64            `json:"VigenciaContrato"`
	NumeroCDP              string             `json:"NumeroCDP"`
	VigenciaCDP            float64            `json:"VigenciaCDP"`
	Mes                    float64            `json:"Mes"`
	DocumentoPersonaId     string             `json:"DocumentoPersonaId"`
	EstadoPagoMensualId    *EstadoPagoMensual `json:"EstadoPagoMensualId"`
	DocumentoResponsableId string             `json:"DocumentoResponsableId"`
	CargoResponsable       string             `json:"CargoResponsable"`
	Ano                    float64            `json:"Ano"`
	Activo                 bool               `json:"Activo"`
	FechaCreacion          time.Time          `json:"FechaCreacion"`
	FechaModificacion      time.Time          `json:"FechaModificacion"`
}
