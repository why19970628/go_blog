package main

import (
	_ "class/models"
	_ "class/routers"
	"github.com/astaxie/beego"
	"strconv"
)

func main() {
	beego.AddFuncMap("PageUp",PageUp)
	beego.AddFuncMap("PageDown",PageDown)
	beego.Run()
}

func PageUp(page int) string {
	pageIndex := page -1
	if pageIndex <1{

	}
	pageIndex1 := strconv.Itoa(pageIndex)
	return pageIndex1
}

func PageDown(page int) string {
	pageIndex := page + 1
	pageIndex1 := strconv.Itoa(pageIndex)
	return pageIndex1
}