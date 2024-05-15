package helpers

func ObtenerDependencias(documento string) (dependencias map[string]interface{}, errorOutput interface{}) {

	defer func() {

		if err := recover(); err != nil {
			errorOutput = map[string]interface{}{
				"Success": true,
				"Status":  502,
				"Message": "Error al consultar las dependencias: " + documento,
				"Error":   err,
			}
			panic(errorOutput)
		}
	}()
	dependencias = make(map[string]interface{})
	dependencias["Dependencias Supervisor"], errorOutput = GetDependenciasSupervisor(documento)
	dependencias["Dependencias Ordenador"], errorOutput = GetDependenciasOrdenadr(documento)
	if dependencias != nil {
		return dependencias, nil
	}

	return nil, errorOutput
}
