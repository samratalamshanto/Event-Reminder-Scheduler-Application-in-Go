package event

import "github.com/gin-gonic/gin"

func ConfigEventRoute(route *gin.Engine) {

	route.Group("/api/v1/event")
	{
		route.GET("/", GetAll)
		route.GET("/{id}", GetById)
		route.POST("/", Create)
		route.PUT("/{id}", Update)
		route.DELETE("/{id}", Delete)
	}
}
