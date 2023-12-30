package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id               uuid.UUID `json:"id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	RegistrationDate time.Time `json:"registration_date"`
}
