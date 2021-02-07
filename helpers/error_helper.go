package helpers

import (
	"omdbapi/config"
	"omdbapi/constants"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type ErrorHelper struct {
}

func ErrorHelperHandler() ErrorHelper {
	return ErrorHelper{}
}

func (handler *ErrorHelper) HTTPResponseError(context *gin.Context, e error, defaultErrorCode int) {

	logger, _ := config.NewLogger()
	switch e.Error() {
	case "record not found":
		defaultErrorCode = constants.ResourceNotFound
		break
	}

	errorConstant := constants.GetErrorConstant(defaultErrorCode)
	context.JSON(errorConstant.HttpCode, gin.H{
		"code":    defaultErrorCode,
		"message": e.Error(),
	})

	if _, ok := e.(*mysql.MySQLError); !ok {
		logger.Error(e.Error())
	}

	panic(e)
}
