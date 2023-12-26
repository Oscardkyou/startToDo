package main

import (
	"fmt"
	"testdk/internal/repository"
	srv "testdk/internal/service"
)

func main() {
	storage := repository.NewInMemmoryRepository()
	service := srv.NewReminderService(storage)
	createdModel := service.ReminderCreate()

	fmt.Printf("%v", createdModel)
}
