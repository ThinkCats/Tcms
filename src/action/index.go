package action

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"tcms/src/dao"
	"fmt"
)

//Index action
func Index(c *gin.Context){

	fmt.Println("index...")
	dao.QueryUser()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "hi",
	})
}

//Ping something
func Ping(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
