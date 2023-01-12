package models

import (
	"time"
)

type PreliquidacionTitan struct {
	Contrato                  string
	Vigencia                  int
	TotalDevengado            float64
	TotalDevengadoConFormato  string
	TotalDescuentos           float64
	TotalDescuentosConFormato string
	TotalPago                 float64
	TotalPagoConFormato       string
	Detalle                   []DetallePreliquidacion
}

type DetallePreliquidacion struct {
	Id                       int
	ContratoPreliquidacionId *ContratoPreliquidacion
	ValorCalculado           float64
	ValorCalculadoConFormato string
	DiasLiquidados           float64
	DiasEspecificos          string
	TipoPreliquidacionId     int
	ConceptoNominaId         *ConceptoNomina
	EstadoDisponibilidadId   int
	Activo                   bool
	FechaCreacion            string
	FechaModificacion        string
	NombreCompleto           string
	Documento                string
	NumeroContrato           string
	VigenciaContrato         int
	Persona                  int
}

type ContratoPreliquidacion struct {
	Id                   int
	ContratoId           *Contrato
	PreliquidacionId     *Preliquidacion
	Cumplido             bool
	Preliquidado         bool
	ResponsableIva       bool
	Dependientes         bool
	Pensionado           bool
	InteresesVivienda    float64
	MedicinaPrepagadaUvt float64
	PensionVoluntaria    float64
	Afc                  float64
	Activo               bool
	FechaCreacion        string
	FechaModificacion    string
}

type ConceptoNomina struct {
	Id                         int
	NombreConcepto             string
	AliasConcepto              string
	NaturalezaConceptoNominaId int
	TipoConceptoNominaId       int
	EstadoConceptoNominaId     int
	Activo                     bool
	FechaCreacion              string
	FechaModificacion          string
}

type Contrato struct {
	Id                int
	NumeroContrato    string
	Vigencia          int
	NombreCompleto    string
	Documento         string
	PersonaId         int
	TipoNominaId      int
	FechaInicio       time.Time
	FechaFin          time.Time
	ValorContrato     float64
	Vacaciones        float64
	DependenciaId     int
	ProyectoId        int
	Cdp               int
	Rp                int
	Unico             bool
	Completo          bool
	Activo            bool
	FechaCreacion     string
	FechaModificacion string
}

type Preliquidacion struct {
	Id                     int
	Descripcion            string
	Mes                    int
	Ano                    int
	EstadoPreliquidacionId int
	NominaId               int
	Activo                 bool
	FechaCreacion          string
	FechaModificacion      string
}
