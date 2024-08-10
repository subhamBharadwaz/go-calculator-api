package models

// AddRequest represents the request payload for the add operation
type Request struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

// AddResponse represents the response payload for the add operation
type Response struct {
	Result float64 `json:"result"`
}
