package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"suyuanweb/models"
	_ "suyuanweb/routers"
	"suyuanweb/sutil"
)

func init() {
	// 注册数据库
	models.RegisterDB()

	beego.AddFuncMap("i18n", i18n.Tr)
	beego.AddFuncMap("isimgpath", sutil.IsImgPath)
	beego.AddFuncMap("isadminversion", sutil.IsAdminVersion)
	beego.AddFuncMap("timeformat", sutil.TimeFormat)
	beego.AddFuncMap("timeformatstyle1", sutil.TimeFormatStyle1)
	beego.AddFuncMap("datexlxs", sutil.DateXlxs)
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true

	beego.Run()
}
