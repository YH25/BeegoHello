package routers

import (
	"HelloBeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//router:路由
    beego.Router("/index", &controllers.MainController{})
    //自己写一个web请求处理
    //http.HandleFunc("/login",处理函数)

}
