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
	if userName==""||userPassWord==""{
		beego.Info("密码或者用户名不能为空")
		return
	}
	userPassWordMd5:=PassWordMd5(userPassWord)//对获取的密码进行MD5加密
	o :=orm.NewOrm()
	user :=models.User{}
	user.Name=userName
	user.PassWord=userPassWordMd5
	_,err:=o.Insert(&user) //Insert返回两个值，但我们只要第二个，即err
	if err !=nil{
		beego.Info("插入失败",err)
		return
	}
	beego.Info("注册成功")
	c.Ctx.Redirect(302, "/login")//重定向
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
	c.TplName="Login.html"
	o:=orm.NewOrm()
	user:=models.User{}
	/*
		根据Id来查询
		user.Id=1
		err:=o.Read(&user)
		if err!=nil{
			beego.Info("查询失败",err)
			return
		}
		beego.Info("查询成功",user.Id)
	*/

	/*
		根据其他属性来进行查询
			user.Name="123"
			err:=o.Read(&user,"Name")
			if err!=nil{
				beego.Info("查询失败",err)
				return
			}
			//id:=user.Id 查询后user被赋予查询到的值
			beego.Info("查询成功",user)
	*/
	userName:=c.GetString("UserName")
	userPassWord:= c.GetString("PassWord")
	if userName==""||userPassWord==""{
		beego.Info("账号或者密码不能为空")
		return
	}
	userPassWordMd5:=PassWordMd5(userPassWord)//对获取的密码进行MD5加密
	user.Name=userName
	err:=o.Read(&user,"Name")
	if err!=nil{
		beego.Info("查询失败",err)
		return
	}
	getPassWord:=user.PassWord
	if getPassWord!=userPassWordMd5{
		beego.Info("登录失败")
		return
	}
	beego.Info("登录成功")
	//后期需要添加cookies和页面跳转

}

func (c *UserChangePasswordController) Get() {
	//对修改密码界面进行控制
	c.TplName = "Index.html"
}
func (c *UserChangePasswordController) Post() {
	//对修改密码逻辑进行判断
	c.TplName = "Index.html"
}
