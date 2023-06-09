package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"test/model"
	"test/pkg/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt = "ajsdijaskdasl122312klsdjka"
	// будем использовать для рассшифровки токена
	signingKey = "kajsdljaskdja332$#"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return a.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (a *AuthService) GenerateToken(login string, password string) (string, error) {
	user, err := a.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(), // время жизни токена
			IssuedAt:  time.Now().Unix(),                     // когда был создан токен
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))

}

func (a *AuthService) ParseToken(accessToken string) (int, error) {

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")

		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims ")

	}

	return claims.UserId, nil
}

func (a *AuthService) IsUserExist(login string) (bool, error) {
	return a.repo.Check(login)
}
