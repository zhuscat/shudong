package models

import (
	"time"
)

// 私信整个会话
type Dialog struct {
	Id        int64
	UserIdOne int64
	UserIdTwo int64
}

// 一条私信的内容
// Read 接收方是否读过这条消息
// false 没有读过
// true 读过
type PrivteLetter struct {
	Id       int64
	FromId   int64
	ToId     int64
	Content  string
	Read     bool
	SendTime time.Time
	DialogId int64
}
