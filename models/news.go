package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type News struct {
	Id           int64
	Title        string    `orm:"size(500)"`   //标题
	Outline      string    `orm:"size(500)"`   //简介
	ShowImg      string    `orm:"size(500)"`   //展示图
	Content      string    `orm:"size(10240)"` //文字内容
	ReleaseState bool      // false下架 true 上架
	Del          bool      //是否删除
	UTime        time.Time //修改时间
	CTime        time.Time //创建时间
	IsHome       bool      //是否首页显示
}

func AddNews(title string, outline string, showimg string, content string) (*News, error) {
	time := time.Now()
	o := orm.NewOrm()
	obj := &News{Title: title, Outline: outline, ShowImg: showimg, Content: content, CTime: time, UTime: time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func GetNewss() ([]News, error) {
	o := orm.NewOrm()
	var objs []News
	_, err := o.Raw("SELECT * FROM news  WHERE del = false ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}

//获得首页显示信息
func GetHomeNewss() ([]News, error) {
	o := orm.NewOrm()
	var objs []News
	_, err := o.Raw("SELECT * FROM news  WHERE del = false AND release_state = true  AND is_home = true ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}

func GetAdminPageNewss(pageSize int,pageNumber int)([]News, error){
	pages := pageSize*(pageNumber-1) 
	sql :=  "SELECT * FROM news  WHERE del = false  ORDER BY id DESC limit ?, ?" 
	o := orm.NewOrm()
	var objs []News
	_, err := o.Raw(sql,pages,pageSize).QueryRows(&objs)
	return objs, err
	
}
//获取数量
func GetAdminPageNewsCount()(int,error){
	o := orm.NewOrm()
	count,err := o.QueryTable("news").Filter("del",false).Count()
	return int(count),err
}


func GetOneNewsFId(id string) (*News, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var objs []News
	_, err = o.Raw("SELECT * FROM news WHERE id = ?  AND del = false ", cid).QueryRows(&objs)
	obj := &News{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}
// 获得上一篇
func GetOlderNews(id string)(*News,error){
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var objs []News
	_, err = o.Raw("SELECT * FROM news WHERE id < ?  AND release_state = true AND del = false ORDER BY id DESC LIMIT 1", cid).QueryRows(&objs)
	obj := &News{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}
//获得下一篇
func GetNewerNews(id string)(*News,error){
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var objs []News
	_, err = o.Raw("SELECT * FROM news WHERE id > ?  AND release_state = true AND del = false ORDER BY id LIMIT 1", cid).QueryRows(&objs)
	obj := &News{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}

func GetReleaseNewss() ([]News, error) {
	o := orm.NewOrm()
	var objs []News
	_, err := o.Raw("SELECT * FROM news  WHERE del = false AND release_state = true ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}

func UpdateNewsState(id string, state bool) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &News{Id: cid}
	obj.ReleaseState = state
	_, err = o.Update(obj, "release_state")
	return err
}
func UpdateNewsHome(id string, ishome bool) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &News{Id: cid}
	obj.IsHome = ishome
	_, err = o.Update(obj, "is_home")
	return err
}

func UpdateNewsInfo(id string, title string, outline string, content string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &News{Id: cid}
	obj.Title = title
	obj.Outline = outline
	obj.Content = content
	_, err = o.Update(obj, "title", "outline", "content")
	return err

}

func UpdateNewsImg(id string, showimg string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &News{Id: cid}
	obj.ShowImg = showimg
	_, err = o.Update(obj, "show_img")
	return err
}
func DelNews(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &News{Id: cid}
	obj.Del = true
	_, err = o.Update(obj, "del")
	return err
}
