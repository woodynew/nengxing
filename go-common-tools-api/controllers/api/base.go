package api

import (
	"go-common-tools-api/apihelpers"

	beego "github.com/beego/beego/v2/server/web"
)

// BaseController operations for Base
type BaseController struct {
	beego.Controller
}

func (c *BaseController) ResultError(code int, msg string) {
	c.Data["json"] = apihelpers.ResultCode(code, msg)
	c.ServeJSON()
}

func (c *BaseController) ResultSuccess(param map[string]string) {
	c.Data["json"] = apihelpers.ResultSuccess(param)
	c.ServeJSON()
}
