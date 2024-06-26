package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// 升级器，用于将 HTTP 连接升级为 WebSocket 连接
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	fmt.Println("WebSocket 服务器启动在 :8080 端口...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("监听错误:", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// 将 HTTP 连接升级为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("升级为 WebSocket 连接时出错:", err)
		return
	}
	defer func() {
		fmt.Println("程序退出")
		conn.Close()
	}()

	for {
		// 读取客户端发送的数据
		messageType, message, err := conn.ReadMessage()
		fmt.Println(11111)
		if err != nil {
			fmt.Println("读取消息时出错:", err)
			break
		}
		fmt.Println(2222222)
		// 根据收到的消息类型处理不同的逻辑
		switch string(message) {
		case "a":
			go handleMessage(conn, messageType, "dsf")
		case "b":
			go handleMessage(conn, messageType, "dsfgdssd")
		case "c":
			go handleMessage(conn, messageType, "sds")
		default:
			fmt.Println("未知消息类型:", string(message))
		}
		fmt.Println(1111)
	}
}
func handleMessage(conn *websocket.Conn, messageType int, response string) {
	for {
		time.Sleep(20 * time.Millisecond)
		// 将消息回发给客户端
		err := conn.WriteMessage(messageType, []byte(response))
		if err != nil {
			// 检测到写入错误（如连接关闭），退出循环
			fmt.Println("发送消息时出错:", err)
			break
		}
	}
}
