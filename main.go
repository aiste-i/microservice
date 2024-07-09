package main

import (
	"context"
	"fmt"

	"github.com/aiste-i/mock-microservice/app"
)

func main() {
	app := app.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
