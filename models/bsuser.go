package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type BsUser struct {
	Id       int64
	Account  string    `orm:"size(100)"`
	Password string    `orm:"size(100)"`
	Token    string    `orm:"size(100)"`
	UTime    time.Time //修改时间
	CTime    time.Time //创建时间
	Del      bool      //是否删除
}



func GetOneBSUser(account string)(*BsUser,error){
	o := orm.NewOrm()
	var objs []BsUser
	_, err := o.Raw("SELECT * FROM bs_user WHERE account = ? ", account).QueryRows(&objs)
	obj := &BsUser{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}

func UpdateBSUserToken(id int64, token string) error {
	o := orm.NewOrm()
	admin := &BsUser{Id: id}
	admin.Token = token
	_, err := o.Update(admin, "token")
	return err
}