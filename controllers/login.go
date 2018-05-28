package controllers

import (
	"html/template"
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

	
	c.Data["Toolbar"] = false
	c.Data["Titulo"]  = "Login" 

	if !c.Ctx.Input.IsPost() {
		return
	}

	nick := c.GetString("Nick")
	password := c.GetString("Password")

	user, err := lib.Authenticate(nick, password)
	if err != nil || user.Id < 1 {
		return
	}

	c.SetLogin(user)

	c.Redirect(c.URLFor("IndexController.Get"), 303)
}

func (c *LoginController) Logout() {
	c.DelLogin()
	c.Ctx.Redirect(302, c.URLFor("LoginController.Login"))
}

