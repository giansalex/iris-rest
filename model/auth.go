package model

// Auth Login model
type Auth struct {
	Username string `json:"username" xml:"username"`
	Password string `json:"password" xml:"password"`
}
