package config

import (
	"gorm.io/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "ujjj:password123/simplerest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}
