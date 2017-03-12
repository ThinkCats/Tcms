package action

import (
	"fmt"
	"net/http"
	"tcms/src/dao"

	"gopkg.in/gin-gonic/gin.v1"
)

//Index action
func Index(c *gin.Context) {

	fmt.Println("index...")
	dao.QueryUser()
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "hi",
	})
}

//Ping something
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

//NotFound 404 not found
func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.tmpl", gin.H{
		"title": "Not Found",
	})
}

//AuthFail auth error
func AuthFail(c *gin.Context) {
	c.HTML(http.StatusNotFound, "401.tmpl", gin.H{
		"title": "Auth Error",
	})
}
