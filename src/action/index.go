package action

import (
	"fmt"
	"net/http"
	"tcms/src/dao"

	"gopkg.in/gin-gonic/gin.v1"
)

//Index action
func Index(c *gin.Context) {
	user, err := dao.QueryUser("wang")
	if err != nil {
		println(err)
	}
	fmt.Println("get user password:", user.Password)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "hi",
	})
}

//Login login view
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", nil)
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
