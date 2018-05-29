package routers

import (
	ctrl "CMP1066/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router(ctrl.Default.String(), &ctrl.IndexController{})
	beego.Router(ctrl.Login.String(), &ctrl.LoginController{}, "get,post:Login")
	beego.Router(ctrl.Logout.String(), &ctrl.LoginController{}, "get:Logout")
	beego.Router(ctrl.Index.String(), &ctrl.IndexController{})
	beego.Router(ctrl.User.String(), &ctrl.UserController{})
	beego.Router(ctrl.User.String() + "/activate", &ctrl.UserController{}, "post:Activate")
	beego.Router(ctrl.Signup.String(), &ctrl.UserController{}, "get:Signup")
	beego.Router(ctrl.Signup.String(), &ctrl.UserController{}, "post:Post")
}
