package controllers

import (
	"html/template"
	"github.com/astaxie/beego"
	"CMP1066/lib"
)

type LoginController struct {
	MainController
}

func (c *LoginController) Login() {
	
	if c.IsLogin {
		c.Ctx.Redirect(302, c.URLFor("IndexController.Get"))
		return
	}

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	flash := beego.NewFlash()

	c.Data["Toolbar"] = false
	c.Data["Titulo"]  = "Login" 

	if c.Ctx.Input.IsPost() {
		nick := c.GetString("Nick")
		password := c.GetString("Password")
	
		user, err := lib.Authenticate(nick, password)
		if err != nil || user.Id < 1 {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
			return
		}
	
		flash.Store(&c.Controller)
		c.SetLogin(user)
		c.Redirect(c.URLFor("IndexController.Get"), 303)
	}
	return 
}

func (c *LoginController) Logout() {
	c.DelLogin()
	c.Ctx.Redirect(302, c.URLFor("LoginController.Login"))
}

