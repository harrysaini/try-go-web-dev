package models

// Status of req
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Response sent
type Response struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}
