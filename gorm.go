package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func Gorm() {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&ExampleModel{})

	c := cron.New(cron.WithSeconds())

	c.AddFunc("*/3 * * * * *", func() {
		fmt.Println("Hello World every 3 second")
		task(db)
	})

	c.Start()

	select {}
}
