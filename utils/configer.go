package utils

import (
	"github.com/astaxie/beego/config"
)

var Configer, _ = config.NewConfig("ini", "conf/config.conf")
