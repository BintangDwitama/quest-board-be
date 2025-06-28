package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type UserMask struct {
	UUID   uuid.UUID
	UserID int64
}

type Task struct {
	ID          int64
	Description string
	IsComplete  bool
	ETA         time.Time
}
