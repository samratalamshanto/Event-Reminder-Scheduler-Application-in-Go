package event

import (
	"Reminder_Event_Mail_Scheduler_in_Go/pakages/util"
	"github.com/gin-gonic/gin"
	"log"
)

func GetAll(ctx *gin.Context) {
	res, err := GetAllEvents()
	if err != nil {
		log.Println(err)
		util.InternalErrorResponse(ctx, err, "Failed to GetAll Events")
	}
	util.SuccessResponse(ctx, res, "Success")
}

func GetById(ctx *gin.Context) {
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

func Create(ctx *gin.Context) {
	var dto EventDto
	err := ctx.ShouldBind(&dto)
	if err != nil {
		log.Println(err)
		util.ClientErrorResponse(ctx, "Got EventDto Parsing Error")
		return
	}

	res, err := CreateEvent(dto)
	if err != nil {
		log.Println(err)
		util.InternalErrorResponse(ctx, err, "Failed to save EventDto")
		return
	}
	util.CreatedResponse(ctx, &res, "Successfully Created Event")
}

func Update(ctx *gin.Context) {
	var dto EventDto
	err := ctx.ShouldBind(&dto)
	if err != nil {
		log.Println(err)
		util.ClientErrorResponse(ctx, "Got EventDto Parsing Error")
		return
	}

	eventIdParam := ctx.Param("id")
	if util.IsEmpty(eventIdParam) {
		util.ClientErrorResponse(ctx, "EventId is required")
		return
	}

	id, err := util.StringToInt64(eventIdParam)
	if err != nil {
		log.Println(err)
		util.ClientErrorResponse(ctx, err.Error())
	}

	res, err := UpdateEvent(dto, id)
	if err != nil {
		log.Println(err)
		util.InternalErrorResponse(ctx, err, "Failed to update EventDto")
		return
	}
	util.SuccessResponse(ctx, &res, "Successfully Updated Event")
}

func Delete(ctx *gin.Context) {
	eventIdParam := ctx.Param("id")
	if util.IsEmpty(eventIdParam) {
		util.ClientErrorResponse(ctx, "EventId is required")
		return
	}

	id, err := util.StringToInt64(eventIdParam)
	if err != nil {
		log.Println(err)
		util.ClientErrorResponse(ctx, err.Error())
	}

	err = DeleteEvent(id)
	if err != nil {
		log.Println(err)
		util.InternalErrorResponse(ctx, err, "Failed to delete Event")
		return
	}
	util.SuccessResponse(ctx, nil, "Successfully deleted Event")
}
