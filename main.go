package main

import (
	"Reminder_Event_Mail_Scheduler_in_Go/pakages/db"
	"Reminder_Event_Mail_Scheduler_in_Go/pakages/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	route := gin.Default()

	errDB := redis.ConnectRedis()
	if errDB != nil {
		log.Fatal(errDB)
	}

	errRedis := db.ConnectDB()
	if errRedis != nil {
		log.Fatal(errRedis)
	}

	fmt.Println("Reminder Event Mail Scheduler in Go Stared")

	err := route.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
