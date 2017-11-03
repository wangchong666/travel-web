package routers

import (
	"travel-web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index.html", &controllers.MainController{})
	beego.Router("/contact.html", &controllers.MainController{},"*:Contact")
	beego.Router("/about.html", &controllers.MainController{},"*:About")
	beego.Router("/posts.html", &controllers.MainController{},"*:Posts")
	
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/images", "static/images")
}
