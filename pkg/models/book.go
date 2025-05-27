package models

import (
	"gorm.io/gorm"
	"github.com/ujjavalparmar/go-bookstore/pkg/config"
)

var db * gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Author      string `gorm:"type:varchar(255)" json:"author"`
	Publication string `gorm:"type:varchar(255)" json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
