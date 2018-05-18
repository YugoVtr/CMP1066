package controllers

import (
	"CMP1066/models"
	"strconv" 			//Para converter string em int64
	"crypto/sha256"		//Para fazer o hash da senha 
)

type UserController struct {
	MainController
}

func (c *UserController) Get() {
	_, users := models.GetAllUsers()

	c.Data["Form"] = &models.User{Status: true}
	c.Data["Users"] = users
}

func (c *UserController) Post() {
	user := models.User{}
	c.ParseForm(&user)

	if user.Read("Login"); user.Id != 0 { 
		c.Data["json"] = map[string]interface{}{"User": "Username already exists"}
	} else { 
		//SHA-256
		hashSenha := sha256.Sum256([]byte(user.Password))
		user.Password = string(hashSenha[:])
		
		if err := user.Insert(); err != nil {
			c.Data["json"] = map[string]interface{}{"User": user.Id }
		} else { 
			c.Data["json"] = map[string]interface{}{"Error": err }
		}
	}
    c.ServeJSON()
}

func (c *UserController) Delete() {
	input   := c.Ctx.Input.Param(":id")
	user := models.User{}
	id,_:= strconv.ParseInt(input,10,64)
	user.Id = id 
	valid := user.Delete()

    c.Data["json"] = map[string]interface{}{"Success": valid }
    c.ServeJSON()
}