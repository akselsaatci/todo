package domain

import (
	"github.com/google/uuid"
	"time"
)

type Todo struct {
	ID             uuid.UUID `gorm:"type:uuid;"gorm:"primaryKey"`
	UserId         uuid.UUID
	Title          string
	Description    string
	CreateDate     time.Time
	IsDone         bool
	CompletionDate *time.Time
}
