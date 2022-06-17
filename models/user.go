package models

import (
	"mygin.web/database"
	"time"
)

type User struct {
	ID           int
	Name         string
	BorrowRecord map[string]time.Time
	Count        int
}

func (user User) Add() (id int64, err error) {
	result := database.DbBook.Create(&user)
	id = int64(user.ID)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//func (user *User) Update(id int64) (err error) {
//	if err = database.DbBook.First(&book, id).Error; err != nil {
//		return
//	}
//	if book.Status {
//		if err = database.DbBook.Model(&book).Update("status", false).Error; err != nil {
//			return
//		}
//	} else {
//		err = errors.New("图书已被借走")
//	}
//	return
//}
