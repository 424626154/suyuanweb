package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["suyuanweb/controllers:HomeController"] = append(beego.GlobalControllerRouter["suyuanweb/controllers:HomeController"],
		beego.ControllerComments{
			Method: "Banner",
			Router: `/banner`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["suyuanweb/controllers:HomeController"] = append(beego.GlobalControllerRouter["suyuanweb/controllers:HomeController"],
		beego.ControllerComments{
			Method: "Newss",
			Router: `/newss`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["suyuanweb/controllers:HomeController"] = append(beego.GlobalControllerRouter["suyuanweb/controllers:HomeController"],
		beego.ControllerComments{
			Method: "News",
			Router: `/news`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["suyuanweb/controllers:HomeController"] = append(beego.GlobalControllerRouter["suyuanweb/controllers:HomeController"],
		beego.ControllerComments{
			Method: "Products",
			Router: `/products`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["suyuanweb/controllers:HomeController"] = append(beego.GlobalControllerRouter["suyuanweb/controllers:HomeController"],
		beego.ControllerComments{
			Method: "Product",
			Router: `/product`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["suyuanweb/controllers:HomeController"] = append(beego.GlobalControllerRouter["suyuanweb/controllers:HomeController"],
		beego.ControllerComments{
			Method: "Service",
			Router: `/service`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["suyuanweb/controllers:HomeController"] = append(beego.GlobalControllerRouter["suyuanweb/controllers:HomeController"],
		beego.ControllerComments{
			Method: "About",
			Router: `/about`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
