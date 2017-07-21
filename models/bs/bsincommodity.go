package bs

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

//入库登记表数据库
type BsInCommodity struct {
	Id          int64
	CDate       string
	OrderNumber string    `orm:"size(500)"` //订单编号
	Name        string    `orm:"size(500)"` //产品名称
	Unit        string    `orm:"size(100)"` //单位
	Spec        string    `orm:"size(500)"` //规格
	UnitPrice   int       //单价
	Number      int       //数量
	Total       int       //总价 金额
	Operator    string    `orm:"size(100)"`  //经办人
	Remarks     string    `orm:"size(1000)"` //备注
	CTime       time.Time //创建时间
	UTime       time.Time //修改时间
	Del         bool
	OpAccount   string `orm:"size(100)"` //操作人账号
	Year        int    //年
	Month       int    //月
	Day         int    //日
}

//添加入库
func AddInCommodity(cdate string, ordernumber string, name string, unit string, spec string, unitprice string, number string, operator string, remarks string) (*BsInCommodity, error) {
	time := time.Now()
	iunitprice, err := strconv.Atoi(unitprice)
	if err != nil {
		return nil, err
	}
	inumber, err := strconv.Atoi(number)
	if err != nil {
		return nil, err
	}
	total := iunitprice * inumber

	o := orm.NewOrm()
	var obj BsInCommodity
	obj.CDate = cdate
	obj.OrderNumber = ordernumber
	obj.Name = name
	obj.Unit = unit
	obj.Spec = spec //规格
	obj.UnitPrice = iunitprice
	obj.Number = inumber
	obj.Total = total
	obj.Operator = operator
	obj.Remarks = remarks
	obj.CTime = time
	obj.UTime = time

	ts := strings.Split(cdate, "-")
	if len(ts) == 3 {
		year, err := strconv.Atoi(ts[0])
		if err != nil {
			return nil, err
		}
		month, err := strconv.Atoi(ts[1])
		if err != nil {
			return nil, err
		}
		day, err := strconv.Atoi(ts[2])
		if err != nil {
			return nil, err
		}
		obj.Year = year
		obj.Month = month
		obj.Day = day
	} else {
		return nil, errors.New("Date format error")
	}

	_, err = o.Insert(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, err
}

func AddInCommodityFObj(item BsInCommodity) (*BsInCommodity, error) {
	time := time.Now()
	total := item.UnitPrice * item.Number
	o := orm.NewOrm()
	var obj BsInCommodity
	obj.CDate = item.CDate
	obj.OrderNumber = item.OrderNumber
	obj.Name = item.Name
	obj.Unit = item.Unit
	obj.Spec = item.Spec
	obj.UnitPrice = item.UnitPrice
	obj.Number = item.Number
	obj.Total = total
	obj.Operator = item.Operator
	obj.Remarks = item.Remarks
	obj.CTime = time
	obj.UTime = time

	ts := strings.Split(item.CDate, "-")
	if len(ts) == 3 {
		year, err := strconv.Atoi(ts[2])
		if err != nil {
			return nil, err
		}
		month, err := strconv.Atoi(ts[0])
		if err != nil {
			return nil, err
		}
		day, err := strconv.Atoi(ts[1])
		if err != nil {
			return nil, err
		}
		obj.Year = year
		obj.Month = month
		obj.Day = day
	} else {
		return nil, errors.New("Date format error")
	}

	_, err := o.Insert(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, err
}

func GetInCommoditys() ([]BsInCommodity, error) {
	var objs []BsInCommodity
	o := orm.NewOrm()
	_, err := o.QueryTable("bs_in_commodity").Filter("del", "false").All(&objs)
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetOneInCommodityFId(id string) (*BsInCommodity, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var obj BsInCommodity
	err = o.QueryTable("bs_in_commodity").Filter("id", cid).One(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, err
}

func DelInCommodity(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &BsInCommodity{Id: cid}
	obj.Del = true
	_, err = o.Update(obj, "del")
	return err
}

func UpInCommodity(id string, cdate string, ordernumber string, name string, unit string, spec string, unitprice string, number string, operator string, remarks string) error {
	time := time.Now()
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	iunitprice, err := strconv.Atoi(unitprice)
	if err != nil {
		return err
	}
	inumber, err := strconv.Atoi(number)
	if err != nil {
		return err
	}
	total := iunitprice * inumber

	o := orm.NewOrm()
	var obj BsInCommodity
	obj.Id = cid
	obj.CDate = cdate
	obj.OrderNumber = ordernumber
	obj.Name = name
	obj.Unit = unit
	obj.Spec = spec //规格
	obj.UnitPrice = iunitprice
	obj.Number = inumber
	obj.Total = total
	obj.Operator = operator
	obj.Remarks = remarks
	obj.CTime = time
	obj.UTime = time

	ts := strings.Split(cdate, "-")
	if len(ts) == 3 {
		year, err := strconv.Atoi(ts[0])
		if err != nil {
			return err
		}
		month, err := strconv.Atoi(ts[1])
		if err != nil {
			return err
		}
		day, err := strconv.Atoi(ts[2])
		if err != nil {
			return err
		}
		obj.Year = year
		obj.Month = month
		obj.Day = day
	} else {
		return errors.New("Date format error")
	}

	_, err = o.Update(&obj, "in_time", "order_number", "name", "unit", "spec", "unit_price", "number", "total", "operator", "remarks", "u_time", "year", "month", "day")

	return err
}

//获取数量
func GetPageInCommodityCount() (int, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("bs_in_commodity").Filter("del", false).Count()
	return int(count), err
}

func GetPageInCommoditys(pageSize int, pageNumber int) ([]BsInCommodity, error) {
	pages := pageSize * (pageNumber - 1)
	sql := "SELECT * FROM bs_in_commodity  WHERE del = false  ORDER BY id DESC limit ?, ?"
	o := orm.NewOrm()
	var objs []BsInCommodity
	_, err := o.Raw(sql, pages, pageSize).QueryRows(&objs)
	return objs, err

}

func GetInCommoditysFDay(year int, month int, day int) (int, error) {
	o := orm.NewOrm()
	var objs []BsInCommodity
	_, err := o.Raw("SELECT * FROM bs_in_commodity where year = ? and month = ? and day = ?", year, month, day).QueryRows(&objs)
	if err != nil {
		return 0, err
	}
	cnt := len(objs)
	return cnt, nil
}

func GetInCommoditysFMonth(year int, month int) (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("bs_in_commodity").Filter("year", year).Filter("month", month).Count()
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
