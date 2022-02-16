package main

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

type SayInterface interface {
	SayHello()
}

type Girl struct {
	Name string
	Age  int
}

func (g *Girl) SayHello() {
	fmt.Println("girl sayhello")
}

type Team struct {
	Girl *Girl
}

// when change SayInterface to *Girl it will error out if not using fx.As
// func NewGirl() SayInterface {
func NewGirl() *Girl {
	return &Girl{
		Name: "Lucy",
		Age:  18,
	}
}

func NewTeam(say SayInterface) *Team {
	return &Team{}
}

func main() {
	invoke := func(team *Team) {

	}
	app := fx.New(
		// fx.Provide(NewGirl),
		fx.Provide(fx.Annotate(NewGirl, fx.As(new(SayInterface)))),
		fx.Provide(NewTeam),
		fx.Invoke(invoke),
	)
	err := app.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
