package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key"`
	NameLastName  string
	UserName      string
	Email         string
	Password      string
	LastLoginDate time.Time
	IsVerified    bool
}
