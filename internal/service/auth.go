package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hixraid/blog/internal/data/model"
	"github.com/hixraid/blog/internal/data/repository"
	"github.com/hixraid/blog/internal/utils"
)

const (
	signingKey = "123456789"
	tokenTTL   = time.Hour * 24
)

type userClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

type Authorization struct {
	repository repository.UserRepository
}

func NewAuthorization(repos repository.UserRepository) *Authorization {
	return &Authorization{repos}
}

func (s *Authorization) CreateUser(input model.UserInput) (int, error) {
	password, err := utils.HashPassword(input.Password)
	if err != nil {
		return -1, err
	}
	input.Password = password

	return s.repository.Create(input)
}

func (s *Authorization) GenerateToken(email, password string) (string, error) {
	password, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}

	user, err := s.repository.Get(email, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &userClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.UserId,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *Authorization) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(accessToken), nil
	})

	if err != nil {
		return -1, err
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok {
		return -1, errors.New("uncorrected type claims")
	}

	return claims.UserId, nil
}
