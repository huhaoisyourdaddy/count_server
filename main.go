package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count = 0

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "这是第 %d 次被调用\n", count)
	fmt.Printf("这是第 %d 次被调用\n", count)
	fmt.Fprintf(w, "这是一个简单的计数服务器")
}
