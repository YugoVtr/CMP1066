package controllers

import (
	"CMP1066/models"
	."CMP1066/lib"		//Para fazer o hash da senha 
	"html/template"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	MainController
}

func (c *UserController) Get() {
	beego.ReadFromRequest(&c.Controller)

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	var users []*models.User
	models.Users().All(&users)

	var audits []*models.Audit
	models.Audits().All(&audits)

	var user models.User = models.User{}
	if id, erro := c.GetInt64("Id"); erro == nil {
		user.Id = id
		user.Read("Id")
		user.Password = ""
	} 	

	c.Data["Form"]   = &user
	c.Data["Users"]  = users
	c.Data["Audits"] = audits
}

func (c *UserController) Post() {

	var redirect string 
	redirect =  c.URLFor("UserController.Signup")

	valid := validation.Validation{}
	user := models.User{}
	c.ParseForm(&user)

	flash := beego.NewFlash()
	var err error 

	if b, erro := valid.Valid(&user) ; erro != nil || !b {
		form := valid.Errors[0]
		flash.Error(form.Key + " " + form.Message)
		flash.Store(&c.Controller)
		c.Redirect(redirect, 302)
		return
	}

	if user.Id != 0 {
		user.Password = Crypto(user.Password)
		err = user.Update()
	} else {
		user.Status = true
		user.Password = Crypto(user.Password)			
		err = user.Insert()
	}

	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect(redirect, 302)
		return
	}
	
	flash.Notice("Salvo com Sucesso")
	flash.Store(&c.Controller)
	c.Redirect(redirect, 303)
}

func (c *UserController) Delete() {
	flash := beego.NewFlash()

	if id, erro := c.GetInt64("Id"); erro == nil {
		user := models.User{}
		user.Id = id

		if err := user.Delete(); err == nil {
			flash.Warning("Usuário Inativo")
			flash.Store(&c.Controller)
		} else {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
		}

	} else {
		flash.Error(erro.Error())
		flash.Store(&c.Controller)
	}
}

func (c *UserController) Activate() {

	flash := beego.NewFlash()

	if id, erro := c.GetInt64("Id"); erro == nil {
		user := models.User{}
		user.Id = id

		if err := user.Activate(); err == nil {
			flash.Notice("Usuário Ativo")
			flash.Store(&c.Controller)
		} else {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
		}

	} else {
		flash.Error(erro.Error())
		flash.Store(&c.Controller)
	}

}

func (c *UserController) Signup() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["Form"] = &models.User{Status: true}
	c.Data["Action"] = c.getControllerURL(); 
}