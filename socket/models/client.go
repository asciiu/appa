package models

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/asciiu/oldiez/socket/constants/topic"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10240
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  10240,
	WriteBufferSize: 10240,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	GameHub *GameHub

	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	//Send chan []byte
	Send chan []interface{}
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) ReadPump() {
	defer func() {
		c.GameHub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	//c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetReadDeadline(time.Time{})
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	// unmarshal successful reads and log all read errors
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNoStatusReceived) {
				log.Printf("error: %v", err)
			}
			break
		} else {
			message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
			c.GameHub.Broadcast <- message
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	ticker := time.NewTicker(3 * time.Second)

	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case messages, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			// append queued messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				msgs := <-c.Send
				for _, msg := range msgs {
					messages = append(messages, msg)
				}
			}

			if json, err := json.Marshal(messages); err != nil {
				log.Println(err)
			} else {
				w.Write(json)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:

			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			asteroids := make([]interface{}, 0)
			asteroids = append(asteroids,
				Asteroid{
					Topic:   topic.NewAsteroid,
					OrderID: uuid.New().String(),
					Size:    0.01,
				},
			)

			if json, err := json.Marshal(asteroids); err != nil {
				log.Println(err)
			} else {
				if err := c.Conn.WriteMessage(websocket.TextMessage, []byte(json)); err != nil {
					return
				}
			}
		}
	}
}
