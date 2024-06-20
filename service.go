package main

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	pb "awesomeProject/pb" // Update with your actual package path
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPubSubServiceServer
	mu          sync.Mutex
	subscribers map[string]context.CancelFunc
}

func NewServer() *server {
	return &server{
		subscribers: make(map[string]context.CancelFunc),
	}
}

func (s *server) Subscribe(req *pb.SubscribeRequest, stream pb.PubSubService_SubscribeServer) error {
	topic := req.GetTopic()
	log.Printf("Client subscribed to topic: %s", topic)

	ctx, cancel := context.WithCancel(context.Background())
	s.mu.Lock()
	s.subscribers[topic] = cancel
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.subscribers, topic)
		s.mu.Unlock()
	}()
	if topic == "stator" {
		for i := 0; i < 10000000000; i++ { // Simulate sending multiple messages
			select {
			case <-ctx.Done():
				log.Printf("Subscription to topic %s canceled", topic)
				return ctx.Err()
			default:
				message := &pb.Message{
					Id:        int32(i + 1),
					Content:   "Message " + string(i+1) + " from topic " + topic,
					Timestamp: time.Now().Format(time.RFC3339),
				}
				response := &pb.SubscribeResponse{
					Message: message,
				}
				if err := stream.Send(response); err != nil {
					return err
				}
				time.Sleep(20 * time.Millisecond) // Simulate delay
			}
		}
	}
	if topic == "mover" {
		for i := 0; i < 100000000; i++ { // Simulate sending multiple messages
			select {
			case <-ctx.Done():
				log.Printf("Subscription to topic %s canceled", topic)
				return ctx.Err()
			default:
				message := &pb.Message{
					Id:        int32(i + 1),
					Content:   "Message " + string(i+1) + " from topic " + topic,
					Timestamp: time.Now().Format(time.RFC3339),
				}
				response := &pb.SubscribeResponse{
					Message: message,
				}
				if err := stream.Send(response); err != nil {
					return err
				}
				time.Sleep(20 * time.Millisecond) // Simulate delay
			}
		}
	}
	return nil
}

func (s *server) Unsubscribe(ctx context.Context, req *pb.UnsubscribeRequest) (*pb.UnsubscribeResponse, error) {
	topic := req.GetTopic()
	s.mu.Lock()
	cancel, ok := s.subscribers[topic]
	s.mu.Unlock()

	if !ok {
		return &pb.UnsubscribeResponse{
			Message: "No active subscription for topic " + topic,
		}, nil
	}

	cancel()
	return &pb.UnsubscribeResponse{
		Message: "Unsubscribed from topic " + topic,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPubSubServiceServer(s, NewServer())

	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
