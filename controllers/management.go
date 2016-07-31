package controllers

import (
	"fmt"
	"shudong/models"
	"shudong/utils"
)

// ManagementController 是与后台管理相关的控制器
type ManagementController struct {
	BaseController
}

// ShowManagementHome 展示后台页面的首页
func (mc *ManagementController) ShowManagementHome() {
	bookCount, _ := models.GetBookCount()
	commentCount, _ := models.TotalComment()
	userCount, _ := models.TotalUser()
	mc.Data["BookCount"] = bookCount
	mc.Data["commentCount"] = commentCount
	mc.Data["UserCount"] = userCount
	mc.Layout = "layout.tpl"
	mc.TplName = "management-home.html"
}

// Prepare 在 ManagementController 执行之前的验证
func (mc *ManagementController) NestPrepare() {
	if !mc.authAdmin() {
		mc.redirect("/signin")
		return
	}
}

// ManageBooks 管理书籍
func (mc *ManagementController) ManageBooks() {
	// 获取页数
	page, err := mc.GetInt("p")
	pageLimit := 10
	filterStr := mc.GetString(":filter")
	var filter []interface{}
	if filterStr == "" {
		filter = []interface{}{}
	} else if filterStr == "onsale" {
		filter = []interface{}{"onsale", true}
	} else if filterStr == "out-of-stock" {
		filter = []interface{}{"onsale", false}
	} else {
		mc.Abort("404")
		return
	}
	fmt.Println(filter)
	if err != nil || page <= 0 {
		page = 1
	}
	// 获取关键词
	keyword := mc.GetString("search")
	if keyword != "" {
		// 这是老旧的API，感觉不是太好，也只能先用着了
		books, err := models.FindBooks(keyword, pageLimit, pageLimit*(page-1))
		if err != nil {
			mc.Data["Books"] = []*models.Book{}
		}
		mc.Data["Books"] = books
		count, _ := models.GetBookCountWithContent(keyword)
		mc.Data["Page"] = utils.NewPaginator(mc.Ctx.Request, pageLimit, count)
	} else {
		books, count := models.BookGetList(page, pageLimit, filter...)
		if count > 0 {
			mc.Data["Books"] = books
			mc.Data["Page"] = utils.NewPaginator(mc.Ctx.Request, pageLimit, count)
		}
	}
	mc.Layout = "layout.tpl"
	mc.TplName = "book-management2.html"
}

// ManageComments 管理评论
func (mc *ManagementController) ManageComments() {
	// 获取评论
	page, err := mc.GetInt("p")
	pageLimit := 20
	if err != nil || page <= 0 {
		page = 1
	}
	keyword := mc.GetString("search")
	comments, count := models.CommentGetList(page, pageLimit, keyword, []interface{}{})
	mc.Data["Comments"] = comments
	mc.Data["Page"] = utils.NewPaginator(mc.Ctx.Request, pageLimit, count)
	mc.Layout = "layout.tpl"
	mc.TplName = "comments-management2.html"
}

// DeleteComment 删除评论
func (mc *ManagementController) DeleteComment() {
	commentid, err := mc.GetInt64("id")
	out := make(map[string]interface{})
	if err != nil {
		out["success"] = false
		out["info"] = "参数错误"
		mc.jsonResult(out)
		return
	}
	_, err = models.CommentDelete(commentid)
	if err != nil {
		out["success"] = false
		out["info"] = "数据库内部错误或评论不存在"
		mc.jsonResult(out)
		return
	}
	out["success"] = true
	out["info"] = "删除成功"
	mc.jsonResult(out)
	return
}

// ManageUsers 管理用户
func (mc *ManagementController) ManageUsers() {
	// 获取用户
	page, err := mc.GetInt("p")
	pageLimit := 16
	filterStr := mc.GetString(":filter")
	fmt.Println("filter", filterStr)
	var filter []interface{}
	if filterStr == "" {
		filter = []interface{}{}
	} else if filterStr == "normal" {
		filter = []interface{}{"can_comment", true}
	} else if filterStr == "ban" {
		filter = []interface{}{"can_comment", false}
	} else {
		mc.Abort("404")
		return
	}
	if err != nil || page <= 0 {
		page = 1
	}
	users, count := models.UserGetList(page, pageLimit, filter...)
	if count > 0 {
		mc.Data["Users"] = users
	}
	mc.Data["Page"] = utils.NewPaginator(mc.Ctx.Request, pageLimit, count)
	mc.Layout = "layout.tpl"
	mc.TplName = "user-management2.html"
}

// Broadcast 发送广播
func (mc *ManagementController) Broadcast() {
	out := make(map[string]interface{})
	content := mc.GetString("content")
	err := models.SendBroadcast(content)
	if err != nil {
		out["success"] = false
	} else {
		out["success"] = true
	}
	mc.jsonResult(out)
}

// SendNotification 给单个人发送通知
func (mc *ManagementController) SendNotification() {
	out := make(map[string]interface{})
	content := mc.GetString("content")
	username := mc.GetString("username")
	user, err := models.GetUserByUsername(username)
	if err != nil {
		out["success"] = false
		out["info"] = "未获取到你要发送给的用户，请确认用户是否存在"
		mc.jsonResult(out)
		return
	}
	ok := models.SendMessage(user.Id, content)
	if !ok {
		out["success"] = false
		out["info"] = "数据库内部错误"
		mc.jsonResult(out)
		return
	} else {
		out["success"] = true
		out["info"] = "发送成功"
		mc.jsonResult(out)
	}
}

// ManageUserCanComment 改变用户是否禁言
func (mc *ManagementController) ManageUserCanComment() {
	out := make(map[string]interface{})
	id, err := mc.GetInt64("id")
	if err != nil {
		out["success"] = false
		out["info"] = "参数错误"
		mc.jsonResult(out)
		return
	}
	user, err := models.GetUser(id)
	if err != nil {
		out["success"] = false
		out["info"] = "未获取到用户，请确认用户是否存在"
		mc.jsonResult(out)
		return
	}
	user.CanComment = !user.CanComment
	_, err = models.UpdateUser(user)
	if err != nil {
		out["success"] = false
		out["info"] = "数据库内部错误"
		mc.jsonResult(out)
		return
	}
	out["success"] = true
	out["info"] = "操作成功"
	mc.jsonResult(out)
}
