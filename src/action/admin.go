package action

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

//AdminHome admin home page
func AdminHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "admin page",
	})
}
