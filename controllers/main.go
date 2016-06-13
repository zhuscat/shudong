package controllers

import (
	"shudong/models"
	"shudong/utils"
)

type MainController struct {
	BaseController
}

// 主页的控制器
func (self *MainController) Get() {
	/** 这段话可以考虑放在一个小函数里面 **/
	pageLimit := 15
	pageNumber, _ := self.GetInt("p")
	if pageNumber <= 0 {
		pageNumber = 1
	}
	pageOffset := (pageNumber - 1) * pageLimit
	/** 这段话结束 **/
	totalNum, err := models.GetOnsaleBookCount()
	if err != nil {
		self.alert("获取书籍总数出错")
	}
	books, err := models.GetOnsaleBooks(pageLimit, pageOffset)
	if err != nil {
		self.alert("获取书籍出错")
	}
	p := utils.NewPaginator(self.Ctx.Request, pageLimit, totalNum)
	self.Data["Books"] = books
	self.Data["Page"] = p
	self.Layout = "layout.tpl"
	self.TplName = "index.tpl"
}
