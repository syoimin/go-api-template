package models

import (
	"time"
)

// モデル
type Sample struct {
	Id        uint      `json:"id" jaFieldName:"Id" gorm:"primary_key"`
	Name      string    `json:"name" jaFieldName:"氏名" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
