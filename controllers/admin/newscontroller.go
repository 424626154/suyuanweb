package admin

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"path"
	"strings"
	"suyuanweb/models"
	"suyuanweb/sutil"
	"time"
	"strconv"
)

type NewsController struct {
	beego.Controller
}

type NewsRes struct {
	Code  int64        `json:"code"` //0 成功 1失败
	News  *models.News `json:"news"`
	Error string       `json:"error"`
}

//新闻列表
func (c *NewsController) Newss() {
	account := sutil.GetAdminAccount(c.Ctx)
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
		count,err := models.GetAdminPageNewsCount()
	if err != nil {
		beego.Error(err)
	}
	newss, err := models.GetAdminPageNewss(pageSize,i_p)
	if err != nil {
		beego.Error(err)
	}
	if err != nil {
		beego.Error(err)
	}
	page := sutil.PageUtil(count,i_p,pageSize)

	c.Data["Page"] = page
	c.Data["Newss"] = newss
	c.Layout = "layout/ladmin.html"
	c.TplName = "admin_layout_content/lc_newss.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/admin_newssjs.html"
}

//添加新闻
func (c *NewsController) AddNews() {

	if c.Ctx.Input.IsGet() {
		account := sutil.GetAdminAccount(c.Ctx)
		c.Data["Account"] = account

		c.Data["Error"] = ""
		c.Layout = "layout/ladmin.html"
		c.TplName = "admin_layout_content/lc_addnews.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/admin_addnewsjs.html"
	}

	if c.Ctx.Input.IsPost() {
		title := c.Input().Get("title")
		outline := c.Input().Get("outline")
		content := c.Input().Get("content")

		if len(title) > 0 && len(outline) > 0 && len(content) > 0 {
			showimg := ""
			_, fh, err := c.GetFile("showimg")
			if err != nil {
				beego.Error(err)
			} else {
				if fh != nil {
					beego.Debug("fh:", fh.Filename)
					tempname := fh.Filename
					t := time.Now().Unix()
					time_str := fmt.Sprintf("%d", t)
					img_uuid := uuid.NewV4().String()
					s := []string{tempname, time_str, img_uuid}
					h := md5.New()
					h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
					showimg = hex.EncodeToString(h.Sum(nil))
					err = c.SaveToFile("showimg", path.Join("imagehosting", showimg))
					if err != nil {
						beego.Error(err)
					}
				}
			}

			_, err = models.AddNews(title, outline, showimg, content)
			if err != nil {
				beego.Error(err)
				RedirectAddNews(c, "参数错误")
				return
			}
			c.Redirect("/admin/newss", 302)
		} else {
			RedirectAddNews(c, "参数错误")
			return
		}

	}
}

//重定向添加新闻
func RedirectAddNews(c *NewsController, err string) {
	c.Data["Error"] = err
	c.Layout = "layout/ladmin.html"
	c.TplName = "admin_layout_content/lc_addnews.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/addnews.html"
}

//新闻详情
func (c *NewsController) News() {
	if c.Ctx.Input.IsGet() {
		id := c.Input().Get("id")
		if len(id) == 0 {
			c.Redirect("/admin/newss", 302)
		}

		i_p := initNewsPage(c)

		op := c.Input().Get("op")
		switch op {
		case "up":
			state := c.Input().Get("state")
			bstate := true
			if state == "true" {
				bstate = false
			}
			err := models.UpdateNewsState(id, bstate)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/news?id="+id+"&p="+i_p, 302)
			return
		case "del":
			err := models.DelNews(id)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/newss&p="+i_p, 302)
			return
		case "home":
			ishome := c.Input().Get("ishome")
			bishome := true
			if ishome == "true" {
				bishome = false
			}
			err := models.UpdateNewsHome(id, bishome)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/news?id="+id+"&p="+i_p, 302)
			return
		}
		beego.Debug(c.Input())
		beego.Debug("id:", id)
		account := sutil.GetAdminAccount(c.Ctx)
		c.Data["Account"] = account

		obj, err := models.GetOneAdmin(account)
		if err != nil {
			beego.Error(err)
		}
		auth := obj.Auth
		c.Data["Auth"] = auth

		news, err := models.GetOneNewsFId(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["News"] = news
		c.Layout = "layout/ladmin.html"
		c.TplName = "admin_layout_content/lc_news.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/admin_newsjs.html"
	}
	if c.Ctx.Input.IsPost() {
		newsres := &NewsRes{Code: 1, Error: "未知错误"}
		op := c.Input().Get("op")
		switch op {
		case "up":
			id := c.Input().Get("id")
			state := c.Input().Get("state")
			bstate := true
			if state == "true" {
				bstate = false
			}
			err := models.UpdateNewsState(id, bstate)
			if err != nil {
				beego.Error(err)
			}
			news, err := models.GetOneNewsFId(id)
			if err != nil {
				beego.Error(err)
			}
			c.Data["News"] = news
			newsres.Code = 0
			newsres.News = news
			newsres_json, err := json.Marshal(newsres)
			if err != nil {
				beego.Error(err)
			}
			beego.Debug("new post op = up json:", newsres_json)
			c.Ctx.WriteString(string(newsres_json))
			return
		}
	}
}

//修改新闻
func (c *NewsController) Upnews() {
	account := sutil.GetAdminAccount(c.Ctx)
	c.Data["Account"] = account

	i_p := initNewsPage(c)

	if c.Ctx.Input.IsGet() {
		id := c.Input().Get("id")
		if len(id) == 0 {
			c.Redirect("/admin/newss?i_p="+i_p, 302)
		}
		news, err := models.GetOneNewsFId(id)
		if err != nil {
			c.Redirect("/admin/newss?i_p="+i_p, 302)
		}
		c.Data["Error"] = ""
		c.Data["News"] = news
		c.Layout = "layout/ladmin.html"
		c.TplName = "admin_layout_content/lc_upnews.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/admin_upnewsjs.html"
	}
	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		title := c.Input().Get("title")
		outline := c.Input().Get("outline")
		content := c.Input().Get("content")
		_, fh, err := c.GetFile("showimg")
		// beego.Debug("上传图片:", fh)
		if err != nil {
			beego.Error(err)
		}
		if len(title) > 0 && len(outline) > 0 && len(content) > 0 {
			showimg := ""
			if fh != nil {
				beego.Debug("fh:", fh.Filename)
				tempname := fh.Filename
				t := time.Now().Unix()
				time_str := fmt.Sprintf("%d", t)
				img_uuid := uuid.NewV4().String()
				s := []string{tempname, time_str, img_uuid}
				h := md5.New()
				h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
				showimg = hex.EncodeToString(h.Sum(nil))
				err = c.SaveToFile("showimg", path.Join("imagehosting", showimg))
				if err != nil {
					beego.Error(err)
				}
			}
			if len(showimg) > 0 {
				err := models.UpdateNewsImg(id, showimg)
				if err != nil {
					beego.Error(err)
				}
			}
			err := models.UpdateNewsInfo(id, title, outline, content)
			if err != nil {
				beego.Error(err)
				c.Redirect("/admin/upnews?id="+id+"&p="+i_p, 302)
			}
			c.Redirect("/admin/news?id="+id+"&p="+i_p, 302)
		} else {
			c.Redirect("/admin/upnews?id="+id+"&p="+i_p, 302)
			return
		}
	}
}


func initNewsPage(c *NewsController) string{
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