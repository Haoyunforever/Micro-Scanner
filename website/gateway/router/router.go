package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()

	ginRouter.LoadHTMLGlob("templates/*")
	ginRouter.Static("/static", "./static")

}
