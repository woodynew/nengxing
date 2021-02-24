package api

import (
	"encoding/json"
	"fmt"
	"go-common-tools-api/apihelpers"
	"go-common-tools-api/models"
	"net/url"
	"os"

	"github.com/tuotoo/qrcode"
)

// QrcodeController operations for Qrcode
type QrcodeController struct {
	BaseController
}

// @Title 识别本地图片二维码
// @Description create object
// @Param	path		path1 	string	true		"文件真实路径"
// @Success 200 {string} apihelpers.JsonResponse
// @Failure 403 body is empty
// @router /decode-file [post]
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

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (o *QrcodeController) Post() {
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *QrcodeController) Get() {
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *QrcodeController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *QrcodeController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *QrcodeController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
