package Http

import (
	CustomErrors "authService/internal/core/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(ctx *gin.Context, err error) {

	customError, ok := err.(*CustomErrors.CustomError)

	if ok {
		ctx.JSON(customError.HttpStatusCode, gin.H{
			"error": customError.Error(),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
	}
}
