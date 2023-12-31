package repository

import (
	"testdk/internal"
	"testdk/internal/domain"
)

type InMemmory struct {
	s map[domain.EntityId]*internal.Reminder
}

func (m *InMemmory) Add(r *internal.Reminder) {
	m.s[r.GetId()] = r

}

func NewInMemmoryRepository() *InMemmory {
	return &InMemmory{
		s: make(map[domain.EntityId]*internal.Reminder, 0),
	}
}

// package main

// type Reminder struct {
// 	id   string
// 	// Добавьте другие поля вашей структуры Reminder
// }

// type InMemory struct {
// 	s map[string]*Reminder
// }

// func (m *InMemory) Add(r *Reminder) {
// 	if m.s == nil {
// 		m.s = make(map[string]*Reminder)
// 	}
// 	m.s[r.id] = r
// }

// func NewInMemoryRepository() *InMemory {
// 	return &InMemory{
// 		s: make(map[string]*Reminder),
// 	}
// }

// func main() {
// 	// Пример использования
// 	repo := NewInMemoryRepository()

// 	reminder1 := &Reminder{id: "1"}
// 	reminder2 := &Reminder{id: "2"}

// 	repo.Add(reminder1)
// 	repo.Add(reminder2)

// 	// Теперь у вас есть напоминания в памяти
// 	// Можете продолжить разрабатывать вашу программу
// }
