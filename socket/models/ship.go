package models

type ShipBoostUpdate struct {
	ClientID string `json:"clientID"`
	Topic    string `json:"topic"`
	Boost    bool   `json:"boost"`
}

type ShipBoostRotation struct {
	ClientID string  `json:"clientID"`
	Topic    string  `json:"topic"`
	Radian   float64 `json:"radian"`
}

type ShipResponse struct {
	Topic    string  `json:"topic"`
	ClientID string  `json:"clientID"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}

func NewShipRequest(clientID, topic string, width, height float64) *ShipResponse {
	return &ShipResponse{
		Topic:    topic,
		ClientID: clientID,
		X:        width / 2,
		Y:        height / 2,
	}
}
