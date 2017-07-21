package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Admin struct {
	Id       int64
	Account  string    `orm:"size(100)"`
	Password string    `orm:"size(100)"`
	Token    string    `orm:"size(100)"`
	Auth     int       //0 默认 最低浏览 1 运营  可修改 删除 2 管理员 可添加账号
	UTime    time.Time //修改时间
	CTime    time.Time //创建时间
	Del      bool      //是否删除
}

// admin db start

func AddAdmin(account string, password string, auth int) (*Admin, error) {
	time := time.Now()
	o := orm.NewOrm()
	obj := &Admin{Account: account, Password: password, Auth: auth, CTime: time, UTime: time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func GetOneAdmin(account string) (*Admin, error) {
	o := orm.NewOrm()
	var admins []Admin
	_, err := o.Raw("SELECT * FROM admin WHERE account = ? ", account).QueryRows(&admins)
	admin := &Admin{}
	if len(admins) > 0 {
		admin = &admins[0]
	}
	return admin, err
}

func GetOneAdminFId(id string) (*Admin, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var admins []Admin
	_, err = o.Raw("SELECT * FROM admin WHERE id = ? ", cid).QueryRows(&admins)
	admin := &Admin{}
	if len(admins) > 0 {
		admin = &admins[0]
	}
	return admin, err
}

func GetOneAdminFromToken(token string) (*Admin, error) {
	o := orm.NewOrm()
	var admins []Admin
	_, err := o.Raw("SELECT * FROM admin WHERE token = ? ", token).QueryRows(&admins)
	admin := &Admin{}
	if len(admins) > 0 {
		admin = &admins[0]
	}
	return admin, err
}

func UpdateAdminToken(id int64, token string) error {
	o := orm.NewOrm()
	admin := &Admin{Id: id}
	admin.Token = token
	_, err := o.Update(admin, "token")
	return err
}

func UpdateAdminPassword(id string, pws string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	admin := &Admin{Id: cid}
	admin.Password = pws
	_, err = o.Update(admin, "password")
	return err
}

func UpdateAdminAuth(id string, auth int) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	admin := &Admin{Id: cid}
	admin.Auth = auth
	_, err = o.Update(admin, "auth")
	return err
}

func GetAdmins() ([]Admin, error) {
	o := orm.NewOrm()
	var objs []Admin
	_, err := o.Raw("SELECT * FROM admin  WHERE del = false ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}

// admin db end
