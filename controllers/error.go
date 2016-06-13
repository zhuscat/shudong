package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (self *ErrorController) Error404() {
	self.Data["content"] = "你掉入了书洞之中"
	self.TplName = "404.tpl"
}
