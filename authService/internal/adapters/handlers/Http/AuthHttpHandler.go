package Http

import (
	Requests "authService/internal/adapters/handlers/Http/Requests"
	"authService/internal/core/dtos"
	CustomErrors "authService/internal/core/errors"
	"authService/internal/core/services"
	"errors"
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
		ctx.AbortWithStatus(400)
		return
	}
	token, err := h.authService.AuthorizeUser(req.Username, req.Password)

	if errors.Is(err, &CustomErrors.UserNotFoundError{}) {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, token)
}
func (h *HTTPHandler) ValidateToken(ctx *gin.Context) {
	req := Requests.ValidateTokenRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	token, err := h.authService.ValidateToken(req.Token)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, token)
}

func (h *HTTPHandler) RegisterUser(ctx *gin.Context) {
	req := dtos.AddUserRequest{}
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	token, err := h.authService.RegisterUser(req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, token)
}
