package routers

import (
	"class/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/abc", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{})
	// 实现了自定义的get方法,
	beego.Router("/login", &controllers.MainController{}, "get:ShowLogin;post:HandleLogin")

	// 文章
	beego.Router("/article/index", &controllers.ArticleController{}, "get:ShowIndex;post:HandleSelect")
	beego.Router("/article/AddArticle", &controllers.ArticleController{}, "get:ShowAdd;post:HandleAdd")
	beego.Router("/article/content", &controllers.ArticleController{}, "get:ShowContent")
	beego.Router("/article/update", &controllers.ArticleController{}, "get:ShowUpdate;post:HandleUpdate")
	beego.Router("/article/delete", &controllers.ArticleController{}, "get:HandleDelete")

	// 文章类型
	beego.Router("/article/articleType", &controllers.ArticleController{}, "get:AddType;post:HandleAddType")

}
