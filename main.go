package main

import (
	"context"
	"fmt"

	"github.com/setkyar/go-order-api/appication"
)

func main() {
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
