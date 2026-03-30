package main

import (
	"log"

	"github.com/DencCPU/gRPCServices/SpotInstrumentService/pkg/apprunner"
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
