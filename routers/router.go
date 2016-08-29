package routers

import (
	"shudong/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// 与用户操作相关的路由
	userRouter()
	// 与书籍操作相关的路由
	bookRouter()
	// 与订单操作相关的路由
	orderRouter()
	// 与消息操作相关的路由（如站内信，私信[还没有做])
	notificationRouter()
	// 与后台管理相关的路由
	managementRouter()
}

func userRouter() {
	// 查看用户信息（非用户本人）
	beego.Router("/user/:userid/?:tab", &controllers.ProfileController{}, "get:ViewUser")
	// 登录
	beego.Router("/signin", &controllers.UserController{}, "*:Signin")
	// 注册
	beego.Router("/signup", &controllers.UserController{}, "*:Signup")
	// 登出
	beego.Router("/signout", &controllers.UserController{}, "*:Signout")
	// 激活账号
	beego.Router("/user/active/:activemessage", &controllers.UserController{}, "get:Active")
	// 请求发送一条激活的信息到邮件
	beego.Router("/user/active", &controllers.UserController{}, "get:RequestActive")
	// 通过邮件重置密码（注：邮件会发送这个链接，然后用户点进去，填写表单后即可重置密码）
	beego.Router("/user/reset/:username/:resettoken", &controllers.UserController{}, "*:ResetPassword")
	// 用户忘记密码
	beego.Router("/forgot", &controllers.UserController{}, "*:ForgetPassword")
	// 用户编辑个人信息
	beego.Router("/edit-profile", &controllers.ProfileController{}, "*:EditProfile")
	// 用户修改密码（此时用户知道其原密码)
	beego.Router("/edit-password", &controllers.ProfileController{}, "*:EditPassword")
	// 用户上传头像
	beego.Router("/upload-avatar", &controllers.ProfileController{}, "post:UploadAvatar")
	// 用户查看个人信息（用户本人）
	beego.Router("/profile/:tab/:subtab", &controllers.ProfileController{}, "get:ShowProfile")
	// ajax 检查用户名
	beego.Router("/user/check-username", &controllers.UserController{}, "post:CheckUsernameAvailable")
	// ajx 检查邮箱
	beego.Router("/user/check-email", &controllers.UserController{}, "post:CheckEmailAvailable")
}

func bookRouter() {
	// 获取书籍详情
	beego.Router("/book/detail/:bookid", &controllers.BookController{}, "*:ShowBookDetail")
	// 发布商品
	beego.Router("/book/publish", &controllers.BookController{}, "*:PublishBook")
	// 上下架商品
	beego.Router("/book/change/:bookid", &controllers.BookController{}, "get:ChangeBook")
	// 编辑图书
	beego.Router("/book/edit/:bookid", &controllers.BookController{}, "*:EditBook")
	// 搜索书籍
	beego.Router("/search", &controllers.SearchController{})
	// 评论商品
	beego.Router("/book/comment", &controllers.CommentController{})
	// 获取评论 此为一个api 用于ajax的get请求
	beego.Router("/comment/get/:bookid", &controllers.CommentController{}, "get:GetCommentHtml")
}

func orderRouter() {
	// 买家下订单
	beego.Router("/order/customer/confirm/:bookid([0-9]+)", &controllers.OrderController{}, "*:AddOrder")
	// 卖家接收订单
	beego.Router("/order/accept", &controllers.OrderController{}, "post:ConfirmOrder")
	// 买家完成订单
	beego.Router("/order/complete", &controllers.OrderController{}, "post:CompleteOrder")
	// 卖家或买家关闭订单
	beego.Router("/order/close", &controllers.OrderController{}, "*:CloseOrder")
}

func notificationRouter() {
	// 站内信
	beego.Router("/message/?:tab", &controllers.MessageController{})
	// 标记一条站内信为已读
	beego.Router("/message/confirm-read", &controllers.MessageController{}, "post:ConfirmRead")
	// 标记所有站内信为已读
	beego.Router("/message/read-all", &controllers.MessageController{}, "get:ReadAll")
	// 查询是否有新信息
	beego.Router("/message/have-new-message", &controllers.MessageController{}, "get:HaveNewMessage")
}

func managementRouter() {
	// 后台管理首页
	beego.Router("/management/", &controllers.ManagementController{}, "get:ShowManagementHome")
	// 后台管理书籍页面
	beego.Router("/management/books/?:filter", &controllers.ManagementController{}, "get:ManageBooks")
	// 后台管理评论页面
	beego.Router("/management/comments", &controllers.ManagementController{}, "get:ManageComments")
	// 后台管理用户页面
	beego.Router("/management/users/?:filter", &controllers.ManagementController{}, "get:ManageUsers")
	// 发送广播
	beego.Router("/management/broadcast", &controllers.ManagementController{}, "post:Broadcast")
	// 发送通知给单个用户
	beego.Router("/management/notification", &controllers.ManagementController{}, "post:SendNotification")
	// 删除评论
	beego.Router("/management/comments/delete", &controllers.ManagementController{}, "post:DeleteComment")
	// 禁言或解除禁言
	beego.Router("/management/users/ban", &controllers.ManagementController{}, "post:ManageUserCanComment")
}

func privateLetterRouter() {
	//进入私信界面
	beego.Router("/privateletter", &controllers.PrivateLetterController{}, "*:GetDialogues")
	//发私信
	beego.Router("/privateletter/:toid", &controllers.PrivateLetterController{}, "*:SendAPrivateLetter")
	//查询是否有新私信
	beego.Router("/privateletter/have-new-letter", &controllers.PrivateLetterController{}, "get:HaveNewPrivateLetter")
}
