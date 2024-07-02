package helpers

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math/big"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_mid_api/models"
)

func sendJson(url string, trequest string, target interface{}, datajson interface{}) error {
	b := new(bytes.Buffer)
	if datajson != nil {
		if err := json.NewEncoder(b).Encode(datajson); err != nil {
			beego.Error(err)
		}
	}
	client := &http.Client{}
	fmt.Print("Json que se le va a apasar a la funcion: ")
	fmt.Println(b)
	req, err := http.NewRequest(trequest, url, b)
	r, err := client.Do(req)
	if err != nil {
		beego.Error("error", err)
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return json.NewDecoder(r.Body).Decode(target)
}

func sendJson3(url string, trequest string, target interface{}, datajson interface{}) error {
	// Convertir datajson en un Formato Json para poderlo enviar como parametro
	b := new(bytes.Buffer)
	if datajson != nil {
		if err := json.NewEncoder(b).Encode(datajson); err != nil {
			fmt.Println(err)
			beego.Error(err)
		}
	}

	// Crear una nueva solicitud POST con el cuerpo del JSON
	req, err := http.NewRequest(trequest, url, b)
	if err != nil {
		fmt.Println("Error al crear la solicitud POST:", err)
		return err
	}

	//Configurar el encabezado Accept
	req.Header.Set("Accept", "application/json")
	// Configurar el encabezado Content-Type
	req.Header.Set("Content-Type", "application/json")

	// Configurar el cliente HTTP con tiempo de espera y tamaño de búfer
	client := &http.Client{
		Timeout: time.Second * 10, // Tiempo de espera máximo de 10 segundos
		Transport: &http.Transport{
			MaxIdleConns:        100, // Número máximo de conexiones inactivas permitidas
			MaxIdleConnsPerHost: 100, // Número máximo de conexiones inactivas permitidas por host
		},
	}

	// Realizar la solicitud POST
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al enviar la solicitud POST:", err)
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return json.NewDecoder(resp.Body).Decode(target)
}

func getJsonTest(url string, target interface{}) (status int, err error) {
	r, err := http.Get(url)
	if err != nil {
		return r.StatusCode, err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return r.StatusCode, json.NewDecoder(r.Body).Decode(target)
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return json.NewDecoder(r.Body).Decode(target)
}

// func getJsonTest(w http.ResponseWriter,r *http.Request){
// 	err := r.ParseForm()
// 	if err != nil {
// 	   log.Fatal("parse form error ",err)
// 	}
// 	// 初始化请求变量结构
// 	formData := make(map[string]interface{})
// 	// 调用json包的解析，解析请求body
// 	json.NewDecoder(r.Body).Decode(&formData)
// 	for key,value := range formData{
// 	   log.Println("key:",key," => value :",value)
// 	}
//  }

func getXml(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return xml.NewDecoder(r.Body).Decode(target)
}

func getJsonWSO2(urlp string, target interface{}) error {
	b := new(bytes.Buffer)
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlp, b)
	req.Header.Set("Accept", "application/json")
	r, err := client.Do(req)
	if err != nil {
		beego.Error("error", err)
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return json.NewDecoder(r.Body).Decode(target)
}

func getJsonWSO2Test(urlp string, target interface{}) (status int, err error) {
	b := new(bytes.Buffer)
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlp, b)
	req.Header.Set("Accept", "application/json")
	r, err := client.Do(req)
	if err != nil {
		beego.Error("error", err)
		return r.StatusCode, err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(nil, err)
		}
	}()

	return r.StatusCode, json.NewDecoder(r.Body).Decode(target)
}

func diff(a, b time.Time) (year, month, day int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	oneDay := time.Hour * 5
	a = a.Add(oneDay)
	b = b.Add(oneDay)
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)

	// Normalize negative values

	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

//CargarReglasBase general
func CargarReglasBase(dominio string) (reglas string, err error) {
	//carga de reglas desde el ruler
	var reglasbase string = ``
	var v []models.Predicado
	err = getJson(beego.AppConfig.String("Urlruler")+"/predicado/?query=Dominio.Nombre:"+dominio+"&limit=-1", &v)
	if err != nil {
		return
	}
	reglasbase = reglasbase + FormatoReglas(v) //funcion general para dar formato a reglas cargadas desde el ruler

	//-----------------------------
	return reglasbase, nil
}

func FormatoReglas(v []models.Predicado) (reglas string) {
	var arregloReglas = make([]string, len(v))
	reglas = ""
	//var respuesta []models.FormatoPreliqu
	for i := 0; i < len(v); i++ {
		arregloReglas[i] = v[i].Nombre
	}

	for i := 0; i < len(arregloReglas); i++ {
		reglas = reglas + arregloReglas[i] + "\n"
	}
	return
}

func FormatMoney(value interface{}, Precision int) string {
	formattedNumber := FormatNumber(value, Precision, ",", ".")
	return FormatMoneyString(formattedNumber, Precision)
}

func FormatMoneyString(formattedNumber string, Precision int) string {
	var format string

	zero := "0"
	if Precision > 0 {
		zero += "." + strings.Repeat("0", Precision)
	}

	format = "%s%v"
	result := strings.Replace(format, "%s", "$", -1)
	result = strings.Replace(result, "%v", formattedNumber, -1)

	return result
}

func FormatNumber(value interface{}, precision int, thousand string, decimal string) string {
	v := reflect.ValueOf(value)
	var x string
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		x = fmt.Sprintf("%d", v.Int())
		if precision > 0 {
			x += "." + strings.Repeat("0", precision)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		x = fmt.Sprintf("%d", v.Uint())
		if precision > 0 {
			x += "." + strings.Repeat("0", precision)
		}
	case reflect.Float32, reflect.Float64:
		x = fmt.Sprintf(fmt.Sprintf("%%.%df", precision), v.Float())
	case reflect.Ptr:
		switch v.Type().String() {
		case "*big.Rat":
			x = value.(*big.Rat).FloatString(precision)

		default:
			panic("Unsupported type - " + v.Type().String())
		}
	default:
		panic("Unsupported type - " + v.Kind().String())
	}

	return formatNumberString(x, precision, thousand, decimal)
}

func formatNumberString(x string, precision int, thousand string, decimal string) string {
	lastIndex := strings.Index(x, ".") - 1
	if lastIndex < 0 {
		lastIndex = len(x) - 1
	}

	var buffer []byte
	var strBuffer bytes.Buffer

	j := 0
	for i := lastIndex; i >= 0; i-- {
		j++
		buffer = append(buffer, x[i])

		if j == 3 && i > 0 && !(i == 1 && x[0] == '-') {
			buffer = append(buffer, ',')
			j = 0
		}
	}

	for i := len(buffer) - 1; i >= 0; i-- {
		strBuffer.WriteByte(buffer[i])
	}
	result := strBuffer.String()

	if thousand != "," {
		result = strings.Replace(result, ",", thousand, -1)
	}

	extra := x[lastIndex+1:]
	if decimal != "." {
		extra = strings.Replace(extra, ".", decimal, 1)
	}

	return result + extra
}

func dias31(fecha_inicio time.Time, fecha_fin time.Time) (dias31 int) {
	count := int(fecha_fin.Sub(fecha_inicio).Hours()/24) + 1
	fmt.Println(count)

	for i := 0; i < count; i++ {
		day := fecha_inicio.AddDate(0, 0, i).Day()
		//fmt.Println(day)
		if day == 31 {
			dias31++
		}
	}
	return
}

func buildQuery(slices []string, columna string) string {

	query := ""

	if len(slices) == 1 {
		query += fmt.Sprintf("%s.in:%v,", columna, slices[0])
	}
	if len(slices) > 1 {
		for i, dato := range slices {
			if i == 0 {
				query += fmt.Sprintf("%s.in:%v|", columna, dato)
			} else if i < len(slices)-1 {
				query += fmt.Sprintf("%s|", dato)
			} else {
				query += fmt.Sprintf("%s,", dato)
			}
		}
		return query
	}
	return query
}

// Funcion para agregar los datos a un slice
func StringToSlice(cadena string) (slice []string) {
	parts := strings.Split(cadena, ",")

	if cadena != "" {
		for _, part := range parts {
			slice = append(slice, part)
		}
	}
	return slice
}

//Funcion para Verificar que se ingresen datos correctos cuando el parametro sean números

func ConvertInt(data []string) (outputError map[string]interface{}) {
	for _, str := range data {
		_, err := strconv.Atoi(str)
		if err != nil && len(data) > 0 {
			outputError = map[string]interface{}{
				"Success": false,
				"Status":  400,
				"Message": "El valor " + str + " no es un número",
				"funcion": "convertInt",
			}
			return outputError
		}
	}
	return nil
}

func capitalizarPrimeraLetra(texto string) string {
	if len(texto) == 0 {
		return texto
	}

	texto = strings.TrimSpace(texto) // Eliminar espacios iniciales y finales
	texto = strings.ToLower(texto)
	return strings.ToUpper(texto[:1]) + texto[1:]
}
