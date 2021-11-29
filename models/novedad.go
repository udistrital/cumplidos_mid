package models

type Novedad struct {
	Contrato       int         `json:"contrato"`
	FechaAdiccion  string      `json:"fechaadicion"`
	FechaCesion    string      `json:"fechacesion"`
	FechaProrroga  string      `json:"fechaprorroga"`
	FechaSolicitud string      `json:"fechasolicitud"`
	TipoNovedad    int         `json:"tiponovedad"`
	Vigencia       int         `json:"vigencia"`
	TiempoProrroga interface{} `json:"tiempoprorroga"`
}

type RespNov struct {
	Type string
	Code string
	Body []Novedad
}
