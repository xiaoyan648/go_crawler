package saver

import (
	"encoding/json"
	"github.com/olivere/elastic"
	"go_crawler/models"
	"go_crawler/util/es"
	"reflect"
	"testing"
)


func Test_save(t *testing.T) {
	//save-----------------------------------------------
	const index = "crawler_test"
	item := models.Item{
			Url:  "https://movie.douban.com/subject/27185018/",
			Type: "movie",
			Id:   "1214814888",
			Source: models.MovieInfo{
				Id:                   0,
				Movie_id:             0,
				Movie_name:           "test",
				Movie_pic:            "test",
				Movie_director:       "test",
				Movie_writer:         "test",
				Movie_country:        "test",
				Movie_language:       "test",
				Movie_main_character: "test",
				Movie_type:           "test",
				Movie_on_time:        "test",
				Movie_span:           "test",
				Movie_grade:          "test",
				Create_time:          "test",
			},
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}
	t.Run("BASE", func(t *testing.T) {
		err := es.Save(client, item, index)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := es.Search(client, index,item.Type,item.Id)
		if err != nil {
			t.Fatal(err)
		}
		var result models.Item
		err = json.Unmarshal(*resp.Source, &result)
		if err != nil {
			t.Fatal(err)
		}
		//json类型和对象类型
		//want {http://movie.com/u/1214814888 movie 1214814888 map[Create_time:test Id:0 Movie_country:test Movie_director:test Movie_grade:test Movie_id:0 Movie_language:test Movie_main_character:test Movie_name:test Movie_on_time:test Movie_pic:test Movie_span:test Movie_type:test Movie_writer:test]};
		// expected {http://movie.com/u/1214814888 movie 1214814888 {0 0 test test test test test test test test test test test test}}
		JsonItem, _ := FromJsonObj(item)

		//断言,struct里全是基本类型可以直接==
		//现在item里有movieinfo的复杂类型，使用deepEqueal比较
		if !reflect.DeepEqual(JsonItem, result) {
			t.Errorf("got %v; expected %v", result, JsonItem)
		}
		err = es.Delete(client, index,item.Type,item.Id)
		if err != nil {
			t.Fatal(err)
		}
	})

}
func FromJsonObj(o interface{}) (models.Item, error) {
	var target models.Item
	s, err := json.Marshal(o)
	if err != nil {
		return target, err
	}
	err = json.Unmarshal(s, &target)
	return target, err

}
