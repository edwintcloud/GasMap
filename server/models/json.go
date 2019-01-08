package models

// ResponseError model
type ResponseError struct {
	Error string `json:"error"`
}

// ResponseMsg model
type ResponseMsg struct {
	Message string `json:"message"`
}
