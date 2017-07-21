package bs

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
)

type BsAllCommodity struct {
	Id           int64
	Name         string `orm:"size(500)"` //名称
	Spec         string `orm:"size(500)"` //规格
	Unit         string `orm:"size(100)"` //单位
	StartNumber  int    //期初数量
	StartTotal   int    //期初金额
	EndNumber    int    //期末数量
	EndTotal     int    //期末金额
	CDate        string `orm:"size(100)"` //日期
	InNumber     int    //购进数量
	InUnitPrice  int    //购进单价
	InTotal      int    //购进金额
	OutNumber    int    //发出数量
	OutUnitPrice int    //发出单价
	OutTotal     int    //发出金额
	OpAccount    string `orm:"size(100)"` //操作账号
	Year         int    //年
	Month        int    //月
	Day          int    //日
	Del          bool
	UTime        time.Time //修改时间
	CTime        time.Time //创建时间
}

func MergeAllCommodity(name string, spec string, unit string, cdate string, innumber int, inunitprice int, outnumber int, outunitprice int) (*BsAllCommodity, error) {
	time := time.Now()

	o := orm.NewOrm()
	var obj BsAllCommodity
	obj.Name = name
	obj.Spec = spec
	obj.Unit = unit
	obj.CTime = time
	obj.UTime = time
	obj.CDate = cdate
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

	startnumber, starttotal, err := GetStartNumberAndTotal(name, spec, unit, obj.Year, obj.Month, obj.Day)
	if err != nil {
		return nil, err
	}
	obj.StartNumber = startnumber
	obj.StartTotal = starttotal

	temp_obj, err := GetOneAllCommodity(name, spec, unit, obj.Year, obj.Month, obj.Day)
	if err != nil {
		return nil, err
	}
	if temp_obj.Id > 0 {
		obj.InNumber = innumber
		obj.InUnitPrice = inunitprice
		if temp_obj.InNumber > 0 {
			obj.InNumber += temp_obj.InNumber
		}
		if temp_obj.InUnitPrice > 0 {
			obj.InUnitPrice = temp_obj.InUnitPrice
		}
		obj.OutNumber = outnumber
		obj.OutUnitPrice = outunitprice
		if temp_obj.OutNumber > 0 {
			obj.OutNumber += temp_obj.OutNumber
		}
		if temp_obj.OutUnitPrice > 0 {
			obj.OutUnitPrice = temp_obj.OutUnitPrice
		}
	} else {
		obj.InNumber = innumber
		obj.InUnitPrice = inunitprice
		obj.InTotal = innumber * inunitprice

		obj.OutNumber = outnumber
		obj.OutUnitPrice = outunitprice
		obj.OutTotal = outnumber * outunitprice

	}

	obj.InTotal = obj.InNumber * obj.InUnitPrice
	obj.OutTotal = obj.OutNumber * obj.OutUnitPrice
	obj.EndNumber = obj.StartNumber - obj.InNumber + obj.OutNumber
	obj.EndTotal = obj.StartTotal - obj.InTotal + obj.OutTotal
	if temp_obj.Id > 0 {
		obj.Id = temp_obj.Id
		_, err = o.Update(&obj, "in_number", "in_unit_price", "in_total", "out_number", "out_unit_price", "out_total", "u_time", "end_number", "end_total")
		if err != nil {
			beego.Error(err)
			return nil, err
		}
	} else {
		_, err = o.Insert(&obj)
		if err != nil {
			beego.Error(err)
			return nil, err
		}
	}
	return &obj, err
}

func AddAllCommodity(name string, spec string, unit string, cdate string, startnumber int, starttotal int, innumber int, inunitprice int, outnumber int, outunitprice int) (*BsAllCommodity, error) {
	time := time.Now()
	o := orm.NewOrm()
	var obj BsAllCommodity
	obj.Name = name
	obj.Spec = spec
	obj.Unit = unit
	obj.CTime = time
	obj.UTime = time
	obj.CDate = cdate
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

	obj.StartNumber = startnumber
	obj.StartTotal = starttotal

	obj.InNumber = innumber
	obj.InUnitPrice = inunitprice
	obj.OutNumber = outnumber
	obj.OutUnitPrice = outunitprice
	obj.InTotal = obj.InNumber * obj.InUnitPrice
	obj.OutTotal = obj.OutNumber * obj.OutUnitPrice
	obj.EndNumber = obj.StartNumber + obj.InNumber - obj.OutNumber
	obj.EndTotal = obj.StartTotal + obj.InTotal - obj.OutTotal

	_, err := o.Insert(&obj)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	return &obj, err
}

func UpAllCommodity(id string, name string, spec string, unit string, cdate string, startnumber int, starttotal int, innumber int, inunitprice int, outnumber int, outunitprice int) (*BsAllCommodity, error) {
	time := time.Now()
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var obj BsAllCommodity
	obj.Id = cid
	obj.Name = name
	obj.Spec = spec
	obj.Unit = unit
	obj.CTime = time
	obj.UTime = time
	obj.CDate = cdate
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

	obj.StartNumber = startnumber
	obj.StartTotal = starttotal

	obj.InNumber = innumber
	obj.InUnitPrice = inunitprice
	obj.OutNumber = outnumber
	obj.OutUnitPrice = outunitprice
	obj.InTotal = obj.InNumber * obj.InUnitPrice
	obj.OutTotal = obj.OutNumber * obj.OutUnitPrice
	obj.EndNumber = obj.StartNumber + obj.InNumber - obj.OutNumber
	obj.EndTotal = obj.StartTotal + obj.InTotal - obj.OutTotal
	_, err = o.Update(&obj, "name", "spec", "unit", "c_time", "start_number", "start_total", "in_number", "in_unit_price", "in_total", "out_number", "out_unit_price", "out_total", "u_time", "end_number", "end_total")
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	return &obj, err
}

func AddAllCommodityFObj(item BsAllCommodity) (*BsAllCommodity, error) {
	time := time.Now()
	o := orm.NewOrm()
	var obj BsAllCommodity
	obj.Name = item.Name
	obj.Spec = item.Spec
	obj.Unit = item.Unit
	obj.CTime = time
	obj.UTime = time
	obj.CDate = item.CDate
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

	obj.StartNumber = item.StartNumber
	obj.StartTotal = item.StartTotal

	obj.InNumber = item.InNumber
	obj.InUnitPrice = item.InUnitPrice
	obj.OutNumber = item.OutNumber
	obj.OutUnitPrice = item.OutUnitPrice
	obj.InTotal = item.InTotal
	obj.OutTotal = item.OutTotal
	obj.EndNumber = item.EndNumber
	obj.EndTotal = item.EndTotal

	_, err := o.Insert(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, err
}

func GetStartNumberAndTotal(name, spec, unit string, year, month, day int) (int, int, error) {
	o := orm.NewOrm()
	var objs []BsAllCommodity
	sql := "SELECT * FROM bs_all_commodity WHERE name = ? AND spec = ? AND unit = ? AND year <= ?  AND (month < ? OR month = ? AND day < ?) AND del = false ORDER BY id ASC "
	_, err := o.Raw(sql, name, spec, unit, year, month, month, day).QueryRows(&objs)
	if err != nil {
		beego.Error(err)
		return 0, 0, err
	}
	var obj BsAllCommodity
	if len(objs) > 0 {
		obj = objs[0]
	}
	beego.Debug("GetStartNumberAndTotal obj:", obj.Id)
	return obj.EndNumber, obj.EndTotal, nil
}

func GetOneAllCommodity(name, spec, unit string, year, month, day int) (*BsAllCommodity, error) {
	o := orm.NewOrm()
	var objs []BsAllCommodity
	sql := "SELECT * FROM bs_all_commodity WHERE name = ? AND spec = ? AND unit = ? AND year = ?  AND month = ? AND day = ? AND del = false ORDER BY id ASC "
	_, err := o.Raw(sql, name, spec, unit, year, month, day).QueryRows(&objs)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	var obj BsAllCommodity
	if len(objs) > 0 {
		obj = objs[0]
	}
	return &obj, nil
}

func GetOneAllCommodityFId(id string) (*BsAllCommodity, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var obj BsAllCommodity
	err = o.QueryTable("bs_all_commodity").Filter("id", cid).One(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, err
}

func GetAllCommoditys() ([]BsAllCommodity, error) {
	var objs []BsAllCommodity
	o := orm.NewOrm()
	_, err := o.QueryTable("bs_all_commodity").Filter("del", "false").All(&objs)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	return objs, nil
}

func DelAllCommodity(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &BsAllCommodity{Id: cid}
	obj.Del = true
	_, err = o.Update(obj, "del")
	return err
}

//获取数量
func GetPageAllCommodityCount() (int, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("bs_all_commodity").Filter("del", false).Count()
	return int(count), err
}

func GetPageAllCommoditys(pageSize int, pageNumber int) ([]BsAllCommodity, error) {
	pages := pageSize * (pageNumber - 1)
	sql := "SELECT * FROM bs_all_commodity  WHERE del = false  ORDER BY id DESC limit ?, ?"
	o := orm.NewOrm()
	var objs []BsAllCommodity
	_, err := o.Raw(sql, pages, pageSize).QueryRows(&objs)
	return objs, err
}

func GetAllCommoditysFDay(year int, month int, day int) (int, error) {
	o := orm.NewOrm()
	var objs []BsAllCommodity
	_, err := o.Raw("SELECT * FROM bs_all_commodity where year = ? and month = ? and day = ?", year, month, day).QueryRows(&objs)
	if err != nil {
		return 0, err
	}
	cnt := len(objs)
	return cnt, nil
}

func GetAllCommoditysFMonth(year int, month int) (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("bs_all_commodity").Filter("year", year).Filter("month", month).Count()
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
