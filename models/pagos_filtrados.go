package models

type PagosFiltrados struct {
	NombreDependencia    string
	Rubro                string
	DocumentoContratista string
	NombreContratista    string
	Vigencia             int
	Ano                  int
	Mes                  int
	Estado               *EstadoPagoMensual
}
