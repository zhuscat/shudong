package controllers

import (
	"net/url"
	"shudong/models"
	"shudong/utils"
	"text/template"
)

type SearchController struct {
	BaseController
}

// url: /search?wd=title&p=1
func (self *SearchController) Get() {
	wd := self.GetString("wd")
	if wd == "" {
		self.alert("请输入关键词")
	}
	unescapeKeyword, err := url.QueryUnescape(wd)
	if err != nil {
		self.alert("wd参数有错误")
	}
	keyword := template.HTMLEscapeString(unescapeKeyword)
	page, _ := self.GetInt("p")
	if page <= 0 {
		page = 1
	}
	limit := 15
	offset := (page - 1) * limit
	if books, err := models.FindBooks(keyword, limit, offset); err == nil {
		totalNum, _ := models.GetBookCountWithContent(keyword)
		p := utils.NewPaginator(self.Ctx.Request, limit, totalNum)
		self.Data["Page"] = p
		self.Data["Keyword"] = wd
		self.Data["Books"] = books
		self.Layout = "layout.tpl"
		self.TplName = "search.tpl"
	} else {
		self.alert("出现错误")
	}
}
