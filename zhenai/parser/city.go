package parser

import (
	"regexp"
	"crawler/engine"
)

var CityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
// user sex regexp, userInfoDetail haven't sex info, so fetch sex in userList page
var sexRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
// 城市页面用户解析器
func ParseCity(bytes []byte) engine.ParseResult {
	submatch := CityRe.FindAllSubmatch(bytes, -1)
	gendermatch := sexRe.FindAllSubmatch(bytes, -1)

	result := engine.ParseResult{}
	for k, item := range submatch {
		name := string(item[2])
		gender := string(gendermatch[k][1])
		result.Items = append(result.Items, "User:"+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:string(item[1]),
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name, gender)
			},
		})
	}
	return result
}