package bs

import (
	"github.com/astaxie/beego"
	"strconv"
	"suyuanweb/models"
	"suyuanweb/models/bs"
	"suyuanweb/sutil"
	"time"
)

//企业管理系统
type BSHomeController struct {
	beego.Controller
}

func (c *BSHomeController) Index() {
	account := sutil.GetBSUserAccount(c.Ctx)
	c.Data["Account"] = account

	c.Layout = "layout/lbs.html"
	c.TplName = "blc/home/index.html"
}

func (c *BSHomeController) Login() {
	if c.Ctx.Input.IsGet() {
		c.Data["Error"] = ""
		c.TplName = "bs/blogin.html"
	}
	if c.Ctx.Input.IsPost() {
		account := c.Input().Get("account")
		password := c.Input().Get("password")
		if len(account) > 0 && len(password) > 0 {
			admin, err := models.GetOneBSUser(account)
			if err != nil {
				beego.Debug(err)
				c.Data["Error"] = err.Error()
			} else {
				if admin.Id > 0 {
					if admin.Password == password {
						token := sutil.CreatAdminToken(account)
						beego.Debug("token:", token)
						err := models.UpdateBSUserToken(admin.Id, token)
						if err != nil {
							c.Data["Error"] = err.Error()
						} else {
							sutil.SaveBSUserToken(token, c.Ctx)
							sutil.SaveBSUserAccount(account, c.Ctx)
							c.Redirect("/bs", 302)
						}

					} else {
						c.Data["Error"] = "密码错误"
					}
				} else {
					c.Data["Error"] = "账号不存在"
				}

			}
		} else {

		}
		c.TplName = "bs/blogin.html"
	}

}

func (c *BSHomeController) Logout() {
	sutil.SaveBSUserToken("", c.Ctx)
	c.Redirect("/bs/login", 302)
	return
}

func (c *BSHomeController) Templates() {
	account := sutil.GetBSUserAccount(c.Ctx)
	c.Data["Account"] = account

	search_str := c.Input().Get("search")
	objs := make([]models.BsTemplate, 0)
	if len(search_str) > 0 {
		temp_objs, err := models.GetBsHomeSsearchTemplates(search_str)
		if err != nil {
			beego.Error(err)
		}
		objs = temp_objs
	} else {
		temp_objs, err := models.GetBsHomeTemplates()
		if err != nil {
			beego.Error(err)
		}
		objs = temp_objs
	}

	beego.Debug(objs)
	c.Data["Temaplates"] = objs
	c.Layout = "layout/lbs.html"
	c.TplName = "blc/home/templates.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/bs/home/templatesjs.html"
}

/******入库******/
// 入库登记表
func (c *BSHomeController) InCommoditys() {
	account := sutil.GetBSUserAccount(c.Ctx)
	c.Data["Account"] = account

	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		p := c.Input().Get("p")
		if len(id) > 0 {
			err := bs.DelInCommodity(id)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/bs/incommoditys?p="+p, 302)
		}
		return
	}

	// objs,err := bs.GetWareregs()
	// if err != nil {
	// 	beego.Error()
	// }
	// c.Data["Wareregs"] = objs

	p := c.Input().Get("p")
	i_p := 1
	if len(p) > 0 {
		temp_p, err := strconv.Atoi(p)
		if err != nil {
			beego.Error(err)
		} else {
			i_p = temp_p
		}
	}

	pageSize := 10
	count, err := bs.GetPageInCommodityCount()
	if err != nil {
		beego.Error(err)
	}
	products, err := bs.GetPageInCommoditys(pageSize, i_p)
	beego.Debug("p:", p, "products:", len(products))
	if err != nil {
		beego.Error(err)
	}
	page := sutil.PageUtil(count, i_p, pageSize)

	c.Data["Wareregs"] = products
	c.Data["Page"] = page

	templet, err := bs.GetCommodityTemplet()
	if err != nil {
		beego.Error(err)
	}
	c.Data["InTemplet"] = templet.InCommodityTemplet
	c.Layout = "layout/lbs.html"
	c.TplName = "blc/home/incommoditys.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/bs/home/incommoditysjs.html"
}

//添加入库登记表
func (c *BSHomeController) AddInCommodity() {
	account := sutil.GetBSUserAccount(c.Ctx)
	c.Data["Account"] = account
	if c.Ctx.Input.IsGet() {
		names, err := bs.GetOrderCommodityNames()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Names"] = names

		specs, err := bs.GetOrderCommoditySpecs()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Specs"] = specs

		c.Layout = "layout/lbs.html"
		c.TplName = "blc/home/addincommodity.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/bs/home/addincommodityjs.html"
	}
	if c.Ctx.Input.IsPost() {
		time := c.Input().Get("time")
		ordernumber := c.Input().Get("ordernumber")
		name := c.Input().Get("name")
		unit := c.Input().Get("unit")
		spec := c.Input().Get("spec")
		unitprice := c.Input().Get("unitprice")
		number := c.Input().Get("number")
		// total := c.Input().Get("total")
		operator := c.Input().Get("operator")
		remarks := c.Input().Get("remarks")
		merge := c.Input().Get("merge")
		// beego.Debug("time:",time,"ordernumber:",ordernumber,"name:",name,"spec:",spec,"unitprice:",unitprice,"number:",number,"total:",total,"operator:",operator,"remarks:",remarks)
		if len(time) > 0 && len(number) > 0 {
			_, err := bs.AddInCommodity(time, ordernumber, name, spec, unit, unitprice, number, operator, remarks)
			if err != nil {
				beego.Debug(err)
			}
			if merge == "on" {
				iunitprice, err := strconv.Atoi(unitprice)
				if err != nil {
					beego.Error(err)
				}
				inumber, err := strconv.Atoi(number)
				if err != nil {
					beego.Error(err)
				}
				_, err = bs.MergeAllCommodity(name, spec, unit, time, inumber, iunitprice, 0, 0)
				if err != nil {
					beego.Debug(err)
				}
			}
			c.Redirect("/bs/incommoditys", 302)
		} else {

		}
	}
}

func (c *BSHomeController) UpInCommodity() {
	account := sutil.GetBSUserAccount(c.Ctx)
	c.Data["Account"] = account
	i_str := initCommodityPage(c)
	if c.Ctx.Input.IsGet() {
		names, err := bs.GetOrderCommodityNames()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Names"] = names

		specs, err := bs.GetOrderCommoditySpecs()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Specs"] = specs

		id := c.Input().Get("id")
		if len(id) > 0 {
			obj, err := bs.GetOneInCommodityFId(id)
			if err != nil {
				beego.Error(err)
			}
			c.Data["Warare"] = obj
		} else {
			c.Redirect("/bs/incommoditys", 302)
		}

		c.Layout = "layout/lbs.html"
		c.TplName = "blc/home/upincommodity.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/bs/home/upincommodityjs.html"
	}
	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		time := c.Input().Get("time")
		ordernumber := c.Input().Get("ordernumber")
		name := c.Input().Get("name")
		unit := c.Input().Get("unit")
		spec := c.Input().Get("spec")
		unitprice := c.Input().Get("unitprice")
		number := c.Input().Get("number")
		// total := c.Input().Get("total")
		operator := c.Input().Get("operator")
		remarks := c.Input().Get("remarks")

		if len(id) > 0 && len(time) > 0 && len(number) > 0 {
			err := bs.UpInCommodity(id, time, ordernumber, name, spec, unit, unitprice, number, operator, remarks)
			if err != nil {
				beego.Debug(err)
			}
			c.Redirect("/bs/incommoditys?p="+i_str, 302)
		} else {

		}
	}
}

/******入库 end ******/
func (c *BSHomeController) Export() {
	account := sutil.GetBSUserAccount(c.Ctx)
	c.Data["Account"] = account
	p := c.Input().Get("p")
	i_p := 1
	if len(p) > 0 {
		temp_p, err := strconv.Atoi(p)
		if err != nil {
			beego.Error(err)
		} else {
			i_p = temp_p
		}
	}

	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		err := bs.DelExports(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/bs/export?p="+p, 302)
		return
	}
	pageSize := 10
	count, err := bs.GetPageExportCount()
	if err != nil {
		beego.Error(err)
	}
	products, err := bs.GetPageExports(pageSize, i_p)
	beego.Debug("p:", p, "products:", len(products))
	if err != nil {
		beego.Error(err)
	}
	page := sutil.PageUtil(count, i_p, pageSize)

	c.Data["Exports"] = products
	c.Data["Page"] = page

	c.Layout = "layout/lbs.html"
	c.TplName = "blc/home/export.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/bs/home/exportjs.html"
}

/******出库******/
// 入库登记表
func (c *BSHomeController) OutCommoditys() {
	account := sutil.GetBSUserAccount(c.Ctx)
	c.Data["Account"] = account

	op := c.Input().Get("op")
	p := c.Input().Get("p")
	switch op {
	case "del":
		id := c.Input().Get("id")

		if len(id) > 0 {
			err := bs.DelOutCommodity(id)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/bs/outcommoditys?p="+p, 302)
		}
		return
	}

	i_p := 1
	if len(p) > 0 {
		temp_p, err := strconv.Atoi(p)
		if err != nil {
			beego.Error(err)
		} else {
			i_p = temp_p
		}
	}

	pageSize := 10
	count, err := bs.GetPageOutCommodityCount()
	if err != nil {
		beego.Error(err)
	}
	products, err := bs.GetPageOutCommoditys(pageSize, i_p)
	beego.Debug("p:", p, "products:", len(products))
	if err != nil {
		beego.Error(err)
	}
	page := sutil.PageUtil(count, i_p, pageSize)

	c.Data["OutCommoditys"] = products
	c.Data["Page"] = page

	templet, err := bs.GetCommodityTemplet()
	if err != nil {
		beego.Error(err)
	}
	c.Data["OutTemplet"] = templet.OutCommodityTemplet

	c.Layout = "layout/lbs.html"
	c.TplName = "blc/home/outcommoditys.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/bs/home/outcommoditysjs.html"
}

//添加入库登记表
func (c *BSHomeController) AddOutCommodity() {
	account := sutil.GetBSUserAccount(c.Ctx)
	c.Data["Account"] = account
	if c.Ctx.Input.IsGet() {
		names, err := bs.GetOrderCommodityNames()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Names"] = names

		specs, err := bs.GetOrderCommoditySpecs()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Specs"] = specs

		c.Layout = "layout/lbs.html"
		c.TplName = "blc/home/addoutcommodity.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/bs/home/addoutcommodityjs.html"
	}
	if c.Ctx.Input.IsPost() {
		time := c.Input().Get("time")
		ordernumber := c.Input().Get("ordernumber")
		name := c.Input().Get("name")
		unit := c.Input().Get("unit")
		spec := c.Input().Get("spec")
		unitprice := c.Input().Get("unitprice")
		number := c.Input().Get("number")
		// total := c.Input().Get("total")
		operator := c.Input().Get("operator")
		remarks := c.Input().Get("remarks")
		merge := c.Input().Get("merge")
		// beego.Debug("time:",time,"ordernumber:",ordernumber,"name:",name,"spec:",spec,"unitprice:",unitprice,"number:",number,"total:",total,"operator:",operator,"remarks:",remarks)
		if len(time) > 0 && len(number) > 0 {
			_, err := bs.AddOutCommodity(time, ordernumber, name, spec, unit, unitprice, number, operator, remarks)
			if err != nil {
				beego.Debug(err)
			}
			if merge == "on" {
				iunitprice, err := strconv.Atoi(unitprice)
				if err != nil {
					beego.Error(err)
				}
				inumber, err := strconv.Atoi(number)
				if err != nil {
					beego.Error(err)
				}
				_, err = bs.MergeAllCommodity(name, spec, unit, time, 0, 0, inumber, iunitprice)
				if err != nil {
					beego.Debug(err)
				}
			}
			c.Redirect("/bs/outcommoditys", 302)
		} else {

		}
	}
}

func (c *BSHomeController) UpOutCommodity() {
	account := sutil.GetBSUserAccount(c.Ctx)
	c.Data["Account"] = account
	i_str := initCommodityPage(c)
	if c.Ctx.Input.IsGet() {
		names, err := bs.GetOrderCommodityNames()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Names"] = names

		specs, err := bs.GetOrderCommoditySpecs()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Specs"] = specs

		id := c.Input().Get("id")
		if len(id) > 0 {
			obj, err := bs.GetOneOutCommodityFId(id)
			if err != nil {
				beego.Error(err)
			}
			c.Data["Warare"] = obj
		} else {
			c.Redirect("/bs/incommoditys", 302)
		}

		c.Layout = "layout/lbs.html"
		c.TplName = "blc/home/upoutcommodity.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/bs/home/upoutcommodityjs.html"
	}
	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		time := c.Input().Get("time")
		ordernumber := c.Input().Get("ordernumber")
		name := c.Input().Get("name")
		unit := c.Input().Get("unit")
		spec := c.Input().Get("spec")
		unitprice := c.Input().Get("unitprice")
		number := c.Input().Get("number")
		// total := c.Input().Get("total")
		operator := c.Input().Get("operator")
		remarks := c.Input().Get("remarks")

		if len(id) > 0 && len(time) > 0 && len(number) > 0 {
			err := bs.UpOutCommodity(id, time, ordernumber, name, unit, spec, unitprice, number, operator, remarks)
			if err != nil {
				beego.Debug(err)
			}
			c.Redirect("/bs/outcommoditys?p="+i_str, 302)
		} else {

		}
	}
}

func initCommodityPage(c *BSHomeController) string {
	p := c.Input().Get("p")
	i_p := 1
	i_str := "1"
	if len(p) > 0 {
		temp_p, err := strconv.Atoi(p)
		if err != nil {
			beego.Error(err)
		} else {
			i_p = temp_p
		}
		i_str = p
	}
	c.Data["Page"] = i_p
	return i_str
}

/******出库 end*****/

func (c *BSHomeController) AllCommoditys() {
	account := sutil.GetBSUserAccount(c.Ctx)
	c.Data["Account"] = account

	op := c.Input().Get("op")
	p := c.Input().Get("p")
	switch op {
	case "del":
		id := c.Input().Get("id")

		if len(id) > 0 {
			err := bs.DelAllCommodity(id)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/bs/allcommoditys?p="+p, 302)
		}
		return
	}

	i_p := 1
	if len(p) > 0 {
		temp_p, err := strconv.Atoi(p)
		if err != nil {
			beego.Error(err)
		} else {
			i_p = temp_p
		}
	}

	pageSize := 10
	count, err := bs.GetPageAllCommodityCount()
	if err != nil {
		beego.Error(err)
	}
	products, err := bs.GetPageAllCommoditys(pageSize, i_p)
	beego.Debug("p:", p, "products:", len(products))
	if err != nil {
		beego.Error(err)
	}
	page := sutil.PageUtil(count, i_p, pageSize)

	c.Data["AllCommoditys"] = products
	c.Data["Page"] = page

	c.Layout = "layout/lbs.html"
	c.TplName = "blc/home/allcommodity.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/bs/home/allcommodityjs.html"
}

func (c *BSHomeController) AddAllCommodity() {
	account := sutil.GetBSUserAccount(c.Ctx)
	c.Data["Account"] = account

	if c.Ctx.Input.IsGet() {
		names, err := bs.GetOrderCommodityNames()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Names"] = names

		specs, err := bs.GetOrderCommoditySpecs()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Specs"] = specs

		c.Layout = "layout/lbs.html"
		c.TplName = "blc/home/addallcommodity.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/bs/home/addallcommodityjs.html"
	}
	if c.Ctx.Input.IsPost() {
		cdate := c.Input().Get("cdate")
		name := c.Input().Get("name")
		spec := c.Input().Get("spec")
		unit := c.Input().Get("unit")
		startnumber := c.Input().Get("startnumber")
		starttotal := c.Input().Get("starttotal")
		inunitprice := c.Input().Get("inunitprice")
		innumber := c.Input().Get("innumber")
		outunitprice := c.Input().Get("outunitprice")
		outnumber := c.Input().Get("outnumber")
		if len(cdate) > 0 && len(startnumber) > 0 && len(starttotal) > 0 {
			istartnumber, err := strconv.Atoi(startnumber)
			if err != nil {
				beego.Error(err)
			}
			istarttotal, err := strconv.Atoi(starttotal)
			if err != nil {
				beego.Error(err)
			}
			iinunitprice, err := strconv.Atoi(inunitprice)
			if err != nil {
				beego.Error(err)
			}
			iinnumber, err := strconv.Atoi(innumber)
			if err != nil {
				beego.Error(err)
			}
			ioutunitprice, err := strconv.Atoi(outunitprice)
			if err != nil {
				beego.Error(err)
			}
			ioutnumber, err := strconv.Atoi(outnumber)
			if err != nil {
				beego.Error(err)
			}
			if iinunitprice > 0 && iinnumber > 0 || ioutunitprice > 0 && ioutnumber > 0 {
				_, err = bs.AddAllCommodity(name, spec, unit, cdate, istartnumber, istarttotal, iinunitprice, iinnumber, ioutunitprice, ioutnumber)
				if err != nil {
					beego.Debug(err)
				}
			}
			c.Redirect("/bs/allcommoditys", 302)
		} else {

		}
	}
}

func (c *BSHomeController) UpAllCommodity() {
	account := sutil.GetBSUserAccount(c.Ctx)
	c.Data["Account"] = account

	i_str := initCommodityPage(c)
	if c.Ctx.Input.IsGet() {
		names, err := bs.GetOrderCommodityNames()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Names"] = names

		specs, err := bs.GetOrderCommoditySpecs()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Specs"] = specs

		id := c.Input().Get("id")
		if len(id) > 0 {
			obj, err := bs.GetOneAllCommodityFId(id)
			if err != nil {
				beego.Error(err)
			}
			c.Data["Commodity"] = obj
		} else {
			c.Redirect("/bs/allcommoditys", 302)
		}

		c.Layout = "layout/lbs.html"
		c.TplName = "blc/home/upallcommodity.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/bs/home/upallcommodityjs.html"
	}
	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		cdate := c.Input().Get("cdate")
		name := c.Input().Get("name")
		spec := c.Input().Get("spec")
		unit := c.Input().Get("unit")
		startnumber := c.Input().Get("startnumber")
		starttotal := c.Input().Get("starttotal")
		inunitprice := c.Input().Get("inunitprice")
		innumber := c.Input().Get("innumber")
		outunitprice := c.Input().Get("outunitprice")
		outnumber := c.Input().Get("outnumber")
		if len(cdate) > 0 && len(startnumber) > 0 && len(starttotal) > 0 {
			istartnumber, err := strconv.Atoi(startnumber)
			if err != nil {
				beego.Error(err)
			}
			istarttotal, err := strconv.Atoi(starttotal)
			if err != nil {
				beego.Error(err)
			}
			iinunitprice, err := strconv.Atoi(inunitprice)
			if err != nil {
				beego.Error(err)
			}
			iinnumber, err := strconv.Atoi(innumber)
			if err != nil {
				beego.Error(err)
			}
			ioutunitprice, err := strconv.Atoi(outunitprice)
			if err != nil {
				beego.Error(err)
			}
			ioutnumber, err := strconv.Atoi(outnumber)
			if err != nil {
				beego.Error(err)
			}
			if iinunitprice > 0 && iinnumber > 0 || ioutunitprice > 0 && ioutnumber > 0 {
				_, err := bs.UpAllCommodity(id, name, spec, unit, cdate, istartnumber, iinunitprice, istarttotal, iinnumber, ioutunitprice, ioutnumber)
				if err != nil {
					beego.Debug(err)
				}
			}
			c.Redirect("/bs/allcommoditys?p="+i_str, 302)
		} else {

		}
	}
}

func (c *BSHomeController) CommodityStatistics() {
	account := sutil.GetBSUserAccount(c.Ctx)
	c.Data["Account"] = account

	t := time.Now()
	year := t.Year() - 2000
	monthindata := make([]int, 12)
	monthoutdata := make([]int, 12)
	monthalldata := make([]int, 12)
	for i := 0; i < 12; i++ {
		in, err := bs.GetInCommoditysFMonth(year, int(i+1))
		if err != nil {
			beego.Error(err)
		}
		monthindata[i] = int(in)
		out, err := bs.GetOutCommoditysFMonth(year, int(i+1))
		if err != nil {
			beego.Error(err)
		}
		monthoutdata[i] = int(out)
		all, err := bs.GetAllCommoditysFMonth(year, int(i+1))
		if err != nil {
			beego.Error(err)
		}
		monthalldata[i] = int(all)
	}

	c.Data["InCommMonthData"] = monthindata
	c.Data["OutCommMonthData"] = monthoutdata
	c.Data["AllCommMonthData"] = monthalldata

	t1 := t
	t2 := t
	t3 := t
	t4 := t
	t5 := t
	t6 := t
	t7 := t
	week := t.Weekday().String()
	switch week {
	case "Monday": //1
		t1 = t
		t2 = t.AddDate(0, 0, 1)
		t3 = t.AddDate(0, 0, 2)
		t4 = t.AddDate(0, 0, 3)
		t5 = t.AddDate(0, 0, 4)
		t6 = t.AddDate(0, 0, 5)
		t7 = t.AddDate(0, 0, 6)
		break
	case "Tuesday": //2
		t1 = t.AddDate(0, 0, -1)
		t2 = t
		t3 = t.AddDate(0, 0, 1)
		t4 = t.AddDate(0, 0, 2)
		t5 = t.AddDate(0, 0, 3)
		t6 = t.AddDate(0, 0, 4)
		t7 = t.AddDate(0, 0, 5)
		break
	case "Wednesday": //3
		t1 = t.AddDate(0, 0, -2)
		t2 = t.AddDate(0, 0, -1)
		t3 = t
		t4 = t.AddDate(0, 0, 1)
		t5 = t.AddDate(0, 0, 2)
		t6 = t.AddDate(0, 0, 3)
		t7 = t.AddDate(0, 0, 4)
		break
	case "Thursday": //4
		t1 = t.AddDate(0, 0, -3)
		t2 = t.AddDate(0, 0, -2)
		t3 = t.AddDate(0, 0, -1)
		t4 = t
		t5 = t.AddDate(0, 0, 1)
		t6 = t.AddDate(0, 0, 2)
		t7 = t.AddDate(0, 0, 3)
		break
	case "Friday": //5
		t1 = t.AddDate(0, 0, -4)
		t2 = t.AddDate(0, 0, -3)
		t3 = t.AddDate(0, 0, -2)
		t4 = t.AddDate(0, 0, -1)
		t5 = t
		t6 = t.AddDate(0, 0, 1)
		t7 = t.AddDate(0, 0, 2)
		break
	case "Saturday": //6
		t1 = t.AddDate(0, 0, -5)
		t2 = t.AddDate(0, 0, -4)
		t3 = t.AddDate(0, 0, -3)
		t4 = t.AddDate(0, 0, -2)
		t5 = t.AddDate(0, 0, -1)
		t6 = t
		t7 = t.AddDate(0, 0, 1)
		break
	case "Sunday": //7
		t1 = t.AddDate(0, 0, -6)
		t2 = t.AddDate(0, 0, -5)
		t3 = t.AddDate(0, 0, -4)
		t4 = t.AddDate(0, 0, -3)
		t5 = t.AddDate(0, 0, -2)
		t6 = t.AddDate(0, 0, -1)
		t7 = t
		break
	}
	months := make([]int, 7)
	days := make([]int, 7)
	months[0] = int(t1.Month())
	days[0] = t1.Day()
	months[1] = int(t2.Month())
	days[1] = t2.Day()
	months[2] = int(t3.Month())
	days[2] = t3.Day()
	months[3] = int(t4.Month())
	days[3] = t4.Day()
	months[4] = int(t5.Month())
	days[4] = t5.Day()
	months[5] = int(t6.Month())
	days[5] = t6.Day()
	months[6] = int(t7.Month())
	days[6] = t7.Day()

	indata := make([]int, 7)
	outdata := make([]int, 7)
	alldata := make([]int, 7)
	for i := 0; i < 7; i++ {
		in, err := bs.GetInCommoditysFDay(year, months[i], days[i])
		if err != nil {
			beego.Error(err)
		}
		indata[i] = int(in)
		out, err := bs.GetInCommoditysFDay(year, months[i], days[i])
		if err != nil {
			beego.Error(err)
		}
		outdata[i] = int(out)
		all, err := bs.GetInCommoditysFDay(year, months[i], days[i])
		if err != nil {
			beego.Error(err)
		}
		alldata[i] = int(all)
	}

	c.Data["InCommData"] = indata
	c.Data["OutCommData"] = outdata
	c.Data["AllCommData"] = alldata

	c.Layout = "layout/lbs.html"
	c.TplName = "blc/home/commstatistics.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/bs/home/commstatisticsjs.html"
}
