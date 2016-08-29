package controllers

import (
	//"container/list"
	//"time"

	"shudong/models"
	"shudong/utils"
)

type PrivateLetterController struct {
	BaseController
}

//url:/privateletter
//main page of privateletter
func (self *PrivateLetterController) GetDialogues() {
	if !self.auth() {
		next := self.Ctx.Request.URL.String()
		url := "/signin?next=" + next
		self.redirect(url)
	}

	if self.isPost() {
		username := self.GetString("username")
		content := self.GetString("content")
		to, _ := models.GetUserByUsername(username)
		toId := to.Id
		models.SendPrivateLetter(self.userId, toId, content)
	}

	user, _ := models.GetUser(self.userId)
	dialogs, _, _ := models.GetAllDialog(self.userId)
	toUsers := models.GetAllToUsers(self.userId, dialogs)
	lastLetters := models.GetLastPrivateLetters(self.userId, toUsers)

	self.Data["Dialogs"] = dialogs
	self.Data["ToUsers"] = toUsers
	self.Data["LastLetters"] = lastLetters
	self.Data["User"] = user
	self.Layout = "layout.tpl"
	self.TplName = "privateletter.tpl"
}

//url:/privateletter/:toid
//send a privateletter
func (self *PrivateLetterController) SendAPrivateLetter() {
	if !self.auth() {
		next := self.Ctx.Request.URL.String()
		url := "/signin?next=" + next
		self.redirect(url)
	}
	user, _ := models.GetUser(self.userId)
	toId, _ := self.GetInt64(":toid")
	toUser, _ := models.GetUser(toId)

	var content string
	if self.isPost() {
		content = self.GetString("content")
		models.SendPrivateLetter(self.userId, toId, content)
	}

	pageNumber, _ := self.GetInt("p")
	if pageNumber <= 0 {
		pageNumber = 1
	}
	pageLimit := 10
	pageOffset := (pageNumber - 1) * pageLimit
	var totalNum int64
	letters, _ := models.GetPrivateLetters(self.userId, toId, pageLimit, pageOffset)
	models.ReadPrivateLetters(self.userId, toId, letters)
	totalNum = models.GetPrivateLetterCount(self.userId, toId)
	p := utils.NewPaginator(self.Ctx.Request, pageLimit, totalNum)


	self.Data["Letters"] = letters
	self.Data["User"] = user
	self.Data["ToUser"] = toUser
	self.Data["Page"] = p
	self.Data["Content"] = content
	self.Layout = "layout.tpl"
	self.TplName = "send_letter.tpl"
}

//url:/privateletter/have-new-letter
//check if there is a new letter
func (self *PrivateLetterController) HaveNewPrivateLetter() {
	out := make(map[string]interface{})
	if !self.auth() {
		out["new"] = false
		self.jsonResult(out)
	}
	out["new"] = models.HaveNewPrivateLetter(self.userId)
	self.jsonResult(out)
}

//url:/privateletter/readmore
//read previous letters
