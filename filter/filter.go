package filter

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"regexp"
	"suyuanweb/models"
	"suyuanweb/sutil"
)

// 官网管理后台控制器
var FilterAdmin = func(ctx *context.Context) {
	islogin, token := sutil.IsLogin(ctx)
	res_url := ctx.Request.RequestURI
	if res_url != "/admin/login" {
		if islogin {
			//根据权限过滤
			obj, err := models.GetOneAdminFromToken(token)
			if err != nil {
				beego.Error(err)
				ctx.Redirect(302, "/admin/login")
			}
			if obj.Token != token {
				ctx.Redirect(302, "/admin/login")
			}
			if obj.Auth == 0 {
				if res_url == "/admin/addadmin" || res_url == "/admin/addproduct" || res_url == "/admin/addnews" || res_url == "/admin/addbanner" {
					ctx.Redirect(302, "/admin/noauth")
				}
			}
		} else {
			ctx.Redirect(302, "/admin/login")
		}
	}
}

var FilterBSAdmin = func(ctx *context.Context) {
	islogin, _ := sutil.IsBSAdminLogin(ctx)
	res_url := ctx.Request.RequestURI
	if res_url != "/bs/admin/login" {
		if islogin {
			//根据权限过滤
		} else {
			ctx.Redirect(302, "/bs/admin/login")
		}
	}
}

var FilterBSUser = func(ctx *context.Context) {
	islogin, _ := sutil.IsBSUserLogin(ctx)
	res_url := ctx.Request.RequestURI
	r := bytes.NewReader([]byte(res_url))
	reg := regexp.MustCompile("/bs/admin/*")
	if reg.MatchReader(r) == false { //过滤掉admin
		if res_url != "/bs/login" {
			if islogin {
				//根据权限过滤
			} else {
				ctx.Redirect(302, "/bs/login")
			}
		}
	}
}

//主页过滤器
var FilterHome = func(ctx *context.Context) {
	res_url := ctx.Request.RequestURI
	isimagehosting := regexp.MustCompile("/imagehosting/*").MatchReader(bytes.NewReader([]byte(res_url)))
	isadmin := regexp.MustCompile("/admin/*").MatchReader(bytes.NewReader([]byte(res_url)))
	req := ctx.Request
	addr := req.RemoteAddr // "IP:port" "192.168.1.150:8889"
	beego.Debug("isimagehosting", isimagehosting, "isadmin:", isadmin)
	if isimagehosting == false && isadmin == false {
		_, err := models.AddPageLog(res_url, addr)
		if err != nil {
			beego.Error(err)
		}

	}

}
