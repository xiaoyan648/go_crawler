package es

import (
	"context"
	"errors"
	"github.com/olivere/elastic"
	"go_crawler/models"
)

//保存item
func Save(client *elastic.Client, item models.Item, index string) (err error){

	if item.Type == "" {
		return errors.New("must supply Type ..")
	}
	indexService := client.Index(). //存储数据，可以添加或者修改，要看id是否存在
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id) //没有的话默认生成
	}
	_, err = indexService.Do(context.Background())
	if err != nil {
		return err;
	}
	return nil;
}
//清空表
func ClearTable(){

}
//delete
func Search(client *elastic.Client,index string, typ string, id string) (*elastic.GetResult, error){
	//从ElasticSearch中获取，根据id
	resp, err := client.Get().
		Index(index).
		Type(typ).
		Id(id).Do(context.Background())

	if err != nil {
		return nil, err;
	}
	return resp, nil;

}
func Delete(client *elastic.Client,index string, typ string, id string) error{
	_, err := client.Delete().
		Index(index).
		Type(typ).
		Id(id).
		Do(context.Background())
	if err != nil {
		return err
	}
	return err
}
