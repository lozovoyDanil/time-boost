package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id               uuid.UUID `json:"id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	EmailConfirmed   bool      `json:"email_confirmed"`
	Password         string    `json:"password"`
	Salt             string    `json:"salt"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	RegistrationDate time.Time `json:"registration_date"`
}
