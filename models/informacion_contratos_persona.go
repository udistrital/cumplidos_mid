package models

type InformacionContratosPersona struct {
	ContratosPersonas struct {
		ContratoPersona []struct {
			EstadoContrato struct {
				Id     string `json:"id"`
				Nombre string `json:"nombre"`
			} `json:"estado_contrato"`
			NumeroContrato string `json:"numero_contrato"`
			TipoContrato   struct {
				Id     string `json:"id"`
				Nombre string `json:"nombre"`
			} `json:"tipo_contrato"`
			Vigencia string `json:"vigencia"`
		} `json:"contrato_persona"`
	} `json:"contratos_personas"`
}
