package model

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	OwnerId      uuid.UUID `json:"owner_id"`
	CreationDate time.Time `json:"creation_date"`
}
