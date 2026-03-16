package main

import (
	apprunner "Academy/gRPCServices/OrderService/pkg/appruner"
	"log"
)

func main() {
	app, err := apprunner.FxAppRunner()
	if err != nil {
		log.Fatal(err)
	}
	app.Run()
}
