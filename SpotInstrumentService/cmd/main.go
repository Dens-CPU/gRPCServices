package main

import (
	"Academy/gRPCServices/SpotInstrumentService/pkg/apprunner"
	"log"
)

func main() {
	err := apprunner.AppRunner()
	if err != nil {
		log.Fatal(err)
	}
}
