package models

import (
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego"
	"strconv"
	"time"
)
//BS后台管理系统管理员
type BsAdmin struct {
	Id       int64
	Account  string    `orm:"size(100)"`
	Password string    `orm:"size(100)"`
	Token    string    `orm:"size(100)"`
	Auth     int       //0 默认 最低浏览 1 员工 浏览 下载 2 管理员 最高权限
	UTime    time.Time //修改时间
	CTime    time.Time //创建时间
	Del      bool      //是否删除
}


func AddBSAdmin(account string, password string, auth int) (*BsAdmin, error){
	time := time.Now()
	o := orm.NewOrm()
	obj := &BsAdmin{Account: account, Password: password, Auth: auth, CTime: time, UTime: time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func GetOneBSAdmin(account string)(*BsAdmin,error){
	o := orm.NewOrm()
	var admins []BsAdmin
	_, err := o.Raw("SELECT * FROM bs_admin WHERE account = ? ", account).QueryRows(&admins)
	admin := &BsAdmin{}
	if len(admins) > 0 {
		admin = &admins[0]
	}
	return admin, err
}

func GetOneBSAdminFId(id string)(*BsAdmin,error){
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var admins []BsAdmin
	_, err = o.Raw("SELECT * FROM bs_admin WHERE id = ? ", cid).QueryRows(&admins)
	admin := &BsAdmin{}
	if len(admins) > 0 {
		admin = &admins[0]
	}
	return admin, err
}

func GetBSAdmins()([]BsAdmin, error){
	o := orm.NewOrm()
	var objs []BsAdmin
	_, err := o.Raw("SELECT * FROM bs_admin  WHERE del = false ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}


func UpdateBSAdminToken(id int64, token string) error {
	o := orm.NewOrm()
	admin := &BsAdmin{Id: id}
	admin.Token = token
	_, err := o.Update(admin, "token")
	return err
}


func UpdateBSAdminPassword(id string, pws string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	admin := &BsAdmin{Id: cid}
	admin.Password = pws
	_, err = o.Update(admin, "password")
	return err
}

func UpdateBSAdminAuth(id string, auth int) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	admin := &BsAdmin{Id: cid}
	admin.Auth = auth
	_, err = o.Update(admin, "auth")
	return err
}


