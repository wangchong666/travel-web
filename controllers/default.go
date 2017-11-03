package controllers

import (
	"github.com/astaxie/beego"
)
var local = "en/"
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	c.Data["isHome"] = true
	c.Layout = local+"layout.html"
	c.TplName = local+"index.html"
	c.LayoutSections = make(map[string]string)
    c.LayoutSections["banner"] = local+"banner-home.html"
}

func (c *MainController) About() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	c.Data["isAbout"] = true
	c.Layout = local+"layout.html"
	c.TplName = local+ "about.html"
	c.LayoutSections = make(map[string]string)
    c.LayoutSections["banner"] =local+ "banner-others.html"
}


func (c *MainController) Contact() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	c.Data["isContact"] = true
	c.Layout = local+"layout.html"
	c.TplName = local+"contact.html"
	c.LayoutSections = make(map[string]string)
    c.LayoutSections["banner"] = local+"banner-others.html"
}

func (c *MainController) Posts() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	c.Data["isPosts"] = true
	c.Layout = local+"layout.html"
	c.TplName = local+"posts.html"
	c.LayoutSections = make(map[string]string)
    c.LayoutSections["banner"] = local+"banner-others.html"
}
