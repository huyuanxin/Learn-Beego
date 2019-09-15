package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
	"server/models"
)

type UserRegisterController struct {
	beego.Controller
}

type UserLoginController struct {
	beego.Controller
}

type UserChangePasswordController struct {
	beego.Controller
}

func PassWordMd5(str string) string  {
	//对密码进行加密
	//存在被解密的可能，进行加盐(未完成)
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *UserRegisterController) Post() {
	//对注册逻辑进行判断
	c.TplName="Register.html"
	userName:=c.GetString("UserName")
	userPassWord:= c.GetString("PassWord")
	userPassWordMd5:=PassWordMd5(c.GetString("PassWord"))//对获取的密码进行MD5加密
	if userName==""||userPassWord==""{
		beego.Info("密码或者用户名不能为空")
		return
	}
	o :=orm.NewOrm()
	user :=models.User{}
	user.Name=userName
	user.PassWord=userPassWordMd5
	_,err:=o.Insert(&user) //Insert返回两个值，但我们只要第二个，即err
	c.Ctx.Redirect(302, "/login")//重定向
	if err !=nil{
		beego.Info("插入失败",err)
		return
	}
}

func (c *UserRegisterController) Get() {
	//对注册页面进行控制
	c.TplName="Register.html"
}

func (c *UserLoginController) Get() {
	//对登录页面进行控制
	c.TplName="Login.html"
}

func (c *UserLoginController) Post() {
	//对登录逻辑进行判断

}
