package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["mynotes/dataServer/controllers:MainController"] = append(beego.GlobalControllerRouter["mynotes/dataServer/controllers:MainController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/staticblock/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mynotes/dataServer/controllers:ObjectController"] = append(beego.GlobalControllerRouter["mynotes/dataServer/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mynotes/dataServer/controllers:ObjectController"] = append(beego.GlobalControllerRouter["mynotes/dataServer/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mynotes/dataServer/controllers:ObjectController"] = append(beego.GlobalControllerRouter["mynotes/dataServer/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mynotes/dataServer/controllers:ObjectController"] = append(beego.GlobalControllerRouter["mynotes/dataServer/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mynotes/dataServer/controllers:ObjectController"] = append(beego.GlobalControllerRouter["mynotes/dataServer/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
