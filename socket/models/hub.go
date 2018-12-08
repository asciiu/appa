// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package models

import (
	"encoding/json"
	"log"

	"github.com/asciiu/oldiez/socket/constants/topic"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			var msg interface{}
			if err := json.Unmarshal(message, &msg); err != nil {
				log.Println(err)
			} else {
				// add the client ID to all client requests
				m := msg.(map[string]interface{})

				switch m["topic"] {
				case topic.ShipSetup:
					var shipSetup ShipSetupRequest
					json.Unmarshal(message, &shipSetup)

					shipResponse := NewShipRequest(shipSetup.ClientID, shipSetup.Topic, shipSetup.ScreenWidth, shipSetup.ScreenHeight)
					if res, err := json.Marshal(shipResponse); err != nil {
						log.Println(err)
					} else {
						h.broadcast(res)
					}

				case topic.ShipBoost:
					var boost ShipBoostUpdate
					json.Unmarshal(message, &boost)
					if res, err := json.Marshal(boost); err != nil {
						log.Println(err)
					} else {
						h.broadcast(res)
					}

				default:
					log.Println("what?")
				}
			}
		}
	}
}

func (h *Hub) broadcast(message []byte) {
	for client := range h.Clients {
		select {
		case client.Send <- message:
		default:
			close(client.Send)
			delete(h.Clients, client)
		}
	}
}
