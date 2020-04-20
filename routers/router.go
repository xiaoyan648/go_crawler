package routers

import (
	"go_crawler/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//自动路由的设置 beego.AutoRouter(&controllers.FlashController{})自动生成
	//自动设置路由，controller上的注释来生成文档，
	//beego.Include(&controllers.FlashController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/movie", &controllers.CrawlMovieController{})
	beego.Router("/crawl/v2", &controllers.CrawlMovieController{},"get:CrawlMovieV1")
	beego.Router("/crawl/v1", &controllers.CrawlMovieController{},"get:CrawlMovieV2")
}
