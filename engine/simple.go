package engine

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/micro/go-micro/util/log"
	"go_crawler/models"
	"time"
)

type SimpleEngine struct {}

func(e *SimpleEngine) Run(seeds ...models.Request) {
	var requests []models.Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	mset := mapset.NewSet()
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Info("Fetching ",r.Url)
		//去重
		if mset.Contains(r.Url) {
			continue
		}
		parseResult, err := worker(r)
		if err != nil {
			log.Error(err)
			continue
		}
		requests = append(requests, parseResult.Requests...)
		log.Info(parseResult.Item)
		//to db
		id, err := models.AddMovieToDB(parseResult.Item)
		if err != nil {
			log.Error("AddMovieToDB: ",err)
		}
		fmt.Println("id-->", id)
		//访问过了，加入去重集合
		mset.Add(r.Url)
		//休眠1s，防止ip禁用
		time.Sleep(time.Second)
	}
}



