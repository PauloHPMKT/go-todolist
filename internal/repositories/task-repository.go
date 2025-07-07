package repositories

import (
	"context"
	"time"

	"github.com/PauloHPMKT/go-todolist/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
  Create(ctx context.Context, task *entities.Task) (primitive.ObjectID, error)
  GetAll(ctx context.Context) ([]entities.Task, error)
  Update(ctx context.Context, id primitive.ObjectID , task *entities.Task) error
  Delete(ctx context.Context, id primitive.ObjectID) error
}

type taskRepository struct { // é quem faz conexao com a collection
  collection *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) TaskRepository {
  return &taskRepository{
    collection: db.Collection("go-todolist"),
  }
}

func (r *taskRepository) Create(ctx context.Context, task *entities.Task) (primitive.ObjectID, error) {
  task.CreatedAt = time.Now()
  result, err := r.collection.InsertOne(ctx, task)
  if err != nil {
    return primitive.NilObjectID, err
  }
  return result.InsertedID.(primitive.ObjectID), nil
}

func (r *taskRepository) GetAll(ctx context.Context) ([]entities.Task, error) {
  cursor, err := r.collection.Find(ctx, bson.M{})
  if err != nil {
    return nil, err
  }
  defer cursor.Close(ctx)

  var tasks []entities.Task
  if err := cursor.All(ctx, &tasks); err != nil {
    return nil, err
  }
  return tasks, nil
}

func (r *taskRepository) Update(ctx context.Context, id primitive.ObjectID, task *entities.Task) error {
  task.UpdatedAt = &time.Time{}
  task.UpdatedAt = &time.Time{}
  filter := bson.M{"_id": id}
  update := bson.M{"$set": task}

  _, err := r.collection.UpdateOne(ctx, filter, update)
  return err
}

func (r *taskRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
  _, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
  return err
}
