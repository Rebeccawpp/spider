package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main()  {
//-----version-1
	//engine.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParseFunc:parser.ParseCityList,
	//})

//-----version-2 concurrent
	e := engine.ConcurrentEngine{ // config spider engine 配置爬虫引擎
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 50,
	}
	e.Run(engine.Request{ // config spider target information 配置爬虫目标信息
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc:parser.ParseCityList,
	})
}