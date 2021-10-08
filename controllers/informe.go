package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/cumplidos_mid/helpers"
	"github.com/udistrital/cumplidos_mid/models"
)

//  InformeController operations for Informe
type InformeController struct {
	beego.Controller
}

// URLMapping ...
func (c *InformeController) URLMapping() {
	c.Mapping("Post", c.PostInforme)
	c.Mapping("GetOne", c.GetInforme)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Informe
// @Param	body		body 	models.Informe	true		"body for Informe content"
// @Success 201 {int} models.Informe
// @Failure 403 body is empty
// @router / [post]
func (c *InformeController) PostInforme() {
	var v models.Informe
	//var v map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	fmt.Println("Informe al llegar", v)
	if response, err := helpers.AddInforme(v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = response
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Informe by contrato,vigencia,mes y anio
// @Param	contrato		path 	string	true		"The key for staticblock"
// @Param	vigencia		path 	string	true		"The key for staticblock"
// @Param	mes		path 	string	true		"The key for staticblock"
// @Param	anio		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Informe
// @Failure 403 :id is empty
// @router /:contrato/:vigencia/:mes/:anio [get]
func (c *InformeController) GetInforme() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "InformeController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	contrato := c.GetString(":contrato")
	vigencia := c.GetString(":vigencia")
	mes := c.GetString(":mes")
	anio := c.GetString(":anio")

	if len(vigencia) > 4 || len(mes) > 2 || len(anio) > 4 {
		panic(map[string]interface{}{"funcion": "GetInforme", "err": "Error en los parametros de ingreso", "status": "400"})
	}

	if informe, err := helpers.Informe(contrato, vigencia, mes, anio); (err == nil) || (len(informe) != 0) {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": informe}
	} else {
		panic(err)
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
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
	// var fields []string
	// var sortby []string
	// var order []string
	// var query = make(map[string]string)
	// var limit int64 = 10
	// var offset int64

	// // fields: col1,col2,entity.col3
	// if v := c.GetString("fields"); v != "" {
	// 	fields = strings.Split(v, ",")
	// }
	// // limit: 10 (default is 10)
	// if v, err := c.GetInt64("limit"); err == nil {
	// 	limit = v
	// }
	// // offset: 0 (default is 0)
	// if v, err := c.GetInt64("offset"); err == nil {
	// 	offset = v
	// }
	// // sortby: col1,col2
	// if v := c.GetString("sortby"); v != "" {
	// 	sortby = strings.Split(v, ",")
	// }
	// // order: desc,asc
	// if v := c.GetString("order"); v != "" {
	// 	order = strings.Split(v, ",")
	// }
	// // query: k:v,k:v
	// if v := c.GetString("query"); v != "" {
	// 	for _, cond := range strings.Split(v, ",") {
	// 		kv := strings.SplitN(cond, ":", 2)
	// 		if len(kv) != 2 {
	// 			c.Data["json"] = errors.New("Error: invalid query key/value pair")
	// 			c.ServeJSON()
	// 			return
	// 		}
	// 		k, v := kv[0], kv[1]
	// 		query[k] = v
	// 	}
	// }

	// l, err := models.GetAllInforme(query, fields, sortby, order, offset, limit)
	// if err != nil {
	// 	c.Data["json"] = err.Error()
	// } else {
	// 	c.Data["json"] = l
	// }
	c.ServeJSON()
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
	// idStr := c.Ctx.Input.Param(":id")
	// id, _ := strconv.ParseInt(idStr, 0, 64)
	// v := models.Informe{Id: id}
	// json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	// if err := models.UpdateInformeById(&v); err == nil {
	// 	c.Data["json"] = "OK"
	// } else {
	// 	c.Data["json"] = err.Error()
	// }
	// c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Informe
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *InformeController) Delete() {
	// idStr := c.Ctx.Input.Param(":id")
	// id, _ := strconv.ParseInt(idStr, 0, 64)
	// if err := models.DeleteInforme(id); err == nil {
	// 	c.Data["json"] = "OK"
	// } else {
	// 	c.Data["json"] = err.Error()
	// }
	// c.ServeJSON()
}
