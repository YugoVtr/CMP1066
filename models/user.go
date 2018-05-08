package models

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id int
    Login string
    Password string
}

func init() {
	orm.RegisterModel(new(User))
    orm.RegisterDriver("sqlite3", orm.DRSqlite)
    orm.RegisterDataBase("default", "sqlite3", "file:models/DbTables-v1.0.0.db")
}

func GetUser(id int) string{
	o := orm.NewOrm()
	o.Using("default")
	user := new(User)
	user.Id = id;
	o.Read(user)

	return user.Login
	

}