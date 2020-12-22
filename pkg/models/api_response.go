package models

// APIResponse struct represents error messages sent out by the api handler to
// client
type APIResponse struct {
	StatusCode  int         `json:"code"`
	Description string      `json:"desc,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Links       *Links      `json:"links,omitempty"`
	Errors      []Error     `json:"errors,omitempty"`
}

// Error struct for any input field errors
type Error struct {
	Code  int    `json:"code"`
	Field string `json:"field,omitempty"`
	Error string `json:"error"`
}

// Links type
type Links struct {
	First string `json:"first,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
	Last  string `json:"last,omitempty"`
}
