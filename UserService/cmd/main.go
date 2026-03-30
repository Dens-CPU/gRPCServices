package main

import (
	"log"

	"github.com/DencCPU/gRPCServices/UserService/pkg/apprunner"
)

func main() {
	err := apprunner.AppRunner()
	if err != nil {
		log.Fatal(err)
	}
}
