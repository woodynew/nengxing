package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// AppController operations for App
type AppController struct {
	beego.Controller
}

type JSONS struct {
	//必须的大写开头
	Code string
	Msg  string
	User []string `json:"user_info"` //key重命名,最外面是反引号
}

// URLMapping ...
func (c *AppController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create App
// @Param	body		body 	models.App	true		"body for App content"
// @Success 201 {object} models.App
// @Failure 403 body is empty
// @router / [post]
func (c *AppController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get App by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.App
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AppController) GetOne() {
	mystruct := &JSONS{"100", "获取成功",
		[]string{"maple", "18"}}
	c.Data["json"] = &mystruct
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get App
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.App
// @Failure 403
// @router / [get]
func (c *AppController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the App
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.App	true		"body for App content"
// @Success 200 {object} models.App
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AppController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the App
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AppController) Delete() {

}
