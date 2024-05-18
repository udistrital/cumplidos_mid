package models

type DocumentosCumplidos struct {
	Enlance string `json:"Enlace"`
	Nombre  string `json:"Nombre"`
	Nuxeo   struct {
		Thumb struct {
			Name string `json:"name"`
			Data string `json:"data"`
		} `json:"thumb:thumbnail"`
		Content struct {
			Name string `json:"name"`
			Data string `json:"data"`
		} `json:"file:content"`
		File string `json:"file"`
	} `json:"Nuxeo"`
}
