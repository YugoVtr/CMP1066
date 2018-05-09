package models

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id int
    Login string
	Password string
	Status bool
}

func init() {
	orm.RegisterModel(new(User))
    orm.RegisterDriver("sqlite3", orm.DRSqlite)
    orm.RegisterDataBase("default", "sqlite3", "file:models/DbTables-v1.0.0.db")
}

func GetUser(id int) User { 
	o := orm.NewOrm()
	user := User{Id:id}
	o.Read(&user)
	return user
}

func GetAllUsers() (int64, []*User) {
	o := orm.NewOrm()
	var users []*User
	var count int64
	count,_ = o.QueryTable("user").All(&users)
	return count, users
}