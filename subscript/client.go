package subscript

import (
	pb "awesomeProject/pb" // Update with your actual package path
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"sync"
	"time"
)

func subscribe(client pb.PubSubServiceClient, topic string) {

	//fmt.Println(GetPubSubClient().cancels[topic])
	// 使用逗号 ok 习惯用法
	if _, ok := GetPubSubClient().cancels[topic]; ok {
		fmt.Println(topic + "该主题已被其他客户端订阅")
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := client.Subscribe(ctx, &pb.SubscribeRequest{Topic: topic})
	GetPubSubClient().cancels[topic] = cancel
	if err != nil {
		log.Fatalf("could not subscribe: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		if resp == nil {
			break
		}
	}
}

func Subscript() {

	// 启动协程

	go subscribe(GetPubSubClient().client, "mover")
	//go subscribe(GetPubSubClient().client, "mover")
	//go subscribe(client, "stator")
	//go subscribe(client, "mover")
	// Simulate some other work
	// 等待一段时间，模拟服务端流接收
	//unsubscribes(client, "example")
	// 获取并取消订阅 "stator"

}
func UnSubcript() {
	cancelFunc := instance.cancels["mover"]
	if cancelFunc != nil {
		cancelFunc()
		delete(instance.cancels, "mover")
		fmt.Println("Subscription 'stator' cancelled")
	} else {
		fmt.Println("Cancel function for 'stator' is nil")
	}
}

type PubSubClient struct {
	client  pb.PubSubServiceClient
	conn    *grpc.ClientConn
	cancels map[string]context.CancelFunc
}

var instance *PubSubClient
var once sync.Once

func GetPubSubClient() *PubSubClient {
	once.Do(func() {
		// 设置连接超时时间为5秒
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// 使用 grpc.Dial 连接 gRPC 服务
		conn, err := grpc.DialContext(ctx, ":50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		md := make(map[string]context.CancelFunc)
		instance = &PubSubClient{
			client:  pb.NewPubSubServiceClient(conn),
			conn:    conn,
			cancels: md,
		}
	})
	return instance
}
