package routers

import (
	"CMP1066/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/user/?:id", &controllers.UserController{})
}
