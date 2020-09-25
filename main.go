package main

import (
	_ "HelloBeego/routers"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//定义config变量，接收并赋值为全局配置变量
	config := beego.AppConfig
	//获取配置选项
	appName := config.String("appname")
	fmt.Println("项目应用名称:",appName)
	port,err := config.Int("httpport")
	if err != nil {
		//配置信息解析错误
		panic("项目配置信息解析错误，清查验后重试")
	}
	fmt.Println("应用监听端口",port)
	driver := config.String("db_driver")//数据驱动
	dbUser := config.String("db_user")//数据可用户名
	dbPassword := config.String("da_password")//密码
	dbIP  := config.String("db_ip")
	dbName := config.String("db_name")

	db, err := sql.Open(driver,dbUser+":"+dbPassword+"@tcp("+dbIP+")/"+dbName+"?charset=utf8")
	if err != nil {
		panic("数据链接打开失败，请重试")
	}
	fmt.Println(db)
	beego.Run()
}

