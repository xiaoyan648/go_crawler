package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"regexp"
	"strings"
	"time"
)
type MovieInfo struct {
	Id int64
	Movie_id int64  //用url   "url": "/subject/26985839/", 26985839id
	Movie_name string
	Movie_pic string
	Movie_director string
	Movie_writer string
	Movie_country string
	Movie_language string
	Movie_main_character string
	Movie_type string
	Movie_on_time string
	Movie_span string
	Movie_grade string
	Create_time string
}

type User struct {
	Id int
	Username string
	Password string
}
type Article struct {
	A_Id int	`orm:"pk;auto"`	//默认id为主键，字段名不为id则手动加上注解
	Title string	`orm:"size(20)"`
	Content string	`orm:"size(500)"`
	Img string	`orm:"size(50);null"` //图片可以不上传，为null
	Time time.Time	`orm:"type(datetime);auto_now_add"`	//因为time包里就一个Time类型，但数据库不止一个，所以指定
														//auto_now 是指最新更新的时间，auto_now_add是指插入的时间
	Count int	`orm:"default(0)"`
}
func init() {
	//orm.Debug = true
	orm.RegisterDataBase("default","mysql","root:root@tcp(127.0.0.1:3306)/crawl?charset=utf8")
	orm.RegisterModel(new(User),new(Article),new(MovieInfo))
	orm.RunSyncdb("default",false,true)
}
//导演名称
func GetMovieDirector(movieHtml string) string{
	if movieHtml == ""{
		return ""
	}
	reg:=regexp.MustCompile(`<a.*?rel="v:directedBy">(.*?)</a>`)
	result:=reg.FindAllStringSubmatch(movieHtml,-1)

	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])

}
//电影名称
func GetMovieName(movieHtml string)string{
	if movieHtml == ""{
		return ""
	}

	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

//主演
func GetMovieMainCharacters(movieHtml string)string{
	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	mainCharacters := ""
	for _,v := range result{
		mainCharacters += v[1] + "/"
	}
	if len(result) == 0 {
		return ""
	}
	return mainCharacters
}
//电影评分
func GetMovieGrade(movieHtml string)string{
	reg := regexp.MustCompile(`<strong.*?property="v:average">(.*?)</strong>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}
//电影分类
func GetMovieGenre(movieHtml string)string{
	reg := regexp.MustCompile(`<span.*?property="v:genre">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	movieGenre := ""
	for _,v := range result{
		movieGenre += v[1] + "/"
	}
	return strings.Trim(movieGenre,"/")
}

//上映时间
func GetMovieOnTime(movieHtml string) string{
	reg := regexp.MustCompile(`<span.*?property="v:initialReleaseDate".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

//电影时长
func GetMovieRunningTime(movieHtml string) string{
	reg := regexp.MustCompile(`<span.*?property="v:runtime".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

func GetMovieInfo(movieHtml string) *MovieInfo {
	var movieInfo MovieInfo
	movieInfo.Movie_name = GetMovieName(movieHtml)
	movieInfo.Movie_director = GetMovieDirector(movieHtml)
	movieInfo.Movie_main_character = GetMovieMainCharacters(movieHtml)
	movieInfo.Movie_type = GetMovieGenre(movieHtml)
	movieInfo.Movie_on_time = GetMovieOnTime(movieHtml) // 上映时间：2016-09-14(中国大陆)
	movieInfo.Movie_grade = GetMovieGrade(movieHtml)
	movieInfo.Movie_span = GetMovieRunningTime(movieHtml)
	//movieInfo.Movie_on_time = movieInfo.Movie_on_time[0:strings.Index(models.GetMovieOnTime(movieHtml), "(")] //上映时间：2016-09-14
	movieInfo.Create_time = time.Now().Format("2006-1-2 15:04:05")
	return &movieInfo
}

//获取相关电影urls
func GetMovieUrls(movieHtml string) []string{
	//reg := regexp.MustCompile(`<a.*?href="(https://movie.douban.com/.*?)"`)
	reg := regexp.MustCompile(`<a href="(https://movie.douban.com/subject/.*?/\?from=subject-page)"`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	movieSets := make([]string,0,len(result))
	for i := range result {
		movieSets = append(movieSets, result[i][1])
	}
	return movieSets
}
//爬取热门电影
//


//DB option
//添加到数据库中
func AddMovieToDB(info interface{}) (int64, error){
	o := orm.NewOrm()
	movieInfo, ok := info.(*MovieInfo)
	if !ok {
		return -1, errors.New("dataType need models.MovieInfo")
	}
	n, err := o.Insert(movieInfo)
	if err != nil {
		return -1, err
	}
	return n, nil
}
