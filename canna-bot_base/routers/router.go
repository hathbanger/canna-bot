package routers

import (
	"canna-bot/canna-bot_base/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
