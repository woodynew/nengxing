// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"go-common-tools-api/controllers"
	controllers_api "go-common-tools-api/controllers/api"

	context "github.com/beego/beego/v2/server/web/context"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/qrcode",
			beego.NSInclude(
				&controllers_api.QrcodeController{},
			),

			// beego.NSInclude(
			// 	&controllers_api.ObjectController{},
			// ),

			// beego.NSRouter("/decode-file", &controllers_api.QrcodeController{}, "get,post:DecodeFile"),
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/app", &controllers.AppController{}, "get:GetOne")

	beego.Get("/hello", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
}
