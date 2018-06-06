package controllers

import (
	"github.com/astaxie/beego"
	"CMP1066/models"
	"regexp"
	"os"
	str "strings"
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
	c.Data["Toolbar"] = true

	//Default layout
	c.Layout = "layout/default.html"

    //Custom layout
	url := getControllerURL(c)
 
	//Default VIEW
	var file string = "view" + url + ".html"
	base := baseURL() + "/views/"

	if _, err := os.Stat(base + file); err == nil {
		c.TplName = file
	}

	c.LayoutSections = make(map[string]string)

	// CSS
	file = "html_head" + url + ".html"
	if _, err := os.Stat(base + file); err == nil {
		c.LayoutSections["HtmlHead"] = file
	}

	// JS
	file = "scripts" + url + ".html"
	if _, err := os.Stat(base + file); err == nil {
		c.LayoutSections["Scripts"] = file
	}
	
	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}

	c.Data["Url"] = url
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

func baseURL() string { 
	base, _ := os.Getwd() 
	base = str.Replace(base,"\\","/",-1 )
	return base
}