package api

import (
	"fmt"
	"go-common-tools-api/apihelpers"
	"net/url"
	"os"

	"github.com/tuotoo/qrcode"
)

// QrcodeController operations for Qrcode
type QrcodeController struct {
	BaseController
}

// URLMapping ...
func (c *QrcodeController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// DecodeFile ...
// @Title DecodeFile
// @Description create Qrcode
// @Param	body		body 	models.Qrcode	true		"body for Qrcode content"
// @Success 201 {object} models.Qrcode
// @Failure 403 body is empty
// @router / [decode-file]
func (c *QrcodeController) DecodeFile() {
	path := c.GetString("path")
	fmt.Println(path)
	// fmt.Println(url.QueryEscape(path))
	// fmt.Println(url.QueryUnescape(path))
	path, _ = url.QueryUnescape(path)

	// "E%3A%2Fxampp%2Fhtdocs%2Flaravel%2Faliyun_code%2Ftelecom_emall%2Fdev%2Femall_api%2Fpublic%2Fwxfile%2Fxcxcode%2F321.jpg"
	// file, err := os.Open("E:/xampp/htdocs/laravel/aliyun_code/telecom_emall/dev/emall_api/public/wxfile/xcxcode/321.jpg")
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)

		c.ResultError(apihelpers.CODE_ERROR, "文件地址错误")
		return
	}

	defer file.Close()
	//识别二维码
	qr, err := qrcode.Decode(file)
	if err != nil {
		fmt.Println(err.Error())
		c.ResultError(apihelpers.CODE_ERROR, err.Error())
		return
	}
	fmt.Println(qr.Content)

	param := map[string]string{
		"text": qr.Content,
	}

	c.ResultSuccess(param)
}

// Post ...
// @Title Create
// @Description create Qrcode
// @Param	body		body 	models.Qrcode	true		"body for Qrcode content"
// @Success 201 {object} models.Qrcode
// @Failure 403 body is empty
// @router / [post]
func (c *QrcodeController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Qrcode by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Qrcode
// @Failure 403 :id is empty
// @router /:id [get]
func (c *QrcodeController) GetOne() {
}

// GetAll ...
// @Title GetAll
// @Description get Qrcode
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Qrcode
// @Failure 403
// @router / [get]
func (c *QrcodeController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Qrcode
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Qrcode	true		"body for Qrcode content"
// @Success 200 {object} models.Qrcode
// @Failure 403 :id is not int
// @router /:id [put]
func (c *QrcodeController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Qrcode
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *QrcodeController) Delete() {

}
