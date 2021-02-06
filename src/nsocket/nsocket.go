package nsocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving websockets ...")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	loop(conn)
}

func loop(conn *websocket.Conn) {
	for {
		messageType, r, err := conn.NextReader()
		if err != nil {
			return
		}
		w, err := conn.NextWriter(messageType)

		if err != nil {
			fmt.Println(err)
		}

		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(err)
		}
		if err := w.Close(); err != nil {
			fmt.Println(err)
		}
	}
}
