package controllers

import (
	"crypto/md5"
	"encoding/hex"
	// "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"io"
	"net/url"
	"os"
	"path"
	"strings"
	"suyuanweb/sutil"
	"time"
)

// 图床服务器
type ImageHostingController struct {
	beego.Controller
}

func (c *ImageHostingController) Get() {
	beego.Debug("ImageHostingController Get")
	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:])
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	f, err := os.Open(filePath)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	defer f.Close()

	_, err = io.Copy(c.Ctx.ResponseWriter, f)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
}
// post上传 返回图片地址
func (c *ImageHostingController) Post() {

	// token := c.Input().Get("token")
	// beego.Debug("c:", c)
	beego.Debug("imagehosting post")
	image_name := ""
	var attachment string
	file_patn := "file"
	beego.Debug("file_patn:", file_patn)
	_, fh, err := c.GetFile(file_patn)
	// beego.Debug("上传图片:", fh)
	if err != nil {
		beego.Error(err)
	}
	// beego.Debug("fh:", fh)
	if fh != nil {
		beego.Debug("fh:", fh.Filename)
		attachment = fh.Filename
		t := time.Now().Unix()
		time_str := fmt.Sprintf("%d", t)
		img_uuid := uuid.NewV4().String()
		s := []string{attachment, time_str, img_uuid}
		h := md5.New()
		h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
		image_name = hex.EncodeToString(h.Sum(nil))
		beego.Info(image_name) // 输出加密结果
		err = c.SaveToFile(file_patn, path.Join("imagehosting", image_name))
		if err != nil {
			beego.Error(err)
		}
	}
	img_rul := sutil.IsImgPath(image_name)
	c.Ctx.WriteString(img_rul)
}
