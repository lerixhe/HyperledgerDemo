package routers

import (
	"HyperledgerDemo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/test", &controllers.MainController{}, "get:GetValue")
}
