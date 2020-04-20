package parser

import (
	"github.com/micro/go-micro/util/log"
	"go_crawler/models"
	"strings"
)

//rename douban
func ParseMovieInfo(html string, url string) models.ParseResult{
	var result models.ParseResult
	//取得id
	//https://movie.douban.com/subject/27185018/ "https://movie.douban.com/subject/26985839/?from=subject-page"
	us := strings.Split(url, "/")
	id := us[len(us)-2]
	log.Info("id---------------- : ", id)
	urls := models.GetMovieUrls(html)
	for i,_ := range urls {
		request := models.Request{
			Url:       urls[i] ,
			ParserFunc: ParseMovieInfo,
		}
		result.Requests = append(result.Requests, request)
	}
	movieInfo := models.GetMovieInfo(html)
	result.Item = models.Item{
		Url:    url,
		Type:   "movie",
		Id:     id,
		Source: movieInfo,
	}
	return result
	
}



