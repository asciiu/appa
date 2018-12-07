package models

import (
	"github.com/asciiu/oldiez/common/constants/response"
)

type ShipSetupRequest struct {
	ClientID     string `json:"clientID"`
	ScreenWidth  uint32 `json:"width"`
	ScreenHeight uint32 `json:"height"`
}

type ShipResponse struct {
	Type     string `json:"type"`
	ClientID string `json:"clientID"`
	X        uint32 `json:"x"`
	Y        uint32 `json:"y"`
}

func NewShipRequest(clientID string, width, height uint32) *ShipResponse {
	return &ShipResponse{
		Type:     response.SetupResponse,
		ClientID: clientID,
		X:        width / 2,
		Y:        height / 2,
	}
}
