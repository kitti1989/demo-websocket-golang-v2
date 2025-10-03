package main

//import (
//	"demo/service/src"
//	"fmt"
//	"net/http"
//)

//func main() {
//	http.HandleFunc("/ping", src.PingHandler)
//	http.HandleFunc("/hello", src.HelloHandler)
//	http.HandleFunc("/model", src.CallModel)

//	fmt.Println("Server running at http://localhost:8080")
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		panic(err)
//	}
//}

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all origins
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("Received:", string(msg))
		conn.WriteMessage(websocket.TextMessage, []byte("Hello from server"))
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected")

	// Main loop for reading messages
	for {
		// ReadMessageType (1 for text, 2 for binary data)
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read failed:", err)
			break
		}

		switch messageType {
		case websocket.BinaryMessage:
			// 'message' is a []byte containing the video frame data
			log.Printf("Received video frame of size: %d bytes\n", len(message))

			// **TODO: Implement frame processing logic here**
			// E.g., Decode the frame, save it to a file, or forward it.
			// For live streaming, you might process it in real-time.
			confirmation := "Server received frame successfully!\nframe of size: %d bytes\n"

			// WriteMessage takes the message type (Text or Binary) and the data ([]byte)
			err = conn.WriteMessage(websocket.TextMessage, []byte(confirmation))
			if err != nil {
				log.Println("Write failed:", err)
				break
			}
			log.Println("Sent confirmation back to client.")

		case websocket.TextMessage:
			log.Printf("Received text message: %s\n", message)
		}
	}
	log.Println("Client disconnected")
}

func main() {
	http.HandleFunc("/ws", handler)
	http.HandleFunc("/wsHandler", wsHandler)
	http.ListenAndServe(":8080", nil)
}
