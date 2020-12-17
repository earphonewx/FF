package model

import (
	"ff/g"
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	Title string `gorm:"type:varchar(20);comment:'书签标题'" json:"title"`
	Url   string `gorm:"type:varchar(100);comment:'书签url';NOT NULL;UNIQUE" json:"url"`
	Type  string `gorm:"type:varchar(20)" json:"type"`
}

func (Bookmark) TableName() string {
	return "bookmark"
}


func GetBookmark(offset int, limit int, maps interface{}) (res []Bookmark, err error) {
	err = g.DB.Model(&Bookmark{}).Order("id").
		Where(maps).Offset(offset).Limit(limit).Find(&res).Error
	return
}

func GetBookmarkCount(maps interface{}) (count int, err error) {
	err = g.DB.Model(&Bookmark{}).Where(maps).Count(&count).Error
	return
}

func AddBookmark(bookmark interface{}) (err error) {
	// 防止前端指定id创建
	err = g.DB.Omit("id").Create(bookmark).Error
	return
}

func EditBookmark(id uint, data interface{}) (err error) {
	err = g.DB.Model(&Bookmark{}).Where("id = ?", id).Omit("id").Updates(data).Error
	return
}

func DeleteBookmark(id int) (err error) {
	err = g.DB.Where("id = ?", id).Delete(&Bookmark{}).Error
	return
}
