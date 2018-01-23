package main

import (
	"github.com/astaxie/beego"
	_ "you_game_go/database"
	_ "you_game_go/models"
	_ "you_game_go/routers"
)

func main() {
	beego.Run()
}
