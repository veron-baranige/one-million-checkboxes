package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

// buffer: temporary storage area in memory used to hold data while it is transffered from one place to another
// buffer size impacts how data is read. if a message exceeds the buffer it will be read in chunks
// a larger buffer may improve performance for large messages but uses more memory.
var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		if os.Getenv("ENV") == "DEV" {
			return true
		}
		return r.Header.Get("Origin") == os.Getenv("ORIGIN")
	},
}

// responsible for sending & receiving messages thourgh the ws
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		log.Printf("Recieved message: %s", message)

		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}
