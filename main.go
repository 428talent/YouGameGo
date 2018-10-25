package main

import (
	"github.com/astaxie/beego"
	_ "yougame.com/yougame-server/auth"
	_ "yougame.com/yougame-server/database"
	_ "yougame.com/yougame-server/models"
	_ "yougame.com/yougame-server/routers"
	_ "yougame.com/yougame-server/security"
)

func main() {
	beego.SetStaticPath("/static","static")
	beego.Run()
}
