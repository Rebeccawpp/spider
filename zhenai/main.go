package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

//const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func main()  {

	//re := regexp.MustCompile(cityListRe)
	//param := re.FindAllSubmatch("http://www.zhenai.com/zhenghun", -1)


	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc:parser.ParseCityList,
	})
}