package controllers

import (
	"class/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"strconv"
	"time"
)

type ArticleController struct {
	beego.Controller
}

// 处理文章列表选择类型的查询
func (c *ArticleController) HandleSelect() {
	typeName := c.GetString("select")

	if typeName == "" {
		beego.Info("下拉框数据失败")
	}
	o := orm.NewOrm()
	var articles []models.Article

	// 惰性查询 需要 多表查询 RelatedSel("ArticleType")
	o.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).All(&articles)
	beego.Info(articles)
	//c.Data["article"] = articles
	//
	//c.TplName = "update.html"
}

// 显示首页内容
func (c *ArticleController) ShowIndex() {
	o := orm.NewOrm()
	var articles []models.Article
	qs := o.QueryTable("Article")
	//qs.All(&articles)
	pageIndex := c.GetString("pageIndex")
	pageIndex1 := 1
	pageIndex1, _ = strconv.Atoi(pageIndex)
	if pageIndex1 < 1 {
		pageIndex1 = 1
	}

	count, err := qs.RelatedSel("ArticleType").Count()

	pageSize := 2
	start := pageSize * (pageIndex1 - 1)
	qs.Limit(pageSize, start).RelatedSel("ArticleType").All(&articles)

	// 分页取整 +1
	pageCount := float64(count) / float64(pageSize)
	pageCount = math.Ceil(pageCount)

	if err != nil {
		beego.Info("查询所有文章信息出错")
		return
	}
	// 分页: 判断当前页是否为首页或尾页
	var FirstPage bool = false
	var LastPage bool = false
	//
	if pageIndex1 == 1 {
		FirstPage = true
	}
	if pageIndex1 == int(pageCount) {
		LastPage = true
	}

	// 获取文章类型
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)
	c.Data["types"] = types
	beego.Info("count = ", count)

	typeName := c.GetString("select")

	var articlebytype []models.Article
	// 判断是否由类型传递
	if typeName == "" {
		beego.Info("下拉框数据失败")
		qs.Limit(pageSize, start).RelatedSel("ArticleType").All(&articlebytype)
	} else {
		// 惰性查询 需要 多表查询 RelatedSel("ArticleType")
		qs.Limit(pageSize, start).RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).All(&articlebytype)
		beego.Info(articles)
	}

	//beego.Info(articles)
	c.Data["FirstPage"] = FirstPage
	c.Data["LastPage"] = LastPage
	c.Data["count"] = count
	c.Data["pageCount"] = pageCount
	c.Data["articles"] = articlebytype
	c.Data["pageIndex"] = pageIndex1
	c.TplName = "index.html"
}

// 添加文章界面ShowAdd
func (c *ArticleController) ShowAdd() {
	o := orm.NewOrm()
	// 获取文章类型
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)
	c.Data["types"] = types
	c.TplName = "add.html"
}

//处理添加文章界面数据
func (c *ArticleController) HandleAdd() {
	//1.拿到数据
	artiName := c.GetString("articleName")
	artiContent := c.GetString("content")
	typeName := c.GetString("select")
	if typeName == "" {
		beego.Info("类型返回结果错误")
		return
	}

	f, h, err := c.GetFile("uploadname")
	defer f.Close()

	//1.要限定格式
	fileext := path.Ext(h.Filename)
	if fileext != ".jpg" && fileext != "png" {
		beego.Info("上传文件格式错误")
		return
	}
	//2.限制大小
	if h.Size > 50000000 {
		beego.Info("上传文件过大")
		return
	}

	//3.需要对文件重命名，防止文件名重复
	filename := time.Now().Format("2006-01-02 15:04:05") + fileext //6-1-2 3:4:5
	if err != nil {
		beego.Info("上传文件失败")
		return
	} else {
		c.SaveToFile("uploadname", "./static/img/"+filename)
	}

	//2.判断数据是否合法
	if artiContent == "" || artiName == "" {
		beego.Info("添加文章数据错误")
		return
	}
	//3.插入数据
	o := orm.NewOrm()
	arti := models.Article{}
	arti.ArtiName = artiName
	arti.Acontent = artiContent
	arti.Aimg = "./static/img/" + filename

	//查找type对象
	var artiType models.ArticleType
	artiType.TypeName = typeName
	beego.Info(artiType)
	err = o.Read(&artiType, "TypeName")
	if err != nil {
		beego.Info("数据库获取类型错误")
		return
	}
	arti.ArticleType = &artiType

	_, err = o.Insert(&arti)
	if err != nil {
		beego.Info("插入数据库错误")
		return
	}

	//4.返回文章界面
	c.Redirect("/article/index", 302)
}

//显示内容详情页面
func (c *ArticleController) ShowContent() {
	//1.获取文章ID
	id, err := c.GetInt("id")
	//beego.Info("id is ",id)
	if err != nil {
		beego.Info("获取文章ID错误", err)
		return
	}
	//2.查询数据库获取数据
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	beego.Info(arti)
	err = o.Read(&arti)
	if err != nil {
		beego.Info("查询错误", err)
		return
	}
	//3.传递数据给试图
	c.Data["article"] = arti

	c.TplName = "content.html"

}

//显示更新内容详情页面
func (c *ArticleController) ShowUpdate() {
	//1.获取文章ID
	id, err := c.GetInt("id")
	beego.Info("id is ", id)
	if err != nil {
		beego.Info("获取文章ID错误", err)
		return
	}
	//2.查询数据库获取数据
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	err = o.Read(&arti)
	if err != nil {
		beego.Info("查询错误", err)
		return
	}
	//3.传递数据给试图
	c.Data["article"] = arti

	c.TplName = "update.html"

}

//处理更新业务数据
func (c *ArticleController) HandleUpdate() {
	//1.拿到数据
	id, _ := c.GetInt("id")
	artiName := c.GetString("articleName")
	content := c.GetString("content")
	f, h, err := c.GetFile("uploadname")
	var filename string
	if err != nil {
		beego.Info("上传文件失败")
		return
	} else {
		defer f.Close()

		//1.要限定格式
		fileext := path.Ext(h.Filename)
		if fileext != ".jpg" && fileext != "png" {
			beego.Info("上传文件格式错误")
			return
		}
		//2.限制大小
		if h.Size > 50000000 {
			beego.Info("上传文件过大")
			return
		}

		//3.需要对文件重命名，防止文件名重复
		filename = time.Now().Format("2006-01-02 15:04:05") + fileext //6-1-2 3:4:5
		c.SaveToFile("uploadname", "./static/img/"+filename)
	}

	//2.对数据进行一个处理
	if artiName == "" || content == "" {
		beego.Info("更新数据获取失败")
		return
	}

	//3.更新操作
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	err = o.Read(&arti)
	if err != nil {
		beego.Info("查询数据错误")
		return
	}
	arti.ArtiName = artiName
	arti.Acontent = content
	arti.Aimg = "./static/img/" + filename

	_, err = o.Update(&arti, "ArtiName", "Acontent", "Aimg")
	if err != nil {
		beego.Info("更新数据显示错误")
		return
	}
	//4.返回列表页面
	c.Redirect("/article/index", 302)
}

//删除操作
func (c *ArticleController) HandleDelete() {
	//1.拿到数据
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("获取id数据错误")
		return
	}

	//2.执行删除操作
	o := orm.NewOrm()
	arti := models.Article{Id: id}
	err = o.Read(&arti)
	if err != nil {
		beego.Info("查询错误")
		return
	}
	o.Delete(&arti)

	//3.返回列表页面
	c.Redirect("/article/index", 302)
}

//文章类型添加
func (c *ArticleController) AddType() {
	c.TplName = "type/add.html"
}

//文章类型添加操作
func (c *ArticleController) HandleAddType() {
	title := c.GetString("title")
	beego.Info(title)
	if title == "" {
		c.Ctx.WriteString("分类名不能为空")
		return
	}

	//1. orm对象
	o := orm.NewOrm()
	//2. 有一个要查的对象
	art := models.ArticleType{}

	//3. 给结构体赋值
	art.TypeName = title
	//art.ArticleType = artType
	//4. 插入
	re, err := o.Insert(&art) //返回插入的id 以及错误
	if err != nil {
		c.Ctx.WriteString("添加分类失败:")
		println("报错了:", err)
		return
	}
	beego.Info(re)
	if re > 0 {
		//c.Ctx.WriteString("添加成功")
		c.Redirect("/article/index", 302)
	} else {
		c.Ctx.WriteString("添加失败")
	}
}

////文章类型删除
//func (c *ArticleController)DelTypePost()  {
//	id,_ :=c.GetInt("id")
//	beego.Info(id)
//	if id > 0{
//
//		//1. orm对象
//		o := orm.NewOrm()
//		//2. 有一个要查的对象
//		art := models.ArticleType{}
//		art.Id = id
//		err := o.Read(&art)
//		if err != nil{
//			c.Ctx.WriteString("没有该分类")
//		}
//		num,err:=o.Delete(&art)
//		if err != nil{
//			c.Ctx.WriteString("删除出错")
//		}
//		if num == 1{
//			c.Ctx.WriteString("删除成功")
//		}else{
//			c.Ctx.WriteString("删除失败")
//		}
//
//	}else{
//
//	}
//}
