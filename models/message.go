package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 系统消息
// Read 接收方是否读过这条消息
// false 没有读过
// true 读过
type Message struct {
	Id       int64
	UserId   int64
	Content  string
	Read     bool
	SendTime time.Time
}

func SendMessage(uid int64, content string) bool {
	if user, _ := GetUser(uid); user != nil {
		message := Message{UserId: user.Id, Content: content, SendTime: time.Now()}
		_, err := orm.NewOrm().Insert(&message)
		if err == nil {
			return true
		} else {
			return false
		}
	}
	return false
}

func HaveNewMessage(uid int64) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("message")

	count, _ := qs.Filter("user_id", uid).Filter("read", false).Count()

	if count > 0 {
		return true
	} else {
		return false
	}
}

func GetMessage(id int64) (*Message, error) {
	message := Message{Id: id}
	if err := orm.NewOrm().Read(&message); err != nil {
		return nil, err
	} else {
		return &message, nil
	}
}

func GetMessageCount(uid int64) int64 {
	count, err := orm.NewOrm().QueryTable("message").Filter("user_id", uid).Count()
	if err != nil {
		return 0
	} else {
		return count
	}
}

func GetMessageCountWithStatus(uid int64, read bool) int64 {
	count, err := orm.NewOrm().QueryTable("message").Filter("user_id", uid).Filter("read", read).Count()
	if err != nil {
		return 0
	} else {
		return count
	}
}

func GetMessages(uid int64, limit int, offset int) ([]*Message, error) {
	var messages []*Message
	_, err := orm.NewOrm().QueryTable("message").Filter("UserId", uid).OrderBy("-SendTime").
		Limit(limit, offset).All(&messages)
	return messages, err
}

func GetMessageWithStatus(uid int64, read bool, limit int, offset int) ([]*Message, error) {
	var messages []*Message
	_, err := orm.NewOrm().QueryTable("message").Filter("UserId", uid).Filter("read", read).OrderBy("-SendTime").
		Limit(limit, offset).All(&messages)
	return messages, err
}

func ReadMessage(uid int64, mid int64) error {
	message := Message{Id: mid}
	err := orm.NewOrm().Read(&message)
	if err != nil {
		return err
	}
	message.Read = true
	_, err = orm.NewOrm().Update(&message, "Read")
	if err != nil {
		return err
	}
	return nil
}

func ReadAllMessages(uid int64) error {
	_, err := orm.NewOrm().Raw("UPDATE `message` SET `read` = 1 WHERE `user_id` = ?", uid).Exec()
	return err
}
