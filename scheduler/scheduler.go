package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit (request engine.Request)  {
	//create goroutine for each Request 为每一个Request创建goroutine
	go func() {
		s.workerChan <- request
	}()
}

// send the initRequest to Scheduler 把初始请求发送给Scheduler
func (s *SimpleScheduler) ConfigMasterWorkerChan (in chan engine.Request)  {
	s.workerChan = in
}