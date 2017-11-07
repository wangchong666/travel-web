package controllers

import (
	"github.com/astaxie/beego"
	"travel-web/models"
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

func (c *MainController) Post() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	name := c.GetString(":name")
	post,err := models.GetPost(name)
	if err != nil{
		c.Abort("404")
	}
	c.Data["isPosts"] = true
	c.Data["post"] = post
	c.Layout = local+"layout.html"
	c.TplName = local+"post.html"
	c.LayoutSections = make(map[string]string)
    c.LayoutSections["banner"] = local+"banner-others.html"
}

func (c *MainController) Posts() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	page ,err:= c.GetInt(":page")
	pageSize,_ := beego.AppConfig.Int("page-size")
	if(err != nil){
		page = 1
	}
	count := models.GetPostsCount()
	posts := models.GetPosts(page-1,pageSize)
	
	total := count/pageSize
	if count%pageSize>0{
		total++
	}
	hasPrev := page >3
	hasNext := page < total-3
	start := page-2
	end:=page+2
	if start<1{
		end+=1-start
		start=1
	}
	if end >= total{
		start-=end-total+1
		end = total-1
	}

	var pages = make([]int, 0, 5)
	for i:=start;i<=end;i++{
		pages= append(pages,i)
	}
	c.Data["prev"] = page-3
	c.Data["next"] = page+3
	c.Data["pages"] = pages
	c.Data["postsCount"] = count
	c.Data["posts"] = posts
	c.Data["isPosts"] = true
	c.Data["hasPrev"] = hasPrev
	c.Data["hasNext"] = hasNext
	c.Layout = local +"layout.html"
	c.TplName = local +"posts.html"
	c.LayoutSections = make(map[string]string)
    c.LayoutSections["banner"] = local+"banner-others.html"
}

