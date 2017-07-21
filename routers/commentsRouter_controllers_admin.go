package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["suyuanweb/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["suyuanweb/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "Admin",
			Router: `/admin`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["suyuanweb/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["suyuanweb/controllers/admin:AdminController"],
		beego.ControllerComments{
			Method: "AdminLogout",
			Router: `/admin/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
