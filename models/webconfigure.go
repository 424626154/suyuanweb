package models


import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type WebCnfigure struct {
	Id           int64
	ShowNews bool //新闻
	EntService bool //企业服务
	Del          bool      //是否删除
	UTime        time.Time //修改时间
	CTime        time.Time //创建时间
}

func GetWebCnfigure()(*WebCnfigure, error) {
	o := orm.NewOrm()
	var objs []WebCnfigure
	_, err := o.Raw("SELECT * FROM web_cnfigure WHERE del = false ").QueryRows(&objs)
	obj := &WebCnfigure{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}

func UpShowNews(id string, show bool) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	time := time.Now()
	obj := &WebCnfigure{Id: cid}
	obj.ShowNews = show
	obj.UTime = time
	obj.CTime = time
	o := orm.NewOrm()
	var objs []WebCnfigure
	_, err = o.Raw("SELECT * FROM web_cnfigure WHERE id = ? AND del = false ",cid).QueryRows(&objs)
	if err != nil {
		return err
	}
	if len(objs) > 0 {
		_, err = o.Update(obj, "show_news")
		if err != nil {
				return err
			}
	}else{
		_, err = o.Insert(obj)
		if err != nil {
			return err
		}
			}

	return err
}

func UpEntService(id string, show bool) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	time := time.Now()
	obj := &WebCnfigure{Id: cid}
	obj.EntService = show
	obj.UTime = time
	obj.CTime = time
	o := orm.NewOrm()
	var objs []WebCnfigure
	_, err = o.Raw("SELECT * FROM web_cnfigure WHERE id = ? AND del = false ",cid).QueryRows(&objs)
	if err != nil {
		return err
	}
	if len(objs) > 0 {
		_, err = o.Update(obj, "ent_service")
		if err != nil {
				return err
			}
	}else{
		_, err = o.Insert(obj)
		if err != nil {
			return err
		}
			}

	return err
}