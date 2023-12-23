package main

func main() {

	var a Reminder = Reminder{
		id: "John Deppy",
	}
	storage := NewInMemmoryRepository()
	service := NewReminderService(storage)
}
