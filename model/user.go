package model

type User struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Balance  int    `json:"balance"`
	Version  int    `json:version`
}

// User represents a user in the system
type UserSocial struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Location string `json:"location"`
}
