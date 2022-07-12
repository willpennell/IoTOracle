package main

import (
	pb "IoToracle/probuf"
	"context"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Did not connect: &v", err)
	}
	defer conn.Close()
	c := pb.NewOracleNodeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	color.Cyan("Client connected at: %v", address)

	r, err := c.SendRequest(ctx, &pb.RequestToNode{
		IotId:         "rp4",
		RequestId:     1,
		RequireResult: true,
	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(time.Second * 5)
	color.Green("%v\n%v\n%v\n%v\n", r.IotId, r.RequestId, r.RequiredResult, r.ActualResult)

}
