package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var DbBook *gorm.DB

func init() {
	var err error
	DbBook, err = gorm.Open("mysql", "root:Sun456564+0.0@tcp(127.0.0.1:3306)/library_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
		log.Fatal(err.Error())
	}
	if DbBook.Error != nil {
		fmt.Printf("database error %v", DbBook.Error)
	}
}
