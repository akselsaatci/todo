package Http

import (
	"authService/internal/adapters/handlers"
	Requests "authService/internal/adapters/handlers/Http/Requests"
	"authService/internal/core/dtos"
	CustomErrors "authService/internal/core/errors"
	"authService/internal/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	authService services.AuthService
}

func NewHTTPHandler(authService services.AuthService) *HTTPHandler {
	return &HTTPHandler{
		authService: authService,
	}
}

func (h *HTTPHandler) AuthorizeUser(ctx *gin.Context) {

	req := Requests.AuthorizeUserRequest{}

	err := ctx.ShouldBind(&req)
	if err != nil {
		customErr := &CustomErrors.CustomError{
			Message:        err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		}
		HandleError(ctx, customErr)
		return
	}

	token, err := h.authService.AuthorizeUser(req.Username, req.Password)

	if err != nil {
		HandleError(ctx, err)
		return
	}

	response := handlers.BuildResponse(http.StatusOK, token)
	ctx.JSON(http.StatusOK, response)
}
func (h *HTTPHandler) ValidateToken(ctx *gin.Context) {
	req := Requests.ValidateTokenRequest{}
	err := ctx.Bind(&req)

	if err != nil {
		customErr := &CustomErrors.CustomError{
			Message:        err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		}
		HandleError(ctx, customErr)
		return
	}

	token, err := h.authService.ValidateToken(req.Token)

	if err != nil {
		HandleError(ctx, err)
		return
	}

	response := handlers.BuildResponse(http.StatusOK, token)
	ctx.JSON(http.StatusOK, response)
}

func (h *HTTPHandler) RegisterUser(ctx *gin.Context) {
	req := dtos.AddUserRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		customErr := &CustomErrors.CustomError{
			Message:        err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		}
		HandleError(ctx, customErr)
		return
	}

	token, err := h.authService.RegisterUser(req)

	if err != nil {
		//Shouldnt be error here log this
		HandleError(ctx, err)
		return
	}

	response := handlers.BuildResponse(http.StatusOK, token)
	ctx.JSON(http.StatusOK, response)
}

func (h *HTTPHandler) GetPublicKey(ctx *gin.Context) {
	//Should be able to call only internally i suppose

	return

}
