package controllers

import (
	"github.com/astaxie/beego"
	"travel-web/models"
	"travel-web/dto"
)
type PostsController struct {
	beego.Controller
}

func (c *PostsController) Get() {
	// c.Data["Website"] = "beego.me"
	page,_ := c.GetInt("_page",1)
	pageSize,_:=c.GetInt("_limit",10)
	posts := models.GetPosts(page-1,pageSize)
	total := models.GetPostsCount()
	c.Data["json"] = dto.PageData{Data:posts,Total:total}
	c.ServeJSON();
}
