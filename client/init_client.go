package client

import (
	"blog/model"
	"log"

	"github.com/MonkeyIsMe/MyTool/client"
	"gorm.io/gorm"
)

var MysqlCli *gorm.DB

func InitClient() error {
	var err error
	proxy := client.MysqlProxy{
		"root",
		"421992850",
		"127.0.0.1",
		"3306",
		"myblog",
	}
	MysqlCli, err = client.NewClient(proxy)
	if err != nil {
		log.Fatalf("Init DB Client Error %+v", err)
		return err
	}

	err = model.SetClient(MysqlCli)
	if err != nil {
		log.Fatalf("Set DB Client Error %+v", err)
		return err
	}

	return nil
}
