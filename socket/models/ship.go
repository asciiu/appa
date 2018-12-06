package models

type ShipSetupRequest struct {
	ClientID     string `json:"clientID"`
	ScreenWidth  uint32 `json:"width"`
	ScreenHeight uint32 `json:"height"`
}

type Ship struct {
	ClientID string `json:"clientID"`
	X        uint32 `json:"x"`
	Y        uint32 `json:"y"`
}

func NewShip(clientID string, width, height uint32) *Ship {
	return &Ship{
		ClientID: clientID,
		X:        width / 2,
		Y:        height / 2,
	}
}
