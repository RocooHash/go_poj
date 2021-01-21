package router

import (
	"github.com/gin-gonic/gin"
	"go_poj/router/middle"
	"go_poj/service"
	"net/http"
)

func InitRouter(g *gin.Engine) {
	middlewares := []gin.HandlerFunc{}
	g.Use(gin.Recovery())
	g.Use(middle.NoCache)
	g.Use(middle.Options)
	g.Use(middle.Secure)
	g.Use(middlewares...)
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	router := g.Group("/user")
	{
		router.POST("/addUser", service.AddUser)       //添加用户
		router.POST("/selectUser", service.SelectUser) //查询用户
	}
}
