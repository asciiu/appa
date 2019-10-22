package controllers

// A ResponseSuccess will always contain a status of "successful".
// This response may or may not include data encapsulating the user information.
// swagger:model responseError
type ResponseError struct {
	Status   string   `json:"status"`
	Messages []string `json:"messages"`
}