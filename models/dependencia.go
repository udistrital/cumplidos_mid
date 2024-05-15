package models

import "encoding/xml"

type Dependencia struct {
	Id                  int
	Nombre              string
	TelefonoDependencia string
	CorreoElectronico   string
}

type DependenciasXmln struct {
	XMLName      xml.Name          `xml:"DependenciasSic"`
	Dependencias []DependenciaXmln `xml:"Dependencia"`
}

type DependenciaXmln struct {
	EsfCodigoDep    string `xml:"ESFCODIGODEP"`
	EsfDepEncargada string `xml:"ESFDEPENCARGADA"`
}

type DependenciaSimple struct {
	Id     int
	Nombre string
}
