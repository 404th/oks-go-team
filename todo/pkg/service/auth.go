package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/404th/todo/model"
	"github.com/404th/todo/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	signingKey = "f7567t8su9di0fosdfhusdfijosfd"
	salt       = "68fed7gwefhmwe8fwef8wefhwe8f"
	timeTTL    = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (ar AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generateHashPassword(user.Password)
	return ar.repo.CreateUser(user)
}

// token generator for sign in method
func (ar AuthService) GenerateToken(username, password string) (string, error) {
	user, err := ar.repo.GetUser(username, generateHashPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(timeTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

// helper func --> password hasher
func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
