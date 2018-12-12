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

type Asteroid struct {
	OrderID string  `json:"orderID"`
	Topic   string  `json:"topic"`
	Size    float64 `json:"size"`
}

type Ship struct {
	ClientID  string  `json:"clientID"`
	Topic     string  `json:"topic"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	VelocityX float64 `json:"velocityX"`
	VelocityY float64 `json:"velocityY"`
	Width     float64 `json:"width"`
	Height    float64 `json:'height"`
	Image     string  `json:"image"`
	Heading   float64 `json:"heading"`
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
