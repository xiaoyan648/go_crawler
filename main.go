package main

import (
	_ "go_crawler/models"
	_ "go_crawler/routers"
	"github.com/astaxie/beego"
)


func main() {
	beego.Run()
}

