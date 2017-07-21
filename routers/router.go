package routers

import (
	"github.com/astaxie/beego"
	"os"
	"suyuanweb/controllers"
	"suyuanweb/controllers/admin"
	"suyuanweb/controllers/bs"
	"suyuanweb/controllers/home"
	"suyuanweb/filter"
)

func init() {
	beego.Router("/", &home.HomeController{}, "*:Home")
	beego.Router("/banner", &home.HomeController{}, "*:Banner")
	beego.Router("/newss", &home.HomeController{}, "*:Newss")
	beego.Router("/news", &home.HomeController{}, "*:News")
	beego.Router("/products", &home.HomeController{}, "*:Products")
	beego.Router("/product", &home.HomeController{}, "*:Product")
	beego.Router("/service", &home.HomeController{}, "*:Service")
	beego.Router("/about", &home.HomeController{}, "*:About")

	beego.Router("/admin", &admin.AdminController{}, "*:Admin")
	beego.Router("/admin/login", &admin.AdminController{}, "*:AdminLogin")
	beego.Router("/admin/logout", &admin.AdminController{}, "*:Logout")
	beego.Router("/admin/noauth", &admin.AdminController{}, "*:Noauth")
	beego.Router("/admin/admins", &admin.AdminAccountController{}, "*:Admins")
	beego.Router("/admin/addadmin", &admin.AdminAccountController{}, "*:AddAdmin")
	beego.Router("/admin/upadminpwd", &admin.AdminAccountController{}, "*:UpAdminPwd")
	beego.Router("/admin/upadminauth", &admin.AdminAccountController{}, "*:UpAdminAuth")
	beego.Router("/admin/updatelog", &admin.AdminController{}, "*:UpdateLog")
	beego.Router("/admin/configure", &admin.AdminController{}, "*:Configure")
	beego.Router("/admin/pageview", &admin.AdminController{}, "*:PageView")

	beego.Router("/admin/products", &admin.ProductController{}, "*:Products")
	beego.Router("/admin/addproduct", &admin.ProductController{}, "*:AddProduct")
	beego.Router("/admin/product", &admin.ProductController{})
	beego.Router("/admin/upproduct", &admin.ProductController{}, "*:UpProduct")
	beego.Router("/admin/newss", &admin.NewsController{}, "*:Newss")
	beego.Router("/admin/addnews", &admin.NewsController{}, "*:AddNews")
	beego.Router("/admin/news", &admin.NewsController{}, "*:News")
	beego.Router("/admin/upnews", &admin.NewsController{}, "*:Upnews")
	beego.Router("/admin/banners", &admin.ImageBannerController{}, "*:ImageBanners")
	beego.Router("/admin/addbanner", &admin.ImageBannerController{}, "*:AddImageBanner")
	beego.Router("/admin/banner", &admin.ImageBannerController{}, "*:ImageBanner")
	beego.Router("/admin/upbanner", &admin.ImageBannerController{}, "*:UpImageBanner")
	beego.Router("/admin/productajax", &admin.ProductController{}, "*:ProductAjax")

	beego.Router("/bs/admin", &bs.BsAdminController{}, "*:Index")
	beego.Router("bs/admin/login", &bs.BsAdminController{}, "*:Login")
	beego.Router("bs/admin/logout", &bs.BsAdminController{}, "*:Logout")
	beego.Router("bs/admin/addadmin", &bs.BsAdminController{}, "*:AddAdmin")
	beego.Router("bs/admin/admins", &bs.BsAdminController{}, "*:Admins")
	beego.Router("bs/admin/upapwd", &bs.BsAdminController{}, "*:Upapwd")
	beego.Router("bs/admin/upaauth", &bs.BsAdminController{}, "*:Upaauth")
	beego.Router("/bs/admin/addtemplate", &bs.BsAdminController{}, "*:AddTemplate")
	beego.Router("/bs/admin/templates", &bs.BsAdminController{}, "*:Templates")
	beego.Router("/bs/admin/template", &bs.BsAdminController{}, "*:Template")
	beego.Router("/bs/admin/uptemplate", &bs.BsAdminController{}, "*:UpTemplate")
	beego.Router("/bs/admin/commoditynames", &bs.BsAdminController{}, "*:CommodityNames")
	beego.Router("/bs/admin/commodityajax", &bs.BsAdminController{}, "*:CommodityAjax")
	beego.Router("/bs/admin/commodityspecs", &bs.BsAdminController{}, "*:CommoditySpecs")
	beego.Router("/bs/admin/commoditytemplates", &bs.BsAdminController{}, "*:CommodityTemplates")

	beego.Router("/bs", &bs.BSHomeController{}, "*:Index")
	beego.Router("/bs/login", &bs.BSHomeController{}, "*:Login")
	beego.Router("/bs/logout", &bs.BSHomeController{}, "*:Logout")
	beego.Router("/bs/templates", &bs.BSHomeController{}, "*:Templates")
	beego.Router("/bs/incommoditys", &bs.BSHomeController{}, "*:InCommoditys")
	beego.Router("/bs/addincommodity", &bs.BSHomeController{}, "*:AddInCommodity")
	beego.Router("/bs/upincommodity", &bs.BSHomeController{}, "*:UpInCommodity")
	beego.Router("/bs/export", &bs.BSHomeController{}, "*:Export")
	beego.Router("/bs/outcommoditys", &bs.BSHomeController{}, "*:OutCommoditys")
	beego.Router("/bs/addoutcommodity", &bs.BSHomeController{}, "*:AddOutCommodity")
	beego.Router("/bs/upoutcommodity", &bs.BSHomeController{}, "*:UpOutCommodity")
	beego.Router("/bs/allcommoditys", &bs.BSHomeController{}, "*:AllCommoditys")
	beego.Router("/bs/addallcommodity", &bs.BSHomeController{}, "*:AddAllCommodity")
	beego.Router("/bs/upallcommodity", &bs.BSHomeController{}, "*:UpAllCommodity")
	beego.Router("/bs/commstatistics", &bs.BSHomeController{}, "*:CommodityStatistics")

	beego.Router("/writexlsx", &controllers.XLSXController{}, "*:WriteXlsx")
	beego.Router("/readxlsx", &controllers.XLSXController{}, "*:ReadXlsx")
	beego.Router("/xlsx/eincommodity", &controllers.XLSXController{}, "*:ExportInCommodity")
	beego.Router("/xlsx/eoutcommodity", &controllers.XLSXController{}, "*:ExportOutCommodity")
	beego.Router("/xlsx/eallcommodity", &controllers.XLSXController{}, "*:ExportAllCommodity")
	beego.Router("/xlsx/iincommodity", &controllers.XLSXController{}, "*:ImportInCommodity")
	beego.Router("/xlsx/ioutcommodity", &controllers.XLSXController{}, "*:ImportOutCommodity")
	beego.Router("/xlsx/iallcommodity", &controllers.XLSXController{}, "*:ImportAllCommodity")

	//后台过滤器
	beego.InsertFilter("/admin/*", beego.BeforeRouter, filter.FilterAdmin)
	beego.InsertFilter("/bs/admin/*", beego.BeforeRouter, filter.FilterBSAdmin)
	beego.InsertFilter("/bs/*", beego.BeforeRouter, filter.FilterBSUser)
	beego.InsertFilter("/*", beego.BeforeRouter, filter.FilterHome)

	// 附件处理
	os.Mkdir("imagehosting", os.ModePerm)
	beego.Router("/imagehosting/:all", &controllers.ImageHostingController{})
	beego.Router("/imagehosting", &controllers.ImageHostingController{})

	os.Mkdir("bsfilehosting", os.ModePerm)
	beego.Router("/bsfilehosting/:all", &controllers.BsFileHostingController{})
	beego.Router("/bsfilehosting", &controllers.BsFileHostingController{})

	// beego.AutoRouter(&controllers.HomeController{})
}
