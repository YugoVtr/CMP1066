package controllers

import (
	"CMP1066/models"
	."CMP1066/lib"		//Para fazer o hash da senha 
	"html/template"
	//"github.com/astaxie/beego"
)

type UserController struct {
	MainController
}

func (c *UserController) Get() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	var users []*models.User
	models.Users().All(&users)

	var user models.User = models.User{}
	if id, erro := c.GetInt64("Id"); erro == nil {
		user.Id = id
		user.Read("Id")
		user.Password = ""
	} 

	c.Data["Form"] = &user
	c.Data["Users"] = users
}

func (c *UserController) Post() {
	c.Data["json"] = map[string]interface{}{"User": "Success"}
	user := models.User{}
	c.ParseForm(&user)

	if user.Id != 0 {
		user.Password = Crypto(user.Password)
		user.Update()
	} else if models.Users().Filter("Nick", user.Nick).Exist() {
		c.Data["json"] = map[string]interface{}{"User": "Nick already exists"}
	} else {
		user.Status = true
		user.Password = Crypto(user.Password)			
		user.Insert()
	}
	
	c.SetLogin(&user)
	c.Redirect(c.URLFor("IndexController.Get"), 303)
}

func (c *UserController) Delete() {
	if id, erro := c.GetInt64("Id"); erro == nil {
		user := models.User{}
		user.Id = id
		valid := user.Delete()
	
		c.Data["json"] = map[string]interface{}{"Success": valid }
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"Error": erro }
		c.ServeJSON()
	}
}

func (c *UserController) Activate() {

	if id, erro := c.GetInt64("Id"); erro == nil {
		user := models.User{}
		user.Id = id
		valid := user.Activate()
	
		c.Data["json"] = map[string]interface{}{"Success": valid }
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"Error": erro }
		c.ServeJSON()
	}
}

func (c *UserController) Signup() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["Form"] = &models.User{Status: true}
}