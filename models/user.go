package models

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id int64			`form:"Id,hidden,<label></label>"`
    Login string		`form:"Login,text,<label>Usuario</label>"`
	Password string		`form:"Password,password,<label>Senha</label>"`
	Status bool			`form:"Status,checkbox,<label>Ativo</label>"`
}

func init() {
	orm.RegisterModel(new(User))
    orm.RegisterDriver("sqlite3", orm.DRSqlite)
    orm.RegisterDataBase("default", "sqlite3", "file:models/DbTables-v1.0.0.db")
}

func GetUser(id int64) User { 
	o := orm.NewOrm()
	user := User{Id:id}
	o.Read(&user)
	return user
}

func GetUserByLogin(login string) User { 
	o := orm.NewOrm()
	user := User{Login:login}
	o.Read(&user,"Login")
	return user
}

func GetAllUsers() (int64, []*User) {
	o := orm.NewOrm()
	var users []*User
	var count int64
	count,_ = o.QueryTable("user").All(&users)
	return count, users
}

func AddOne(user User) int64 {
	o := orm.NewOrm()
	
	id, _ := o.Insert(&user)
	return id
}

func Update(user User) int64 {
	o := orm.NewOrm()
	num, _:= o.Update(&user)
	return num
}

func Delete(userId int64) int64{
	user := User{Id: userId, Status: false}
	o := orm.NewOrm()
	num, _:= o.Update(&user,"Status")
	return num
}