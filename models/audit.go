package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Audit struct {
	Id int64			`form:"Id,hidden,<label></label>"`
    User *User			`orm:"null;rel(one)"`
    Date time.Time
}

func init() {
	orm.RegisterModel(new(Audit))
}

func (audit *Audit) Read(fields ...string) error {
	if err := orm.NewOrm().Read(audit, fields...); err != nil {
		return err
	}
	return nil
}

func Audits() orm.QuerySeter {
	var table Audit
	return orm.NewOrm().QueryTable(table).OrderBy("Date")
}

func (audit *Audit) Insert() error {
	now := time.Now()
	audit.Date = now.Add(-3*time.Hour)
	if _, err := orm.NewOrm().Insert(audit); err != nil {
		return err
	}
	return nil	
}