package event

import (
	"Reminder_Event_Mail_Scheduler_in_Go/pakages/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GetEvents(ctx *gin.Context) {
	res, err := GetAllEvents()
	if err != nil {
		log.Println(err)
		util.InternalErrorResponse(ctx, err, "Failed to GetAll Events")
	}
	util.SuccessResponse(ctx, res, "Success")
}

func GetEventById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if util.IsEmpty(idParam) {
		util.ClientErrorResponse(ctx, "id is required")
		return
	}
	id, err := util.StringToInt64(idParam)
	if err != nil {
		util.ClientErrorResponse(ctx, err.Error())
		return
	}

	res, err := GetEvent(id)
	if err != nil {
		log.Println(err)
		util.InternalErrorResponse(ctx, err, "Failed to GetAll Events")
	}
	util.SuccessResponse(ctx, res, "Success")
}
