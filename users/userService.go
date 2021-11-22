package users

import (
	"context"
	"errors"

	"backend/database"
	"backend/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func UniqueEmail(email string) error {
	filter := bson.M{"email": email}
	count, _ := database.UserCollection.CountDocuments(context.Background(), filter)
	if count > 0 {
		return errors.New("Email already exist")
	}
	return nil
}

func FindByEmail(email string) (User, error) {
	u := new(User)
	filter := bson.M{"email": email}
	if err := database.UserCollection.FindOne(context.Background(), filter).Decode(u); err != nil {
		return *u, errors.New("User not found")
	}
	return *u, nil
}

func Insert(userNew UserNew, rol string) (string, error) {
	u := new(User)
	u.Email = userNew.Email
	u.Password = userNew.Password
	u.Rol = rol
	u.Active = true
	validate := utils.NewValidator()
	if err := validate.Struct(u); err != nil {
		return "", utils.ValidatorErrors(err)
	}
	_ = setPassword(u, u.Password)
	userInsertion, err := database.UserCollection.InsertOne(context.Background(), u)
	if err != nil {
		return "", errors.New("Error on user insert")
	}
	userId := userInsertion.InsertedID.(primitive.ObjectID).Hex()
	return userId, nil
}

func setPassword(user *User, password string) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Error on set password")
	}
	user.Password = string(hash)
	return nil
}
