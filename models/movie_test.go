package models

import (
	"github.com/astaxie/beego/httplib"
	"testing"
)

func TestGetMovieUrls(t *testing.T) {
	//html, err := fetcher.Fetch("https://movie.douban.com/subject/25827935/")
	rsp := httplib.Get("https://movie.douban.com/subject/25827935/")
	html, _ := rsp.String()
	t.Log(html)
	t.Log(GetMovieUrls(html))

}
