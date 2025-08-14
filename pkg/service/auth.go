package service

import (
	"crypto/sha1"
	"fmt"

	todo "github.com/amenshenin/go1"
	"github.com/amenshenin/go1/pkg/repository"
)

const salt = "assasasadgfdh12245ytgffsafs"

type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
