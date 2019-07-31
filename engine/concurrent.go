package engine

import "log"

// 并发引擎
type ConcurrentEngine struct {
	Scheduler Scheduler  // 任务调度器
	WorkerCount int  // 任务并发数量
}
// 任务调度器
type Scheduler interface {
	Submit(request Request) // 提交任务
	ConfigMasterWorkerChan(chan Request) // 配置初始请求任务
}

func (e *ConcurrentEngine) Run (seeds ...Request)  {
	in := make(chan Request)  // scheduler的输入
	out := make(chan ParseResult) // worker的输出
	e.Scheduler.ConfigMasterWorkerChan(in) // 把初始请求提交给scheduler

	// create goRoutine
	for i:=0; i<e.WorkerCount; i++ {
		createWorker(in, out)
	}
	// engine commit the request to Scheduler： engine把请求任务提交给scheduler
	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}
	itemCount := 0
	for {
		// accept the result parsed by Worker：接受worker的解析结果
		result := <-out
		for _,item := range result.Items{
			log.Printf("Got item: #%d, %v\n", itemCount, item)
			itemCount++
		}
		// send the Request parsed by Worker to Scheduler：把worker解析出的Request送给scheduler
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

// create任务，调用worker，分发goroutine
func createWorker(in chan Request, out chan ParseResult)  {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}