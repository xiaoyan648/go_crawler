package saver

import (
	"github.com/micro/go-micro/util/log"
	"github.com/olivere/elastic"
	"go_crawler/models"
	"go_crawler/util/es"
)


func ItemSaver() chan models.Item{
	//一个save chan 使用一个 Client
	//关闭内网的sniff
	const index  = "crawler"
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		log.Error("ItemSaver error:",err)
	}

	saver := make(chan models.Item)
	go func() {
		itemCount := 0
		for {
			item := <- saver
			//log.Infof("ItemSaver prepare save %v, %d",item,itemCount)
			log.Infof("ItemSave got %v ,count : %d", item, itemCount)
			err := es.Save(client, item, index)
			if err != nil {
				log.Errorf("ItemSaver error : %v ", err)
			}
			itemCount++
		}
	}()
	return saver
}




