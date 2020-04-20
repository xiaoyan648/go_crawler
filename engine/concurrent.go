package engine

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/micro/go-micro/util/log"
	"go_crawler/models"
	"go_crawler/util"
	"time"
)

//接口
type ReadyNotifier interface {
	WorkerReady(chan models.Request)
}

type Scheduler interface {
	Submit(request models.Request)
	GetWorkerChan()chan models.Request
	ReadyNotifier
	Run()
}
//添加一个chan缓冲，表示队列长度
type ConcurrentEngine struct {
	Scheduler   Scheduler //Sheduler
	WorkerCount int       //worker的数量
	ItemChan chan models.Item //接收目标数据的chan
}

func (e *ConcurrentEngine) Run(seeds ...models.Request) {

	out := make(chan models.ParseResult)
	mset := mapset.NewSet()

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.GetWorkerChan(), out, e.Scheduler) //创建worker
	}
	//参数seeds的request，要分配任务
	for _, r := range seeds {
		//去重,可以把去重容器维护在
		if mset.Contains(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
		mset.Add(r.Url)
	}
	//从out中获取result，对于item就打印即可，对于request，就继续分配
	for {
		var result models.ParseResult
		select {
			case <-time.After(time.Second*20):
				log.Info("process end")
				return
			case result = <-out:
				go func() {e.ItemChan <- result.Item}() //放入不要阻塞，会影响主协程的速度，worker,saver互不影响
		}
		for _, request := range result.Requests {
			if mset.Contains(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
			mset.Add(request.Url)
		}
	}
}

func createWorker(in chan models.Request, out chan models.ParseResult, ready ReadyNotifier) {
	//每一个worker一个request chan
	go func() {
		for{
			ready.WorkerReady(in) //表示这个worker已经没有数据在处理了，放入workerC han
			request := <- in
			/*body, err := fetcher.Fetch(request.Url)
			if err != nil {
				log.Error("createWorker fail : ", err)
			}
			out <- request.ParserFunc(body)*/
			result, err := worker(request)
			if err != nil {
				log.Error("createWorker fail : ", err)
				continue
			}
			out <- result
			util.RandomTimeSleep(3) //暂时采用这种方法，休眠防止ip禁用
		}
	}()
}

/*
func (e *ConcurrentEngine) Run(seeds ...models.Request) {

	//worker公用一个in，out
	in := make(chan models.Request)
	out := make(chan models.ParseResult)
	mset := mapset.NewSet()
	e.Scheduler.RegisterMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out) //创建worker
	}
	//参数seeds的request，要分配任务
	for _, r := range seeds {
		//去重
		if mset.Contains(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
		mset.Add(r.Url)
	}
	//从out中获取result，对于item就打印即可，对于request，就继续分配
	for {
		var result models.ParseResult
		select {
		case <-time.After(time.Second*3):
			log.Info("process end")
			return
		case result = <-out:
			log.Info(result.Item)
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan models.Request, out chan models.ParseResult) {
	go func() {
		for{
			request := <- in
			body, err := fetcher.Fetch(request.Url)
			if err != nil {
				log.Error("createWorker fail : ", err)
				return
			}
			out <- request.ParserFunc(body)
		}
	}()

}*/