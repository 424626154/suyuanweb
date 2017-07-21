package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"suyuanweb/models/bs"
)

func RegisterDB() {
	db_url := "root:suyuan@/suyuandb?charset=utf8&loc=Local"
	cnf, err := config.NewConfig("ini", "conf/suyuanweb.conf")
	if err != nil {
		beego.Error(err)
	}
	syuan_runenv := cnf.String("suyuanweb::syuan_runenv")
	beego.Debug("syuan_runenv:", syuan_runenv)
	if syuan_runenv == "ali" {
		db_url = "root:sbb890503@/suyuandb?charset=utf8&loc=Local"
	} else if syuan_runenv == "suyuan" {
		db_url = "root:suyuanjiuye@/suyuandb?charset=utf8&loc=Local"
	}
	beego.Debug("db_url:", db_url)
	// set default database
	orm.RegisterDataBase("default", "mysql", db_url)
	// register model
	orm.RegisterModel(new(Admin)) //官网管理员表
	orm.RegisterModel(new(Product)) //产品
	orm.RegisterModel(new(News)) //新闻
	orm.RegisterModel(new(ImgBanner)) //图片广告
	orm.RegisterModel(new(WebCnfigure)) //网站配制
	orm.RegisterModel(new(PageLog))//网站log

	//bs
	orm.RegisterModel(new(BsAdmin))    //企业工作系统管理员
	orm.RegisterModel(new(BsUser))     //企业员工
	orm.RegisterModel(new(BsTemplate)) //工作模板
	orm.RegisterModel(new(bs.BsCommodityNameConfig))//商品名称配置
	orm.RegisterModel(new(bs.BsCommoditySpecConfig))//商品规格配置
	orm.RegisterModel(new(bs.BsInCommodity))//入库
	orm.RegisterModel(new(bs.BsExport))//导出记录
	orm.RegisterModel(new(bs.BsOutCommodity))//出库
	orm.RegisterModel(new(bs.BsAllCommodity)) //进出总
	orm.RegisterModel(new(bs.BsCommodityTemplet))//进销存模板
	// create table
	orm.RunSyncdb("default", false, true)
}
