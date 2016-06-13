package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

// TODO:之后移到config里面
var (
	salt1 = Configer.String("slat1")
	salt2 = Configer.String("salt2")
)

func EncPassword(username, pwd string) string {
	h := md5.New()
	io.WriteString(h, pwd)

	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	io.WriteString(h, salt1)
	io.WriteString(h, username)
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last := fmt.Sprintf("%x", h.Sum(nil))

	return last
}

// 看看还有没有改进的空间
func ResetToken() string {
	timeString := strconv.FormatInt(time.Now().Unix(), 10)
	h := md5.New()
	io.WriteString(h, timeString)

	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	io.WriteString(h, salt1)
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last := fmt.Sprintf("%x", h.Sum(nil))

	return last
}
