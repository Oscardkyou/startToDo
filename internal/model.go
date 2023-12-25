package internal

import (
	"testdk/internal/domain"
)

type Reminder struct {
	id domain.EntityId
}

func NewReminder(id domain.EntityId) *Reminder {
	return &Reminder{
		id: id,
	}
}

func (r *Reminder) GetId() domain.EntityId {
	return r.id
}
