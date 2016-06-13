package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"shudong/utils"
	"time"
)

var (
	DATABASE_ERR = errors.New("Database Error")
	LOGICAL_ERR  = errors.New("Logical Error")
	NOTFOUND_ERR = errors.New("Not Found Error")
)

func init() {
	// 设置默认时区
	orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Shanghai")
	// 注册数据库
	orm.RegisterDataBase("default", "mysql", utils.Configer.String("database"), 30)
	// 注册模型
	orm.RegisterModel(new(User), new(Book), new(Order),
		new(Comment), new(Dialog), new(PrivteLetter), new(Message))
	// 创建表
	orm.RunSyncdb("default", false, true)
}
