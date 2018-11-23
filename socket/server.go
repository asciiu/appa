package main

import (
	"log"
	"net/http"
	"time"

	"github.com/asciiu/oldiez/socket/models"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

type WebsocketServer struct {
	connections []*websocket.Conn
}

func NewServer() *WebsocketServer {
	return &WebsocketServer{
		connections: make([]*websocket.Conn, 0),
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Connect handles websocket connections
func (controller *WebsocketServer) Connect(hub *models.Hub, w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	client := &models.Client{
		Conn: conn,
		Send: make(chan []byte, 256),
		Hub:  hub,
	}
	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()

	return nil
}
