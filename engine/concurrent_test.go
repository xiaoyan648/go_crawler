package engine

import (
	"go_crawler/models"
	"go_crawler/parser"
	"go_crawler/saver"
	"go_crawler/scheduler"
	"testing"
	"time"
)

func TestConcurrentEngine_Run(t *testing.T) {
	t.Run("test QueuedScheduler", func(t *testing.T) {
		now := time.Now()
		cc := ConcurrentEngine{
			Scheduler: scheduler.NewQueuedScheduler() ,//SimpleScheduler实现了scheduler接口，*SimpleScheduler，接口指针类型
			WorkerCount:100,
			ItemChan:saver.ItemSaver(),
		}
		req := models.Request{
			Url:        "https://movie.douban.com/subject/25827935/",
			ParserFunc: parser.ParseMovieInfo,
		}

		cc.Run(req)
		t.Log(time.Now().Sub(now))
	})
	t.Run("test SimpleScheduler", func(t *testing.T) {
		now := time.Now()
		cc := ConcurrentEngine{
			Scheduler: scheduler.NewSimpleScheduler() ,//SimpleScheduler实现了scheduler接口，*SimpleScheduler，接口指针类型
			WorkerCount:100,
			ItemChan:saver.ItemSaver(),
		}
		req := models.Request{
			Url:        "https://movie.douban.com/subject/25827935/",
			ParserFunc: parser.ParseMovieInfo,
		}
		cc.Run(req)
		t.Log(time.Now().Sub(now))
	})

	/*
	问题一：再次爬取，无法覆盖，会出现重复数据
	解决方案：
	2.通过redis等维护一个去重集合，重新运行，redis中的去重集合也不会失效。这种是每次爬都只爬新数据，要设置失效时间，因为有的数据
	需要更新。
	2.一直运行这个程序，通过主线程的命令行去接收，如 update  movie douban就是 将电影库里属于豆瓣的表清空，
	redis里的去重集合失效重新爬取。
	update all ...更新所有这样，
	如果只新添数据而不更新以前数据，new movie douban不清空数据，不将redis失效

	对象引用失效，但对象里的属性是引用型变量未失效？
	*/
}

