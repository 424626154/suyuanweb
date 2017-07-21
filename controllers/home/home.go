package home

import (
	"github.com/astaxie/beego"
	// "math"
	"suyuanweb/models"
)

type HomeController struct {
	beego.Controller
}
// 首页
func (c *HomeController) Home() {

	configure := InitConfigure(c)

	banners, err := models.GetReleaseImgBanners()
	if err != nil {
		beego.Error(err)
	}
	if configure.ShowNews {
		//新闻
		newss, err := models.GetHomeNewss()
		if err != nil {
			beego.Error(err)
		}
			c.Data["Newss"] = newss
	}
	c.Data["Banners"] = banners

	//产品
	products, err := models.GetHomeProducts()
	if err != nil {
		beego.Error(err)
	}
	//滚动栏形式
	// items_len := 4
	// last_num := len(products) % items_len
	// grops_num := int(math.Ceil(float64(len(products)) / float64(items_len)))
	// beego.Debug("products:",products,"grops_num:",grops_num)
	// grops := make([][]*models.Product, grops_num)
	// for i := 0; i < grops_num; i++ {
	// 	items := make([]*models.Product, items_len)
	// 	if i == grops_num-1 {
	// 		items = make([]*models.Product, last_num)
	// 	}
	// 	beego.Debug(len(items))
	// 	for j := 0; j < len(items); j++ {
	// 		items[j] = &products[i*items_len+j]
	// 	}
	// 	grops[i] = items
	// }
	// beego.Debug(grops)
	// c.Data["Grops"] = grops
	c.Data["Products"] = products

	c.Layout = "layout/lhome.html"
	c.TplName = "home/home.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/home/homejs.html"
}

// Banner详情
func (c *HomeController) Banner() {
	InitConfigure(c)
	id := c.Input().Get("id")
	if len(id) > 0 {
		banner, err := models.GetOneImgBannerFId(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Banner"] = banner
	}
	c.Layout = "layout/lhome.html"
	c.TplName = "home/banner.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/bannerjs.html"
}

// 新闻列表
func (c *HomeController) Newss() {
	InitConfigure(c)
	objs, err := models.GetReleaseNewss()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Newss"] = objs
	c.Layout = "layout/lhome.html"
	c.TplName = "home/newss.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/home/newssjs.html"
}
// 新闻详情
func (c *HomeController) News() {
	InitConfigure(c)
	id := c.Input().Get("id")
	if len(id) == 0 {
		c.Redirect("/newss", 302)
	}
	obj, err := models.GetOneNewsFId(id)
	if err != nil {
		beego.Error(err)
	}


	older_obj, err := models.GetOlderNews(id)
	if err != nil {
		beego.Error(err)
	}
	newer_obj, err := models.GetNewerNews(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Older"] = older_obj.Id
	c.Data["Newer"] = newer_obj.Id

	c.Data["News"] = obj
	c.Layout = "layout/lhome.html"
	c.TplName = "home/news.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/home/newsjs.html"
}


// 产品列表
func (c *HomeController) Products() {
	InitConfigure(c)
	objs, err := models.GetReleaseProducts()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Newss"] = objs
	c.Layout = "layout/lhome.html"
	c.TplName = "home/products.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/home/productsjs.html"
}

// 产品详情
func (c *HomeController) Product() {
	InitConfigure(c)
	id := c.Input().Get("id")
	if len(id) == 0 {
		c.Redirect("/products", 302)
	}


	obj, err := models.GetOneProductFId(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Product"] = obj

	older_obj, err := models.GetHomeOlderProduct(obj.OrderId)
	if err != nil {
		beego.Error(err)
	}
	newer_obj, err := models.GetHomeNewerProduct(obj.OrderId)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Older"] = older_obj.Id
	c.Data["Newer"] = newer_obj.Id


	c.Layout = "layout/lhome.html"
	c.TplName = "home/product.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/home/productjs.html"
}

// 服务
func (c *HomeController) Service() {
	InitConfigure(c)
	c.Layout = "layout/lhome.html"
	c.TplName = "home/service.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/home/servicejs.html"
}

// 关于
func (c *HomeController) About() {
	InitConfigure(c)
	c.Layout = "layout/lhome.html"
	c.TplName = "home/about.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/home/aboutjs.html"
}


func InitConfigure(c *HomeController) (*models.WebCnfigure){
	configure,err := models.GetWebCnfigure()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Configure"] = configure
	return configure
}