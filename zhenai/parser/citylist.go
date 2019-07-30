package parser

import (
	"regexp"
	"crawler/engine"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9b-z]+)"[^>]*>([^<]+)</a>`
// 解析城市列表
func ParseCityList(bytes []byte) engine.ParseResult  {
	re := regexp.MustCompile(cityListRe)
	// submatch 是[][][]byte类型数据
	// 第一个[]表示匹配到多少条数据，第二个[]表示匹配的数据中要提取的内容
	submatch := re.FindAllSubmatch(bytes, -1)
	result := engine.ParseResult{}
	limit := 10
	for _,item := range submatch {
		//fmt.Println(string(item[2]))
		//fmt.Println(string(item[1]))
		result.Items = append(result.Items, "City:"+string(item[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:string(item[1]),
			ParseFunc:ParseCity,
		})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}

//func NilParseFunc([]byte) engine.ParseResult {
//	return engine.ParseResult{}
//}