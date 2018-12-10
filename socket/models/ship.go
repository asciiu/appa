package models

type ShipBoost struct {
	ClientID string `json:"clientID"`
	Topic    string `json:"topic"`
	Boost    bool   `json:"boost"`
}

type ShipRotation struct {
	ClientID string  `json:"clientID"`
	Topic    string  `json:"topic"`
	Radian   float64 `json:"radian"`
}

type Ship struct {
	Topic    string  `json:"topic"`
	ClientID string  `json:"clientID"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Width    float64 `json:"width"`
	Height   float64 `json:'height"`
	Image    string  `json:"image"`
}

type Message struct {
	ClientID string `json:"clientID"`
	Topic    string `json:"topic"`
}

func NewShipRequest(clientID, topic string, width, height float64) *Ship {
	return &Ship{
		Topic:    topic,
		ClientID: clientID,
		X:        width / 2,
		Y:        height / 2,
	}
}
