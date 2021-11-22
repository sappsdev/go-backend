package auth

type UserResponse struct {
	Avatar string `json:"avatar"`
	Rol    string `json:"rol"`
	Status bool `json:"status"`
}

type ErrorResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token 	string       	`json:"token"`
	Rol    string `json:"rol"`
	Errors 	ErrorResponse `json:"errors"`
}