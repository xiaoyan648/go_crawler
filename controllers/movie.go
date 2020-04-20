package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/micro/go-micro/util/log"
	"go_crawler/engine"
	"go_crawler/models"
	"go_crawler/parser"
	"go_crawler/scheduler"
	"time"
)

type CrawlMovieController struct {
	beego.Controller
}
/*
 获取单个电影的信息
 */
func (c *CrawlMovieController) Get() {
	//sUrl := "https://movie.douban.com/subject/25827935/" //七月与安生
	sUrl := "https://movie.douban.com/subject/26752088/" //我不是药神
	rsp := httplib.Get(sUrl)
	sMovieHtml, err := rsp.String()
	if err != nil {
		log.Error(err)
	}
	movieInfo := models.GetMovieInfo(sMovieHtml)
	//insert to DB
	models.AddMovieToDB(&movieInfo)
	c.Data["json"] = &movieInfo
	c.ServeJSON()
}
/*
 爬取电影和相关电影得信息
 */
func (c *CrawlMovieController) CrawlMovieV1() {
	req := models.Request{
		Url:        "https://movie.douban.com/subject/25827935/",
		ParserFunc: parser.ParseMovieInfo,
	}
	simple := &engine.SimpleEngine{}
	simple.Run(req)
}
/*
并发
 */
func (c *CrawlMovieController) CrawlMovieV2() {
	cc := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{}, //SimpleScheduler实现了scheduler接口，*SimpleScheduler，接口指针类型
		WorkerCount:100,
	}
	req := models.Request{
		Url:        "https://movie.douban.com/subject/25827935/",
		ParserFunc: parser.ParseMovieInfo,
	}
	cc.Run(req)
}
/*
 redis队列（可以解决分布式数据同步问题）
 */
func (c *CrawlMovieController) CrawlMovieV3() {
	//连接到redis
	models.ConnectRedis("127.0.0.1:6379")
	url := "https://movie.douban.com/subject/25827935/" //七月与安生
	//sUrl := "https://movie.douban.com/subject/26752088/" //我不是药神

	//先添加到队列中
	models.PutinQueue(url)
	c.Ctx.WriteString("crawling......")
	for !models.IsQueueEmpty(){

		url = models.PopfromQueue()
		//Url是否应该被访问过
		if models.IsVisit(url) {
			continue
		}
		rsp := httplib.Get(url)
		//设置User-agent以及cookie  403
		//rsp.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:50.0) Gecko/20100101 Firefox/50.0")
		//rsp.Header("Cookie", `bid=gFP9qSgGTfA; __utma=30149280.1124851270.1482153600.1483055851.1483064193.8; __utmz=30149280.1482971588.4.2.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; ll="118221"; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1483064193%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; _pk_id.100001.4cf6=5afcf5e5496eab22.1482413017.7.1483066280.1483057909.; __utma=223695111.1636117731.1482413017.1483055857.1483064193.7; __utmz=223695111.1483055857.6.5.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _vwo_uuid_v2=BDC2DBEDF8958EC838F9D9394CC5D9A0|2cc6ef7952be8c2d5408cb7c8cce2684; ap=1; viewed="1006073"; gr_user_id=e5c932fc-2af6-4861-8a4f-5d696f34570b; __utmc=30149280; __utmc=223695111; _pk_ses.100001.4cf6=*; __utmb=30149280.0.10.1483064193; __utmb=223695111.0.10.1483064193`)

		sMovieHtml, err := rsp.String()
		if err != nil {
			log.Error(err)
		}
		movieInfo := models.GetMovieInfo(sMovieHtml)

		//to db
		id, err := models.AddMovieToDB(movieInfo)
		if err != nil {
			log.Error("AddMovieToDB: ",err)
		}
		fmt.Println("id-->", id)
		//提取该页面的所有电影连接连接
		urls := models.GetMovieUrls(sMovieHtml)
		//放入队列中
		for _, url := range urls {
			models.PutinQueue(url)
			//c.Ctx.WriteString("<br>" + url + "</br>")
		}

		//以访问过的url集合，用于去重
		models.AddToSet(url)
		//休眠1s，防止ip禁用
		time.Sleep(time.Second)
	}
	c.Ctx.WriteString("end of crawl!")
}
