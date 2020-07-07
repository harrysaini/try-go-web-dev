package requests

// CreateUser model
type CreateUser struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}
