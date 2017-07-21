package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type ImgBanner struct {
	Id           int64
	Image        string    `orm:"size(500)"`   //标题
	Link         string    `orm:"size(500)"`   //链接
	Content      string    `orm:"size(10240)"` //文字内容
	ReleaseState bool      // false下架 true 上架
	State int  			//0 不可点击 1外链接 2 自定义文字
	BLink        bool      //是否使用外链
	Del          bool      //是否删除
	UTime        time.Time //修改时间
	CTime        time.Time //创建时间
}

func AddImgBanner(image string, state int, content string, link string) (*ImgBanner, error) {
	time := time.Now()
	o := orm.NewOrm()
	obj := &ImgBanner{Image: image, State: state, Content: content, Link: link, CTime: time, UTime: time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func GetImgBanners() ([]ImgBanner, error) {
	o := orm.NewOrm()
	var objs []ImgBanner
	_, err := o.Raw("SELECT * FROM img_banner  WHERE del = false ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}

func GetOneImgBannerFId(id string) (*ImgBanner, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var objs []ImgBanner
	_, err = o.Raw("SELECT * FROM img_banner WHERE id = ?  AND del = false", cid).QueryRows(&objs)
	obj := &ImgBanner{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}

func GetReleaseImgBanners() ([]ImgBanner, error) {
	o := orm.NewOrm()
	var objs []ImgBanner
	_, err := o.Raw("SELECT * FROM img_banner  WHERE del = false AND release_state = true ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}

func UpdateImgBannerState(id string, state bool) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &ImgBanner{Id: cid}
	obj.ReleaseState = state
	_, err = o.Update(obj, "release_state")
	return err
}
func UpdateImgBannerInfo(id string, state int, content string, link string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &ImgBanner{Id: cid}
	obj.State = state
	obj.Link = link
	obj.Content = content
	_, err = o.Update(obj, "state", "link", "content")
	return err

}

func UpdateImgBannerImg(id string, image string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &ImgBanner{Id: cid}
	obj.Image = image
	_, err = o.Update(obj, "image")
	return err
}
func DelImgBanner(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &ImgBanner{Id: cid}
	obj.Del = true
	_, err = o.Update(obj, "del")
	return err
}
