package routers

import (
	"travel-web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/sitemap.xml", &controllers.MainController{},"*:Sitemap")
	beego.Router("/?:lang/index.html", &controllers.MainController{})
	beego.Router("/?:lang/contact.html", &controllers.MainController{},"*:Contact")
	beego.Router("/?:lang/about.html", &controllers.MainController{},"*:About")
	beego.Router("/?:lang/posts/page/:page", &controllers.MainController{},"*:Posts")
	beego.Router("/?:lang/post/name/:name", &controllers.MainController{},"*:Post")
	beego.Router("/?:lang/posts.html", &controllers.MainController{},"*:Posts")
	
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/images", "static/images")
}
