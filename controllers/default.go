package controllers

import (
	"class/models"
	_ "class/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-playground/locales/om"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	// 插入
	// orm对象
	//o := orm.NewOrm()
	//user := models.User{
	//	Id:   0,
	//	Name: "wang",
	//	Pwd:  "666",
	//}
	//_, err := o.Insert(&user)
	//if err !=nil{
	//	beego.Info("插入失败",err)
	//	return

	// 查询

	//o := orm.NewOrm()
	//user := models.User{}

	//user.Id = 1
	//err := o.Read(&user)

	//user.Name = "wang"
	//err := o.Read(&user,"Name")
	//if err != nil {
	//	beego.Info("查询失败",err)
	//}
	//
	//beego.Info("查询成功", user)
	//fmt.Print("查询：",user)
	//
	//c.Data["data"] = "beego.me"
	//c.TplName = "index.html"


	//o := orm.NewOrm()
	//user := models.User{};
	//user.Id =1
	//err := o.Read(&user)
	//if err == nil {
	//	user.Name = "22"
	//	user.Pwd = "456"
	//	_,err = o.Update(&user)
	//	if err != nil {
	//		beego.Info("更新失败")
	//	}
	//}

	//o := orm.NewOrm()
	//user := models.User{};
	//user.Id =2
	//_,err :=o.Delete(&user)
	//if err != nil {
	//	beego.Info("删除错误")
	//	return
	//}
	//
	//c.Data["data"] = "beego.me"
	c.TplName = "register.html"
}

// 注册
func (c *MainController) Post() {
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	beego.Info(userName,pwd)

	if userName =="" || pwd== ""{
		beego.Info("数据不能为空")
		c.Redirect("/register",302)
		return
	}
	o := orm.NewOrm()
	user := models.User{
		Name: userName,
		Pwd:  pwd,
	}
	_,err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据库失败")
		c.Redirect("/register",302)
		return
	}
	c.Redirect("/login",302)
}


func (c *MainController)ShowLogin() {
	c.TplName = "login.html"
}

func (c *MainController)HandleLogin() {
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	if userName =="" || pwd== ""{
		beego.Info("数据不能为空")
		c.TplName="login.html"
		return
	}
	o := orm.NewOrm()
	user := models.User{}

	user.Name = userName
	err := o.Read(&user,"Name")
	if err != nil {
		beego.Info("查询失败")
		c.TplName = "login.html"
		return
	}
	c.Redirect("/index",302)
}


