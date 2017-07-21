package admin

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"path"
	"strings"
	"suyuanweb/models"
	"suyuanweb/sutil"
	"time"
)

type ImageBannerController struct {
	beego.Controller
}

//Banner列表
func (c *ImageBannerController) ImageBanners() {
	account := sutil.GetAdminAccount(c.Ctx)
	c.Data["Account"] = account

	objs, err := models.GetImgBanners()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Banners"] = objs
	c.Layout = "layout/ladmin.html"
	c.TplName = "admin_layout_content/lc_banners.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/admin_bannersjs.html"
}

//添加Banner
func (c *ImageBannerController) AddImageBanner() {
	if c.Ctx.Input.IsGet() {
		account := sutil.GetAdminAccount(c.Ctx)
		c.Data["Account"] = account

		c.Data["Error"] = ""
		c.Layout = "layout/ladmin.html"
		c.TplName = "admin_layout_content/lc_addbanner.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/admin/admin_addbannerjs.html"
	}

	if c.Ctx.Input.IsPost() {
		rd := c.Input().Get("rd")
		link := c.Input().Get("link")
		content := c.Input().Get("content")
		beego.Debug("rd:", rd, "link:", link, "content:", content)
		if rd == "rd0" && len(link) > 0 || rd == "rd1" && len(content) > 0 || rd == "rd2"{
			showimg := ""
			_, fh, err := c.GetFile("showimg")
			if err != nil {
				beego.Error(err)
				RedirectAddImageBanner(c, err.Error())
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
						RedirectAddImageBanner(c, err.Error())
					}
				}
				state := 0
				if rd == "rd0" {
					state = 1
				}else if rd == "rd1" {
					state = 2
				}
				_, err := models.AddImgBanner(showimg, state, content, link)
				if err != nil {
					beego.Error(err)
					RedirectAddImageBanner(c, err.Error())
				}
				c.Redirect("/admin/banners", 302)
			}
		} else {
			err_str := "参数错误"
			if rd == "rd0" && len(link) == 0 {
				err_str = "连接不能为空"
			}else if rd == "rd1" && len(content) == 0{
				err_str = "内容不能为空"	
			}
			RedirectAddImageBanner(c, err_str)
		}
	}
}

//重定向添加Banner
func RedirectAddImageBanner(c *ImageBannerController, err string) {
	c.Data["Error"] = err
	c.Layout = "layout/ladmin.html"
	c.TplName = "admin_layout_content/lc_addbanner.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/admin/admin_addbannerjs.html"
}

//Banner详情
func (c *ImageBannerController) ImageBanner() {
	id := c.Input().Get("id")
	if len(id) == 0 {
		c.Redirect("/admin/banners", 302)
	}
	op := c.Input().Get("op")
	switch op {
	case "up":
		state := c.Input().Get("state")
		bstate := true
		if state == "true" {
			bstate = false
		}
		err := models.UpdateImgBannerState(id, bstate)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/banner?id="+id, 302)
		return
	case "del":
		err := models.DelImgBanner(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/banners", 302)
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

	banner, err := models.GetOneImgBannerFId(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Banner"] = banner
	c.Layout = "layout/ladmin.html"
	c.TplName = "admin_layout_content/lc_banner.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/admin_bannerjs.html"
}

//Banner新闻
func (c *ImageBannerController) UpImageBanner() {
	if c.Ctx.Input.IsGet() {
		account := sutil.GetAdminAccount(c.Ctx)
		c.Data["Account"] = account

		id := c.Input().Get("id")
		if len(id) == 0 {
			c.Redirect("/admin/banners", 302)
		}
		obj, err := models.GetOneImgBannerFId(id)
		if err != nil {
			c.Redirect("/admin/banners", 302)
		}
		c.Data["Error"] = ""
		c.Data["Banner"] = obj
		c.Layout = "layout/ladmin.html"
		c.TplName = "admin_layout_content/lc_upbanner.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/admin/admin_upbannerjs.html"
	}
	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		rd := c.Input().Get("rd")
		link := c.Input().Get("link")
		content := c.Input().Get("content")
		beego.Debug("rd:", rd, "link:", link, "content:", content)
		if rd == "rd0" && len(link) > 0 || rd == "rd1" && len(content) > 0 || rd == "rd2" {
			_, fh, err := c.GetFile("showimg")
			// beego.Debug("上传图片:", fh)
			if err != nil {
				beego.Error(err)
			}
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
				err := models.UpdateImgBannerImg(id, showimg)
				if err != nil {
					beego.Error(err)
				}
			}
			state := 0
				if rd == "rd0" {
					state = 1
				}else if rd == "rd1" {
					state = 2
				}
			err = models.UpdateImgBannerInfo(id, state, content, link)
			if err != nil {
				beego.Error(err)
				c.Redirect("/admin/upbanner?id="+id, 302)
			}
			c.Redirect("/admin/banner?id="+id, 302)
		} else {
			c.Redirect("/admin/upbanner?id="+id, 302)
		}
	}
}
