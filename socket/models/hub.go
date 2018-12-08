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
	Broadcast chan []interface{}

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []interface{}),
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
		case messages := <-h.Broadcast:
			responses := make([]interface{}, 0)

			for _, msg := range messages {
				m := msg.(map[string]interface{})

				switch m["topic"] {
				case topic.ShipSetup:
					shipResponse := NewShipRequest(
						m["clientID"].(string),
						m["topic"].(string),
						m["screenWidth"].(float64),
						m["screenHeight"].(float64))
					responses = append(responses, shipResponse)

				case topic.ShipBoost:
					boost := ShipBoost{
						ClientID: m["clientID"].(string),
						Topic:    m["topic"].(string),
						Boost:    m["boost"].(bool),
					}
					responses = append(responses, boost)

				case topic.ShipRotation:
					rot := ShipRotation{
						ClientID: m["clientID"].(string),
						Topic:    m["topic"].(string),
						Radian:   m["radian"].(float64),
					}
					responses = append(responses, rot)

				case topic.ShipLaser:
					rot := ShipLaser{
						ClientID: m["clientID"].(string),
						Topic:    m["topic"].(string),
					}
					responses = append(responses, rot)

				default:
					log.Println("what?")
				}
			}

			if res, err := json.Marshal(responses); err != nil {
				log.Println(err)
			} else {
				h.broadcast(res)
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
