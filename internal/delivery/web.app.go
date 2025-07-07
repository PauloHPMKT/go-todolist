package delivery

import (
	"log"

	"github.com/PauloHPMKT/go-todolist/internal/delivery/dependencies"
	"github.com/PauloHPMKT/go-todolist/internal/interfaces/handlers"
	"github.com/PauloHPMKT/go-todolist/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start() {
  container := dependencies.Setup()
  router := gin.Default()
  router.Use(middlewares.CorsMiddleware())

  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

  err := container.Invoke(func(taskHandler *handlers.TaskHandler) {
    // Define routes for task operations
    router.POST("/task", taskHandler.CreateTask)
    router.GET("/tasks", taskHandler.ListTasks)
    router.PATCH("/task/:id", taskHandler.UpdateTask)
    router.DELETE("/task/:id", taskHandler.DeleteTask)

    log.Println("Task routes initialized successfully")
    router.Run(":8080")
  })
  if err != nil {
    log.Fatalf("Failed to initialize task routes: %v", err)
  }
}
