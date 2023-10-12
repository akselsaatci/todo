package ApiHandlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"toDoService/internal/core/customErrors"
	"toDoService/internal/core/dtos/Requests"
	"toDoService/internal/core/ports"
)

// TODO Not validating xDDDDDDDD should validate the user

type HTTPHandler struct {
	TodoService ports.ToDoService
}

func NewHTTPHandler(todoService ports.ToDoService) *HTTPHandler {
	return &HTTPHandler{
		TodoService: todoService,
	}
}
func getTokenFromHeader(ctx *gin.Context) (string, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return "", &customErrors.UnauthorizedError{}
	}

	tokenParts := strings.Split(token, "Bearer ")
	if len(tokenParts) != 2 {
		return "", errors.New("invalid token format")
	}

	return tokenParts[1], nil
}

func (h *HTTPHandler) GetUserTodos(ctx *gin.Context) {
	// should get token from auth header

	token, err := getTokenFromHeader(ctx)
	if err != nil {
		ctx.AbortWithError(401, &customErrors.UnauthorizedError{})
		return
	}

	todos, err := h.TodoService.GetTodoByUserId(token)
	if err != nil {
		if errors.Is(err, &customErrors.UnauthorizedError{}) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

func (h *HTTPHandler) AddTodo(ctx *gin.Context) {
	token, err := getTokenFromHeader(ctx)
	if err != nil {
		ctx.AbortWithError(401, &customErrors.UnauthorizedError{})
		return
	}
	data := &Requests.CreateTodoRequest{}

	err = ctx.Bind(data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.TodoService.CreateTodo(data, token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

func (h *HTTPHandler) UpdateTodo(ctx *gin.Context) {
	token, err := getTokenFromHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data := &Requests.UpdateTodoRequest{}
	err = ctx.Bind(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = h.TodoService.UpdateTodo(data, token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// TODO gonna look which way is more useful
	m := make(map[string]int)

	ctx.JSON(http.StatusOK, m)
	return
}

func (h *HTTPHandler) DeleteTodo(ctx *gin.Context) {
	token, err := getTokenFromHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	todoId, isOk := ctx.GetQuery("id")
	if isOk == false {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	err = h.TodoService.DeleteTodo(todoId, token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}
