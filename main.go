package main

import (
	"HelloBeego/db_mysql"
	_ "HelloBeego/routers"
	"github.com/astaxie/beego"
)

func main() {
	//1、连接数据库
	db_mysql.Connect()

	//2其他配置

	beego.Run()//代码简洁，
}