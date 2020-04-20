package models

//解析后返回的结果
type ParseResult struct {
	Requests []Request
	Item    Item
}

type Request struct {
	Url        string                   //解析出来的URL
	ParserFunc func(string, string) ParseResult //传入html，url，处理这个URL所需要的函数
}


type Item struct {
	Url string //URL
	Type string //存储到ElasticSearch时的type
	Id  string //用户Id
	Source interface{}
}


func NilParser([] byte) ParseResult {
	return ParseResult{}
}