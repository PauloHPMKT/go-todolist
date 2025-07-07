package usecases

import (
	"context"

	"github.com/PauloHPMKT/go-todolist/internal/entities"
	"github.com/PauloHPMKT/go-todolist/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUseCase interface {
  CreateTask(ctx context.Context, task *entities.Task) (primitive.ObjectID, error)
  GetTasks(ctx context.Context) ([]entities.Task, error)
  UpdateTask(ctx context.Context, id primitive.ObjectID , task *entities.Task) error
  DeleteTask(ctx context.Context, id primitive.ObjectID) error
}

type taskUseCase struct {
  repo repositories.TaskRepository
}

func NewTaskUseCase(repo repositories.TaskRepository) TaskUseCase {
  return &taskUseCase{repo: repo}
}

func (taskUseCase *taskUseCase) CreateTask(ctx context.Context, task *entities.Task) (primitive.ObjectID, error) {
  return taskUseCase.repo.Create(ctx, task)
}

func (taskUseCase *taskUseCase) GetTasks(ctx context.Context) ([]entities.Task, error) {
  return taskUseCase.repo.GetAll(ctx)
}

func (taskUseCase *taskUseCase) UpdateTask(ctx context.Context, id primitive.ObjectID, task *entities.Task) error {
  return taskUseCase.repo.Update(ctx, id, task)
}

func (taskUseCase *taskUseCase) DeleteTask(ctx context.Context, id primitive.ObjectID) error {
  return taskUseCase.repo.Delete(ctx, id)
}
