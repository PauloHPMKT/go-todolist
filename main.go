// @title              TodoList API
// @version            1.0
// @description        This is a simple Todo List API using Go and MongoDB.
// @host              localhost:8080
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	_ "github.com/PauloHPMKT/go-todolist/docs"

	"github.com/PauloHPMKT/go-todolist/middlewares"
	"github.com/PauloHPMKT/go-todolist/schemas"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func connectDB() {
  clientOptions := options.Client().ApplyURI("mongodb://localhost:27018")
  client, _ := mongo.Connect(context.TODO(), clientOptions)
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  err := client.Ping(ctx, nil)
  if err != nil {
    panic(err)
  }
  collection = client.Database("go-todolist").Collection("tasks")
  log.Println("Connected to MongoDB!")
}

// @Summary           Create a new task
// @Description       Create a new task with title, description, and due date
// @Tags              tasks
// @Accept            json
// @Produce           json
// @Param             task     body      schemas.Task    true    "Task object"
// @Success           201      {object}  map[string]interface{}
// @Failure           400      {object}  map[string]string "Bad Request"
// @Failure           500      {object}  map[string]string "Internal Server Error"
// @Router            /task    [post]
func createTask(c *gin.Context) {
  var task schemas.Task
  if err := c.ShouldBindJSON(&task); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  task.CreatedAt = time.Now()
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  result, err := collection.InsertOne(ctx, task)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
    return
  }
  c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

// @Summary           List all tasks
// @Description       Retrieve a list of all tasks
// @Tags              tasks
// @Produce           json
// @Success           200      {array}   schemas.Task
// @Failure           500      {object}  map[string]string "Internal Server Error"
// @Router            /tasks   [get]
func listTasks(c *gin.Context) {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  cursor, err := collection.Find(ctx, bson.M{})
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
    return
  }
  defer cursor.Close(ctx)

  var tasks []schemas.Task
  if err := cursor.All(ctx, &tasks); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, tasks)
}

// @Summary           Update a task
// @Description       Update an existing task by ID
// @Tags              tasks
// @Accept            json
// @Produce           json
// @Param             id       path      string    true    "Task ID"
// @Param             task     body      schemas.Task  true    "Task object"
// @Success           200      {object}  map[string]string "Task updated successfully"
// @Failure           400      {object}  map[string]string "Bad Request"
// @Failure           404      {object}  map[string]string "Task not found"
// @Failure           500      {object}  map[string]string "Internal Server Error"
// @Router            /task/{id} [patch]
func updateTaks(c *gin.Context) {
  isParams := c.Params.ByName("id")
  objId, err := primitive.ObjectIDFromHex(isParams)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Invalid task ID",
    })
    return
  }

  var task map[string]interface{}
  if err := c.ShouldBindJSON(&task); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  result, err := collection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{
    "$set": task,
  })
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
    return
  }

  if result.MatchedCount == 0 {
    c.JSON(http.StatusNotFound, gin.H{
      "error": "Task not found",
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "Task updated successfully",
  })
}

// @Summary           Delete a task
// @Description       Delete an existing task by ID
// @Tags              tasks
// @Param             id       path      string    true    "Task ID"
// @Success           200      {object}  map[string]string "Task deleted successfully"
// @Failure           400      {object}  map[string]string "Bad Request"
// @Failure           404      {object}  map[string]string "Task not found"
// @Failure           500      {object}  map[string]string "Internal Server Error"
// @Router            /task/{id} [delete]
func deleteTask(c *gin.Context) {
  isParams := c.Params.ByName("id")
  objId, err := primitive.ObjectIDFromHex(isParams)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Invalid task ID",
    })
    return
  }

  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  result, err := collection.DeleteOne(ctx, bson.M{"_id": objId})
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
    return
  }

  if result.DeletedCount == 0 {
    c.JSON(http.StatusNotFound, gin.H{
      "error": "Task not found",
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "Task deleted successfully",
  })
}

func main() {
  connectDB()
  router := gin.Default()
  router.Use(middlewares.CorsMiddleware())

  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

  router.POST("/task", createTask)
  router.GET("/tasks", listTasks)
  router.PATCH("/task/:id", updateTaks)
  router.DELETE("/task/:id", deleteTask)

  router.Run(":8080")
}
