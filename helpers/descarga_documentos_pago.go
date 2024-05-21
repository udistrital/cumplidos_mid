package helpers

import (
	"archive/zip"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

func DescargarDocumentosSolicitudesPagos(id_pago_mensual string) (outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al descargar el .zip",
				"Funcion": "DescargarDocumentosSolicitudesPagos",
			}
			panic(outputError)
		}
	}()

	//Realizar la solicitud opara obtener los documentos asociados al pago

	//documentos, error := TraerEnlacesDocumentosAsociadosPagoMensual(id_pago_mensual)

	var documentos []models.DocumentosSoporte
	var respuesta_consulta map[string]interface{}

	fmt.Println(beego.AppConfig.String("UrlPruebasApi2") + "/contratos_contratista/documentos_pago_mensual/" + id_pago_mensual)
	if response, error := getJsonTest(beego.AppConfig.String("UrlPruebasApi2")+"/contratos_contratista/documentos_pago_mensual/"+id_pago_mensual, &respuesta_consulta); (response == 200) && (error == nil) {
		if len(respuesta_consulta) > 0 {
			LimpiezaRespuestaRefactor(respuesta_consulta, &documentos)
		}
	} else {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  502,
			"Message": "Error al obtener los documentos del pago",
			"Funcion": "TraerEnlacesDocumentosAsociadosPagoMensual",
			"Error":   error,
		}
		return outputError
	}

	/*
		if error != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al obtener los documentos del pago",
				"Funcion": "TraerEnlacesDocumentosAsociadosPagoMensual",
				"Error":   error,
			}
			return outputError
		}

	*/

	//Crear un archivo ZIP

	zipFile, err := os.Create("documentos_pago.zip")
	if err != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  502,
			"Message": "Error al crear el archivo ZIP",
			"Funcion": "DescargarDocumentosSolicitudesPagos",
		}
		return outputError
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Decodificar los archivos en base64 y agregarlos al .zip
	for i, documento := range documentos {
		pdfData, error := base64.StdEncoding.DecodeString(documento.Archivo.File)
		if error != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al decodificar el archivo base64",
				"Error":   error,
			}
			return outputError
		}

		// Crear una entrada en el ZIP para cada archivo PDF con su nombre específico y un índice único
		fileName := fmt.Sprintf("%s_%d.pdf", filepath.Base(documento.Documento.TipoDocumento.Nombre), i)
		zipEntry, err := zipWriter.Create(fileName)
		if err != nil {
			outputError = map[string]interface{}{
				"Success": false,
				"Status":  502,
				"Message": "Error al crear la entrada en el archivo ZIP",
				"Error":   err.Error(),
			}
			return outputError
		}

		// Escribir los datos del PDF en la entrada del ZIP
		_, err = zipEntry.Write(pdfData)
		if err != nil {
			outputError := map[string]interface{}{
				"Success": false,
				"Status":  502,
				"Message": "Error al escribir el contenido del PDF en el archivo ZIP",
				"Error":   err.Error(),
			}
			return outputError
		}
	}

	return nil
}
