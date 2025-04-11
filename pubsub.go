package main

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Data string `json:"data"`
}

type PubSub struct {
	subs []chan Message
	mu   sync.Mutex // Save variable of lock
}

// Subscribe adds a new subscriber channel
func (ps *PubSub) Subscribe() chan Message {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := make(chan Message, 1)
	ps.subs = append(ps.subs, ch)
	return ch
}

// Publish sends a message to all subscriber channel
func (ps *PubSub) Publish(msg *Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	for _, sub := range ps.subs {
		sub <- *msg
	}
}

func (ps *PubSub) Unsubscribe(ch chan Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	for i, sub := range ps.subs {
		if sub == ch {
			ps.subs = append(ps.subs[:i], ps.subs[i+1:]...)
			close(ch)
			break
		}
	}
}

func Pubsub() {
	app := fiber.New()

	pubsub := &PubSub{}

	app.Post("publisher", func(ctx *fiber.Ctx) error {
		message := new(Message)
		if err := ctx.BodyParser(message); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		pubsub.Publish(message)
		return ctx.JSON(&fiber.Map{
			"message": "add to subscriber",
		})
	})

	sub := pubsub.Subscribe()
	go func() {
		for msg := range sub {
			fmt.Println("Recieve message from sub1: ", msg)
		}
	}()

	sub2 := pubsub.Subscribe()
	go func() {
		for msg := range sub2 {
			fmt.Println("Recieve message from sub2: ", msg)
		}
	}()

	app.Listen(":8080")
}
