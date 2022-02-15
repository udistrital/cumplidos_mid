package models

type ContratoGeneral struct {
	LugarEjecucion struct {
		Sede        string
		Dependencia string
	}

	Supervisor struct {
		DependenciaSupervisor string
	}
}
