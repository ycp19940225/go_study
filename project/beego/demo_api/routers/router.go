// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"demo_api/controllers"

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
	)
	beego.AddNamespace(ns)
	beego.AutoRouter(&controllers.LoginController{})
	beego.AutoRouter(&controllers.CommentsController{})
	beego.AutoRouter(&controllers.RedisTestController{})
	beego.AutoRouter(&controllers.MqDemoController{})
	beego.AutoRouter(&controllers.ChannelDemoController{})
	beego.AutoRouter(&controllers.EsDemoController{})
}
