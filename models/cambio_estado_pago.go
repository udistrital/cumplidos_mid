package models

import "time"

type CambioEstadoPago struct {
	Id                  int `json:"Id"`
	EstadoPagoMensualId struct {
		Id                int       `json:"Id"`
		Nombre            string    `json:"Nombre"`
		Descripcion       string    `json:"Descripcion"`
		CodigoAbreviacion string    `json:"CodigoAbreviacion"`
		NumeroOrden       int       `json:"NumeroOrden"`
		Activo            bool      `json:"Activo"`
		FechaCreacion     time.Time `json:"FechaCreacion"`
		FechaModificacion time.Time `json:"FechaModificacion"`
	} `json:"EstadoPagoMensualId"`
	DocumentoResponsableId string `json:"DocumentoResponsableId"`
	CargoResponsable       string `json:"CargoResponsable"`
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
