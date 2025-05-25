package models

import (
	"context"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type AuthCredentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthRepository interface {
	RegisterUser(ctx context.Context, registerData *AuthCredentials) (*User, error)
	GetUser(ctx context.Context, query interface{}, args ...interface{}) (*User, error)
}

type AuthService interface {
	Login(ctx context.Context, loginData *AuthCredentials)(string, *User, error)
	Register(ctx context.Context, registerData *AuthCredentials)(string, *User, error)
}

//check if a password matches a hash
func MatchesHas(password, hash string)bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err  == nil
}

//check is an email is valid
func IsValidEmail(email string)bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}