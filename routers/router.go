package routers

import (
	"github.com/astaxie/beego"
	"beego-learning/controllers"
)
//路由大小写敏感
func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/Register",&controllers.UserRegisterController{})
	beego.Router("/Login",&controllers.UserLoginController{})
	beego.Router("/ChangePassWord",&controllers.UserChangePasswordController{})
}
