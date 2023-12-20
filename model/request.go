package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User represents a user in the system
type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Location string `json:"location"`
}
