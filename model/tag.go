package model

import (
	"log"

	"github.com/MonkeyIsMe/MyTool/constant"
)

// TableName 返回一个User的表名
func (tag Tag) TableName() string {
	return "t_tag"
}

// InsertTag 插入一个分类标签
func (tag Tag) AddTag() error {
	err := DBClient.Create(&tag).Error
	if err != nil {
		log.Fatalf("Insert Tag Err: [%+v]", err)
		return err
	}
	return nil
}

// DeleteTag 删除一个分类标签
func (tag Tag) DeleteTag() error {
	err := DBClient.Delete(&tag).Error
	if err != nil {
		log.Fatalf("Delete Tag Err: [%+v]", err)
		return err
	}
	return nil
}

// QueryTagPageSize 分页查询tag
func QueryTagPageSize(page int) ([]Tag, error) {
	var tagList []Tag
	err := DBClient.Offset((page - 1) * constant.PageSize).Limit(constant.PageSize).Find(&tagList).Error
	if err != nil {
		log.Fatalf("Query Tag PageSize Err: [%+v]", err)
		return nil, err
	}
	return tagList, nil
}
