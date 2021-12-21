package model

import (
	"log"
)

// TableName 返回一个User的表名
func (user User) TableName() string {
	return "tab_user"
}

// AddUser 添加一个用户
func (user User) AddUser() error {
	err := DBClient.Create(&user).Error
	if err != nil {
		log.Fatalf("Add User error %+v", err)
		return err
	}
	return nil
}

// DeleteUser 删除一个用户
func (user User) DeleteUser() error {
	err := DBClient.Delete(&user).Error
	if err != nil {
		log.Fatalf("Add User error %+v", err)
		return err
	}
	return nil
}
