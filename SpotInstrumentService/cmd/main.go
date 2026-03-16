package main

import (
	"Academy/gRPCServices/SpotInstrumentService/pkg/apprunner"
	"log"
)

func main() {
	app, err := apprunner.FxAppRunner()
	if err != nil {
		log.Fatal(err)
	}
	app.Run()
	// err := apprunner.AppRunner()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
