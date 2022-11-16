package service

import (
	"crypto/sha1"
	"errors"
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

// parsing token and get data
func (ar AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(a *jwt.Token) (interface{}, error) {
		if _, ok := a.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not type *tokenClaims")
	}

	return claims.UserId, nil
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
