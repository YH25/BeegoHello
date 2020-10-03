package routers

import (
	"HelloBeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//用户注册功能接口 http://127.0.0.1:9090/index
	//beego.Router("/index",&controllers.RegisterController{})
	//http://127.0.0.1:9090
    beego.Router("/index", &controllers.MainController{})
    //自己写一个web请求处理
    //http.HandleFunc("/login",处理函数)

}
