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
	mysqlAddress := appConfig.DefaultString("mysql_address", "0.0.0.0")
	mysqlPort := appConfig.DefaultString("mysql_port", "3306")
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/you_game?charset=utf8", mysqlUsername, mysqlPassword, mysqlAddress, mysqlPort)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", connectString)
	orm.RegisterModel(
		new(models.Game),
		new(models.Image),
		new(models.Tag),
		new(models.Good),
		new(models.WishList),
		new(models.CartItem),
		new(models.Order),
		new(models.OrderGood),
		new(models.Wallet),
		new(models.User),
		new(models.Profile),
		new(models.Transaction),
		new(models.UserGroup),
		new(models.Permission),
		new(models.Comment),
		new(models.InventoryItem),
		new(models.GameCollection),
	)
	orm.RunSyncdb("default", false, true)

}
