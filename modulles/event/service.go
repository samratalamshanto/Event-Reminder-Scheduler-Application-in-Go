package event

import (
	"Reminder_Event_Mail_Scheduler_in_Go/modulles/base"
	"Reminder_Event_Mail_Scheduler_in_Go/pakages/db"
)

func GetAllEvents() ([]Event, error) {
	var eventList []Event
	result := db.DB.Where(&Event{Status: TODO}).Find(&eventList)
	return eventList, result.Error
}

func GetEvent(id int64) (Event, error) {
	var event Event
	result := db.DB.Where(&Event{BaseEntity: base.BaseEntity{Id: id}, Status: TODO}).First(&event)
	return event, result.Error
}

func CreateEvent(dto EventDto) (Event, error) {
	var event Event
	mapEventDto(&dto, &event)
	result := db.DB.Save(&event)
	return event, result.Error
}

func UpdateEvent(dto EventDto, id int64) (Event, error) {
	var event Event
	mapEventDto(&dto, &event)
	event.Id = id
	result := db.DB.Save(&event)
	return event, result.Error
}

func DeleteEvent(id int64) error {
	res := db.DB.Delete(&Event{BaseEntity: base.BaseEntity{Id: id}})
	return res.Error
}

func mapEventDto(dto *EventDto, event *Event) {
	event.Status = TODO
	event.Name = dto.Name
	event.Description = dto.Description
	event.ReminderDate = dto.ReminderDate
}
