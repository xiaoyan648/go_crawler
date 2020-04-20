package scheduler

import "go_crawler/models"

type SimpleScheduler struct {
	workerChan chan models.Request
}

func NewSimpleScheduler()  *SimpleScheduler{
	var s SimpleScheduler
	s.workerChan = make(chan models.Request)
	return &s
}

func (s *SimpleScheduler) Submit(request models.Request) {
	//一定要启动协程，因为worker只做数据得输入和过滤数据输出，submit新url和数据得显示都在外部
	//要让外部得 <-out执行得到
	go func() {s.workerChan <- request}()
}
//simple使用同一个通道获取数据
func (s *SimpleScheduler) GetWorkerChan() chan models.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(w chan models.Request) {

}

func (s *SimpleScheduler) Run () {

}