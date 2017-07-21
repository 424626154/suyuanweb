package admin

import (
	"github.com/astaxie/beego"
	"suyuanweb/models"
	"suyuanweb/sutil"
)

type AdminAccountController struct {
	beego.Controller
}

func (c *AdminAccountController) Admins() {
	objs, err := models.GetAdmins()
	if err != nil {
		beego.Error(err)
	}
	account := sutil.GetAdminAccount(c.Ctx)
	obj, err := models.GetOneAdmin(account)
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("objs:", objs)
	c.Data["MyAdmin"] = obj
	c.Data["Admins"] = objs

	c.Data["Account"] = account
	c.Layout = "layout/ladmin.html"
	c.TplName = "admin_layout_content/lc_admins.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/admin_adminsjs.html"
}

func (c *AdminAccountController) AddAdmin() {
	account := sutil.GetAdminAccount(c.Ctx)
	c.Data["Account"] = account
	if c.Ctx.Input.IsGet() {
		c.Data["Error"] = ""
		c.Layout = "layout/ladmin.html"
		c.TplName = "admin_layout_content/lc_addadmin.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/admin_addadminjs.html"
	}
	if c.Ctx.Input.IsPost() {
		account := c.Input().Get("account")
		password := c.Input().Get("password")
		confirmpwd := c.Input().Get("confirmpwd")
		rd := c.Input().Get("rd")

		if len(account) > 0 && len(password) > 0 && len(confirmpwd) > 0 && password == confirmpwd {
			obj, err := models.GetOneAdmin(account)
			if err != nil {
				beego.Error(err)
			}
			if obj.Account == account {
				c.Data["Error"] = "账号已存在"
				c.Layout = "layout/ladmin.html"
				c.TplName = "admin_layout_content/lc_addadmin.html"
				c.LayoutSections = make(map[string]string)
				c.LayoutSections["Scripts"] = "scripts/admin_addadminjs.html"
				return
			}
			auth := 0
			if rd == "rd0" {
				auth = 0
			} else if rd == "rd1" {
				auth = 1
			} else if rd == "rd2" {
				auth = 2
			}
			_, err = models.AddAdmin(account, password, auth)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/admins", 302)
		} else {
			err_str := "参数错误"
			if len(account) == 0 {
				err_str = "账号不能为空"
			} else if len(password) == 0 {
				err_str = "密码不能为空"
			} else if len(confirmpwd) == 0 {
				err_str = "确认密码不能为空"
			} else if password != confirmpwd {
				err_str = "两次输入密码请保持一致"
			}
			c.Data["Error"] = err_str
			c.Layout = "layout/ladmin.html"
			c.TplName = "admin_layout_content/lc_addadmin.html"
			c.LayoutSections = make(map[string]string)
			c.LayoutSections["Scripts"] = "scripts/admin_addadminjs.html"
		}
	}
}

func (c *AdminAccountController) UpAdminPwd() {
	account := sutil.GetAdminAccount(c.Ctx)
	c.Data["Account"] = account
	if c.Ctx.Input.IsGet() {
		id := c.Input().Get("id")
		err_str := c.Input().Get("err")
		if len(id) > 0 {
			obj, err := models.GetOneAdminFId(id)
			if err != nil {
				beego.Error(err)
			}
			beego.Debug("obj:", obj)
			c.Data["Error"] = err_str
			c.Data["Admin"] = obj
			c.Layout = "layout/ladmin.html"
			c.TplName = "admin_layout_content/lc_upadminpwd.html"
			c.LayoutSections = make(map[string]string)
			c.LayoutSections["Scripts"] = "scripts/admin_upadminpwdjs.html"
		} else {
			c.Redirect("/admin/admins", 302)
		}
	}
	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		account := c.Input().Get("account")
		password := c.Input().Get("password")
		newpassword := c.Input().Get("newpassword")
		newconfirmpwd := c.Input().Get("newconfirmpwd")
		if len(id) > 0 && len(account) > 0 && len(password) > 0 && len(newpassword) > 0 && len(newconfirmpwd) > 0 && newpassword == newconfirmpwd {
			obj, err := models.GetOneAdminFId(id)
			if err != nil {
				beego.Error(err)
			}
			if obj.Password != password {
				err_str := "原始密码错误"
				c.Redirect("/admin/upadminpwd?err="+err_str+"&id="+id, 302)
			}
			err = models.UpdateAdminPassword(id, newpassword)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/admins", 302)
		} else {
			err_str := "参数错误"
			if len(password) == 0 {
				err_str = "请输入原始密码"
			} else if len(newpassword) == 0 {
				err_str = "请输入新密码"
			} else if len(newconfirmpwd) == 0 {
				err_str = "请输入确认新密码"
			} else if newpassword != newconfirmpwd {
				err_str = "请确保两次密码一致"
			}
			c.Redirect("/admin/upadminpwd?err="+err_str+"&id="+id, 302)
		}
	}
}

func (c *AdminAccountController) UpAdminAuth() {
	account := sutil.GetAdminAccount(c.Ctx)
	c.Data["Account"] = account
	if c.Ctx.Input.IsGet() {
		id := c.Input().Get("id")
		err_str := c.Input().Get("err")
		if len(id) > 0 {
			obj, err := models.GetOneAdminFId(id)
			if err != nil {
				beego.Error(err)
			}
			beego.Debug("obj:", obj)
			c.Data["Error"] = err_str
			c.Data["Admin"] = obj
			c.Layout = "layout/ladmin.html"
			c.TplName = "admin_layout_content/lc_upadminauth.html"
			c.LayoutSections = make(map[string]string)
			c.LayoutSections["Scripts"] = "scripts/admin_upadminauthjs.html"
		} else {
			c.Redirect("/admin/admins", 302)
		}
	}
	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		rd := c.Input().Get("rd")
		if len(id) > 0 && len(rd) > 0 {
			_, err := models.GetOneAdminFId(id)
			if err != nil {
				beego.Error(err)
			}
			auth := 0
			if rd == "rd0" {
				auth = 0
			} else if rd == "rd1" {
				auth = 1
			} else if rd == "rd2" {
				auth = 2
			}
			err = models.UpdateAdminAuth(id, auth)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/admins", 302)
		} else {
			err_str := "参数错误"
			c.Redirect("/admin/upadminauth?err="+err_str+"&id="+id, 302)
		}

	}
}
