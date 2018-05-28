package controllers

import (
	"github.com/astaxie/beego"
	"CMP1066/models"
	"regexp"
)

type Controller int

const (
	Default Controller = 0 
	Login	Controller = 1
	Logout	Controller = 2
	Index	Controller = 3
	User	Controller = 4 
	Signup	Controller = 5 
)

type MainController struct {
	beego.Controller
	Userinfo *models.User
	IsLogin  bool
}

type NestPreparer interface {
	NestPrepare()
}

func (c *MainController) NestPrepare() {
	url := c.Ctx.Input.URL()
	if  !c.IsLogin && url != Login.String() && url != Signup.String() {
		c.Ctx.Redirect(302, c.LoginPath())
		return
	}
}

func (c *MainController) Prepare() {
	
	c.SetParams()
	
	c.IsLogin = c.GetSession("userinfo") != nil
	if c.IsLogin {
		c.Userinfo = c.GetLogin()
		c.Data["Login"] = false
	} else {
		c.Data["Login"] = true
	}
	
	c.Data["IsLogin"] = c.IsLogin
	c.Data["Userinfo"] = c.Userinfo

	//Default layout
	c.Layout = "layout/default.html"

    //Custom layout
	url := getControllerURL(c)


	//Default view
	c.TplName = "view" + url + ".html"
	c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = "html_head" + url + ".html"
    c.LayoutSections["Scripts"] = "scripts" + url + ".html"
	
	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (c *MainController) LoginPath() string {
	return c.URLFor("LoginController.Login")
}

func (c *MainController) GetLogin() *models.User {
	u := &models.User{Id: c.GetSession("userinfo").(int64)}
	u.Read()
	return u
}

func (c *MainController) DelLogin() {
	c.DelSession("userinfo")
}

func (c *MainController) SetLogin(user *models.User) {
	c.SetSession("userinfo", user.Id)
}

func (c *MainController) SetParams() {
	c.Data["Params"] = make(map[string]string)
	for k, v := range c.Input() {
		c.Data["Params"].(map[string]string)[k] = v[0]
	}
}

func (name Controller) String() string {
    names := [...]string{
        "/", 
		"/login", 
		"/logout",
		"/index",
		"/user",
		"/signup"}
		
	return names[name]
}

func getControllerURL(c *MainController) string {
	url := c.Ctx.Input.URL()
	r := regexp.MustCompile(`/[^/]*`)
    name := r.FindString(url)	

	if name == Default.String() {
		name = Index.String()
	}

	return name
}
