package models

type OrdenesPagoTercero struct {
	OrdenesPago struct {
		Tercero []struct {
			Vigencia        string `json:"vigencia"`
			Ano             string `json:"ano"`
			NumeroDocumento string `json:"numero_documento"`
			Mes             string `json:"mes"`
			Registro        string `json:"registro"`
			Detalle         string `json:"detalle"`
		} `json:"tercero"`
	} `json:"ordenes_pago"`
}
