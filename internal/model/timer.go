package model

import (
	"time"

	"github.com/google/uuid"
)

type TimerType string

type Timer struct {
	Id              uuid.UUID     `json:"id"`
	Name            string        `json:"name"`
	Type            TimerType     `json:"type"`
	StartTime       time.Time     `json:"start_time"`
	Duration        time.Duration `json:"duration"`
	ShortBreak      time.Duration `json:"short_break"`
	LongBreak       time.Duration `json:"long_break"`
	TimersInSession int           `json:"timers_in_session"`
}
