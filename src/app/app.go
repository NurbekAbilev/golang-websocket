package main

import (
	"fmt"
	"net/http"
)

func serveStatic() {
	fmt.Println("Serving static files...")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3000", nil)
}

func main() {
	go serveStatic()
	//go nsocket.ServeWebSocket(http.ResponseWriter(), http.Request{})
	//go nsocket.Example()
}
