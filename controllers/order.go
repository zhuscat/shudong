package controllers

import (
	"shudong/models"
	"strconv"
)

type OrderController struct {
	BaseController
}

// url: /order/customer/confirm/:bookid
// 自己不能买自己的书
// 什么时候能下订单
// 1. 商品还没有下架
// 2. 买家不是卖家
func (self *OrderController) AddOrder() {
	if !self.auth() {
		next := self.Ctx.Request.URL.String()
		url := "/signin?next=" + next
		self.redirect(url)
	}
	if !self.isActive() {
		self.alert("请先到激活你的账号（请到你提供的邮箱查收，如果已经过期，请到个人中心进行激活操作）")
	}
	bookId, err := self.GetInt64(":bookid")
	if err != nil {
		self.alert("参数错误")
	}
	book, err := models.GetBook(bookId)
	if err != nil {
		self.alert("获取书籍失败")
	}
	if self.userId == book.VendorId || book.Onsale == false {
		self.redirect("/book/detail/" + strconv.Itoa(int(bookId)))
	}
	vendor, _ := models.GetUser(book.VendorId)
	user, _ := models.GetUser(self.userId)
	if self.isPost() {
		if book == nil {
			self.alert("获取书籍失败")
		} else {
			addr := user.Place
			if _, err := models.AddOrder(bookId, self.userId, addr); err != nil {
				self.alert("存入数据库失败")
			} else {
				// TODO: 发一条系统站内信给卖家
				self.alert("订单下达成功")
			}
		}
	}
	self.Data["Book"] = book
	self.Data["Vendor"] = vendor
	self.TplName = "confirm-order.tpl"
}

// vendor 确认，此时已经有id了
// url: /order/vendor/confirm
func (self *OrderController) ConfirmOrder() {
	if !self.auth() {
		url := "/signin?next=/"
		self.redirect(url)
	}
	orderId, err := self.GetInt64("orderid")
	if err != nil {
		out := make(map[string]interface{})
		out["orderid"] = orderId
		out["success"] = false
		self.jsonResult(out)
	}
	err = models.ConfirmOrder(orderId, self.userId)
	if err != nil {
		out := make(map[string]interface{})
		out["orderid"] = orderId
		out["success"] = false
		self.jsonResult(out)
	}
	out := make(map[string]interface{})
	out["orderid"] = orderId
	out["success"] = true
	self.jsonResult(out)
}

// POST的方法关闭订单
// url: /order/close
func (self *OrderController) CloseOrder() {
	if !self.auth() {
		url := "/signin?next=/"
		self.redirect(url)
	}
	orderId, err := self.GetInt64("orderid")
	out := make(map[string]interface{})
	out["orderid"] = orderId
	if err != nil {
		out["success"] = false
		self.jsonResult(out)
	}
	err = models.CloseOrder(orderId, self.userId)
	if err != nil {
		out["success"] = false
		self.jsonResult(out)
	}
	out["success"] = true
	self.jsonResult(out)
}

// customer 完成订单
// url: /order/customer/complete/:orderid
func (self *OrderController) CompleteOrder() {
	if !self.auth() {
		url := "/signin?next=/"
		self.redirect(url)
	}
	orderId, err := self.GetInt64("orderid")
	out := make(map[string]interface{})
	if err != nil {
		out["success"] = false
		self.jsonResult(out)
	}
	err = models.CompeleteOrder(orderId, self.userId)
	if err != nil {
		out["success"] = false
		self.jsonResult(out)
	}
	out["success"] = true
	self.jsonResult(out)
}
