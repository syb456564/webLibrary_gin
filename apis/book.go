package apis

import (
	"github.com/gin-gonic/gin"
	"mygin.web/models"
	"net/http"
	"strconv"
)

func ListBooksApi(c *gin.Context) {
	var book models.Book
	var listBook = make(map[string]int)
	result, err := book.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "抱歉未找到相关信息",
		})
		return
	}
	for i := range result {
		_, ok := listBook[result[i].Name]
		if ok {
			if result[i].Status {
				listBook[result[i].Name]++
			}
		} else {
			if result[i].Status {
				listBook[result[i].Name] = 1
			} else {
				listBook[result[i].Name] = 0
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data": listBook,
	})
}

// AddBookApi 添加图书
func AddBookApi(c *gin.Context) {
	var book models.Book
	book.Name = c.Request.FormValue("name")
	book.Author = c.Request.FormValue("author")
	book.Status, _ = strconv.ParseBool(c.Request.FormValue("status"))
	id, err := book.Add()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "添加失败",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "添加成功",
		"data":    id,
	})
}

// UpdateBookStatusApi 归还图书
func UpdateBookStatusApi(c *gin.Context) {
	var book models.Book
	id, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	status, _ := strconv.ParseBool(c.Param("status"))
	err := book.Update(id, status)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
	})
}
func BorrowList(c *gin.Context) {

}
