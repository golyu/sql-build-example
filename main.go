package main

import (
	_ "github.com/golyu/sql-build-example/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	maxIdle := 100
	maxConn := 500
	dbname := beego.AppConfig.String("dbconfig::dbname")
	orm.RegisterDataBase("default", "mysql", dbname, maxIdle, maxConn)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
