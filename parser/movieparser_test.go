package parser

import (
	"go_crawler/models"
	"io/ioutil"
	"testing"
)

func TestParseMovieInfo(t *testing.T) {
		contents,err :=ioutil.ReadFile("movie_data.html")
		if err != nil{
			panic(err)
		}

		result := ParseMovieInfo(string(contents),"https://movie.douban.com/subject/1292722/")
		actual := result.Item
		//actual.Source = actual.Source.(*models.MovieInfo)
		actual.Source.(*models.MovieInfo).Create_time = ""
		expected := models.Item{
			Url:    "https://movie.douban.com/subject/1292722/",
			Type:   "movie",
			Id:     "1292722",
			Source: &models.MovieInfo {
				Id : 0,
				Movie_id : 0,
				Movie_name : "泰坦尼克号 Titanic",
				Movie_pic : "",
				Movie_director : "詹姆斯·卡梅隆",
				Movie_writer : "",
				Movie_country : "",
				Movie_language : "",
				Movie_main_character : "莱昂纳多·迪卡普里奥/凯特·温丝莱特/比利·赞恩/凯西·贝茨/弗兰西丝·费舍/格劳瑞亚·斯图尔特/比尔·帕克斯顿/伯纳德·希尔/大卫·沃纳/维克多·加博/乔纳森·海德/苏茜·爱米斯/刘易斯·阿伯内西/尼古拉斯·卡斯柯恩/阿那托利·萨加洛维奇/丹尼·努齐/杰森·贝瑞/伊万·斯图尔特/艾恩·格拉法德/乔纳森·菲利普斯/马克·林赛·查普曼/理查德·格拉翰/保罗·布赖特威尔/艾瑞克·布里登/夏洛特·查顿/博纳德·福克斯/迈克尔·英塞恩/法妮·布雷特/马丁·贾维斯/罗莎琳·艾尔斯/罗切尔·罗斯/乔纳森·伊万斯-琼斯/西蒙·克雷恩/爱德华德·弗莱彻/斯科特·安德森/马丁·伊斯特/克雷格·凯利/格雷戈里·库克/利亚姆·图伊/詹姆斯·兰开斯特/艾尔莎·瑞雯/卢·帕尔特/泰瑞·佛瑞斯塔/凯文·德·拉·诺伊/",
				Movie_type : "剧情/爱情/灾难",
				Movie_on_time : "1998-04-03(中国大陆)",
				Movie_span : "194分钟",
				Movie_grade : "9.4",
				Create_time : "",
		},
		}
		//没有map chan slice可以直接比较
		if *actual.Source.(*models.MovieInfo) != *expected.Source.(*models.MovieInfo){
			t.Errorf("expected %v , but was %v \n",expected,actual)
		}

}