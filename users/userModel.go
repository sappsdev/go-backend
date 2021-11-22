package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Password string             `json:"password" validate:"required,alphanum,min=8"`
	Email    string             `json:"email" validate:"required,email"`
	Rol      string             `json:"rol"`
	Active   bool               `json:"active"`
	Verified bool               `json:"verified,omitempty"`
	Pin      string             `json:"pin,omitempty"`
}

type UserNew struct {
	Password string `json:"password" validate:"required,alphanum,min=8"`
	Email    string `json:"email" validate:"required,email"`
	Code     string `json:"code,omitempty"`
}

type UserLogin struct {
	Password string `json:"password" validate:"required,alphanum,min=8"`
	Email    string `json:"email" validate:"required,email"`
}