package infra

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase() *mongo.Database {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  clientOptions := options.Client().ApplyURI("mongodb://localhost:27018")
  client, err := mongo.Connect(ctx, clientOptions)

  if err != nil {
    panic(err)
  }

  return client.Database("go-todolist")
}
