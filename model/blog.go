package model

import (
	"log"

	"github.com/MonkeyIsMe/MyTool/constant"
)

// AddBlog 添加一个一篇博客
func (blog Blog) AddBlog() error {
	err := DBClient.Create(&blog).Error
	if err != nil {
		log.Fatalf("Add Blog Err: [%+v]", err)
		return err
	}
	return nil
}

// DeleteUser 删除一个博客
func (blog Blog) DeleteBlog() error {
	err := DBClient.Delete(&blog).Error
	if err != nil {
		log.Fatalf("Delete Blog Err: [%+v]", err)
		return err
	}
	return nil
}

// QueryBlogByName 根据博客名查询博客
func (blog Blog) QueryBlogByName(name string) error {
	err := DBClient.Where("blog_name = ?", name).First(&blog).Error
	if err != nil {
		log.Fatalf("Query Blog By BlogName Err: [%+v]", err)
		return err
	}
	return nil
}

// QueryBlogById 根据ID查询博客
func (blog Blog) QueryBlogByID(id int) error {
	err := DBClient.First(&blog, id).Error
	if err != nil {
		log.Fatalf("Query Blog By ID Err [%+v]", err)
		return err
	}
	return nil
}

// QueryBlogPageSize 分页查询博客
func QueryBlogPageSize(page int) ([]Blog, error) {
	var blogList []Blog
	err := DBClient.Offset((page - 1) * constant.PageSize).Limit(constant.PageSize).Find(&blogList).Error
	if err != nil {
		log.Fatalf("Query Blog List Err: [%+v]", err)
		return nil, err
	}
	return blogList, nil
}
