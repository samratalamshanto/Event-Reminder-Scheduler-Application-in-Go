package event

import "time"

type EventDto struct {
	Name         string    `json:"name"`
	Description  string    `json:"description" `
	ReminderDate time.Time `json:"reminderDate" `
}
