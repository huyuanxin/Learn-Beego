package routers

import (
	"github.com/astaxie/beego"
	"server/controllers"
)
//路由大小写敏感
func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/register",&controllers.UserRegisterController{})
	beego.Router("/login",&controllers.UserLoginController{})
	beego.Router("/changepassword",&controllers.UserLoginController{})
}
