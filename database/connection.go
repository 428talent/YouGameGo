package database

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"yougame.com/yougame-server/models"
)

func init() {
	appConfig, err := config.NewConfig("ini", "./conf/app_local.conf")
	if err != nil {
		beego.Error(err)
	}
	mysqlUsername := appConfig.DefaultString("mysql_username", "root")
	mysqlPassword := appConfig.DefaultString("mysql_password", "root")
	connectString := fmt.Sprintf("%s:%s@/you_game?charset=utf8", mysqlUsername, mysqlPassword)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", connectString)
	orm.RegisterModel(new(models.Game), new(models.Image),new(models.Tag),new(models.Good),new(models.WishList))
	orm.RunSyncdb("default", false, true)

}
