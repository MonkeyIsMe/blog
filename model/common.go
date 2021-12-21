package model

import (
	"fmt"

	"gorm.io/gorm"
)

var DBClient *gorm.DB

// SetClient 初始化DB的client
func SetClient(input *gorm.DB) error {
	if input == nil {
		err := fmt.Errorf("input client is nil")
		return err
	}
	DBClient = input
	return nil
}

type Blog struct {
	BlogID       int
	BlogName     string
	BlogInfo     string
	BlogAbstract string
	TagName      []string
}

type User struct {
	UserID       int
	UserName     string
	UserAccount  string
	UserPassword string
	UserInfo     string
	UserEmail    string
	UserIcon     string
}

type Tag struct {
	TagID   int    `gorm:"column:tag_id"`
	TagName string `gorm:"column:tag_name"`
	TagInfo string `gorm:"column:tag_info"`
}

type Classify struct {
	ClassifyID   int
	ClassifyName int
	ClassifyInfo string
}

type Comment struct {
	CommentID   int
	CommentInfo string
	CreateTime  string
	UserName    string
	IsValid     int
	BlogID      int
}

type Log struct {
	LogID        int
	LogOperation int
	LogDetail    string
}

type Data struct {
	DataID    int
	CreateDay int
	PageView  int64 // 浏览量
}
