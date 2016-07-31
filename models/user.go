package models

import (
	"encoding/base64"
	"errors"
	"fmt"
	"shudong/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 用户
type User struct {
	Id               int64
	Name             string
	Email            string
	Password         string
	Active           bool `orm:"default(false)"`
	Place            string
	Qq               string
	Weixin           string
	PhoneNumber      int
	Avatar           string `orm:"default(default-avatar.png)"`
	CreatedDate      time.Time
	IsAdmin          bool `orm:"default(false)"`
	ActiveMessage    string
	ResetToken       string
	ResetExpiredDate time.Time `orm:"null"`
	ExpiredDate      time.Time
	CanComment       bool `orm:"default(true)"`
}

// Name Email 构成的内容应该是唯一的
func (u *User) TableUnique() [][]string {
	return [][]string{
		[]string{"Name", "Email"},
	}
}

// 输入一个 Id，获取 User
func GetUser(uid int64) (*User, error) {
	user := User{Id: uid}
	err := orm.NewOrm().Read(&user)
	if err == nil {
		return &user, nil
	} else {
		return nil, err
	}
}

// 输入用户名，获取 User
func GetUserByUsername(name string) (*User, error) {
	user := User{Name: name}
	err := orm.NewOrm().Read(&user, "Name")
	if err == nil {
		return &user, nil
	} else {
		return nil, err
	}
}

// GetUserByEmail 输入邮箱，获取 User
func GetUserByEmail(email string) (*User, error) {
	user := User{Email: email}
	err := orm.NewOrm().Read(&user, "Email")
	if err == nil {
		return &user, nil
	}
	return nil, err
}

// 创建一个用户
func AddUser(name string, email string, pwd string) (int64, error) {
	o := orm.NewOrm()
	nameErr := o.Read(&User{Name: name}, "Name")
	emailErr := o.Read(&User{Email: email}, "Email")
	// 如果出现错误，则用户不存在，则插入用户记录
	if nameErr != nil && emailErr != nil {
		encpwd := utils.EncPassword(name, pwd)
		user := User{Name: name, Email: email, Password: encpwd, CreatedDate: time.Now()}
		rawMessage := name + strconv.FormatInt(user.CreatedDate.Unix(), 10)
		fmt.Println("password", encpwd)
		// 加密 暂时使用 base64，之后可能会修改
		user.ActiveMessage = base64.StdEncoding.EncodeToString([]byte(rawMessage))
		// 过期时间
		duration, _ := time.ParseDuration("24h")
		user.ExpiredDate = user.CreatedDate.Add(duration)
		id, err := o.Insert(&user)
		// 成功写入数据库才发送邮件
		if err == nil {
			utils.SendActiveMail(email, user.ActiveMessage)
		}
		return id, err
	} else if nameErr == nil {
		return 0, errors.New("duplicate user name")
	} else {
		return 0, errors.New("duplicate email")
	}
}

// func (self *User) Update(fields ...string) error{
// 	_, err := orm.NewOrm().Update(self, fields...)
// 	return err
// }
//
// func (self *User) Active(activeMessage string) error {
// 	if self.ExpiredDate.Unix() < time.Now().Unix() {
// 		return errors.New("激活信息已过期")
// 	}
// 	self.Active = true
// 	return self.Update("Active")
// }

// 激活用户
func ActiveUser(activeMessage string) error {
	o := orm.NewOrm()
	user := User{ActiveMessage: activeMessage}
	err := o.Read(&user, "ActiveMessage")
	if err != nil {
		return errors.New("user not exist")
	} else {
		// 过期了
		if user.ExpiredDate.Unix() < time.Now().Unix() {
			return errors.New("expired active message")
		} else {
			user.Active = true
			_, err = o.Update(&user, "Active")
			if err != nil {
				return errors.New("fail to active")
			} else {
				return nil
			}
		}
	}
}

// 更新用户激活信息字段
func UpdateUserWithActiveMessage(user *User) (int64, error) {
	duration, _ := time.ParseDuration("1h")
	user.ExpiredDate = time.Now().Add(duration)
	return orm.NewOrm().Update(user, "ActiveMessage", "ExpiredDate")
}

// 设置重置密码的token
func SetUserResetToken(username, email string) bool {
	user, _ := GetUserByUsername(username)
	if user == nil || user.Email != email {
		return false
	}
	user.ResetToken = utils.ResetToken()
	duration, _ := time.ParseDuration("1h")
	user.ResetExpiredDate = time.Now().Add(duration)
	if _, err := orm.NewOrm().Update(user, "ResetToken", "ResetExpiredDate"); err != nil {
		return false
	}
	utils.SendResetMail(user.Email, user.Name, user.ResetToken)
	return true
}

// 重置密码
func ResetUserPassword(username, token, password string) bool {
	user, _ := GetUserByUsername(username)
	if user == nil || user.ResetToken != token || len(password) < 6 {
		return false
	}
	user.Password = utils.EncPassword(username, password)
	if _, err := orm.NewOrm().Update(user, "Password"); err != nil {
		return false
	}
	return true
}

func UpdateUser(user *User) (int64, error) {
	return orm.NewOrm().Update(user)
}

// UserGetList 接受参数 page 页数 pageSize 一页最多的个数 filters 过滤条件
// 如: [can_comment: true]
// 返回 User数组和User的总量
// 给查找用户列表一个统一的接口

func UserGetList(page int, pageSize int, filters ...interface{}) ([]*User, int64) {
	offset := (page - 1) * pageSize
	var users []*User
	query := orm.NewOrm().QueryTable("user")
	fmt.Println(filters)
	if len(filters) >= 2 && len(filters)%2 == 0 {
		l := len(filters)
		for i := 0; i < l; i += 2 {
			query = query.Filter(filters[i].(string), filters[i+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-CreatedDate").Limit(pageSize, offset).All(&users)
	return users, total
}

// TotalUser 获取用户总数
func TotalUser() (int64, error) {
	return orm.NewOrm().QueryTable("user").Count()
}
