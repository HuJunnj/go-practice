package main

import (
	pb "awesomeProject/pb" // Update with your actual package path
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func subscribe(client pb.PubSubServiceClient, topic string) {
	/*ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()*/

	stream, err := client.Subscribe(context.Background(), &pb.SubscribeRequest{Topic: topic})
	if err != nil {
		log.Fatalf("could not subscribe: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive message: %v", err)
		}
		if resp == nil {
			break
		}
		log.Printf("Received message: ID=%d, Content=%s, Timestamp=%s",
			resp.GetMessage().GetId(),
			resp.GetMessage().GetContent(),
			resp.GetMessage().GetTimestamp(),
		)
	}
}

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPubSubServiceClient(conn)

	go subscribe(client, "example")
	//go subscribe(client, "stator")
	//go subscribe(client, "mover")
	// Simulate some other work
	// time.Sleep(30 * time.Second)

	//unsubscribes(client, "example")
	select {}
	fmt.Println("dsf")
}
func unsubscribes(client pb.PubSubServiceClient, topic string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.Unsubscribe(ctx, &pb.UnsubscribeRequest{Topic: topic})
	if err != nil {
		log.Fatalf("could not unsubscribe: %v", err)
	}

	log.Printf("Unsubscribe response: %s", resp.GetMessage())
}
