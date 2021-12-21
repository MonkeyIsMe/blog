package model

import (
	"log"

	"github.com/MonkeyIsMe/MyTool/constant"
)

// TableName 返回一个User的表名
func (log Log) TableName() string {
	return "t_log"
}

// AddLog 添加一条日志
func (l Log) AddLog() error {
	err := DBClient.Create(&l).Error
	if err != nil {
		log.Fatalf("Add Blog Err: [%+v]", err)
		return err
	}
	return nil
}

// DeleteLog 删除一条日志
func (l Log) DeleteLog() error {
	err := DBClient.Create(&l).Error
	if err != nil {
		log.Fatalf("Add Blog Err: [%+v]", err)
		return err
	}
	return nil
}

// QueryLogPageSize 分页查询日志
func QueryLogPageSize(page int) ([]Log, error) {
	var logList []Log
	err := DBClient.Offset((page - 1) * constant.PageSize).Limit(constant.PageSize).Find(&logList).Error
	if err != nil {
		log.Fatalf("Query Tag PageSize Err: [%+v]", err)
		return nil, err
	}
	return logList, nil
}
