package engine

import (
	"go_crawler/fetcher"
	"go_crawler/models"
)

func worker(request models.Request) (models.ParseResult, error) {
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		return models.ParseResult{}, err
	}
	return request.ParserFunc(body,request.Url), nil
}

