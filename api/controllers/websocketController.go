package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

type WebsocketController struct {
	connections []*websocket.Conn
}

func NewWebsocketController() *WebsocketController {
	return &WebsocketController{
		connections: make([]*websocket.Conn, 0),
	}
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// Connect handles websocket connections
func (controller *WebsocketController) Connect(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	if err := ws.WriteMessage(websocket.TextMessage, []byte("ready!")); err != nil {
		log.Println("write:", err)
		return err
	}
	i := len(controller.connections)
	controller.connections = append(controller.connections, ws)

	// block until client closes
	if _, _, err := ws.ReadMessage(); err != nil {
		// client closes this will read: websocket: close 1005 (no status)
		log.Println("read error: ", err)
	}

	// remove the connection from the connect pool
	controller.connections = append(controller.connections[:i], controller.connections[i+1:]...)
	return nil
}

func (controller *WebsocketController) Ticker() {
	for {
		time.Sleep(5 * time.Second)

		// send events to all connected clients
		for _, conn := range controller.connections {
			//json, err := json.Marshal(events)
			//if err != nil {
			//log.Println("marchall error: ", err)
			//}
			//conn.WriteMessage(websocket.TextMessage, json)
			conn.WriteMessage(websocket.TextMessage, []byte("tick"))
		}
	}
}
