package models

import (
	"encoding/json"
	"time"
)

type EstadoPagoMensualReferencia struct {
	Id                int       `json:"Id"`
	Nombre            string    `json:"Nombre"`
	Descripcion       string    `json:"Descripcion"`
	CodigoAbreviacion string    `json:"CodigoAbreviacion"`
	NumeroOrden       int       `json:"NumeroOrden"`
	Activo            bool      `json:"Activo"`
	FechaCreacion     time.Time `json:"FechaCreacion"`
	FechaModificacion time.Time `json:"FechaModificacion"`
}

func (e *EstadoPagoMensualReferencia) UnmarshalJSON(data []byte) error {
	var id int
	if err := json.Unmarshal(data, &id); err == nil {
		e.Id = id
		return nil
	}

	type estadoPagoMensualReferencia EstadoPagoMensualReferencia
	var estado estadoPagoMensualReferencia
	if err := json.Unmarshal(data, &estado); err != nil {
		return err
	}

	*e = EstadoPagoMensualReferencia(estado)
	return nil
}

type CambioEstadoPago struct {
	Id                     int                         `json:"Id"`
	EstadoPagoMensualId    EstadoPagoMensualReferencia `json:"EstadoPagoMensualId"`
	DocumentoResponsableId string                      `json:"DocumentoResponsableId"`
	CargoResponsable       string                      `json:"CargoResponsable"`
	PagoMensualId          struct {
		Id                     int         `json:"Id"`
		NumeroContrato         string      `json:"NumeroContrato"`
		VigenciaContrato       int         `json:"VigenciaContrato"`
		Mes                    int         `json:"Mes"`
		DocumentoPersonaId     string      `json:"DocumentoPersonaId"`
		EstadoPagoMensualId    interface{} `json:"EstadoPagoMensualId"`
		DocumentoResponsableId string      `json:"DocumentoResponsableId"`
		CargoResponsable       string      `json:"CargoResponsable"`
		Ano                    int         `json:"Ano"`
		Activo                 bool        `json:"Activo"`
		FechaCreacion          time.Time   `json:"FechaCreacion"`
		FechaModificacion      time.Time   `json:"FechaModificacion"`
		NumeroCDP              string      `json:"NumeroCDP"`
		VigenciaCDP            int         `json:"VigenciaCDP"`
	} `json:"PagoMensualId"`
	Activo            bool      `json:"Activo"`
	FechaCreacion     time.Time `json:"FechaCreacion"`
	FechaModificacion time.Time `json:"FechaModificacion"`
	NombreEstado      string
	DescripcionEstado string
	NombreResponsable string
}
