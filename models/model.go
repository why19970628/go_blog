package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int
	Name string
	Pwd string
	// 多对多
	Article []*Article `orm:"rel(m2m)"`
}

//文章结构体
type Article struct {
	Id int `orm:"pk;auto"`
	ArtiName string `orm:"size(20)"`    // 文章名字
	Atime time.Time `orm:"auto_now"`    // 发布时间
	Acount int `orm:"default(0);null"`  // 阅读量
	Acontent string  					// 文章内容
	Aimg string							// 图片
	Atype string						// 类型
	ArticleType *ArticleType `orm:"rel(fk)"` // 设置外键
	User []*User `orm:"reverse(many)"` //
}

//文章类型
type ArticleType struct {
	Id int
	TypeName string `orm:"size(20)"`
	//一对多
	Articles []*Article `orm:"reverse(many)"`

}

func init()  {
	// 注册驱动
	//orm.RegisterDriver("mysql", orm.DR_MySQL)
	// 注册默认数据库
	// 备注：此处第一个参数必须设置为“default”（因为我现在只有一个数据库），否则编译报错说：必须有一个注册DB的别名为 default
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterModel(new(User),new(Article),new(ArticleType))
	// //自动建表                         重新
	orm.RunSyncdb("default",false,true)
}