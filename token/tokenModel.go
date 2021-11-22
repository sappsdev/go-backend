package token

type UserAuth struct {
	ID    string `bson:"id" json:"id"`
	Email string `json:"email"`
	Rol   string `json:"rol"`
}