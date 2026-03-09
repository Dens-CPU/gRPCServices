package main

import (
	apprunner "Academy/gRPCServices/OrderService/pkg/appruner"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	
	err := apprunner.AppRunner(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
