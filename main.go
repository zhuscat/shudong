package main

import (
	"shudong/controllers"
	_ "shudong/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
