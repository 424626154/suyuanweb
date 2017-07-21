package bs
import (
	"github.com/astaxie/beego/orm"
	"time"
	"strconv"
)
type BsExport struct{
	Id int64
	Title string `orm:"size(500)"`
	Path string `orm:"size(500)"`
	OpAccount string `orm:"size(100)"`//操作人账号
	CTime time.Time
	UTime time.Time
	Del bool
}


func AddBsExport( title,path string)(*BsExport,error){
	time := time.Now()
	o := orm.NewOrm()
	obj := &BsExport{Title: title, Path:path, CTime:time, UTime: time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func GetBsExports()([]BsExport,error){
	var objs []BsExport
	o := orm.NewOrm()
	_, err := o.QueryTable("bs_export").Filter("del", "false").All(&objs)
	if err != nil {
		return nil,err
	}
	return objs,nil
}

func DelExports(id string) error{
		cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &BsExport{Id: cid}
	obj.Del = true
	_, err = o.Update(obj, "del")
	return err
}






func GetPageExportCount()(int,error){
	o := orm.NewOrm()
	count,err := o.QueryTable("bs_export").Filter("del",false).Count()
	return int(count),err
}

func GetPageExports(pageSize int,pageNumber int)([]BsExport, error){
	pages := pageSize*(pageNumber-1)
	sql :=  "SELECT * FROM bs_export  WHERE del = false  ORDER BY id DESC limit ?, ?"
	o := orm.NewOrm()
	var objs []BsExport
	_, err := o.Raw(sql,pages,pageSize).QueryRows(&objs)
	return objs, err

}