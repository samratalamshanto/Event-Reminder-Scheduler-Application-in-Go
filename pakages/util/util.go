package util

import (
	"Reminder_Event_Mail_Scheduler_in_Go/modulles/base"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SuccessResponse(ctx *gin.Context, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, base.CommonResponse{
		Code:    200,
		Message: msg,
		Data:    data,
		MSG:     msg,
	})
}

func CreatedResponse(ctx *gin.Context, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, base.CommonResponse{
		Code:    201,
		Message: msg,
		Data:    data,
		MSG:     msg,
	})
}

func InternalErrorResponse(ctx *gin.Context, err error, msg string) {
	ctx.JSON(http.StatusInternalServerError, base.CommonResponse{
		Code:    500,
		Message: msg,
		MSG:     msg,
		Error:   err.Error(),
		Err:     err.Error(),
	})
}

func ClientErrorResponse(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, base.CommonResponse{
		Code:    400,
		Message: msg,
		MSG:     msg,
	})
}

func IsEmpty(data interface{}) bool {
	switch data.(type) {
	case nil:
		return true
	case string:
		return data == ""
	default:
		return false
	}
}

func StringToInt64(str string) (int64, error) {
	i, err := StringToInt(str)
	return int64(i), err
}

func StringToInt(str string) (int, error) {
	i, err := strconv.Atoi(str)
	return i, err
}
