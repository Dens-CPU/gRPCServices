package main

import (
	spotAPI "Academy/gRPCServices/Protobuf/gen/spot_service"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient(":8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}

	client := spotAPI.NewSpotInstrumentServiceClient(conn)

	resp, err := client.ViewMarket(context.Background(), &spotAPI.ViewReq{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.EnableMarkets)
}
