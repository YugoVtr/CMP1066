package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	
)

type User struct {
	Id int64			`form:"Id,hidden,<label></label>"`
    Nick string			`orm:"size(64);unique" form:"Nick,text,<label>Usuario</label>" valid:"Required"`
	Password string		`orm:"size(128)" form:"Password,password,<label>Senha</label>" valid:"Required;MinSize(6)"`
	Status bool			`form:"Status,hidden,<label display='none'></label>"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (user *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(user, fields...); err != nil {
		return err
	}
	return nil
}

func Users() orm.QuerySeter {
	var table User
	return orm.NewOrm().QueryTable(table).OrderBy("Id")
}

func (user *User) Insert() error {
	if _, err := orm.NewOrm().Insert(user); err != nil {
		return err
	}
	return nil	
}

func (user *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(user, fields...); err != nil {
		return err
	}
	return nil
}

func (user *User) Delete() error {
	user.Status = false
	return user.Update("Status")
}

func (user *User) Activate() error {
	user.Status = true
	return user.Update("Status")
}
