package api

import (
	beego "github.com/beego/beego/v2/server/web"
)

// BaseController operations for Base
type BaseController struct {
	beego.Controller
}

// html
func (c *BaseController) Html(str string) {
	c.Ctx.WriteString(str) //self.GetControllerAndAction()
}
