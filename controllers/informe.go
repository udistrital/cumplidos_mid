package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
)

// InformeController operations for Informe
type InformeController struct {
	beego.Controller
}

// URLMapping ...
func (c *InformeController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Informe
// @Param	body		body 	models.Informe	true		"body for Informe content"
// @Success 201 {object} models.Informe
// @Failure 403 body is empty
// @router / [post]
func (c *InformeController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Informe by num_documento
// @Param	num_documento	path 	string	true		"numero documento contratista"
// @Param	contrato	path 	string	true		"numero de contrato"
// @Param	vigencia	path 	string	true		"vigencia del contrato"
// @Param	anio	path 	string	true		"añio del cumplido"
// @Param	mes	path 	string	true		"mes del cumplido"
// @Success 200 {object} models.Informe
// @Failure 403 :num_documento is empty
// @router /:num_documento/:contrato/:vigencia/:anio/:mes [get]
func (c *InformeController) GetOne() {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "ContratosContratistaController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	fmt.Println("Numero documento:", c.GetString(":num_documento"))
	fmt.Println("Contrato:", c.GetString(":contrato"))
	fmt.Println("Vigencia:", c.GetString(":vigencia"))
	fmt.Println("Año:", c.GetString(":anio"))
	fmt.Println("Mes:", c.GetString(":mes"))

	num_documento := c.GetString(":num_documento")
	contrato := c.GetString(":contrato")
	vigencia := c.GetString(":vigencia")
	anio := c.GetString(":anio")
	mes := c.GetString(":mes")

	_, err := strconv.Atoi(num_documento)
	_, err2 := strconv.Atoi(contrato)
	_, err3 := strconv.Atoi(vigencia)
	_, err4 := strconv.Atoi(anio)
	_, err5 := strconv.Atoi(mes)

	if err != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || len(num_documento) < 2 {
		fmt.Println(err)
		fmt.Println(err2)
		fmt.Println(err3)
		panic(map[string]interface{}{"funcion": "GetOne", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if informacion_informe, err := helpers.InformacionInforme(num_documento, contrato, vigencia); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": informacion_informe}
	} else {
		panic(err)
	}
	c.ServeJSON()

}

// GetAll ...
// @Title GetAll
// @Description get Informe
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Informe
// @Failure 403
// @router / [get]
func (c *InformeController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Informe
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Informe	true		"body for Informe content"
// @Success 200 {object} models.Informe
// @Failure 403 :id is not int
// @router /:id [put]
func (c *InformeController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Informe
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *InformeController) Delete() {

}
