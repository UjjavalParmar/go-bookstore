package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var (
	db * gorm.DB
)

func Connect(){
	ds := "root:rootpassword@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(ds), &gorm.Config{})
	if err != nil{
		panic(err.Error())
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}
