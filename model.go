package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

// Example model represents a database model
type ExampleModel struct {
	gorm.Model
	Name string
}

// task is the function to be executed on a schedule
func task(db *gorm.DB) {
	now := time.Now()
	log.Println("Task is being run... ", now)

	//Example database operation
	newRecord := ExampleModel{
		Name: fmt.Sprintf("Record at %s", now.Format(time.RFC850)),
	}
	db.Create(&newRecord)
}
