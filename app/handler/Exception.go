package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Exception struct {
	status  int
	message string
	param   string
}

func WrapException(eType int, param string) *Exception {
	exception := &Exception{}
	switch eType {
	case 404:
		exception.status = http.StatusNotFound
		exception.message = param + " not found"
		break
	case 401:
		exception.status = http.StatusUnauthorized
		exception.message = "Unauthorized"
		break
	case 403:
		exception.status = http.StatusForbidden
		exception.message = "Forbidden"
		break
	}

	exception.param = param
	return exception
}

func (exception *Exception) ExceptionHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		if exception.status > 0 {
			fmt.Println(exception)
		} else {
			fmt.Println("Exception Status is 0")
		}
		// exception := &exception{}

		// if errorType > 0 {
		// 	context.JSON(exception.wrapException(errorType, param))
		// }
	}
}

func GetMessage(exceptionStatus int, param string) string {
	switch exceptionStatus {
	case 404:
		return param + " not found"
	case 401:
		return "Unauthorized"
	case 403:
		return "Forbidden"
	}
	return "No Exception Message Found"
}

func SetException(exceptionStatus int, param string) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		exception := &Exception{}
		exception.status = exceptionStatus
		exception.param = param
		exception.message = GetMessage(exceptionStatus, param)

		fmt.Println(exception)

		if exception.status > 0 {
			context.AbortWithStatusJSON(exceptionStatus, exception)
		}
	}
}
