package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Status       string        `json:"status"`
	Priority     int           `json:"priority"`
	Deadline     time.Time     `json:"deadline"`
	CreationDate time.Time     `json:"creation_date"`
	TimeSpent    time.Duration `json:"time_spent"`
	ProjectId    uuid.UUID     `json:"project_id"`
}
