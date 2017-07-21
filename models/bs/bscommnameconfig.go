package bs
import (
	"github.com/astaxie/beego/orm"
	"time"
	"strconv"
)
//商品名称配置表
type BsCommodityNameConfig struct{
	Id int64
	Name string  `orm:"size(500)"`
	Order int64
	CTime time.Time
	UTime time.Time
	Del bool
}


func AddCommodityName( name string)(*BsCommodityNameConfig,error){
	time := time.Now()
	o := orm.NewOrm()
	obj := &BsCommodityNameConfig{Name: name, CTime: time, UTime: time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func UpCommodityName(id string,name string)error{
	time := time.Now()
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &BsCommodityNameConfig{Id: cid}
	obj.Name = name
	obj.UTime = time
	_, err = o.Update(obj, "name","u_time")
	return err
}

func UpCommodityNameOrder(id string,order string) error{
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
	obj := &BsCommodityNameConfig{Id: cid}
	obj.Order = iorder
	obj.UTime = time
	_, err = o.Update(obj, "order","u_time")
	return err
}


func GetCommodityNames()([]BsCommodityNameConfig ,error){
	var objs []BsCommodityNameConfig
	o := orm.NewOrm()
	_, err := o.QueryTable("bs_commodity_name_config").Filter("del", "false").All(&objs)
	if err != nil {
		return nil,err
	}
	return objs,nil
}

func GetOrderCommodityNames()([]BsCommodityNameConfig ,error){
	var objs []BsCommodityNameConfig
	o := orm.NewOrm()
	_, err := o.QueryTable("bs_commodity_name_config").Filter("del", "false").OrderBy("order").All(&objs)
	if err != nil {
		return nil,err
	}
	return objs,nil
}


func DelCommodityName(id string) error{
		cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &BsCommodityNameConfig{Id: cid}
	obj.Del = true
	_, err = o.Update(obj, "del")
	return err
}