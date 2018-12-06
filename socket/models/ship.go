package models

type Ship struct {
	ClientID string `json:"clientID"`
	X        uint   `json:"x"`
	Y        uint   `json:"y"`
}
