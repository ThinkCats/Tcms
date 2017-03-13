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

	//admin auth check
	authMiddleWare := middleware.CheckAdminAuth()
	r.POST("/login", authMiddleWare.LoginHandler)
	admin := r.Group("/admin")
	admin.Use(authMiddleWare.MiddlewareFunc())
	{
		admin.GET("/", action.AdminHome)
	}

	//rest auth check
	rest := r.Group("/rest")
	rest.Use(middleware.CheckRestAuth)
	{
		rest.GET("/ping", action.Ping)
	}

	r.Run(":4000")
}
