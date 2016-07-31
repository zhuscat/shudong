package controllers

import (
	"fmt"
	"net/url"
	"shudong/models"
	"shudong/utils"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	BaseController
}

// Signin 登录
func (self *UserController) Signin() {
	if self.auth() {
		self.redirect("/")
	}
	if self.isPost() {
		username := strings.ToLower(self.GetString("username"))
		password := self.GetString("password")
		user, _ := models.GetUserByUsername(username)
		if user == nil {
			self.alert("获取用户失败")
		} else {
			if user.Password == utils.EncPassword(username, password) {
				// 登陆成功
				self.SetSession("userid", user.Id)
				self.SetSession("username", user.Name)
				next, _ := url.QueryUnescape(self.GetString("next"))
				if len(next) > 0 {
					self.redirect(next)
				} else {
					self.redirect("/")
				}
			} else {
				// 密码错误
				self.alert("密码错误")
			}
		}
	}
	self.TplName = "signin.tpl"
}

// Signup 注册
func (self *UserController) Signup() {
	if self.auth() {
		self.redirect("/")
	}
	if self.isPost() {
		username := strings.ToLower(self.GetString("username"))
		email := strings.ToLower(self.GetString("email"))
		password := self.GetString("password")
		passwordAgain := self.GetString("passwordagain")

		if password != passwordAgain {
			self.alert("两次密码输入不同")
		}

		if !utils.IsValidUsername(username) {
			self.alert("非法用户名")
		}
		if !utils.IsValidPassword(password) {
			self.alert("非法密码")
		}
		if !utils.IsValidEmail(email) {
			self.alert("非法邮箱")
		}

		uid, err := models.AddUser(username, email, password)

		if err == nil {
			self.SetSession("userid", uid)
			self.SetSession("username", username)
			self.alert("注册成功，请到你的邮箱查收邮件并激活，只有激活之后才能发布和购买书籍哦")
		} else {
			// 数据库存储出现错误
			//self.redirect("/signup")
			self.alert("数据库存储出现错误")
		}
	}
	self.TplName = "signup.html"
}

// Signout 登出
func (self *UserController) Signout() {
	self.SetSession("username", nil)
	self.SetSession("userid", nil)

	self.DelSession("username")
	self.DelSession("userid")

	self.redirect("/")
}

// Active 激活账号
func (self *UserController) Active() {
	if !self.auth() {
		self.alert("请在登录状态下访问该链接")
	}
	activeMessage := self.Ctx.Input.Param(":activemessage")
	if models.ActiveUser(activeMessage) == nil {
		self.alert("激活成功")
	} else {
		self.alert("激活码有误")
	}
}

// Active 请求发送激活消息
func (self *UserController) RequestActive() {
	if self.auth() {
		user, err := models.GetUserByUsername(self.userName)
		if err != nil {
			self.alert("出现错误，请稍后再试")
		}
		if user.Active == true {
			self.alert("你已经激活过了，不需要再激活了")
		} else {
			currentTime := time.Now()
			activeMessage := self.userName + strconv.FormatInt(currentTime.Unix(), 10)
			user.ActiveMessage = activeMessage
			duration, _ := time.ParseDuration("24h")
			user.ExpiredDate = currentTime.Add(duration)
			_, err := models.UpdateUserWithActiveMessage(user)
			if err == nil {
				utils.SendActiveMail(user.Email, user.ActiveMessage)
				self.alert("请到邮箱查收")
			} else {
				self.alert("出现了错误，请稍后再试")
			}
		}
	}
}

// ResetPassword 重置密码（通过邮箱获取连接的方式）
// url: /user/reset/:username/:resettoken
func (self *UserController) ResetPassword() {
	username := self.GetString(":username")
	resettoken := self.GetString(":resettoken")
	if self.isPost() {
		password := self.GetString("password")
		passwordAgain := self.GetString("passwordagain")
		if password != passwordAgain {
			self.alert("两次输入密码不同")
		}
		if models.ResetUserPassword(username, resettoken, password) {
			self.alert("重置成功")
		} else {
			self.alert("重置失败")
		}
	}
	user, _ := models.GetUserByUsername(username)
	if user != nil && user.ResetToken == resettoken && user.ResetExpiredDate.Unix() > time.Now().Unix() {
		self.TplName = "reset_password.tpl"
	} else {
		self.Ctx.Output.Body([]byte("invalid"))
		fmt.Println(user.ResetExpiredDate.Unix(), time.Now().Unix())
		self.StopRun()
		return
	}
}

// ForgetPassword 忘记密码
// url: /forgot
func (self *UserController) ForgetPassword() {
	if self.isPost() {
		username := self.GetString("username")
		email := self.GetString("email")
		if ok := models.SetUserResetToken(username, email); ok {
			self.alert("请到邮件查收")
		} else {
			self.alert("出现错误，请稍后再试")
		}
	}
	self.TplName = "forgot.tpl"
}

// 检查用户名是否可以注册
// 未测试过是否有用
// 不知道在不存在的用户的时候会不会发生 err
// TODO: new api
func (self *UserController) CheckUsernameAvailable() {
	username := self.GetString("username")
	out := make(map[string]interface{})
	if len(username) < 5 || len(username) > 20 {
		out["success"] = false
		out["info"] = "用户名长度为5-20个字符"
		self.jsonResult(out)
		return
	}
	_, err := models.GetUserByUsername(username)
	if err != nil {
		out["success"] = true
		out["info"] = "可以使用该用户名"
		self.jsonResult(out)
		return
	} else {
		out["success"] = false
		out["info"] = "用户名已被注册"
		self.jsonResult(out)
		return
	}
}

// CheckEmailAvailable 检查邮箱是否可用
func (self *UserController) CheckEmailAvailable() {
	email := self.GetString("email")
	out := make(map[string]interface{})
	if ok := utils.IsValidEmail(email); !ok {
		out["success"] = false
		out["info"] = "邮箱不合法"
		self.jsonResult(out)
		return
	}
	_, err := models.GetUserByEmail(email)
	if err != nil {
		out["success"] = false
		out["info"] = "邮箱已经被注册过了"
		self.jsonResult(out)
		return
	}
	out["success"] = true
	out["info"] = "可以使用该邮箱"
	self.jsonResult(out)

}

// 设置用户禁言
// TODO: new api
func (self *UserController) SetUserCanComment() {
	if !self.authAdmin() {
		self.redirect("/")
		return
	}
	uidStr := self.GetString("id")
	out := make(map[string]interface{})
	canComment, err := self.GetBool("can_comment")
	if err != nil {
		out["success"] = false
		out["info"] = "操作方法错误"
		self.jsonResult(out)
		return
	}
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		out["success"] = false
		out["info"] = "参数错误"
		self.jsonResult(out)
		return
	}
	user, err := models.GetUser(uid)
	if err != nil || user.Id == 0 {
		out["success"] = false
		out["info"] = "获取用户名错误"
		self.jsonResult(out)
		return
	}
	user.CanComment = canComment
	_, err = models.UpdateUser(user)
	if err != nil {
		out["success"] = false
		out["info"] = "数据库内部错误"
		self.jsonResult(out)
		return
	}
	out["success"] = true
	out["info"] = "操作成功"
	self.jsonResult(out)
	return
}
