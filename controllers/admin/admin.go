package admin

import (
	"github.com/astaxie/beego"
	"suyuanweb/models"
	"suyuanweb/sutil"
	"time"
)

//后台
type AdminController struct {
	beego.Controller
}

//后台首页
func (c *AdminController) Admin() {
	account := sutil.GetAdminAccount(c.Ctx)
	c.Data["Account"] = account
	c.Layout = "layout/ladmin.html"
	c.TplName = "admin_layout_content/lc_admin.html"
}

// 登出
func (c *AdminController) Logout() {
	sutil.SaveAdminToken("", c.Ctx)
	c.Redirect("/admin/login", 302)
	return
}

// 管理员登录
func (c *AdminController) AdminLogin() {
	if c.Ctx.Input.IsGet() {
		c.Data["Error"] = ""
		c.TplName = "admin/alogin.html"
		return
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Home Post")
		account := c.Input().Get("account")
		password := c.Input().Get("password")
		// autologin := c.Input().Get("autologin")
		// beego.Debug(account, password, autologin)

		if len(account) > 0 && len(password) > 0 {
			admin, err := models.GetOneAdmin(account)
			if err != nil {
				beego.Debug(err)
				c.Data["Error"] = err.Error()
			} else {
				if admin.Id > 0 {
					if admin.Password == password {
						token := sutil.CreatAdminToken(account)
						beego.Debug("token:", token)
						err := models.UpdateAdminToken(admin.Id, token)
						if err != nil {
							c.Data["Error"] = err.Error()
						} else {
							sutil.SaveAdminToken(token, c.Ctx)
							sutil.SaveAdminAccount(account, c.Ctx)
							c.Redirect("/admin", 302)
							return
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
		c.TplName = "admin/alogin.html"
		return
	}
}

//更细日志
func (c *AdminController) UpdateLog() {
	c.TplName = "admin/aupdatelog.html"
}

func (c *AdminController) Noauth() {
	account := sutil.GetAdminAccount(c.Ctx)
	c.Data["Account"] = account
	c.Layout = "layout/ladmin.html"
	c.TplName = "admin_layout_content/lc_noauth.html"
}

//官网配置
func (c *AdminController) Configure() {
	account := sutil.GetAdminAccount(c.Ctx)

	op := c.Input().Get("op")
	switch op {
	case "up":
		id := c.Input().Get("id")
		news_check := c.Input().Get("news_check")
		entservice_check := c.Input().Get("entservice_check")
		bshownews := false
		if news_check == "true" {
			bshownews = true
		}
		err := models.UpShowNews(id, bshownews)
		if err != nil {
			beego.Error(err)
		}
		bentservice := false
		if entservice_check == "true" {
			bentservice = true
		}
		err = models.UpEntService(id, bentservice)
		if err != nil {
			beego.Error(err)
		}

		c.Redirect("/admin/configure", 302)
		return
	}

	cnfigure, err := models.GetWebCnfigure()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Cnfigure"] = cnfigure

	c.Data["Account"] = account
	c.Layout = "layout/ladmin.html"
	c.TplName = "admin_layout_content/lc_configure.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/admin/admin_configurejs.html"

}

//访问量统计
func (c *AdminController) PageView() {
	account := sutil.GetAdminAccount(c.Ctx)

	objs, err := models.GetPageLogs()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Count"] = len(objs)
	t := time.Now()
	year := t.Year()
	monthdata := make([]int, 12)
	for i := 0; i < 12; i++ {
		con, err := models.GetPageLogsFMonth(year, int(i+1))
		if err != nil {
			beego.Error(err)
		}
		monthdata[i] = int(con)
	}
	beego.Debug("monthdata:", monthdata)
	c.Data["MonthData"] = monthdata
	// beego.Debug("week:", t.Weekday()) //"Wednesday"
	// day := t.AddDate(0, 0, 1).Day()

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

	month1 := int(t1.Month())
	day1 := t1.Day()
	month2 := int(t2.Month())
	day2 := t2.Day()
	month3 := int(t3.Month())
	day3 := t3.Day()
	month4 := int(t4.Month())
	day4 := t4.Day()
	month5 := int(t5.Month())
	day5 := t5.Day()
	month6 := int(t6.Month())
	day6 := t6.Day()
	month7 := int(t7.Month())
	day7 := t7.Day()

	weekdata := make([]int, 7)
	week1, err := models.GetPageLogsFDay(year, month1, day1)
	if err != nil {
		beego.Error(err)
	}
	weekdata[0] = int(week1)

	week2, err := models.GetPageLogsFDay(year, month2, day2)
	if err != nil {
		beego.Error(err)
	}
	weekdata[1] = int(week2)

	week3, err := models.GetPageLogsFDay(year, month3, day3)
	if err != nil {
		beego.Error(err)
	}
	weekdata[2] = int(week3)
	beego.Debug("year", year, "month3:", month3, "day3:", day3)
	week4, err := models.GetPageLogsFDay(year, month4, day4)
	if err != nil {
		beego.Error(err)
	}
	weekdata[3] = int(week4)

	week5, err := models.GetPageLogsFDay(year, month5, day5)
	if err != nil {
		beego.Error(err)
	}
	weekdata[4] = int(week5)

	week6, err := models.GetPageLogsFDay(year, month6, day6)
	if err != nil {
		beego.Error(err)
	}
	weekdata[5] = int(week6)

	week7, err := models.GetPageLogsFDay(year, month7, day7)
	if err != nil {
		beego.Error(err)
	}
	weekdata[6] = int(week7)

	beego.Debug("WeekData:", weekdata)
	c.Data["WeekData"] = weekdata

	c.Data["Account"] = account
	c.Layout = "layout/ladmin.html"
	c.TplName = "admin_layout_content/lc_pageview.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/admin/admin_pageviewjs.html"
}
