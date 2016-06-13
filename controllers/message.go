package controllers

import (
	"shudong/models"
	"shudong/utils"
)

type MessageController struct {
	BaseController
}

// url: /messages/?:tab
// read unread
func (self *MessageController) Get() {
	if !self.auth() {
		next := self.Ctx.Request.URL.String()
		url := "/signin?next=" + next
		self.redirect(url)
	}
	tab := self.GetString(":tab")
	var messages []*models.Message
	// 进行分页的操作
	pageNumber, _ := self.GetInt("p")
	if pageNumber <= 0 {
		pageNumber = 1
	}
	pageLimit := 10
	pageOffset := (pageNumber - 1) * pageLimit
	var totalNum int64
	if tab == "" {
		//
		messages, _ = models.GetMessages(self.userId, pageLimit, pageOffset)
		totalNum = models.GetMessageCount(self.userId)
	} else if tab == "read" {
		//
		messages, _ = models.GetMessageWithStatus(self.userId, true, pageLimit, pageOffset)
		totalNum = models.GetMessageCountWithStatus(self.userId, true)
	} else if tab == "unread" {
		//
		messages, _ = models.GetMessageWithStatus(self.userId, false, pageLimit, pageOffset)
		totalNum = models.GetMessageCountWithStatus(self.userId, false)
	} else {
		self.Abort("404")
	}
	p := utils.NewPaginator(self.Ctx.Request, pageLimit, totalNum)
	user, _ := models.GetUser(self.userId)
	self.Data["Messages"] = messages
	self.Data["User"] = user
	self.Data["Page"] = p
	self.Data["Tab"] = tab
	self.Layout = "layout.tpl"
	self.TplName = "messages.tpl"
}

// url: /message/confirm-read
// 只接收POST
func (self *MessageController) ConfirmRead() {
	if !self.auth() {
		self.redirect("/signin?next=/messages")
	}
	messageId, err := self.GetInt64("messageid")
	if err != nil {
		// 参数错误
		out := make(map[string]interface{})
		out["success"] = false
		self.jsonResult(out)
	}
	err = models.ReadMessage(self.userId, messageId)
	if err != nil {
		out := make(map[string]interface{})
		out["success"] = false
		self.jsonResult(out)
	}
	out := make(map[string]interface{})
	out["success"] = true
	self.jsonResult(out)
}

// url: /message/read-all
// 接收GET
func (self *MessageController) ReadAll() {
	if !self.auth() {
		self.redirect("/signin?next=/message")
	}
	err := models.ReadAllMessages(self.userId)
	if err != nil {
		out := make(map[string]interface{})
		out["success"] = false
		self.jsonResult(out)
	}
	out := make(map[string]interface{})
	out["success"] = true
	self.jsonResult(out)
}

// url: /message/have-new-message
// GET
func (self *MessageController) HaveNewMessage() {
	out := make(map[string]interface{})
	if !self.auth() {
		out["new"] = false
		self.jsonResult(out)
	}
	out["new"] = models.HaveNewMessage(self.userId)
	self.jsonResult(out)
}
