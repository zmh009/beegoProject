package controllers

import (
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/context"
)

type LoginController struct {
    beego.Controller
}

func (this *LoginController) Get() {
    isExit := this.Input().Get("exit") == "1"
    fmt.Println("isExit:", isExit, "exit:", this.Input().Get("exit"))
    if isExit {
        fmt.Println("exit")
        this.Data["IsLogin"] = false
        this.Ctx.SetCookie("uname", "", -1, "/")
        this.Ctx.SetCookie("pwd", "", -1, "/")
        this.Redirect("/", 302)
        return
    }
    this.TplName = "login.html"
}

func (this *LoginController) Post() {
    fmt.Println("this post")
    uname := this.Input().Get("uname")
    pwd := this.Input().Get("pwd")
    autoLogin := this.Input().Get("autoLogin") == "on"

    if  beego.AppConfig.String("uname") == uname &&
        beego.AppConfig.String("passwd") == pwd {

            maxAge := 0
            if autoLogin {
                maxAge = 1 << 31 - 1
            }
            this.Ctx.SetCookie("uname", uname, maxAge, "/")
            this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
    }

    this.Redirect("/", 302)
    return
}

func checkAccount(ctx *context.Context) bool {
    uname := ctx.GetCookie("uname")
    pwd := ctx.GetCookie("pwd")


    fmt.Println("this checkAccount, uname:", uname,", pwd:", pwd)
    return beego.AppConfig.String("uname") == uname &&
        beego.AppConfig.String("passwd") == pwd
}