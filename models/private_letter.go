package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

//const TimeFormat = "2006-01-02 15:04:05"

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

//main page
//Get all dialogs from a user

func GetAllDialog(from int64) ([]*Dialog, error, error) {
	var dialogues1 []*Dialog
	var dialogues2 []*Dialog
	var dialogues []*Dialog
	_, err1 := orm.NewOrm().QueryTable("dialog").Filter("UserIdOne", from).All(&dialogues1)
	_, err2 := orm.NewOrm().QueryTable("dialog").Filter("UserIdTwo", from).All(&dialogues2)
	dialogues = append(dialogues1, dialogues2...)
	return dialogues, err1, err2
}

func GetAllToUsers(from int64, dialogs []*Dialog) []*User {
	var toUsers []*User
	var temp []int64
	for cnt := 0; cnt < len(dialogs); cnt++ {
		if from == dialogs[cnt].UserIdOne {
			temp = append(temp, dialogs[cnt].UserIdTwo)
		}
		if from == dialogs[cnt].UserIdTwo {
			temp = append(temp, dialogs[cnt].UserIdOne)
		}
		toUser, _ := GetUser(temp[cnt])
		toUsers = append(toUsers, toUser)
	}
	return toUsers
}

func GetLastPrivateLetters(from int64, to []*User) []PrivteLetter {
	var allLastPrivateLetters []PrivteLetter
	for cnt := 0; cnt < len(to); cnt++ {
		var lastPrivateLetter PrivteLetter
		dialogId := GetDialogId(from, to[cnt].Id)
		_, err := orm.NewOrm().QueryTable("privte_letter").Filter("DialogId", dialogId).OrderBy("-SendTime").Limit(1).All(&lastPrivateLetter)
		if err == nil {
			allLastPrivateLetters = append(allLastPrivateLetters, lastPrivateLetter)
		}
	}
	return allLastPrivateLetters
}

//a dialog with 2 users
//Send a private letter
func SendPrivateLetter(from int64, to int64, content string) bool {
	user1, _ := GetUser(from)
	user2, _ := GetUser(to)
	if user1 != nil && user2 != nil && user1.Id != user2.Id {
		//create a new letter
		privateLetter := PrivteLetter{FromId: user1.Id, ToId: user2.Id, Content: content, Read: false, SendTime: time.Now(), DialogId: GetDialogId(from, to)}
		_, err := orm.NewOrm().Insert(&privateLetter)
		if err == nil {
			return true
		} else {
			return false
		}
	}
	return false
}

func GetDialogId(uid1 int64, uid2 int64) int64 {
	//If the dialog exists,return the dialogId
	//else create a new dialog and return the new id
	//Uid1<Uid2
	if uid1 < uid2 {

	} else {
		temp := uid1
		uid1 = uid2
		uid2 = temp
	}
	var dialogue Dialog
	var newDialog Dialog
	newDialog.UserIdOne = uid1
	newDialog.UserIdTwo = uid2
	o := orm.NewOrm()
	err := o.QueryTable("dialog").Filter("UserIdOne", uid1).Filter("UserIdTwo", uid2).One(&dialogue)
	if err == orm.ErrNoRows {
		id, _ := o.Insert(&newDialog)
		return id
	}
	return dialogue.Id
}

func GetPrivateLetters(from int64, to int64, limit int, offset int) ([]*PrivteLetter, error) {
	var letters []*PrivteLetter
	dialog := GetDialogId(from, to)
	_, err1 := orm.NewOrm().QueryTable("privte_letter").Filter("DialogId", dialog).OrderBy("-SendTime").
		Limit(limit, offset).All(&letters)
	return letters, err1
}

func GetPrivateLetterCount(from int64, to int64) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable("privte_letter")

	count, err := qs.Filter("DialogId", GetDialogId(from, to)).Count()

	if err != nil {
		return 0
	} else {
		return count
	}
}

func ReadPrivateLetters(from int64, to int64, letters []*PrivteLetter) bool {
	for cnt := 0; cnt < len(letters); cnt++ {
		letters[cnt].Read = true
		_, err := orm.NewOrm().QueryTable("privte_letter").Filter("FromId", to).Filter("ToId", from).Update(orm.Params{
			"Read": true})
		if err != nil {
			return false
		}
	}
	return true
}

//If there are new letters
func HaveNewPrivateLetter(to int64) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("privte_letter")

	count, _ := qs.Filter("ToId", to).Filter("Read", false).Count()

	if count > 0 {
		return true
	} else {
		return false
	}
}
