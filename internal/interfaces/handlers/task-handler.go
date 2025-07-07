package handlers

import (
	"net/http"

	"github.com/PauloHPMKT/go-todolist/internal/entities"
	"github.com/PauloHPMKT/go-todolist/internal/usecases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskHandler struct {
  usecase usecases.TaskUseCase
}

func NewTaskHandler(usecase usecases.TaskUseCase) *TaskHandler {
  return &TaskHandler{usecase: usecase}
}

func (handler *TaskHandler) CreateTask(c *gin.Context) {
  var task entities.Task
  if err := c.ShouldBindJSON(&task); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Invalid request payload",
    })
    return
  }
  id, err := handler.usecase.CreateTask(c.Request.Context(), &task)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": "Failed to create task",
    })
    return
  }
  c.JSON(http.StatusCreated, gin.H{
    "id": id.Hex(),
    "message": "Task created successfully",
  })
}

func (handler *TaskHandler) ListTasks(c *gin.Context) {
  tasks, err := handler.usecase.GetTasks(c.Request.Context())
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": "Failed to retrieve tasks",
    })
    return
  }
  c.JSON(http.StatusOK, tasks)
}

func (handler *TaskHandler) UpdateTask(c *gin.Context) {
  idParam := c.Param("id")
  id, err := primitive.ObjectIDFromHex(idParam)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Invalid task ID",
    })
    return
  }

  var task entities.Task
  if err := c.ShouldBindJSON(&task); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Invalid request payload",
    })
    return
  }

  if err := handler.usecase.UpdateTask(c.Request.Context(), id, &task); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": "Failed to update task",
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "Task updated successfully",
  })
}

func (handler *TaskHandler) DeleteTask(c *gin.Context) {
  idParam := c.Param("id")
  id, err := primitive.ObjectIDFromHex(idParam)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Invalid task ID",
    })
    return
  }

  if err := handler.usecase.DeleteTask(c.Request.Context(), id); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": "Failed to delete task",
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "Task deleted successfully",
  })
}
