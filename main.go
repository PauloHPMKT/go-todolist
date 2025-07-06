package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/PauloHPMKT/go-todolist/schemas"
	"github.com/gin-gonic/gin"
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

  router.POST("/task", createTask)
  router.GET("/tasks", listTasks)
  router.PATCH("/task/:id", updateTaks)
  router.DELETE("/task/:id", deleteTask)

  router.Run(":8080")
}
