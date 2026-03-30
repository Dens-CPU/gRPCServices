package main

import (
	"log"

	"github.com/DencCPU/gRPCServices/OrderService/pkg/apprunner"
)

func main() {
	app, err := apprunner.FxAppRunner()
	if err != nil {
		log.Fatal(err)
	}
	app.Run()
}
