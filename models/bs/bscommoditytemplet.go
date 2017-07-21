package bs
import (
	"github.com/astaxie/beego/orm"
	"time"
	// "strconv"
)
//商品名称配置表
type BsCommodityTemplet struct{
	Id int64
	InCommodityTemplet string  `orm:"size(500)"`
	OutCommodityTemplet string `orm:"size(500)"`
	AllCommodityTemplet string `orm:"size(500)"`
	CTime time.Time
	UTime time.Time
}


func AddInCommodityTemplet(path string)error{
	time := time.Now()
	o := orm.NewOrm()
	var objs []BsCommodityTemplet
	_, err := o.QueryTable("bs_commodity_templet").All(&objs)
	if err != nil {
		return err
	}
	if err != nil{
		return err
	}
	var obj BsCommodityTemplet
	obj.InCommodityTemplet = path
	if len(objs) > 0 {
		obj.Id = objs[0].Id
		obj.UTime = time
		_, err = o.Update(&obj, "in_commodity_templet","u_time")
		if err != nil{
			return err
		}
	}else{
		obj.CTime = time
		obj.UTime = time
		_, err = o.Insert(&obj)
		if err != nil{
			return err
		}
	}
	return nil
}

func AddOutCommodityTemplet(path string)error{
	time := time.Now()
	o := orm.NewOrm()
	var objs []BsCommodityTemplet
	_, err := o.QueryTable("bs_commodity_templet").All(&objs)
	if err != nil {
		return err
	}
	if err != nil{
		return err
	}
	var obj BsCommodityTemplet
	obj.OutCommodityTemplet = path
	if len(objs) > 0 {
		obj.Id = objs[0].Id
		obj.UTime = time
		_, err = o.Update(&obj, "out_commodity_templet","u_time")
		if err != nil{
			return err
		}
	}else{
		obj.CTime = time
		obj.UTime = time
		_, err = o.Insert(&obj)
		if err != nil{
			return err
		}
	}
	return nil
}
func AddAllCommodityTemplet(path string)error{
	time := time.Now()
	o := orm.NewOrm()
	var objs []BsCommodityTemplet
	_, err := o.QueryTable("bs_commodity_templet").All(&objs)
	if err != nil {
		return err
	}
	if err != nil{
		return err
	}
	var obj BsCommodityTemplet
	obj.AllCommodityTemplet = path
	if len(objs) > 0 {
		obj.Id = objs[0].Id
		obj.UTime = time
		_, err = o.Update(&obj, "all_commodity_templet","u_time")
		if err != nil{
			return err
		}
	}else{
		obj.CTime = time
		obj.UTime = time
		_, err = o.Insert(&obj)
		if err != nil{
			return err
		}
	}
	return nil
}
func GetCommodityTemplet() (*BsCommodityTemplet,error){
	o := orm.NewOrm()
	var objs []BsCommodityTemplet
	_, err := o.Raw("SELECT * FROM bs_commodity_templet").QueryRows(&objs)
	obj := &BsCommodityTemplet{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}