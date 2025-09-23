package models

import "time"

type InformacionContratosPersona struct {
	ContratosPersonas struct {
		ContratoPersona []struct {
			NumeroContrato string    `json:"numero_contrato"`
			Vigencia       string    `json:"vigencia"`
			Vigencia_cdp   string    `json:"vigencia_cdp"`
			NumeroCDP      string    `json:"cdp"`
			FechaInicio    time.Time `json:"fecha_inicio"`
			FechaFin       time.Time `json:"fecha_fin"`
		} `json:"contrato_persona"`
	} `json:"contratos_personas"`
}
