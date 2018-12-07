package models

type ShipSetupRequest struct {
	ClientID     string `json:"clientID"`
	Topic        string `json:"topic"`
	ScreenWidth  uint32 `json:"width"`
	ScreenHeight uint32 `json:"height"`
}

type ShipResponse struct {
	Topic    string `json:"topic"`
	ClientID string `json:"clientID"`
	X        uint32 `json:"x"`
	Y        uint32 `json:"y"`
}

func NewShipRequest(clientID, topic string, width, height uint32) *ShipResponse {
	return &ShipResponse{
		Topic:    topic,
		ClientID: clientID,
		X:        width / 2,
		Y:        height / 2,
	}
}
