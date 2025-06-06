package main

import (
	_ "Reminder_Event_Mail_Scheduler_in_Go/docs" // 👈 Required for Swagger
	"Reminder_Event_Mail_Scheduler_in_Go/pakages/corn_jobs"
	"Reminder_Event_Mail_Scheduler_in_Go/pakages/db"
	"Reminder_Event_Mail_Scheduler_in_Go/pakages/redis"
	"Reminder_Event_Mail_Scheduler_in_Go/pakages/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// @title Reminder Event Mail Scheduler API
// @version 1.0
// @description Sends reminder emails using Go
// @host localhost:8080
// @BasePath /api/v1
func main() {
	errDB := redis.ConnectRedis()
	if errDB != nil {
		log.Fatal(errDB)
	}

	errRedis := db.ConnectDB()
	if errRedis != nil {
		log.Fatal(errRedis)
	}

	route := gin.Default()

	router.ConfigRoutes(route)

	corn_jobs.Config()

	fmt.Println("Reminder Event Mail Scheduler in Go Stared")

	err := route.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
