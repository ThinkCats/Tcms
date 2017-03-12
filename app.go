package main

import (
	"tcms/src/action"
	"tcms/src/middleware"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	mains()
}

func mains() {

	r := gin.Default()
	r.LoadHTMLGlob("template/*")

	r.GET("/", action.Index)

	rest := r.Group("/rest")
	rest.Use(middleware.CheckRestAuth)
	{
		rest.GET("/ping", action.Ping)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
