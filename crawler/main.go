package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/persist"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler_distributed/config"
)

func main() {
	// 1. Single Task Edition
	// engine.SimpleEngine{}.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })

	// 2. Concurrent Edition
	// e := engine.ConcurrentEngine{
	// 	Scheduler:   &scheduler.SimpleScheduler{},
	// 	WorkerCount: 100,
	// }

	// 3. Queue Scheduler Edition
	// e := engine.ConcurrentEngine{
	// 	Scheduler:   &scheduler.QueuedScheduler{},
	// 	WorkerCount: 100,
	// }
	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })

	// 4. Page
	itemChan, err := persist.ItemSaver("crawler_dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:      &scheduler.QueuedScheduler{},
		WorkerCount:    100,
		ItemChan:       itemChan,
		RequestProcess: engine.Worker,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList, config.ParseCityList),
	})

}
