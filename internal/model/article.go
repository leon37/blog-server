package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`           // 文章标题
	Desc          string `json:"desc"`            // 文章简述
	CoverImageUrl string `json:"cover_image_url"` // 封面图片地址
	Content       string `json:"content"`         // 文章内容
	State         uint8  `json:"state"`           // 状态：0为禁用、1为启用
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Get(db *gorm.DB) (*Article, error) {
	article := &Article{}
	err := db.Where("id = ? AND is_del = ?", a.ID, 0).Find(article).Error
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a Article) Count(db *gorm.DB) (int, error) {
	var count int
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)
	err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", a.State)
	if err = db.Where("is_del = ?", 0).Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	err := db.Model(a).Where("id = ? AND is_del = ?", a.ID, 0).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", a.ID, 0).Delete(&a).Error
}
