package certificacionesHelper

import "net/http"
	

func TestGetSolicitudesPagoMensual(t *testing.T) {

	if response, err = http.Post("http://localhost:8080/v1/solicitudes_pagos");

}
