package main

import (
    "fmt"

	_ "canna-bot/canna-bot_base/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
    fmt.Println("Server now runnings on 8080!!")
}

