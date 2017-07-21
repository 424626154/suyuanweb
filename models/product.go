package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Product struct {
	Id           int64
	Name         string    `orm:"size(500)"` //品名
	TasteType    string    `orm:"size(500)"` //香型
	Weight       string    `orm:"size(100)"` //净含量
	Alcoholic    int       `orm:"size(100)"` //酒精度
	ShowImg      string    `orm:"size(500)"` //展示图
	Brief 			 string 	 `orm:"size(5000)"` // 产品简介
	ProductState bool      // false下架 true 上架
	Del          bool      //是否删除
	UTime        time.Time //修改时间
	CTime        time.Time //创建时间
	IsHome       bool      //是否首页显示
	OrderId 			int64 //排序
}

// 添加产品
func AddProduct(name string, tasteType string, weight string, alcoholic int, showImg string,brief string) (*Product, error) {
	time := time.Now()
	o := orm.NewOrm()
	obj := &Product{Name: name, TasteType: tasteType, Weight: weight, Alcoholic: alcoholic,Brief:brief, ShowImg: showImg, CTime: time, UTime: time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

// 获得全部产品
func GetAllProducts() ([]Product, error) {
	o := orm.NewOrm()
	var objs []Product
	_, err := o.Raw("SELECT * FROM product  WHERE del = false ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}

//获得已上架产品
func GetReleaseProducts() ([]Product, error) {
	o := orm.NewOrm()
	var objs []Product
	_, err := o.Raw("SELECT * FROM product  WHERE del = false AND product_state = true ORDER BY order_id DESC").QueryRows(&objs)
	return objs, err
}

//获得主页产品
func GetHomeProducts() ([]Product, error) {
	o := orm.NewOrm()
	var objs []Product
	_, err := o.Raw("SELECT * FROM product  WHERE del = false AND product_state = true AND is_home = true ORDER BY order_id DESC").QueryRows(&objs)
	return objs, err
}
//获得后台分页的产品数据
// mysql分页
//  需用到的参数:
//  pageSize 每页显示多少条数据
//  pageNumber 页数 从客户端传来
//  totalRecouds 表中的总记录数 select count (*) from 表名
//  totalPages 总页数
//  totalPages=totalRecouds%pageSize==0?totalRecouds/pageSize:totalRecouds/pageSize+1
//  pages 起始位置
//  pages= pageSize*(pageNumber-1)
//  SQL语句:
//  select * from 表名 limit pages, pageSize;
//  mysql 分页依赖于关键字 limit 它需两个参数:起始位置和pageSize
//  起始位置=页大小*(页数-1)
//  起始位置=pageSize*(pageNumber -1)
 // pageSize 每页显示多少条数据
 // pageNumber 页数 从客户端传来
func GetAdminPageProducts(pageSize int,pageNumber int)([]Product, error){
	pages := pageSize*(pageNumber-1)
	sql :=  "SELECT * FROM product  WHERE del = false  ORDER BY id DESC limit ?, ?"
	o := orm.NewOrm()
	var objs []Product
	_, err := o.Raw(sql,pages,pageSize).QueryRows(&objs)
	return objs, err

}
//获取数量
func GetAdminPageProductCount()(int,error){
	o := orm.NewOrm()
	count,err := o.QueryTable("product").Filter("del",false).Count()
	return int(count),err
}

// 通过id获得
func GetOneProductFId(id string) (*Product, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var objs []Product
	_, err = o.Raw("SELECT * FROM product WHERE id = ? ", cid).QueryRows(&objs)
	obj := &Product{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}
// 获得上一篇
func GetOlderProduct(id string)(*Product,error){
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var objs []Product
	_, err = o.Raw("SELECT * FROM product WHERE id < ?  AND product_state = true AND del = false ORDER BY id DESC LIMIT 1", cid).QueryRows(&objs)
	obj := &Product{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}
//获得下一篇
func GetNewerProduct(id string)(*Product,error){
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var objs []Product
	_, err = o.Raw("SELECT * FROM product WHERE id > ?  AND product_state = true AND del = false ORDER BY id ASC LIMIT 1", cid).QueryRows(&objs)
	obj := &Product{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}
// 获得上一篇
func GetHomeOlderProduct(order_id int64)(*Product,error){
	o := orm.NewOrm()
	var objs []Product
	_, err := o.Raw("SELECT * FROM product WHERE order_id > ?  AND product_state = true AND del = false ORDER BY order_id ASC  LIMIT 1", order_id).QueryRows(&objs)
	obj := &Product{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}
//获得下一篇
func GetHomeNewerProduct(order_id int64)(*Product,error){
	o := orm.NewOrm()
	var objs []Product
	_, err := o.Raw("SELECT * FROM product WHERE order_id < ?  AND product_state = true AND del = false ORDER BY order_id DESC LIMIT 1", order_id).QueryRows(&objs)
	obj := &Product{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}


// 根据id获得
// func GetProductFId(id string) (*Product, error) {
// 	cid, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		return nil, err
// 	}
// 	o := orm.NewOrm()
// 	obj := &Product{Id: cid}
// 	err = o.Read(obj)
// 	return obj, err
// }

// 修改产品上下架状态
func UpdateProductState(id string, state bool) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &Product{Id: cid}
	obj.ProductState = state
	_, err = o.Update(obj, "product_state")
	return err
}

// 修改是否首页显示
func UpdateProductIsHome(id string, ishome bool) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &Product{Id: cid}
	obj.IsHome = ishome
	_, err = o.Update(obj, "is_home")
	return err
}

func UpdateProductInfo(id string, name string, tasteType string, weight string, alcoholic string , brief string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	ialcoholic, err := strconv.Atoi(alcoholic)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &Product{Id: cid}
	obj.Name = name
	obj.TasteType = tasteType
	obj.Weight = weight
	obj.Alcoholic = ialcoholic
	obj.Brief = brief
	_, err = o.Update(obj, "name", "taste_type", "weight", "alcoholic","brief")
	return err

}

func UpdateProductImg(id string, showimg string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &Product{Id: cid}
	obj.ShowImg = showimg
	_, err = o.Update(obj, "show_img")
	return err
}

func UpdateProductOrder(id string,order string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	iorder, err := strconv.ParseInt(order, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &Product{Id: cid}
	obj.OrderId = iorder
	_, err = o.Update(obj, "order_id")
	return err
}

func GetOneProductFOrder(order string)(*Product,error){
	iorder, err := strconv.ParseInt(order, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	var objs []Product
	_, err = o.Raw("SELECT * FROM product WHERE order_id = ? ", iorder).QueryRows(&objs)
	obj := &Product{}
	if len(objs) > 0 {
		obj = &objs[0]
	}
	return obj, err
}

func DelProduct(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &Product{Id: cid}
	obj.Del = true
	_, err = o.Update(obj, "del")
	return err
}
