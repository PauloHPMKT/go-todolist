package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
  ID          primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
  Title       string              `bson:"title" json:"title"`
  Description string              `bson:"description" json:"description"`
  Completed   bool                `bson:"completed" json:"completed"`
  DueDate     time.Time           `bson:"due_date" json:"due_date"`
  CreatedAt   time.Time           `bson:"created_at" json:"created_at"`
  UpdatedAt   *time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}
