// @APIVersion 1.0.0
// @Title 文档
// @Description 数据服务的文档
// @Contact wangdyqxx@163.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"mynotes/dataServer/controllers"
)

func init() {
	beego.SetStaticPath("/swagger","swagger")
    //beego.Router("/", &controllers.MainController{})
	//beego.Router("/v1.0/object", &controllers.ObjectController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(&controllers.ObjectController{},),
		),
	)
	beego.AddNamespace(ns)
}
