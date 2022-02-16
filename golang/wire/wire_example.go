package main

import (
	"fmt"

	"github.com/google/wire"
)

type Message string
type Greeter struct {
	Message Message
}
type Event struct {
	Greeter Greeter
}

func NewMessage() Message {
	return Message("Hello there!")
}
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}
func (g Greeter) Greet() Message {
	return g.Message
}
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
func main() {
	event := InitializeEvent()
	event.Start()
}
