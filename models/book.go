package models

import (
	"errors"
	"mygin.web/database"
)

type Book struct {
	ID     int    `json:"id"`     //图书编号
	Name   string `json:"name"`   //书名
	Author string `json:"author"` //作者
	Status bool   `json:"status"` //是否借出
}

var Books []Book

// GetAll 获取数据库中的所有图书
func (book *Book) GetAll() (books []Book, err error) {
	if err = database.DbBook.Find(&books).Error; err != nil {
		return
	}
	return
}

func (book *Book) GetBook(id int64) (b Book, err error) {
	if err = database.DbBook.First(&b, id).Error; err != nil {
		return
	}
	return
}

// Add 添加图书
func (book Book) Add() (id int, err error) {
	result := database.DbBook.Create(&book)
	id = book.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// Update 借书
func (book *Book) Update(id int64, status bool) (err error) {
	if err = database.DbBook.First(&book, id).Error; err != nil {
		return
	}
	if !status {
		if book.Status {
			if err = database.DbBook.Model(&book).Update("status", status).Error; err != nil {
				return
			}
		} else {
			err = errors.New("图书已被借走")
		}
	} else {
		if !book.Status {
			if err = database.DbBook.Model(&book).Update("status", status).Error; err != nil {
				return
			}
		} else {
			err = errors.New("图书已存在，请检查要归还图书的id")
		}
	}
	return
}
