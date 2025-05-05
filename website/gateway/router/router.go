package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"html/template"
)

// marshal 函数将任务数据转为 JS 变量
func marshal(v interface{}) template.JS {
	a, _ := json.Marshal(v)
	return template.JS(a)
}

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()

	ginRouter.SetFuncMap(template.FuncMap{
		"marshal": marshal,
	})

	//ginRouter.LoadHTMLGlob("templates/*")
	//ginRouter.Static("/static", "./static")

	v1 := ginRouter.Group("/")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "Request Success!")
		})
	}

	return ginRouter
}
