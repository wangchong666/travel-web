package main

import (
	_ "travel-web/routers"
	_ "travel-web/util"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

