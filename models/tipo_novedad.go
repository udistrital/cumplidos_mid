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
		return "220"
	case TipoNovedadCesion:
		return "219"
	case TipoNovedadTerminacion:
		return "218"
	case TipoNovedadSuspension:
		return "216"
	}
	return "All"
}
