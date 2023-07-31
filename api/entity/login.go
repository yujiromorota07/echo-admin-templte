package entity

// Login - Login struct
type Login struct {
	Email    string
	Password string
}

// LoginToken - login token
type LoginToken struct {
	Token string `json:"token"`
}
