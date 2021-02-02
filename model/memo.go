package model

import (
	"ff/g"
	"time"
)

type Memo struct {
	PubModel
	Content  string    `gorm:"type:varchar(100);comment:'备忘录事项';NOT NULL" json:"content"`
	Deadline time.Time `gorm:"NOT NULL;comment:'截止日期'" json:"deadline"`
}

func (Memo) TableName() string {
	return "memo"
}


func GetMemo(offset int, limit int, maps interface{}) (res []Memo, err error) {
	err = g.DB().Model(&Memo{}).Order("deadline desc").
		Where(maps).Offset(offset).Limit(limit).Find(&res).Error

	return
}

func GetMemoCount(maps interface{}) (count int, err error) {
	err = g.DB().Model(&Memo{}).Where(maps).Count(&count).Error

	return
}

func AddMemo(memo interface{}) (err error) {
	// 防止前端指定id创建
	err = g.DB().Omit("id").Create(memo).Error

	return
}

func EditMemo(id uint, data interface{}) (err error) {
	err = g.DB().Model(&Memo{}).Where("id = ?", id).Omit("id").Updates(data).Error

	return
}

func DeleteMemo(id int) (err error) {
	err = g.DB().Where("id = ?", id).Delete(&Memo{}).Error
	return
}
