package utils

import (
	"fmt"
	"net/smtp"
	"strings"
)

// 之后移动到 config 里面
// 之后需要有一个 config 文件
var (
	username = Configer.String("email")
	password = Configer.String("emailpassword")
	host     = Configer.String("emailhost")
)

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/html; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType +
		"\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}

func SendActiveMail(to, activeMessage string) {
	activeUrl := "localhost:8080/user/active/" + activeMessage
	msg := "欢迎来到书洞，为了更好地使用书洞，请点击一下连接来激活您的账号：\n"
	msg += activeUrl
	err := SendToMail(username, password, host, to, "激活您的账号", msg, "text")
	fmt.Println(err)
}

func SendResetMail(to, usr, token string) {
	resetUrl := "localhost:8080/user/reset/" + usr + "/" + token
	err := SendToMail(username, password, host, to, "重置您的密码", resetUrl, "text")
	fmt.Println(err)
}
