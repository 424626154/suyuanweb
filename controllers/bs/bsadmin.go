package bs

import (
	"github.com/astaxie/beego"
	"suyuanweb/models"
	"suyuanweb/models/bs"
	"suyuanweb/sutil"
	"path"
	"strconv"
	"time"
)

// 工作系统后台


type BsAdminController struct {
	beego.Controller
}


func (c *BsAdminController) Index(){
	account := sutil.GetBSAdminAccount(c.Ctx)
	c.Data["Account"] = account

	c.Layout="layout/lbsa.html"
	c.TplName = "blc/admin/index.html"
}

func (c *BsAdminController) Login(){
	if c.Ctx.Input.IsGet() {
		c.Data["Error"] = ""
		c.TplName = "bs/balogin.html"
		return
	}
	if c.Ctx.Input.IsPost() {
		account := c.Input().Get("account")
		password := c.Input().Get("password")
		if len(account) > 0 && len(password) > 0 {
				admin, err := models.GetOneBSAdmin(account)
				if err != nil {
					beego.Debug(err)
					c.Data["Error"] = err.Error()
				} else {
					if admin.Id > 0 {
						if admin.Password == password {
							token := sutil.CreatAdminToken(account)
							beego.Debug("token:", token)
							err := models.UpdateBSAdminToken(admin.Id, token)
							if err != nil {
								c.Data["Error"] = err.Error()
							} else {
								sutil.SaveBSAdminToken(token, c.Ctx)
								sutil.SaveBSAdminAccount(account, c.Ctx)
								c.Redirect("/bs/admin", 302)
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
			c.TplName = "bs/balogin.html"
	}
}

func (c *BsAdminController) Logout(){
	sutil.SaveBSAdminToken("", c.Ctx)
	c.Redirect("/bs/admin/login", 302)
	return
}


//管理员账号列表
func (c *BsAdminController) Admins(){
	account := sutil.GetBSAdminAccount(c.Ctx)
	c.Data["Account"] = account

	obj, err := models.GetOneBSAdmin(account)
	if err != nil {
		beego.Error(err)
	}
	c.Data["MyAdmin"] = obj
	beego.Debug("myadnmin:",obj)


	objs, err := models.GetBSAdmins()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Admins"] = objs

	c.Layout = "layout/lbsa.html"
	c.TplName = "blc/admin/admins.html"
}
 
func (c *BsAdminController) AddAdmin(){
	account := sutil.GetBSAdminAccount(c.Ctx)
	c.Data["Account"] = account
	if c.Ctx.Input.IsGet() {
		c.Data["Error"] = ""
		c.Layout="layout/lbsa.html"
		c.TplName = "blc/admin/addadmin.html"
	}
	if c.Ctx.Input.IsPost() {
		account := c.Input().Get("account")
		password := c.Input().Get("password")
		confirmpwd := c.Input().Get("confirmpwd")
		rd := c.Input().Get("rd")

		if len(account) > 0 && len(password) > 0 && len(confirmpwd) > 0 && password == confirmpwd {
			obj, err := models.GetOneBSAdmin(account)
			if err != nil {
				beego.Error(err)
			}
			if obj.Account == account {
				RedirectAddAdmin(c,"账号已存在")
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
			_, err = models.AddBSAdmin(account, password, auth)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/bs/admin/admins", 302)
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
			RedirectAddAdmin(c,err_str)
		}
	}
}



func (c *BsAdminController) Upapwd(){
	account := sutil.GetBSAdminAccount(c.Ctx)
	c.Data["Account"] = account

	if c.Ctx.Input.IsGet() {
		id := c.Input().Get("id")
		err_str := c.Input().Get("err")
		if len(id) > 0 {
			obj, err := models.GetOneBSAdminFId(id)
			if err != nil {
				beego.Error(err)
			}
			beego.Debug("obj:", obj)
			c.Data["Error"] = err_str
			c.Data["Admin"] = obj
			c.Layout = "layout/lbsa.html"
			c.TplName = "blc/admin/upapwd.html"
		} else {
			c.Redirect("/bs/admin/admins", 302)
		}
	}
	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		account := c.Input().Get("account")
		password := c.Input().Get("password")
		newpassword := c.Input().Get("newpassword")
		newconfirmpwd := c.Input().Get("newconfirmpwd")
		if len(id) > 0 && len(account) > 0 && len(password) > 0 && len(newpassword) > 0 && len(newconfirmpwd) > 0 && newpassword == newconfirmpwd {
			obj, err := models.GetOneBSAdminFId(id)
			if err != nil {
				beego.Error(err)
			}
			if obj.Password != password {
				err_str := "原始密码错误"
				c.Redirect("/bs/admin/upapwd?err="+err_str+"&id="+id, 302)
			}
			err = models.UpdateBSAdminPassword(id, newpassword)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/bs/admin/admins", 302)
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
			c.Redirect("/bs/admin/upapwd?err="+err_str+"&id="+id, 302)
		}
	}
}

func (c *BsAdminController) Upaauth(){
	account := sutil.GetBSAdminAccount(c.Ctx)
	c.Data["Account"] = account

	if c.Ctx.Input.IsGet() {
		id := c.Input().Get("id")
		err_str := c.Input().Get("err")
		if len(id) > 0 {
			obj, err := models.GetOneBSAdminFId(id)
			if err != nil {
				beego.Error(err)
			}
			beego.Debug("obj:", obj)
			c.Data["Error"] = err_str
			c.Data["Admin"] = obj
			c.Layout = "layout/lbsa.html"
			c.TplName = "blc/admin/upaauth.html"
		} else {
			c.Redirect("/bs/admin/admins", 302)
		}
	}
	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		rd := c.Input().Get("rd")
		if len(id) > 0 && len(rd) > 0 {
			_, err := models.GetOneBSAdminFId(id)
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
			err = models.UpdateBSAdminAuth(id, auth)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/bs/admin/admins", 302)
		} else {
			err_str := "参数错误"
			c.Redirect("/bs/admin/upaauth?err="+err_str+"&id="+id, 302)
		}
	}
}

func (c *BsAdminController) AddTemplate(){
	account := sutil.GetBSAdminAccount(c.Ctx)
	c.Data["Account"] = account

	if c.Ctx.Input.IsGet() {
		c.Data["Error"] = ""
		c.Layout="layout/lbsa.html"
		c.TplName = "blc/admin/addtemplate.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/bs_addtemplatejs.html"
	}
	if c.Ctx.Input.IsPost() {
		title := c.Input().Get("title")
		describe := c.Input().Get("describe")
		beego.Debug("title:",title)
		if len(title) > 0  {
			obj,err := models.GetBsTemplateFTitle(title)
			if err != nil{
				beego.Error(err)
				RedirectAddTemplate(c,err.Error())
				return
			}
			if obj.Title == title {
				RedirectAddTemplate(c,"模板标题已存在")
				return
			}
			file_path := "filepath"
			_, fh, err := c.GetFile(file_path)
			if err != nil {
				beego.Error(err)
				RedirectAddTemplate(c,err.Error())
				return
			}
			beego.Debug("上传文件 err:", err)
			if fh != nil {
				file_name := fh.Filename
				file_type := path.Ext(file_name)
				// tempname := fh.Filename
				 // beego.Debug("file_name:",path.Join("bsfilehosting", file_name))
				// save_file_name := name+file_type
				err = c.SaveToFile(file_path, path.Join("bsfilehosting",file_name))
				 beego.Debug("保存文件 err：",err)
				if err != nil {
					beego.Error(err)
					RedirectAddTemplate(c,err.Error())
					return
				} else {
					//图片保存成功 存储数据
					obj,err := models.AddBsTemplate(title,describe,file_name,file_type)
					if err != nil {
						beego.Error(err)
					}
					c.Data["Template"] = obj
					c.Layout="layout/lbsa.html"
					c.TplName = "blc/admin/addtemplate_success.html"
				}
			} else {
				RedirectAddTemplate(c,"获取文件失败")
				return
			}
		} else {
			RedirectAddTemplate(c,"参数错误")
				return
		}
	}
}

func (c *BsAdminController) Templates(){
	account := sutil.GetBSAdminAccount(c.Ctx)
	c.Data["Account"] = account

	p := c.Input().Get("p")
	i_p := 1
	if len(p) > 0 {
		temp_p, err := strconv.Atoi(p)
		if err != nil{
			beego.Error(err)
		}else{
			i_p = temp_p
		}
	}

	pageSize := 10
	count,err := models.GetBSAdminTemaplatesCount()
	if err != nil {
		beego.Error(err)
	}
	objs, err := models.GetBSAdminTemaplatesPage(pageSize,i_p)
	if err != nil {
		beego.Error(err)
	}
	page := sutil.PageUtil(count,i_p,pageSize)


	c.Data["Temaplates"] = objs
	c.Data["Page"] = page

	c.Layout="layout/lbsa.html"
	c.TplName = "blc/admin/templates.html"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/bs/admin/templatesjs.html"

}


func (c *BsAdminController) Template(){
	account := sutil.GetBSAdminAccount(c.Ctx)
	c.Data["Account"] = account

	p_str := initTemplatePage(c)

	id := c.Input().Get("id")
	obj,err := models.GetBsTemplateFid(id)
	if err != nil {
		beego.Error(err)
	}
	op := c.Input().Get("op")
	switch op {
	case "up":
		state := c.Input().Get("state")
		bstate := true
		if state == "true" {
			bstate = false
		}
		err := models.UpdateBsTemplateState(id, bstate)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/bs/admin/template?id="+id+"&p="+p_str, 302)
		return
	case "del":
		err := models.DelBsTemplate(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/bs/admin/templates?p="+p_str, 302)
		return
	}
	// beego.Debug(obj)
	c.Data["Template"] = obj
	c.Layout="layout/lbsa.html"
	c.TplName = "blc/admin/template.html"
}

func (c *BsAdminController) UpTemplate(){
	account := sutil.GetBSAdminAccount(c.Ctx)
	c.Data["Account"] = account

	p_str := initTemplatePage(c)

	if c.Ctx.Input.IsGet() {
		id := c.Input().Get("id")
		if len(id) == 0 {
			c.Redirect("/bs/admin/templates", 302)
		}
		template, err := models.GetBsTemplateFid(id)
		if err != nil {
			c.Redirect("/bs/admin/template", 302)
		}
		err_data := ""
		err_str := c.Input().Get("err")
		if len(err_str) > 0 {
			err_data = err_str
		}
		c.Data["Error"] = err_data
		c.Data["Template"] = template
		c.Layout = "layout/lbsa.html"
		c.TplName = "blc/admin/uptemplate.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/bs_ptemplatejs.html"
	}
	if c.Ctx.Input.IsPost(){
		id := c.Input().Get("id")
		title := c.Input().Get("title")
		describe := c.Input().Get("describe")
		if len(id) > 0 && len(title) > 0 {
			//文件上传
			file_path := "filepath"
			_, fh, err := c.GetFile(file_path)
			if err != nil {
				beego.Error(err)
			}else{
				if fh != nil {
					file_name := fh.Filename
					file_type := path.Ext(file_name)
					err = c.SaveToFile(file_path, path.Join("bsfilehosting",file_name))
					 beego.Debug("保存文件 err：",err)
					if err != nil {
						beego.Error(err)
					} else {
						//图片保存成功 存储数据
						err := models.UpBsTemplateFile(id,file_name,file_type)
						if err != nil {
							beego.Error(err)
						}
					}
				}
			}
			//内容修改
			err = models.UpBsTemplateInfo(id,title,describe)
			if err != nil {
				beego.Error(err)
				RedirectUpTemplate(c,id,err.Error())
			}
			c.Redirect("/bs/admin/template?id="+id+"&p="+p_str, 302)
		}else{
			if len(id) == 0{
				c.Redirect("/bs/admin/templates", 302)
			}else if len(title) == 0 {
				RedirectUpTemplate(c,id,"模板标题不能为空")
			}
		}
	}
}

func RedirectAddTemplate(c *BsAdminController,err string){
		c.Data["Error"] = err
		c.Layout="layout/lbsa.html"
		c.TplName = "blc/admin/addtemplate.html"
}

func RedirectUpTemplate(c *BsAdminController,id string,err string){
	c.Redirect("/bs/admin/uptemplate?id="+id+"&err="+err, 302)
}

func RedirectAddAdmin(c *BsAdminController,err string){
		c.Data["Error"] = err
		c.Layout="layout/lbsa.html"
		c.TplName = "blc/admin/addadmin.html"
}

func initTemplatePage(c *BsAdminController) string{
	p := c.Input().Get("p")
	i_p := 1
	i_str := "1"
	if len(p) > 0 {
		temp_p, err := strconv.Atoi(p)
		if err != nil{
			beego.Error(err)
		}else{
			i_p = temp_p
		}
		i_str = p
	}
	c.Data["Page"]= i_p	
	return i_str
}


func (c *BsAdminController) CommodityNames(){
	account := sutil.GetBSAdminAccount(c.Ctx)
	c.Data["Account"] = account
	objs,err := bs.GetCommodityNames()
	if err != nil {
		beego.Error(err)
	}
	c.Data["CommodityNames"] = objs
	c.Layout="layout/lbsa.html"
	c.TplName = "blc/admin/commoditynames.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/bs/admin/commoditynamesjs.html"

}

func (c *BsAdminController) CommoditySpecs(){
	account := sutil.GetBSAdminAccount(c.Ctx)
	c.Data["Account"] = account
	objs,err := bs.GetCommoditySpecs()
	if err != nil {
		beego.Error(err)
	}
	c.Data["CommoditySpecs"] = objs
	c.Layout="layout/lbsa.html"
	c.TplName = "blc/admin/commodityspecs.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/bs/admin/commodityspecsjs.html"

}


func (c *BsAdminController)CommodityAjax(){
	op := c.Input().Get("op")
	switch op {
	case "add":
		commodityname := c.Input().Get("name")
		if len(commodityname) > 0 {
				_,err := bs.AddCommodityName(commodityname)
				if err != nil{
					beego.Error(err)
				}
				ajax_json := "{\"code\":0}"
				c.Ctx.WriteString(ajax_json)
		}else{
				json_err := "参数错误"
				ajax_json := "{\"code\":1,\"err\":\""+json_err+"\"}"
				c.Ctx.WriteString(ajax_json)
		}
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) > 0 {
				err := bs.DelCommodityName(id)
				if err != nil{
					beego.Error(err)
				}
				ajax_json := "{\"code\":0}"
				c.Ctx.WriteString(ajax_json)
		}else{
				json_err := "参数错误"
				ajax_json := "{\"code\":1,\"err\":\""+json_err+"\"}"
				c.Ctx.WriteString(ajax_json)
		}
		return
	case "up":
		id := c.Input().Get("id")
		name := c.Input().Get("name")
		if len(id) > 0 && len(name) > 0{
				err := bs.UpCommodityName(id,name)
				if err != nil{
					beego.Error(err)
				}
				ajax_json := "{\"code\":0}"
				c.Ctx.WriteString(ajax_json)
			}else{
				json_err := "参数错误"
				ajax_json := "{\"code\":1,\"err\":\""+json_err+"\"}"
				c.Ctx.WriteString(ajax_json)
			}
		return 
	case "uporder":
		id := c.Input().Get("id")
		order := c.Input().Get("order")
		if len(id) > 0 && len(order) > 0{
				err := bs.UpCommodityNameOrder(id,order)
				if err != nil{
					beego.Error(err)
				}
				ajax_json := "{\"code\":0}"
				c.Ctx.WriteString(ajax_json)
			}else{
				json_err := "参数错误"
				ajax_json := "{\"code\":1,\"err\":\""+json_err+"\"}"
				c.Ctx.WriteString(ajax_json)
			}
		return
	case "addspec":
		commodityspec := c.Input().Get("spec")
		if len(commodityspec) > 0 {
				_,err := bs.AddCommoditySpec(commodityspec)
				if err != nil{
					beego.Error(err)
				}
				ajax_json := "{\"code\":0}"
				c.Ctx.WriteString(ajax_json)
		}else{
				json_err := "参数错误"
				ajax_json := "{\"code\":1,\"err\":\""+json_err+"\"}"
				c.Ctx.WriteString(ajax_json)
		}
		return
	case "delspec":
		id := c.Input().Get("id")
		if len(id) > 0 {
				err := bs.DelCommoditySpec(id)
				if err != nil{
					beego.Error(err)
				}
				ajax_json := "{\"code\":0}"
				c.Ctx.WriteString(ajax_json)
		}else{
				json_err := "参数错误"
				ajax_json := "{\"code\":1,\"err\":\""+json_err+"\"}"
				c.Ctx.WriteString(ajax_json)
		}
		return
	case "upspec":
		id := c.Input().Get("id")
		spec := c.Input().Get("spec")
		if len(id) > 0 && len(spec) > 0{
				err := bs.UpCommoditySpec(id,spec)
				if err != nil{
					beego.Error(err)
				}
				ajax_json := "{\"code\":0}"
				c.Ctx.WriteString(ajax_json)
			}else{
				json_err := "参数错误"
				ajax_json := "{\"code\":1,\"err\":\""+json_err+"\"}"
				c.Ctx.WriteString(ajax_json)
			}
		return 
	case "upspecorder":
		id := c.Input().Get("id")
		order := c.Input().Get("order")
		if len(id) > 0 && len(order) > 0{
				err := bs.UpCommoditySpecOrder(id,order)
				if err != nil{
					beego.Error(err)
				}
				ajax_json := "{\"code\":0}"
				c.Ctx.WriteString(ajax_json)
			}else{
				json_err := "参数错误"
				ajax_json := "{\"code\":1,\"err\":\""+json_err+"\"}"
				c.Ctx.WriteString(ajax_json)
			}
		return		
	}

}


func (c *BsAdminController) CommodityTemplates(){
	account := sutil.GetBSAdminAccount(c.Ctx)
	c.Data["Account"] = account

	if c.Ctx.Input.IsGet() {
		obj,err := bs.GetCommodityTemplet()
		if err != nil{
			beego.Debug(err)
		}
		c.Data["Templet"] = obj

		c.Layout="layout/lbsa.html"
		c.TplName = "blc/admin/commoditytemplates.html"

		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/bs/admin/commoditytemplatesjs.html"

	}
	if c.Ctx.Input.IsPost() {
		op := c.Input().Get("op")
		file_path := "filepath"
		_, fh, err := c.GetFile(file_path)
		if err != nil {
			beego.Error(err)
			json_err := err.Error()
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
			return
		}
		if fh != nil {
			file_name := fh.Filename
			t := time.Now()
			t1 := t.Format("2006-01-02_15-04-05")
			save_file_name := t1 + "_" + file_name
			save_file_path := path.Join("bsfilehosting", save_file_name)
			err := c.SaveToFile(file_path, save_file_path)
			if err != nil {
				beego.Error(err)
				json_err := err.Error()
				ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
				c.Ctx.WriteString(ajax_json)
				return
			}
			if op == "in"{
				err = bs.AddInCommodityTemplet(save_file_path)
				if err != nil{
						beego.Error(err)
						json_err := err.Error()
						ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
						c.Ctx.WriteString(ajax_json)
						return
				}
				json_msg := "导入完成"
				ajax_json := "{\"code\":0,\"msg\":\"" + json_msg + "\"}"
				c.Ctx.WriteString(ajax_json)
			} else if op == "out"{
				err = bs.AddOutCommodityTemplet(save_file_path)
				beego.Debug("111111AddOutCommodityTemplet:",save_file_path)
				if err != nil{
						beego.Error(err)
						json_err := err.Error()
						ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
						c.Ctx.WriteString(ajax_json)
						return
				}
				json_msg := "导入完成"
				ajax_json := "{\"code\":0,\"msg\":\"" + json_msg + "\"}"
				c.Ctx.WriteString(ajax_json)
			}else if op == "all"{
				err = bs.AddAllCommodityTemplet(save_file_path)
				if err != nil{
						beego.Error(err)
						json_err := err.Error()
						ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
						c.Ctx.WriteString(ajax_json)
						return
				}
				json_msg := "导入完成"
				ajax_json := "{\"code\":0,\"msg\":\"" + json_msg + "\"}"
				c.Ctx.WriteString(ajax_json)
			}else{
						json_err := "操作类型错误"
						ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
						c.Ctx.WriteString(ajax_json)
			}
		}
	}
}

