// @title              TodoList API
// @version            1.0
// @description        This is a simple Todo List API using Go and MongoDB.
// @host              localhost:8080
package main

import (
	_ "github.com/PauloHPMKT/go-todolist/docs"

	"github.com/PauloHPMKT/go-todolist/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary           Create a new task
// @Description       Create a new task with title, description, and due date
// @Tags              tasks
// @Accept            json
// @Produce           json
// @Param             task     body      entities.Task    true    "Task object"
// @Success           201      {object}  map[string]interface{}
// @Failure           400      {object}  map[string]string "Bad Request"
// @Failure           500      {object}  map[string]string "Internal Server Error"
// @Router            /task    [post]

// @Summary           List all tasks
// @Description       Retrieve a list of all tasks
// @Tags              tasks
// @Produce           json
// @Success           200      {array}   entities.Task
// @Failure           500      {object}  map[string]string "Internal Server Error"
// @Router            /tasks   [get]

// @Summary           Update a task
// @Description       Update an existing task by ID
// @Tags              tasks
// @Accept            json
// @Produce           json
// @Param             id       path      string    true    "Task ID"
// @Param             task     body      entities.Task  true    "Task object"
// @Success           200      {object}  map[string]string "Task updated successfully"
// @Failure           400      {object}  map[string]string "Bad Request"
// @Failure           404      {object}  map[string]string "Task not found"
// @Failure           500      {object}  map[string]string "Internal Server Error"
// @Router            /task/{id} [patch]

// @Summary           Delete a task
// @Description       Delete an existing task by ID
// @Tags              tasks
// @Param             id       path      string    true    "Task ID"
// @Success           200      {object}  map[string]string "Task deleted successfully"
// @Failure           400      {object}  map[string]string "Bad Request"
// @Failure           404      {object}  map[string]string "Task not found"
// @Failure           500      {object}  map[string]string "Internal Server Error"
// @Router            /task/{id} [delete]


func main() {
  router := gin.Default()
  router.Use(middlewares.CorsMiddleware())

  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

  router.POST("/task", )
  router.GET("/tasks", )
  router.PATCH("/task/:id", )
  router.DELETE("/task/:id", )

  router.Run(":8080")
}
