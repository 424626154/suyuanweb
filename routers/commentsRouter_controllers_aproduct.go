package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["suyuanweb/controllers/aproduct:ProductController"] = append(beego.GlobalControllerRouter["suyuanweb/controllers/aproduct:ProductController"],
		beego.ControllerComments{
			Method: "Upstate",
			Router: `/admin/product/upstate`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
