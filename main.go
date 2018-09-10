package main

import (
    "beegoProject/models"
    _ "beegoProject/routers"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
)

func init() {
    // 初始化注册DB
    models.RegisterDB()
}

func main() {
    orm.Debug = true
    orm.RunSyncdb("default", false, true)
	beego.Run()
}

