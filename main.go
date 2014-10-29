package main

import (
	"github.com/astaxie/beego"
	_ "mcms/bll"     //注册业务层
	_ "mcms/dal"     //注册数据库
	_ "mcms/models"  //注册数据实体
	_ "mcms/routers" //注册路由
)

func main() {
	beego.Run()
}
