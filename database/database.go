package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection(){

	var err error

	const MYSQL = "root:@tcp(127.0.0.1:3306)/db_gofiber?charset=utf8mb4&parseTime=True&loc=Local"
	
	dsn := MYSQL
	
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		panic(err)
	}
	fmt.Println("database connected")
}