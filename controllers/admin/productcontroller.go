package admin

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"path"
	"strconv"
	"strings"
	"suyuanweb/models"
	"suyuanweb/sutil"
	"time"
)

//产品详情
type ProductController struct {
	beego.Controller
}

type ProductRes struct {
	Code  int64  `json:"code"` //0 成功 1失败
	Error string `json:"error"`
	Id    string `json:"id"`
	State bool   `json:"state"`
}

func (c *ProductController) Get() {
	account := sutil.GetAdminAccount(c.Ctx)
	c.Data["Account"] = account

	i_str := initProductPage(c)

	obj, err := models.GetOneAdmin(account)
	if err != nil {
		beego.Error(err)
	}
	auth := obj.Auth
	c.Data["Auth"] = auth

	id := c.Input().Get("id")
	if len(id) > 0 {
		op := c.Input().Get("op")
		switch op {
		case "up":
			state := c.Input().Get("state")
			bstate := true
			if state == "true" {
				bstate = false
			}
			err := models.UpdateProductState(id, bstate)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/product?id="+id+"&p="+i_str, 302)
			return
		case "del":
			err := models.DelProduct(id)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/products?p="+i_str, 302)
			return
		case "home":
			ishome := c.Input().Get("ishome")
			bishome := true
			if ishome == "true" {
				bishome = false
			}
			err := models.UpdateProductIsHome(id, bishome)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/product?id="+id+"&p="+i_str, 302)
			return
		}

		product, err := models.GetOneProductFId(id)
		if err != nil {
			beego.Error(err)
			c.Redirect("/admin/products"+"&p="+i_str, 302)
			return
		} else {
			beego.Debug("product:", product)
			c.Data["Product"] = product
			c.Layout = "layout/ladmin.html"
			c.TplName = "admin_layout_content/lc_product.html"
			c.LayoutSections = make(map[string]string)
			c.LayoutSections["Scripts"] = "scripts/admin/admin_productjs.html"
		}

	} else {
		c.Redirect("/admin/products?p="+i_str, 302)
		return
	}
}

//产品列表
func (c *ProductController) Products() {
	account := sutil.GetAdminAccount(c.Ctx)
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

	pageSize := 10
	count, err := models.GetAdminPageProductCount()
	if err != nil {
		beego.Error(err)
	}
	products, err := models.GetAdminPageProducts(pageSize, i_p)
	beego.Debug("p:", p, "products:", len(products))
	if err != nil {
		beego.Error(err)
	}
	page := sutil.PageUtil(count, i_p, pageSize)

	c.Data["Products"] = products
	c.Data["Page"] = page
	c.Layout = "layout/ladmin.html"
	c.TplName = "admin_layout_content/lc_products.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "scripts/admin/admin_products.html"
}

// 添加产品
func (c *ProductController) AddProduct() {
	if c.Ctx.Input.IsGet() {
		account := sutil.GetAdminAccount(c.Ctx)
		c.Data["Account"] = account
		AddproductHtml(c, "")
	}

	if c.Ctx.Input.IsPost() {
		name := c.Input().Get("name")
		tastetype := c.Input().Get("tastetype")
		weight := c.Input().Get("weight")
		alcoholic := c.Input().Get("alcoholic")
		brief := c.Input().Get("brief")
		beego.Debug("is addproduct post name:", name, "tastetype:", tastetype, "weight:", weight, "alcoholic:", alcoholic)
		beego.Debug("name:", len(name) > 0)
		beego.Debug("tastetype:", len(tastetype) > 0)
		beego.Debug("weight:", len(weight) > 0)
		beego.Debug("alcoholic:", len(alcoholic) > 0)
		if len(name) > 0 && len(tastetype) > 0 && len(weight) > 0 && len(alcoholic) > 0 {
			_, fh, err := c.GetFile("showimg")
			// beego.Debug("上传图片:", fh)
			if err != nil {
				beego.Error(err)
				AddproductHtml(c, err.Error())
				return
			}
			if fh != nil {
				beego.Debug("fh:", fh.Filename)
				tempname := fh.Filename
				t := time.Now().Unix()
				time_str := fmt.Sprintf("%d", t)
				img_uuid := uuid.NewV4().String()
				s := []string{tempname, time_str, img_uuid}
				h := md5.New()
				h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
				showimg := hex.EncodeToString(h.Sum(nil))
				beego.Info(showimg) // 输出加密结果
				beego.Info(path.Join("showimg", showimg))
				err = c.SaveToFile("showimg", path.Join("imagehosting", showimg))
				if err != nil {
					beego.Error(err)
					AddproductHtml(c, err.Error())
					return
				} else {
					//图片保存成功 存储数据
					ialcoholic, err := strconv.Atoi(alcoholic)
					if err != nil {
						beego.Error(err)
						AddproductHtml(c, err.Error())
						return
					}
					product, err := models.AddProduct(name, tastetype, weight, ialcoholic, showimg, brief)
					if err != nil {
						beego.Error(err)
						AddproductHtml(c, err.Error())
						return
					} else {
						c.Data["Product"] = product
						c.Layout = "layout/ladmin.html"
						c.TplName = "admin_layout_content/lc_addproduct_success.html"
						return
					}
				}
			} else {
				c.Data["Error"] = "获取文件失败"
				AddproductHtml(c, "获取文件失败")
				return
			}

			AddproductHtml(c, "位置错误")
		} else {
			AddproductHtml(c, "参数错误")
		}
	}
}

func AddproductHtml(c *ProductController, err string) {
	c.Data["Error"] = err
	c.Layout = "layout/ladmin.html"
	c.TplName = "admin_layout_content/lc_addproduct.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["AHead"] = "admin/ahead_addproduct.html"
	c.LayoutSections["Scripts"] = "scripts/admin/admin_addproduct.html"
}

//修改产品
func (c *ProductController) UpProduct() {
	account := sutil.GetAdminAccount(c.Ctx)
	c.Data["Account"] = account

	i_str := initProductPage(c)

	if c.Ctx.Input.IsGet() {

		id := c.Input().Get("id")
		if len(id) == 0 {
			c.Redirect("/admin/products?p="+i_str, 302)
			return
		}
		product, err := models.GetOneProductFId(id)
		if err != nil {
			c.Redirect("/admin/products?p="+i_str, 302)
			return
		}

		c.Data["Error"] = ""
		c.Data["Product"] = product
		c.Layout = "layout/ladmin.html"
		c.TplName = "admin_layout_content/lc_upproduct.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Scripts"] = "scripts/admin_upproductjs.html"
	}
	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		beego.Debug("post id:", id)
		name := c.Input().Get("name")
		tastetype := c.Input().Get("tastetype")
		weight := c.Input().Get("weight")
		alcoholic := c.Input().Get("alcoholic")
		brief := c.Input().Get("brief")
		_, fh, err := c.GetFile("showimg")
		// beego.Debug("上传图片:", fh)
		if err != nil {
			beego.Error(err)
		}
		if len(name) > 0 && len(tastetype) > 0 && len(weight) > 0 && len(alcoholic) > 0 {
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
					c.Redirect("/admin/upproduct?id="+id+"&p="+i_str, 302)
					return
				}
			}
			if len(showimg) > 0 {
				err := models.UpdateProductImg(id, showimg)
				if err != nil {
					beego.Error(err)
					c.Redirect("/admin/upproduct?id="+id+"&p="+i_str, 302)
					return
				}
			}
			err := models.UpdateProductInfo(id, name, tastetype, weight, alcoholic, brief)
			if err != nil {
				beego.Error(err)
				c.Redirect("/admin/upproduct?id="+id+"&p="+i_str, 302)
				return
			}
			c.Redirect("/admin/product?id="+id+"&p="+i_str, 302)
			return
		} else {
			c.Redirect("/admin/upproduct?id="+id+"&p="+i_str, 302)
			return
		}
	}
}

 
func (c *ProductController) ProductAjax() {
	op := c.Input().Get("op")
	switch op {
	case "uporder":
		id := c.Input().Get("id")
		order := c.Input().Get("order")
		if len(id) > 0 && len(order) > 0 {
			obj,err := models.GetOneProductFOrder(order)
			if err != nil {
				beego.Error(err)
			}
			cid, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				beego.Error(err)
			}
			if obj.Id > 0 && obj.Id != cid {
				json_err := "排序ID已存在"
				ajax_json := "{\"code\":1,\"id\":"+id+",\"err\":\""+json_err+"\"}"
				c.Ctx.WriteString(ajax_json)
				return
			}else{
				err := models.UpdateProductOrder(id, order)
				if err != nil {
					beego.Error(err)
				}
			}

		}
		c.Ctx.WriteString("{\"code\":0,\"id\":"+id+",\"order\":"+order+"}")
		return
	}
}

func initProductPage(c *ProductController) string {
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
