package main

import (
	"awesomeProject/subscript"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!") // 输出到客户端的内容
}
func handlers(w http.ResponseWriter, r *http.Request) {
	subscript.Subscript()
	fmt.Fprintf(w, "Hello, World!") // 输出到客户端的内容
}
func handlerb(w http.ResponseWriter, r *http.Request) {
	subscript.UnSubcript() // 输出到客户端的内容
}

func main() {
	http.HandleFunc("/", handler)             // 设置访问的路由
	http.HandleFunc("/subscript", handlers)   // 设置访问的路由
	http.HandleFunc("/unsubscript", handlerb) // 设置访问的路由
	fmt.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil) // 设置监听的端口
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
