package helpers

import (
	"archive/zip"
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
)

func downloadZip(idPago int, zipName string) (outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al descargar el Zip",
				"Funcion": "downloadZip",
			}
			panic(outputError)
		}
	}()

	//Realizar la solicitud opara obtener los documentos desde la Api

	/*

			resp, err := http.Get("URL_DE_TU_API")
		    if err != nil {
		        fmt.Println("Error al hacer la solicitud HTTP:", err)
		        return
		    }
		    defer resp.Body.Close()

	*/

	var respuesta_peticion map[string]interface{}
	var documentos models.DocumentosCumplidos

	if response, err := getJsonTest(beego.AppConfig.String("UrlCrudCumplidos")+"/document?query=Id.in:"+strconv.Itoa(idPago)+",Activo:true", &respuesta_peticion); (err == nil) && (response == 200) {
		LimpiezaRespuestaRefactor(respuesta_peticion, &documentos)
	} else {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  502,
			"Message": "Error al hacer la solicitud y decodificar el JSON",
			"Funcion": "downloadAndAddToZip",
		}
		return outputError
	}

	//Crear un archivo ZIP

	zipFile, err := os.Create(zipName + ".zip")
	if err != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  502,
			"Message": "Error al crear el archivo ZIP",
			"Funcion": "downloadAndAddToZip",
		}
		return outputError
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Descargar y agregar cada documento al archivo ZIP

	downloadFileAndAddToZip(documentos.Nuxeo.Thumb.Data, documentos.Nuxeo.Thumb.Name, zipWriter)
	downloadFileAndAddToZip(documentos.Nuxeo.Content.Data, documentos.Nuxeo.Content.Name, zipWriter)

	//Decodificar el archivo en base64 y agregarlo al .zip
	pdfData, error := base64.StdEncoding.DecodeString(documentos.Nuxeo.File)
	if error != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  502,
			"Message": "Error al decodificar el archivo base64",
			"Error":   error,
		}
		return outputError
	}

	_, err = zipFile.Write(pdfData)
	if err != nil {
		outputError := map[string]interface{}{
			"Success": false,
			"Status":  502,
			"Message": "Error al escribir el contenido del PDF en el archivo ZIP",
			"Error":   err.Error(),
		}
		return outputError
	}

	return nil

}

func downloadFileAndAddToZip(url, fileName string, zipWriter *zip.Writer) (outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"Succes":  false,
				"Status":  502,
				"Message": "Error al descargar el documento y agregarlo al .zip",
				"Funcion": "downloadAndAddToZip",
			}
			panic(outputError)
		}
	}()
	// Solicitar el documento

	resp, err := http.Get(url)
	if err != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  502,
			"Message": "Error al obtener el documento",
			"Error":   err,
		}
		return outputError
	}
	defer resp.Body.Close()

	// Crear un archivo en el archivo ZIP

	zipFile, error := zipWriter.Create(fileName)
	if error != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  502,
			"Message": "Error al crear el archivo en el archivo ZIP",
			"Error":   error,
		}
		return outputError
	}

	// Copiar el contenido del documento al archivo en el archivo ZIP

	_, error1 := io.Copy(zipFile, resp.Body)
	if error1 != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  502,
			"Message": "Error al copiar el contenido del documento al archivo ZIP",
			"Error":   error1,
		}
		return outputError
	}

	return nil
}
