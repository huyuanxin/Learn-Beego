package routers

import (
	"beego-learning/controllers"
	"github.com/astaxie/beego"
)
//路由大小写敏感
func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/Register",&controllers.UserRegisterController{})
	beego.Router("/Login",&controllers.UserLoginController{})
	beego.Router("/ChangePassWord",&controllers.UserChangePasswordController{})
}
