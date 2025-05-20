package controller

import (
	"errors"
	"time"

	"user-service/model"
	"user-service/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Repo      repository.IUserRepo
	JWTSecret string
}

func (c *UserController) Register(user *model.User, password string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.PasswordHash = string(hash)

	return c.Repo.Create(user)
}

func (c *UserController) Login(id, password string) (string, error) {
	u, err := c.Repo.FindByIdentifier(id)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) != nil {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(2 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(c.JWTSecret))
}
