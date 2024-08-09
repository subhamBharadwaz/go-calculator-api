package models

// AddRequest represents the request payload for the add operation
type AddRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

// AddResponse represents the response payload for the add operation
type AddResponse struct {
	Result int `json:"result"`
}
