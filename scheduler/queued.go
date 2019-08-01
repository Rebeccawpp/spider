package scheduler

import "crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request // Request channel
	workerChan chan chan engine.Request // worker channel,其中每一个worker是一个chan engine.Request 类型
}

// 提交请求任务到requestChannel
func (s *QueuedScheduler) Submit (request engine.Request) {
	s.requestChan <- request
}
func (s *QueuedScheduler) ConfigMasterWorkerChan (chan engine.Request)  {
	panic("implement me")
}
// 通知，有一个worker可以接收request
func (s *QueuedScheduler) WorkerReady (w chan engine.Request)  {
	s.workerChan <- w
}
func (s *QueuedScheduler) Run()  {
	// 生成channel
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		// 创建请求队列和工作队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeWorker chan engine.Request
			var activeRequest engine.Request
			// 当requestQ和workerQ同时有数据时
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-s.requestChan: // 当requestChan 收到数据
				requestQ = append(requestQ, r)
			case w := <-s.workerChan: // 当workerChan 收到数据
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest: //当请求队列和任务队列都不为空时，给任务队列分配任务
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]



			}
		}
	}()
}