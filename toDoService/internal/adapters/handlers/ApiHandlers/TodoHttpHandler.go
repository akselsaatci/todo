package ApiHandlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toDoService/internal/core/dtos/Requests"
	"toDoService/internal/core/ports"
)

type HTTPHandler struct {
	TodoService ports.ToDoService
}

func NewHTTPHandler(todoService ports.ToDoService) *HTTPHandler {
	return &HTTPHandler{
		TodoService: todoService,
	}
}

func (h *HTTPHandler) GetUserTodos(ctx *gin.Context) {

	req := Requests.GetUserTodosRequest{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.AbortWithStatus(400)
		return
	}
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	todos, nerr := h.TodoService.GetTodoByUserId(&req)

	if nerr != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, todos)
}
