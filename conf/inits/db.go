package inits

import (
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
)

func init() {
	dbname := "default"	
	runmode := beego.AppConfig.String("runmode")
	datasource := beego.AppConfig.String("datasource")

	switch runmode {
	case "prod":
		orm.RegisterDriver("sqlite3", orm.DRSqlite)
		orm.RegisterDataBase(dbname, "sqlite3", datasource)
	case "dev":
		orm.Debug = true
		fallthrough
	default:
		orm.RegisterDriver("sqlite3", orm.DRSqlite)
		orm.RegisterDataBase(dbname, "sqlite3", datasource)
	}

	// Drop table and re-create | Print log.
	force, verbose := false, true

	// Error.
	err := orm.RunSyncdb(dbname, force, verbose)
	if err != nil {
		panic(err)
	}

}
