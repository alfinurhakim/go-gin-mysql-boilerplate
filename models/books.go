package models

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	ID              int64          `json:"id" gorm:"primary_key;auto_increment;not_null"` //big integer
	BookNumber      string         `json:"book_number" gorm:"size:255"`
	BookTitle       string         `json:"book_title" gorm:"size:255"`
	Author          string         `json:"author" gorm:"size:255"`
	PublicationYear int64          `json:"publication_year" gorm:"size:40;"`
	Publisher       string         `json:"publisher" gorm:"type:text"`
	CreatedBy       int64          `json:"created_by" gorm:"size:20"`
	UpdatedBy       int64          `json:"updated_by" gorm:"size:20"`
	CreatedAt       time.Time      `json:"created_at"` //timestamps
	UpdatedAt       time.Time      `json:"updated_at"` //timestamps
	DeletedAt       gorm.DeletedAt `json:"deleted_at"` //timestamps
}

func (Books) TableName() string {
	return "books"
}
