package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

type User struct {
	Id uint64
	Name string
	PassWord string
}

func init()  {
	//设置数据库信息
	orm.RegisterDataBase("default","mysql","root:root@tcp(localhost:3306)/test?charset=utf8")
	//映射Model数据
	orm.RegisterModel(new(User))
	//建表
	orm.RunSyncdb("default",false,true)
}
