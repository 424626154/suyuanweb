package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type PageLog struct {
	Id    int64
	Url   string     `orm:"size(500)"`
	Addr  string     `orm:"size(100)"`
	CTime time.Time  //创建时间
	Mark  bool       //是否已经导入
	Year  int        //年
	Month time.Month //月
	Day   int        //日
}

func AddPageLog(url string, addr string) (*PageLog, error) {
	time := time.Now()
	year := time.Year()
	month := time.Month()
	day := time.Day()
	o := orm.NewOrm()
	obj := &PageLog{Url: url, Addr: addr, CTime: time, Year: year, Month: month, Day: day}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func GetPageLogs() ([]PageLog, error) {
	o := orm.NewOrm()
	var objs []PageLog
	_, err := o.Raw("SELECT * FROM page_log ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}

func GetPageLogsFMonth(year int, month int) (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("page_log").Filter("year", year).Filter("month", month).Count()
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func GetPageLogsFDay(year int, month int, day int) (int, error) {
	// beego.Debug("year:", year, "month:", month, "day:", day)
	o := orm.NewOrm()
	// cnt, err := o.QueryTable("page_log").Filter("year", year).Filter("month", month).Filter("day", day).Count()
	// SELECT count(*) FROM suyuandb.page_log where year = 2017 and month = 7 and day = 12;
	var objs []PageLog
	_, err := o.Raw("SELECT * FROM page_log where year = ? and month = ? and day = ?", year, month, day).QueryRows(&objs)
	if err != nil {
		return 0, err
	}
	cnt := len(objs)
	return cnt, nil
}
