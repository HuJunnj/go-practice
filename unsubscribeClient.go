package main

import (
	"awesomeProject/pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {

		}
	}(conn)

	client := pb.NewPubSubServiceClient(conn)

	unsubscribe(client, "example")
}
func unsubscribe(client pb.PubSubServiceClient, topic string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.Unsubscribe(ctx, &pb.UnsubscribeRequest{Topic: topic})
	if err != nil {
		log.Fatalf("could not unsubscribe: %v", err)
	}

	log.Printf("Unsubscribe response: %s", resp.GetMessage())
}
