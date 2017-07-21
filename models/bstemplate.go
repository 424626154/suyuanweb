package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"strconv"
)

type BsTemplate struct {
	Id       int64
	Title  string    `orm:"size(500)"`
	Describe  string    `orm:"size(500)"`
	FileName string `orm:"size(500)"`
	FileType string `orm:"size(500)"`
	UTime    time.Time //修改时间
	CTime    time.Time //创建时间
	Del      bool      //是否删除
	ReleaseState bool // false下架 true 上架
}
// 添加模板
func AddBsTemplate(title string ,describe string,filename string,filetype  string)(*BsTemplate ,error){
	time := time.Now()
	o := orm.NewOrm()
	obj := &BsTemplate{Title:title,Describe:describe,FileName:filename, FileType:filetype,CTime: time, UTime: time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
// 根据ID获得模板
func GetBsTemplateFid(id string)(*BsTemplate,error){
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var objs []BsTemplate
	_, err = o.Raw("SELECT * FROM bs_template WHERE id = ?  AND del = false ", cid).QueryRows(&objs)
	obj := &BsTemplate{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}
// 根据标题获得模板
func GetBsTemplateFTitle( title string)(*BsTemplate,error){
	o := orm.NewOrm()
	var objs []BsTemplate
	_, err := o.Raw("SELECT * FROM bs_template WHERE title = ? ", title).QueryRows(&objs)
	obj := &BsTemplate{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}
// 获得所有模板
func GetBsTemplates()([]BsTemplate,error){
	o := orm.NewOrm()
	var objs []BsTemplate
	_, err := o.Raw("SELECT * FROM bs_template  WHERE del = false ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}

func GetBsHomeTemplates()([]BsTemplate,error){
	o := orm.NewOrm()
	var objs []BsTemplate
	_, err := o.Raw("SELECT * FROM bs_template  WHERE del = false AND release_state = true ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}

func GetBsHomeSsearchTemplates(search string)([]BsTemplate,error){
	o := orm.NewOrm()
	var objs []BsTemplate
	_, err := o.Raw("SELECT * FROM bs_template  WHERE del = false AND release_state = true AND title LIKE  ? ORDER BY id DESC","%"+search+"%").QueryRows(&objs)
	return objs, err
}


func GetBSAdminTemaplatesCount()(int,error){
	o := orm.NewOrm()
	count,err := o.QueryTable("bs_template").Filter("del",false).Count()
	return int(count),err
}



func GetBSAdminTemaplatesPage(pageSize int,pageNumber int)([]BsTemplate, error){
	pages := pageSize*(pageNumber-1) 
	sql :=  "SELECT * FROM bs_template  WHERE del = false  ORDER BY id DESC limit ?, ?" 
	o := orm.NewOrm()
	var objs []BsTemplate
	_, err := o.Raw(sql,pages,pageSize).QueryRows(&objs)
	return objs, err
	
}

// 删除模板
func DelBsTemplate(id string)(error){
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &BsTemplate{Id: cid}
	obj.Del = true
	_, err = o.Update(obj, "del")
	return err
}


// 修改模板状态
func UpdateBsTemplateState(id string,state bool)(error){
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &BsTemplate{Id: cid}
	obj.ReleaseState = state
	_, err = o.Update(obj, "release_state")
	return err
}
// 修改模板文件
func UpBsTemplateFile(id string,filename string,filetype string) error{
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &BsTemplate{Id: cid}
	obj.FileName = filename
	obj.FileType = filetype
	_, err = o.Update(obj, "file_name","file_type")
	return err
}
// 修改模板信息
func UpBsTemplateInfo(id string,title string,describe string) error{
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &BsTemplate{Id: cid}
	obj.Title = title
	obj.Describe = describe
	_, err = o.Update(obj, "title","describe")
	return err
}