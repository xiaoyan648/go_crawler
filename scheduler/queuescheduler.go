package scheduler

import (
	"go_crawler/models"
)

type QueuedScheduler struct {
	requestChan chan models.Request
	workerChan  chan chan models.Request
	//Mset mapset.Set 这样的话每个调度器都需要添加一个set，去实现相应功能，代码耦合，提出这块作为修饰类
}
/*type QueuedSchedulerDecorate struct {
	Mset mapset.Set
	QueuedScheduler QueuedScheduler
}*/
func NewQueuedScheduler() *QueuedScheduler{
	var s QueuedScheduler
	s.workerChan = make(chan chan models.Request)
	s.requestChan = make(chan models.Request)
	//s.Mset = mapset.NewSet()
	return &s
}

func (s *QueuedScheduler) Submit(request models.Request) {
	//去重,是在submit前后进行固定操作，也就是可以使用修饰模式，可以把去重容器维护在
	s.requestChan <- request
}
//queuescheduler是使用多个输入通道
func (s *QueuedScheduler) GetWorkerChan() chan models.Request{
	return make(chan models.Request)
}

func (s *QueuedScheduler) WorkerReady(w chan models.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	/*s.workerChan = make(chan chan models.Request)
	s.requestChan = make(chan models.Request)*/
	go func() {
		var requestQ [] models.Request
		var workerQ [] chan models.Request
		for {
			var activeRequest models.Request
			var activeWorker chan models.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
