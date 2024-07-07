package helpers

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/cumplidos_mid/models"
	"github.com/udistrital/utils_oas/request"
)

func DescargarDocumentosSolicitudesPagos(id_pago_mensual string) (DocumentosZip models.DescargaDocumentos, outputError map[string]interface{}) {
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

	var respuesta_peticion map[string]interface{}
	var pagos_mensuales []models.PagoMensual
	documentos, error := TraerEnlacesDocumentosAsociadosPagoMensual(id_pago_mensual)

	if error != nil {
		outputError = map[string]interface{}{
			"Succes":  false,
			"Status":  502,
			"Message": "Error al obtener los documentos del pago",
			"Funcion": "TraerEnlacesDocumentosAsociadosPagoMensual",
			"Error":   error,
		}
		return DocumentosZip, outputError
	} else if len(documentos) == 0 {
		return DocumentosZip, nil
	}

	//Crear un archivo ZIP

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

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
			return DocumentosZip, outputError
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
			return DocumentosZip, outputError
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
			return DocumentosZip, outputError
		}
	}

	// Cerrar el writer del ZIP
	err := zipWriter.Close()
	if err != nil {
		outputError := map[string]interface{}{
			"Success": false,
			"Status":  502,
			"Message": "Error al cerrar el archivo ZIP",
			"Error":   err.Error(),
		}
		return DocumentosZip, outputError
	}

	DocumentosZip.File = base64.StdEncoding.EncodeToString(buf.Bytes())

	//fmt.Println(beego.AppConfig.String("UrlCrudCumplidos") + "/pago_mensual/?query=Id:" + id_pago_mensual)
	if response, err := request.GetJsonTest2(beego.AppConfig.String("UrlCrudCumplidos")+"/pago_mensual/?query=Id:"+id_pago_mensual, &respuesta_peticion); (err == nil) && (response == 200) {
		if respuesta_peticion != nil {
			LimpiezaRespuestaRefactor(respuesta_peticion, &pagos_mensuales)
		}
	}

	var informacion_contrato_contratista models.InformacionContratoContratista
	for _, pago_mensual := range pagos_mensuales {
		informacion_contrato_contratista, error = GetInformacionContratoContratista(pago_mensual.NumeroContrato, strconv.FormatFloat(pago_mensual.VigenciaContrato, 'f', 0, 64))

		if error == nil {
			DocumentosZip.Nombre = informacion_contrato_contratista.InformacionContratista.NombreCompleto + "_" + pago_mensual.NumeroContrato + "_" + informacion_contrato_contratista.InformacionContratista.Documento.Numero + "_" + strconv.FormatFloat(pago_mensual.Ano, 'f', 0, 64) + "_" + strconv.FormatFloat(pago_mensual.Mes, 'f', 0, 64)
		} else {
			outputError := map[string]interface{}{
				"Success": false,
				"Status":  502,
				"Message": "Error al Buscar los datos del contratista",
			}
			return DocumentosZip, outputError
		}
	}

	return DocumentosZip, nil
}
