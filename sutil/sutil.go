package sutil

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/context"
	"suyuanweb/models"
	"time"
	"strings"
	"strconv"
)

//根据账号创建token规则，适用于官网后台，企业管理后台
func CreatAdminToken(account string) string {
	timestamp := time.Now()
	token := account + "_" + timestamp.Format("2006-01-02 15:04:05")
	return token
}

func SaveAdminToken(token string, ctx *context.Context) {
	maxAge := 1<<31 - 1
	ctx.SetCookie("admin_token", token, maxAge, "/")
}

func IsAdminToken(ctx *context.Context) string {
	ck, err := ctx.Request.Cookie("admin_token")
	if err != nil {
		beego.Error(err)
		return ""
	}
	return ck.Value
}

// 官网后台是否登录
func IsLogin(ctx *context.Context) (bool, string) {
	token := IsAdminToken(ctx)
	beego.Debug("is login token :", token)
	if len(token) > 0 {
		return true, token
	} else {
		return false, ""
	}
}

//保存用户账号
func SaveAdminAccount(account string, ctx *context.Context) {
	maxAge := 1<<31 - 1
	ctx.SetCookie("admin_account", account, maxAge, "/")
}

//获得用户账号
func GetAdminAccount(ctx *context.Context) string {
	ck, err := ctx.Request.Cookie("admin_account")
	if err != nil {
		beego.Error(err)
		return ""
	}
	return ck.Value
}

/******BS系统******/

func SaveBSAdminToken(token string, ctx *context.Context) {
	maxAge := 1<<31 - 1
	ctx.SetCookie("bs_admin_token", token, maxAge, "/")
}

func IsBSAdminToken(ctx *context.Context) string {
	ck, err := ctx.Request.Cookie("bs_admin_token")
	if err != nil {
		beego.Error(err)
		return ""
	}
	return ck.Value
}

// 企业工作系统管理后台验证登录
func IsBSAdminLogin(ctx *context.Context) (bool, string) {
	token := IsBSAdminToken(ctx)
	beego.Debug("is login token :", token)
	if len(token) > 0 {
		return true, token
	} else {
		return false, ""
	}
}

//保存BS用户账号
func SaveBSAdminAccount(account string, ctx *context.Context) {
	maxAge := 1<<31 - 1
	ctx.SetCookie("bs_admin_account", account, maxAge, "/")
}

//获得BS用户账号
func GetBSAdminAccount(ctx *context.Context) string {
	ck, err := ctx.Request.Cookie("bs_admin_account")
	if err != nil {
		beego.Error(err)
		return ""
	}
	return ck.Value
}

func IsBSUserLogin(ctx *context.Context) (bool, string) {
	token := IsBSUserToken(ctx)
	if len(token) > 0 {
		return true, token
	} else {
		return false, ""
	}
}

func IsBSUserToken(ctx *context.Context) string {
	ck, err := ctx.Request.Cookie("bs_user_token")
	if err != nil {
		beego.Error(err)
		return ""
	}
	return ck.Value
}

func SaveBSUserToken(token string, ctx *context.Context) {
	maxAge := 1<<31 - 1
	ctx.SetCookie("bs_user_token", token, maxAge, "/")
}

func SaveBSUserAccount(account string, ctx *context.Context) {
	maxAge := 1<<31 - 1
	ctx.SetCookie("bs_user_account", account, maxAge, "/")
}

func GetBSUserAccount(ctx *context.Context) string {
	ck, err := ctx.Request.Cookie("bs_user_account")
	if err != nil {
		beego.Error(err)
		return ""
	}
	return ck.Value
}

//图床服务器路径
func IsImgPath(imgid string) (imgpath string) {
	url := "/imagehosting/"
	return fmt.Sprintf("%s%s", url, imgid)
}

//后台开发版本
func IsAdminVersion() (version string) {
	admin_version := "1.0.0_Bate"
	cnf, err := config.NewConfig("ini", "conf/suyuanweb.conf")
	if err != nil {
		beego.Error(err)
	} else {
		admin_version = cnf.String("suyuanweb::admin_version")
	}
	return admin_version
}

// 格式化时间
func TimeFormat(ctime time.Time) (out string) {
	minute := 60
	hour := minute * 60
	day := hour * 24
	month := day * 30
	year := month * 12
	now := time.Now().Unix()

	in := ctime.Unix()
	diffValue := now - in
	if diffValue < 0 {
		//若日期不符则弹出窗口告之
	}
	yearC := diffValue / int64(year)
	monthC := diffValue / int64(month)
	weekC := diffValue / int64((7 * day))
	dayC := diffValue / int64(day)
	hourC := diffValue / int64(hour)
	minC := diffValue / int64(minute)
	result := ""

	if yearC >= 1 {
		result = time.Unix(in, 0).Format("2006-01-02 15:04:05")
	} else if monthC >= 1 {
		result = fmt.Sprintf("%d个月前", monthC)
	} else if weekC >= 1 {
		result = fmt.Sprintf("%d周前", weekC)
	} else if dayC >= 1 {
		result = fmt.Sprintf("%d天前", dayC)
	} else if hourC >= 1 {
		result = fmt.Sprintf("%d小时前", hourC)
	} else if minC >= 1 {
		result = fmt.Sprintf("%d分钟前", minC)
	} else {
		result = "刚刚发表"
	}
	return result
}

func TimeFormatStyle1(ctime time.Time) (out string) {
	now := time.Now().Unix()
	in := ctime.Unix()
	diffValue := now - in
	if diffValue < 0 {
		//若日期不符则弹出窗口告之
	}
	result := time.Unix(in, 0).Format("2006-01-02 15:04:05")
	return result
}


func DateXlxs(cdate string)(out string){
	ts := strings.Split(cdate, "-")
	year := 16
	month := 1
	day := 2
	if len(ts) == 3 {
		tyear,err := strconv.Atoi(ts[2])
		if err != nil {
			beego.Error(err)
		}
		year = tyear
		tmonth,err := strconv.Atoi(ts[0])
		if err != nil {
			beego.Error(err)
		}
		month = tmonth
		tday,err := strconv.Atoi(ts[1])
		if err != nil {
			beego.Error(err)
		}
		day = tday
	}
	return "20"+strconv.Itoa(year)+"-"+strconv.Itoa(month)+"-"+strconv.Itoa(day)
}

// 分页
// func PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {
//     tp := count / pageSize
//     if count % pageSize > 0 {
//         tp = count / pageSize + 1
//     }
//     return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
// }

// count 数据数量
// pageNo 当前页数
// pageSize 每页显示条数
func PageUtil(count int, pageNo int, pageSize int) models.Page {

	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return models.Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp}
}
