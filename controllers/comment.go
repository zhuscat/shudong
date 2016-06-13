package controllers

import (
	"bytes"
	"html/template"
	"shudong/models"
	"strconv"
)

type CommentController struct {
	BaseController
}

func (self *CommentController) Get() {
	self.TplName = "comment.tpl"
}

func (self *CommentController) Post() {
	// 发表评论的时候要求的信息 当前用户 书籍
	if !self.auth() {
		self.redirect("/signin")
	}

	bookId, err := self.GetInt64("bookid")
	if err != nil {
		self.alert("书籍参数错误")
	}
	userId := self.userId
	book, _ := models.GetBook(bookId)
	if book == nil {
		self.alert("获取书籍失败")
	}
	vendorId := book.VendorId
	content := self.GetString("content")
	comment := models.NewComment()
	comment.BookId = book.Id
	comment.UserId = userId
	comment.VendorId = vendorId
	comment.Content = content
	_, err = models.CommentAdd(comment)
	if err != nil {
		self.alert("评论失败")
	} else {
		self.redirect("/book/detail/" + strconv.Itoa(int(bookId)) + "?p=1")
	}
}

//url: /comment/get/:bookid?p=1
func (self *CommentController) GetCommentHtml() {
	pageLimit := 10
	bookId, err := self.GetInt64(":bookid")
	if err != nil {
		self.Ctx.Output.Body([]byte("获取评论失败"))
		self.StopRun()
	}
	pageNumber, err := self.GetInt("p")
	if err != nil {
		self.Ctx.Output.Body([]byte("页数错误"))
		self.StopRun()
	}
	pageOffset := func() int {
		if pageNumber <= 0 {
			return 0
		}
		return (pageNumber - 1) * pageLimit
	}()
	comments, err := models.GetComments(bookId, pageLimit, pageOffset)
	if err != nil {
		self.Ctx.Output.Body([]byte("获取评论失败"))
		self.StopRun()
	}
	html := ""
	for _, comment := range comments {
		html += commentDiv(comment)
	}
	self.Ctx.Output.Body([]byte(html))
}

func commentDiv(comment *models.Comment) string {
	var buf bytes.Buffer
	t, _ := template.ParseFiles("views/comment-div.tpl")
	t.Execute(&buf, comment)
	return buf.String()
}
