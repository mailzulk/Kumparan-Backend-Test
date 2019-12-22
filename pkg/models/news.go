package models

import (
	"news-2/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// News - The model we use on this repo
// We only specify Author and Body, because ID and CreatedAt is automatically created
type News struct {
	gorm.Model
	Author string `json:"author"`
	Body   string `json:"body"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&News{})
}

// CreateNews - What to do if the function to create a new news called
func (b *News) CreateNews() *News {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// GetAllNews -  What to do if the function to get all news called
func GetAllNews() []News {
	var News []News
	db.Find(&News)
	return News
}

// GetNewsByID - What to do if the function to get news by ID called
func GetNewsByID(ID int64) (*News, *gorm.DB) {
	var getNews News
	db := db.Where("ID = ?", ID).Find(&getNews)
	return &getNews, db
}

// DeleteNews - WHat to do if the function to delete news by ID called
func DeleteNews(ID int64) News {
	var news News
	db.Where("ID = ?", ID).Delete(news)
	return news
}
