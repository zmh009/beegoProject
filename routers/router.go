package routers

import (
	"beegoProject/controllers"
	"github.com/astaxie/beego"
)

func init() {

	// 注册路由
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/categoy", &controllers.CategoyController{})
	beego.Router("/login", &controllers.LoginController{})
}
