package controllers

import (
	"errors"
	"net/url"
	"shudong/models"
	"shudong/utils"
	"time"
)

type BookController struct {
	BaseController
}

func (self *BookController) validateOwnerAndGetBook(bookid int64) (*models.Book, error) {
	book, err := models.GetBook(bookid)
	if err != nil {
		return nil, err
	}
	if book.VendorId != self.userId {
		return nil, errors.New("The user is not the owner of the book")
	}
	return book, nil
}

func (self *BookController) PublishBook() {
	if !self.auth() {
		next := url.QueryEscape(self.Ctx.Request.URL.String())
		url := "/signin?next=" + next
		self.redirect(url)
	}
	if !self.isActive() {
		self.alert("请先到激活你的账号（请到你提供的邮箱查收，如果已经过期，请到个人中心进行激活操作）")
	}
	if self.isPost() {
		f, h, err := self.GetFile("picture")
		var filename string
		if err != nil {
			// 出现错误
			self.alert("上传文件失败")
		} else {
			filename = time.Now().Format("2006-01-02-03-04-05") + h.Filename
			err = self.SaveToFile("picture", "./static/www/book/"+filename)
			if err != nil {
				self.alert("存储文件失败")
			}
		}
		defer f.Close()
		title := self.GetString("title")
		author := self.GetString("author")
		publisher := self.GetString("publisher")
		price, err := self.GetFloat("price")
		if err != nil {
			self.alert("价格必须输入数字")
		}
		isbn := self.GetString("isbn")
		description := self.GetString("description")
		// 添加图书
		book := models.NewBook()
		book.Title = title
		book.Author = author
		book.Publisher = publisher
		book.Price = price
		book.Isbn = isbn
		book.VendorId = self.userId
		book.Description = description
		book.Picture = filename
		_, err = models.BookAdd(book)
		if err == nil {
			self.alert("发布成功")
		} else {
			self.alert("发布失败")
		}
	}
	self.Layout = "layout.tpl"
	self.TplName = "publish_book.tpl"
}

func (self *BookController) ShowBookDetail() {
	// 这里需要增加错误检测
	bookId, err := self.GetInt64(":bookid")
	if err != nil {
		self.Abort("404")
	}
	book, err := models.GetBook(bookId)
	if err != nil {
		self.Abort("404")
	}
	otherBooks, err := models.GetRecommendBookWithUserId(book.VendorId, book.Id)
	if err != nil {
		self.alert("获取推荐书籍失败")
	}
	pageLimit := 10
	showComment := false
	pageOffset := func() int {
		page, _ := self.GetInt("p")
		if page <= 0 {
			return 0
		}
		showComment = true
		return (page - 1) * pageLimit
	}()
	comments, err := models.GetComments(book.Id, pageLimit, pageOffset)
	if err != nil {
		self.alert("获取评论失败")
	}
	totalNum, err := models.GetCommentCount(book.Id)
	if err != nil {
		self.alert("获取评论条数失败")
	}
	var user *models.User
	if self.auth() {
		user, err = models.GetUser(self.userId)
		if err != nil {
			self.alert("获取用户失败")
		}
	}
	p := utils.NewPaginator(self.Ctx.Request, pageLimit, totalNum)
	self.Data["User"] = user
	self.Data["Book"] = book
	self.Data["OtherBooks"] = otherBooks
	self.Data["Comments"] = comments
	self.Data["Page"] = p
	self.Data["ShowComment"] = showComment
	self.Layout = "layout.tpl"
	self.TplName = "detail.tpl"
}

func (self *BookController) EditBook() {
	if !self.auth() {
		next := self.Ctx.Request.URL.String()
		url := "/signin?next=" + next
		self.redirect(url)
	}
	bookid, err := self.GetInt64(":bookid")
	if err != nil {
		self.Abort("404")
	}
	book, err := self.validateOwnerAndGetBook(bookid)
	if err != nil {
		self.alert("非法的用户或未成功获取书籍")
	}
	// 这个时候就可以获取到图书了，可以进行图书的编辑了
	if self.isPost() {
		f, h, err := self.GetFile("picture")
		// 有文件上传过来
		if err == nil || f != nil {
			filename := time.Now().Format("2006-01-02-03-04-05") + h.Filename
			err = self.SaveToFile("picture", "./static/www/book/"+filename)
			if err != nil {
				self.alert("存储失败")
			}
			book.Picture = filename
			f.Close()
		}
		title := self.GetString("title")
		author := self.GetString("author")
		publisher := self.GetString("publisher")
		price, err := self.GetFloat("price")
		desc := self.GetString("description")
		if err != nil {
			self.alert("价格必须要是数字")
		}
		isbn := self.GetString("isbn")
		book.Title = title
		book.Author = author
		book.Publisher = publisher
		book.Price = price
		book.Isbn = isbn
		book.Description = desc
		book.UpdatedTime = time.Now()
		_, err = models.UpdateBook(book)
		if err != nil {
			self.alert("更新失败")
		} else {
			self.alert("更新成功")
			self.StopRun()
		}
	}
	self.Data["Book"] = book
	self.Layout = "layout.tpl"
	self.TplName = "edit-book.tpl"
}

// url: /book/change/:bookid
func (self *BookController) ChangeBook() {
	if !self.auth() {
		next := self.Ctx.Request.URL.String()
		url := "/signin?next=" + next
		self.redirect(url)
	}
	bookid, err := self.GetInt64(":bookid")
	if err != nil {
		out := make(map[string]interface{})
		out["success"] = false
		out["msg"] = "书籍参数错误"
		self.jsonResult(out)
	}
	book, err := self.validateOwnerAndGetBook(bookid)
	if err != nil {
		out := make(map[string]interface{})
		out["success"] = false
		out["msg"] = "验证失败或书籍不存在"
		self.jsonResult(out)
	}
	book.Onsale = !book.Onsale
	_, err = models.UpdateBook(book)
	if err != nil {
		out := make(map[string]interface{})
		out["success"] = false
		out["msg"] = "操作失败"
		self.jsonResult(out)
	} else {
		out := make(map[string]interface{})
		out["success"] = true
		out["msg"] = "操作成功"
		self.jsonResult(out)
	}
}
