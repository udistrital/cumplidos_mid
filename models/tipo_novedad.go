package models

type TipoNovedad int64

const (
	TipoNovedadOtrosi TipoNovedad = iota
	TipoNovedadCesion
	TipoNovedadTerminacion
	TipoNovedadSuspension
	TipoNovedadTodas
)

func (tN TipoNovedad) String() string {
	switch tN {
	case TipoNovedadOtrosi:
		return "TipoNovedad:8"
	case TipoNovedadCesion:
		return "TipoNovedad:2"
	case TipoNovedadTerminacion:
		return "TipoNovedad:5"
	case TipoNovedadSuspension:
		return "TipoNovedad:1"
	}
	return "All"
}
