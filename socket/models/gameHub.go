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
type GameHub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []interface{}

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	// clientID -> ship
	Players map[string]*Ship
}

func NewGameHub() *GameHub {
	return &GameHub{
		Broadcast:  make(chan []interface{}),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Players:    make(map[string]*Ship, 0),
	}
}

func (h *GameHub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			// send current players to new client
			ships := make([]*Ship, 0)
			for _, ship := range h.Players {
				ships = append(ships, ship)
			}
			if len(ships) > 0 {
				if res, err := json.Marshal(ships); err != nil {
					log.Println(err)
				} else {
					client.Send <- res
				}
			}

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
			// TODO you need to remove the player from h.Players here

		case messages := <-h.Broadcast:
			responses := make([]interface{}, 0)

			for _, msg := range messages {
				m := msg.(map[string]interface{})
				clientID := m["clientID"].(string)

				switch m["topic"] {
				case topic.PlayerRegister:
					playerShip := NewShipRequest(
						clientID,
						m["topic"].(string),
						m["screenWidth"].(float64),
						m["screenHeight"].(float64),
					)
					h.Players[clientID] = playerShip
					responses = append(responses, playerShip)

				case topic.PlayerUnregister:
					delete(h.Players, clientID)
					responses = append(responses, Message{
						ClientID: clientID,
						Topic:    m["topic"].(string),
					})

				case topic.ShipBoost:
					boost := ShipBoost{
						ClientID: clientID,
						Topic:    m["topic"].(string),
						Boost:    m["boost"].(bool),
					}
					responses = append(responses, boost)

				case topic.ShipCoordinates:
					playerShip := h.Players[clientID]
					playerShip.X = m["x"].(float64)
					playerShip.Y = m["y"].(float64)

				case topic.ShipHeading:
					playerShip := h.Players[clientID]
					playerShip.Heading = m["heading"].(float64)

				case topic.ShipRotation:
					rot := ShipRotation{
						ClientID: clientID,
						Topic:    m["topic"].(string),
						Radian:   m["radian"].(float64),
					}
					responses = append(responses, rot)

				case topic.ShipLaser:
					responses = append(responses, Message{
						ClientID: clientID,
						Topic:    m["topic"].(string),
					})

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

func (h *GameHub) broadcast(message []byte) {
	for client := range h.Clients {
		select {
		case client.Send <- message:
		default:
			close(client.Send)
			delete(h.Clients, client)
		}
	}
}
