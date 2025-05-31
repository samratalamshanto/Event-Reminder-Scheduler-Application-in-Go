package event

import (
	"Reminder_Event_Mail_Scheduler_in_Go/modulles/base"
	"time"
)

type Event struct {
	base.BaseEntity
	Name         string    `json:"name" gorm:"type:text; not null"`
	Description  string    `json:"description" gorm:"type:text;"`
	ReminderDate time.Time `json:"reminderDate" gorm:"type:datetime; not null"`
	Status       string    `json:"status" gorm:"type:varchar(100); not null"`
}

const (
	TODO       = "TODO"
	COMPLETED  = "COMPLETED"
	InProgress = "InProgress"
	PENDING    = "PENDING"
)
