package model

import (
	"time"

	"github.com/google/uuid"
)

type WorkSession struct {
	Id        uuid.UUID `json:"id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	TaskId    uuid.UUID `json:"task_id"`
	UserId    uuid.UUID `json:"user_id"`
	Notes     string    `json:"notes"`
}
