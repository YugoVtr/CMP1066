package main

import (
	_ "CMP1066/routers"
	_ "CMP1066/conf/inits"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

