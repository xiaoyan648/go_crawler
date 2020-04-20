package engine

import (
	"go_crawler/models"
	"go_crawler/parser"
	"testing"
)

func TestRun(t *testing.T) {
	req := models.Request{
		Url:        "https://movie.douban.com/subject/25827935/",
		ParserFunc: parser.ParseMovieInfo,
	}
	s := &SimpleEngine{}
	s.Run(req)
}
