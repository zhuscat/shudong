package controllers

import (
	"shudong/models"

	"github.com/astaxie/beego"
)

// BaseController 将会是该项目中所有其他控制器的父类
// 该类包含了一些基础功能，如用户验证
type BaseController struct {
	beego.Controller
	userId   int64
	userName string
	user     *models.User
}

func (self *BaseController) Prepare() {
	username, _ := self.GetSession("username").(string)
	userid, _ := self.GetSession("userid").(int64)
	self.userName = username
	self.userId = userid
	if len(username) > 0 {
		self.Data["Login"] = true
	} else {
		self.Data["Login"] = false
	}
}

// 判断用户是否已经登陆
// 如果登陆了的话，返回 true，并且设置 userId 和 userName
func (self *BaseController) auth() bool {
	// username, ok := self.GetSession("username").(string)
	// if !ok {
	// 	return false
	// }
	// userid, ok := self.GetSession("userid").(int64)
	// if !ok {
	// 	return false
	// }
	// self.userName = username
	// self.userId = userid
	// return true
	if len(self.userName) > 0 {
		return true
	} else {
		return false
	}
}

// 检查是不是管理员
// TODO: new api
func (self *BaseController) authAdmin() bool {
	if len(self.userName) > 0 && self.user.IsAdmin == true {
		return true
	} else {
		return false
	}
}

func (self *BaseController) isActive() bool {
	user, err := models.GetUser(self.userId)
	if err != nil {
		return false
	}
	if user == nil {
		return false
	}
	if user.Active == true {
		return true
	}
	return false
}

func (self *BaseController) isPost() bool {
	return self.Ctx.Request.Method == "POST"
}

func (self *BaseController) isGet() bool {
	return self.Ctx.Request.Method == "GET"
}

// 重定向
func (self *BaseController) redirect(url string) {
	self.Ctx.Redirect(302, url)
	self.StopRun()
}

// 显示信息
func (self *BaseController) alert(message string) {
	self.Data["Redirect"] = self.Ctx.Request.Referer()
	self.Data["Alert"] = message
	self.TplName = "alert.tpl"
	self.Render()
	self.StopRun()
}

func (self *BaseController) jsonResult(out interface{}) {
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}
