package main

import (
	"context"
	app "gowit-task/internal/ticket-app"
	"log"
)

type root interface {
	Register() error
	Resolve(ctx context.Context) error
	Release() error
}

func main() {
	var app root = app.NewApp()
	err := app.Register()
	if err != nil {
		log.Panicln(err)
		panic("fail to register app")
	}
	ctx := context.Background()

	if err = app.Resolve(ctx); err != nil {
		log.Printf("cant resolve root: %v", err)
	}

	if err = app.Release(); err != nil {
		log.Fatalf("fail to release app, got %v", err)
	}

	log.Println("app finish successful")
}
