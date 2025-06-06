package router

import (
	"Reminder_Event_Mail_Scheduler_in_Go/modulles/event"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigRoutes(route *gin.Engine) {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	//route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	event.ConfigEventRoute(route)
}
