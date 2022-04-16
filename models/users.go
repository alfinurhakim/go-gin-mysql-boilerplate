package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        int64          `json:"id" gorm:"primary_key;auto_increment;not_null"` //big integer
	Name      string         `json:"name" gorm:"size:255"`
	UserName  string         `json:"user_name" gorm:"size:45"`
	Password  string         `json:"password" gorm:"type:text"`
	Token     string         `json:"token" gorm:"type:text"`
	CreatedBy int64          `json:"created_by" gorm:"size:20"`
	UpdatedBy int64          `json:"updated_by" gorm:"size:20"`
	CreatedAt time.Time      `json:"created_at"` //timestamps
	UpdatedAt time.Time      `json:"updated_at"` //timestamps
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //timestamps
}

func (Users) TableName() string {
	return "users"
}
