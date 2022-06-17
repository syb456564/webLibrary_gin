package main

import (
	_ "github.com/go-sql-driver/mysql"
	"mygin.web/database"
	"mygin.web/router"
)

func main() {
	defer database.DbBook.Close()
	r := router.InitRouter()
	r.Run(":8080")
}
