package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/404th/todo/model"
	"github.com/404th/todo/pkg/repository"
)

const salt = "68fed7gwefhmwe8fwef8wefhwe8f"

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

// helper func --> password hasher
func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
