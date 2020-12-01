package models

import (
	"time"
)

type PagoMensual struct {
	Id                     int                
	NumeroContrato         string             
	VigenciaContrato       float64            
	Mes                    float64            
	DocumentoPersonaId     string             
	EstadoPagoMensualId    *EstadoPagoMensual 
	DocumentoResponsableId string        
	CargoResponsable       string          
	Ano                    float64     
	Activo                 bool              
	FechaCreacion          time.Time        
	FechaModificacion      time.Time         
}
