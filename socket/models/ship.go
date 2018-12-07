package models

type ShipSetupRequest struct {
	ClientID     string `json:"clientID"`
	Topic        string `json:"topic"`
	ScreenWidth  uint32 `json:"screenWidth"`
	ScreenHeight uint32 `json:"screenHeight"`
}

type ShipBoostUpdate struct {
	ClientID string `json:"clientID"`
	Topic    string `json:"topic"`
	Boost    uint32 `json:"boost"`
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
