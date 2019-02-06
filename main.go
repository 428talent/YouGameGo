package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"time"
	_ "yougame.com/yougame-server/database"
	"yougame.com/yougame-server/log"
	_ "yougame.com/yougame-server/log"
	_ "yougame.com/yougame-server/mail"
	_ "yougame.com/yougame-server/models"
	_ "yougame.com/yougame-server/routers"
	_ "yougame.com/yougame-server/security"
	_ "yougame.com/yougame-server/transaction"
)

func main() {
	beego.SetStaticPath("/static", "static")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	orm.Debug = true
	log.LogClient.Channel <- &log.LogPayload{
		Message: "repository starting",
		Level:   "info",
		Time:    time.Now(),
	}
	beego.Run()

}
