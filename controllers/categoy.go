package controllers

import "github.com/astaxie/beego"

type CategoyController struct {
    beego.Controller
}

func (this *CategoyController) Get () {
    this.TplName = "categoy.html"
    this.Data["IsCategoy"] = true
}