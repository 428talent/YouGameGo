package database

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"yougame.com/yougame-server/models"
)

func init() {
	appLocalConfig, err := config.NewConfig("ini", "./conf/app_local.conf")
	if err != nil {
		beego.Error(err)
	}

	mysqlAddress := os.Getenv("APPLICATION_MYSQL_HOST")
	if len(mysqlAddress) == 0{
		mysqlAddress = appLocalConfig.DefaultString("mysql_address", "0.0.0.0")
	}
	mysqlUsername := os.Getenv("APPLICATION_MYSQL_USERNAME")
	if len(mysqlUsername) == 0{
		mysqlUsername = appLocalConfig.DefaultString("mysql_username", "root")
	}
	mysqlPassword := os.Getenv("APPLICATION_MYSQL_PASSWORD")
	if len(mysqlPassword) == 0{
		mysqlPassword = appLocalConfig.DefaultString("mysql_password", "root")
	}
	mysqlPort := os.Getenv("APPLICATION_MYSQL_PORT")
	if len(mysqlPort) == 0{
		mysqlPort = appLocalConfig.DefaultString("mysql_port", "3306")
	}
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
		new(models.Follow),
	)
	orm.RunSyncdb("default", false, true)

}
