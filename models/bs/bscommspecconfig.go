package bs

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)
//Specifications
type BsCommoditySpecConfig struct{
	Id int64
	Spec string  `orm:"size(500)"`
	Order int64
	CTime time.Time
	UTime time.Time
	Del bool
}

func AddCommoditySpec( spec string)(*BsCommoditySpecConfig,error){
	time := time.Now()
	o := orm.NewOrm()
	obj := &BsCommoditySpecConfig{Spec: spec, CTime: time, UTime: time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func UpCommoditySpec(id string,spec string)error{
	time := time.Now()
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &BsCommoditySpecConfig{Id: cid}
	obj.Spec = spec
	obj.UTime = time
	_, err = o.Update(obj, "spec","u_time")
	return err
}


func UpCommoditySpecOrder(id string,order string) error{
	time := time.Now()
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	iorder, err := strconv.ParseInt(order, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &BsCommoditySpecConfig{Id: cid}
	obj.Order = iorder
	obj.UTime = time
	_, err = o.Update(obj, "order","u_time")
	return err
}

func GetCommoditySpecs()([]BsCommoditySpecConfig ,error){
	var objs []BsCommoditySpecConfig
	o := orm.NewOrm()
	_, err := o.QueryTable("bs_commodity_spec_config").Filter("del", "false").All(&objs)
	if err != nil {
		return nil,err
	}
	return objs,nil
}

func GetOrderCommoditySpecs()([]BsCommoditySpecConfig ,error){
	var objs []BsCommoditySpecConfig
	o := orm.NewOrm()
	_, err := o.QueryTable("bs_commodity_spec_config").Filter("del", "false").OrderBy("order").All(&objs)
	if err != nil {
		return nil,err
	}
	return objs,nil
}


func DelCommoditySpec(id string) error{
		cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &BsCommoditySpecConfig{Id: cid}
	obj.Del = true
	_, err = o.Update(obj, "del")
	return err
}