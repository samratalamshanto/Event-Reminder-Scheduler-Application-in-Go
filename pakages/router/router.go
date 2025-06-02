package router

import (
	"Reminder_Event_Mail_Scheduler_in_Go/modulles/event"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(route *gin.Engine) {
	event.ConfigEventRoute(route)
}
