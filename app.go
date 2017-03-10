package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"tcms/src/action"
)

func main(){
	mains()
}

func mains() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.GET("/", action.Index)
	r.GET("/ping", action.Ping)
	r.Run() // listen and serve on 0.0.0.0:8080
}
