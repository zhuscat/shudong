package controllers

import (
	"shudong/models"
	"shudong/utils"
	"time"
)

type ProfileController struct {
	BaseController
}

const (
	publishedBookActive = `
		<li class="menu-area-item menu-area-item-selected"><a href="/profile/published/all">发布的书</a></li>
		<li class="menu-area-item"><a href="/profile/sale/all">卖的</a></li>
		<li class="menu-area-item"><a href="/profile/buy/all">买的</a></li>`
	publishedBookAll = `<li class="bar-item bar-item-selected"><a href="/profile/published/all">所有</a></li>
		<li class="bar-item"><a href="/profile/published/onsale">正在卖的</a></li>
		<li class="bar-item"><a href="/profile/published/out-of-stock">已下架的</a></li>`
	publishedBookOnsale = `<li class="bar-item"><a href="/profile/published/all">所有</a></li>
		<li class="bar-item bar-item-selected"><a href="/profile/published/onsale">正在卖的</a></li>
		<li class="bar-item"><a href="/profile/published/out-of-stock">已下架的</a></li>`
	publishedBookOutOfStock = `<li class="bar-item"><a href="/profile/published/all">所有</a></li>
		<li class="bar-item"><a href="/profile/published/onsale">正在卖的</a></li>
		<li class="bar-item bar-item-selected"><a href="/profile/published/out-of-stock">已下架的</a></li>`
	saleActive = `<li class="menu-area-item"><a href="/profile/published/all">发布的书</a></li>
		<li class="menu-area-item menu-area-item-selected"><a href="/profile/sale/all">卖的</a></li>
		<li class="menu-area-item"><a href="/profile/buy/all">买的</a></li>
		`
	saleAll = `<li class="bar-item bar-item-selected"><a href="/profile/sale/all">所有</a></li>
		<li class="bar-item"><a href="/profile/sale/request">已请求</a></li>
		<li class="bar-item"><a href="/profile/sale/response">已接受</a></li>
		<li class="bar-item"><a href="/profile/sale/complete">已完成</a></li>
		<li class="bar-item"><a href="/profile/sale/close">已关闭</a></li>`
	saleRequest = `<li class="bar-item"><a href="/profile/sale/all">所有</a></li>
		<li class="bar-item bar-item-selected"><a href="/profile/sale/request">已请求</a></li>
		<li class="bar-item"><a href="/profile/sale/response">已接受</a></li>
		<li class="bar-item"><a href="/profile/sale/complete">已完成</a></li>
		<li class="bar-item"><a href="/profile/sale/close">已关闭</a></li>`
	saleResponse = `<li class="bar-item bar-item-selected"><a href="/profile/sale/all">所有</a></li>
		<li class="bar-item"><a href="/profile/sale/request">已请求</a></li>
		<li class="bar-item bar-item-selected"><a href="/profile/sale/response">已接受</a></li>
		<li class="bar-item"><a href="/profile/sale/complete">已完成</a></li>
		<li class="bar-item"><a href="/profile/sale/close">已关闭</a></li>`
	saleComplete = `<li class="bar-item bar-item-selected"><a href="/profile/sale/all">所有</a></li>
		<li class="bar-item"><a href="/profile/sale/request">已请求</a></li>
		<li class="bar-item"><a href="/profile/sale/response">已接受</a></li>
		<li class="bar-item bar-item-selected"><a href="/profile/sale/complete">已完成</a></li>
		<li class="bar-item"><a href="/profile/sale/close">已关闭</a></li>`
	saleClose = `<li class="bar-item bar-item-selected"><a href="/profile/sale/all">所有</a></li>
		<li class="bar-item"><a href="/profile/sale/request">已请求</a></li>
		<li class="bar-item"><a href="/profile/sale/response">已接受</a></li>
		<li class="bar-item"><a href="/profile/sale/complete">已完成</a></li>
		<li class="bar-item bar-item-selected"><a href="/profile/sale/close">已关闭</a></li>`
	buyActive = `
		<li class="menu-area-item"><a href="/profile/published/all">发布的书</a></li>
		<li class="menu-area-item"><a href="/profile/sale/all">卖的</a></li>
		<li class="menu-area-item menu-area-item-selected"><a href="/profile/buy/all">买的</a></li>
		`
	buyAll = `<li class="bar-item bar-item-selected"><a href="/profile/buy/all">所有</a></li>
		<li class="bar-item"><a href="/profile/buy/request">已请求</a></li>
		<li class="bar-item"><a href="/profile/buy/response">已接受</a></li>
		<li class="bar-item"><a href="/profile/buy/complete">已完成</a></li>
		<li class="bar-item"><a href="/profile/buy/close">已关闭</a></li>`
	buyRequest = `<li class="bar-item"><a href="/profile/buy/all">所有</a></li>
		<li class="bar-item bar-item-selected"><a href="/profile/buy/request">已请求</a></li>
		<li class="bar-item"><a href="/profile/buy/response">已接受</a></li>
		<li class="bar-item"><a href="/profile/buy/complete">已完成</a></li>
		<li class="bar-item"><a href="/profile/buy/close">已关闭</a></li>`
	buyResponse = `<li class="bar-item"><a href="/profile/buy/all">所有</a></li>
		<li class="bar-item"><a href="/profile/buy/request">已请求</a></li>
		<li class="bar-item bar-item-selected"><a href="/profile/buy/response">已接受</a></li>
		<li class="bar-item"><a href="/profile/buy/complete">已完成</a></li>
		<li class="bar-item"><a href="/profile/buy/close">已关闭</a></li>`
	buyComplete = `<li class="bar-item"><a href="/profile/buy/all">所有</a></li>
		<li class="bar-item"><a href="/profile/buy/request">已请求</a></li>
		<li class="bar-item"><a href="/profile/buy/response">已接受</a></li>
		<li class="bar-item bar-item-selected"><a href="/profile/buy/complete">已完成</a></li>
		<li class="bar-item"><a href="/profile/buy/close">已关闭</a></li>`
	buyClose = `<li class="bar-item"><a href="/profile/buy/all">所有</a></li>
		<li class="bar-item"><a href="/profile/buy/request">已请求</a></li>
		<li class="bar-item"><a href="/profile/buy/response">已接受</a></li>
		<li class="bar-item"><a href="/profile/buy/complete">已完成</a></li>
		<li class="bar-item bar-item-selected"><a href="/profile/buy/close">已关闭</a></li>`
)

// url: /profile/:tab/:subtab
// 示例
// /profile/published/all
// /profile/published/onsale
// /profile/published/out-of-stock
// /profile/sale/all close complete request response
// /profile/buy/all close complete request response
func (self *ProfileController) ShowProfile() {
	if !self.auth() {
		next := self.Ctx.Request.URL.String()
		url := "/signin?next=" + next
		self.redirect(url)
	}
	tab := self.GetString(":tab")
	subtab := self.GetString(":subtab")
	pageParam, _ := self.GetInt("p")
	pageLimit := 10
	pageOffset := func() int {
		if pageParam <= 0 {
			return 0
		}
		return (pageParam - 1) * pageLimit
	}()
	topMenubar := ""
	rightMenubar := ""
	var totalNum int64
	if tab == "published" {
		rightMenubar = publishedBookActive
		if subtab == "all" {
			topMenubar = publishedBookAll
			books, _ := models.FindBooksWithUserId(self.userId, pageLimit, pageOffset)
			totalNum, _ = models.GetBookCountWithUserId(self.userId)
			self.Data["Books"] = books
		} else if subtab == "onsale" {
			topMenubar = publishedBookOnsale
			books, _ := models.FindBooksWithUserIdAndStatus(self.userId, true, pageLimit, pageOffset)
			totalNum, _ = models.GetBookCountWithUserIdAndStatus(self.userId, true)
			self.Data["Books"] = books
		} else if subtab == "out-of-stock" {
			topMenubar = publishedBookOutOfStock
			books, _ := models.FindBooksWithUserIdAndStatus(self.userId, false, pageLimit, pageOffset)
			totalNum, _ = models.GetBookCountWithUserIdAndStatus(self.userId, false)
			self.Data["Books"] = books
		}
	} else if tab == "sale" {
		rightMenubar = saleActive
		if subtab == "all" {
			topMenubar = saleAll
			orders, _ := models.FindOrderWithVendorId(self.userId, pageLimit, pageOffset)
			totalNum = models.GetOrderCount(-1, self.userId, 0)
			self.Data["Orders"] = orders
		} else if subtab == "request" {
			topMenubar = saleRequest
			orders, _ := models.FindOrderWithVendorIdAndStatus(self.userId, 0, pageLimit, pageOffset)
			self.Data["Orders"] = orders
			totalNum = models.GetOrderCount(0, self.userId, 0)
		} else if subtab == "response" {
			topMenubar = saleResponse
			orders, _ := models.FindOrderWithVendorIdAndStatus(self.userId, 1, pageLimit, pageOffset)
			totalNum = models.GetOrderCount(1, self.userId, 0)
			self.Data["Orders"] = orders
		} else if subtab == "complete" {
			topMenubar = saleComplete
			orders, _ := models.FindOrderWithVendorIdAndStatus(self.userId, 2, pageLimit, pageOffset)
			totalNum = models.GetOrderCount(2, self.userId, 0)
			self.Data["Orders"] = orders
		} else if subtab == "close" {
			topMenubar = saleClose
			orders, _ := models.FindOrderWithVendorIdAndStatus(self.userId, 3, pageLimit, pageOffset)
			totalNum = models.GetOrderCount(3, self.userId, 0)
			self.Data["Orders"] = orders
		}
	} else if tab == "buy" {
		rightMenubar = buyActive
		if subtab == "all" {
			topMenubar = buyAll
			orders, _ := models.FindOrderWithCustomerId(self.userId, pageLimit, pageOffset)
			totalNum = models.GetOrderCount(-1, 0, self.userId)
			self.Data["Orders"] = orders
		} else if subtab == "request" {
			topMenubar = buyRequest
			orders, _ := models.FindOrderWithCustomerIdAndStatus(self.userId, 0, pageLimit, pageOffset)
			totalNum = models.GetOrderCount(0, 0, self.userId)
			self.Data["Orders"] = orders
		} else if subtab == "response" {
			topMenubar = buyResponse
			orders, _ := models.FindOrderWithCustomerIdAndStatus(self.userId, 1, pageLimit, pageOffset)
			totalNum = models.GetOrderCount(1, 0, self.userId)
			self.Data["Orders"] = orders
		} else if subtab == "complete" {
			topMenubar = buyComplete
			orders, _ := models.FindOrderWithCustomerIdAndStatus(self.userId, 2, pageLimit, pageOffset)
			totalNum = models.GetOrderCount(2, 0, self.userId)
			self.Data["Orders"] = orders
		} else if subtab == "close" {
			topMenubar = buyClose
			orders, _ := models.FindOrderWithCustomerIdAndStatus(self.userId, 3, pageLimit, pageOffset)
			totalNum = models.GetOrderCount(3, 0, self.userId)
			self.Data["Orders"] = orders
		}
	}
	user, _ := models.GetUser(self.userId)
	self.Data["User"] = user
	self.Data["Menubar"] = topMenubar
	self.Data["RightMenubar"] = rightMenubar
	p := utils.NewPaginator(self.Ctx.Request, 10, totalNum)
	self.Data["Page"] = p
	self.Layout = "layout.tpl"
	self.TplName = "profile.tpl"
}

// url: /edit-profile
// TODO: 未验证错误
func (self *ProfileController) EditProfile() {
	if !self.auth() {
		next := self.Ctx.Request.URL.String()
		url := "/signin?next=" + next
		self.redirect(url)
	}

	if self.isPost() {
		phoneNumber, err := self.GetInt("phone-number")
		qqNumber := self.GetString("qq")
		weixinNumber := self.GetString("weixin")
		address := self.GetString("address")
		if err != nil {
			self.Ctx.Output.Body([]byte("表单有错误"))
			self.StopRun()
		}
		user, _ := models.GetUser(self.userId)
		user.PhoneNumber = phoneNumber
		user.Qq = qqNumber
		user.Weixin = weixinNumber
		user.Place = address
		_, err = models.UpdateUser(user)
		if err != nil {
			self.alert("数据库存储失败")
		} else {
			self.redirect("/profile/published/all")
		}
	}
	user, _ := models.GetUser(self.userId)
	self.Data["Avatar"] = user.Avatar
	self.Data["Username"] = user.Name
	self.Data["Email"] = user.Email
	self.Data["Phone"] = user.PhoneNumber
	self.Data["QQ"] = user.Qq
	self.Data["Weixin"] = user.Weixin
	self.Layout = "layout.tpl"
	self.TplName = "edit-profile.tpl"
}

// url: /edit-password
func (self *ProfileController) EditPassword() {
	//originalpassword
	//newpassword
	//newpasswordagain
	if !self.auth() {
		next := self.Ctx.Request.URL.String()
		url := "/signin?next=" + next
		self.redirect(url)
	}

	if self.isPost() {
		originalpassword := self.GetString("originalpassword")
		newpassword := self.GetString("newpassword")
		newpasswordagain := self.GetString("newpasswordagain")
		if newpassword != newpasswordagain {
			self.Ctx.Output.Body([]byte("表单错误"))
		}
		encOriPwd := utils.EncPassword(self.userName, originalpassword)
		user, err := models.GetUser(self.userId)
		if err != nil {
			self.Ctx.Output.Body([]byte("获取用户失败"))
		}
		if user.Password == encOriPwd {
			encNewPwd := utils.EncPassword(self.userName, newpassword)
			user.Password = encNewPwd
			_, err = models.UpdateUser(user)
			if err != nil {
				self.Ctx.Output.Body([]byte("更新密码失败"))
			} else {
				self.Ctx.Output.Body([]byte("更新密码成功"))
			}
		}
	}
	self.Layout = "layout.tpl"
	self.TplName = "edit-password.tpl"
}

// url: /upload-avatar 上传头像的表单提交到这里
func (self *ProfileController) UploadAvatar() {
	if !self.auth() {
		self.redirect("/signin")
	}
	f, h, err := self.GetFile("avatar")
	defer f.Close()
	if err != nil {
		self.Ctx.Output.Body([]byte("上传头像失败"))
		self.StopRun()
	}
	filename := time.Now().Format("2006-01-02-03-04-05") + h.Filename
	err = self.SaveToFile("avatar", "./static/www/avatar/"+filename)
	if err != nil {
		self.Ctx.Output.Body([]byte("上传头像失败"))
		self.StopRun()
	}
	user, _ := models.GetUser(self.userId)
	if user != nil {
		user.Avatar = filename
		_, err = models.UpdateUser(user)
	}
	if err != nil {
		self.Ctx.Output.Body([]byte("数据库错误"))
		self.StopRun()
	}
	self.redirect("/edit-profile")
}

// url: /user/:userid/?:tab
// 这个用来查看用户的个人信息（作为非本人查看）
// 之后应该可以整合到ShowProfile里面去
func (self *ProfileController) ViewUser() {
	userId, err := self.GetInt64(":userid")
	if err != nil {
		self.Ctx.Output.Body([]byte("参数错误"))
		self.StopRun()
	}
	if self.userId == userId {
		self.redirect("/profile/published/all")
	}
	/** 这段话可以考虑放在一个小函数里面 **/
	pageLimit := 10
	pageNumber, _ := self.GetInt("p")
	if pageNumber <= 0 {
		pageNumber = 1
	}
	pageOffset := (pageNumber - 1) * pageLimit
	/** 这段话结束 **/
	otherUser, err := models.GetUser(userId)
	if err != nil {
		self.Ctx.Output.Body([]byte("读取用户出错"))
		self.StopRun()
	}
	// 获取:tab参数
	tab := self.GetString(":tab")
	var totalNum int64
	var books []*models.Book
	if tab == "" {
		//
		totalNum, err = models.GetBookCountWithUserId(self.userId)
		if err != nil {
			self.Ctx.Output.Body([]byte("读取个数错误"))
			self.StopRun()
		}
		books, err = models.FindBooksWithUserId(userId, pageLimit, pageOffset)
		if err != nil {
			self.Ctx.Output.Body([]byte("读取书籍错误"))
			self.StopRun()
		}
	} else if tab == "onsale" {
		//
		totalNum, err = models.GetBookCountWithUserIdAndStatus(self.userId, true)
		if err != nil {
			self.Ctx.Output.Body([]byte("读取个数错误"))
			self.StopRun()
		}
		books, err = models.FindBooksWithUserIdAndStatus(userId, true, pageLimit, pageOffset)
		if err != nil {
			self.Ctx.Output.Body([]byte("读取书籍错误"))
			self.StopRun()
		}
	} else if tab == "out-of-stock" {
		//
		totalNum, err = models.GetBookCountWithUserIdAndStatus(self.userId, false)
		if err != nil {
			self.Ctx.Output.Body([]byte("读取个数错误"))
			self.StopRun()
		}
		books, err = models.FindBooksWithUserIdAndStatus(userId, false, pageLimit, pageOffset)
		if err != nil {
			self.Ctx.Output.Body([]byte("读取书籍错误"))
			self.StopRun()
		}
	} else {
		// 404 NOT FOUND
		self.Abort("404")
	}
	p := utils.NewPaginator(self.Ctx.Request, pageLimit, totalNum)
	self.Data["Page"] = p
	self.Data["Books"] = books
	self.Data["OtherUser"] = otherUser
	self.Data["Tab"] = tab
	self.Layout = "layout.tpl"
	self.TplName = "user.tpl"
}
