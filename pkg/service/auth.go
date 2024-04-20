package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mbredikhin/snippets"
	"github.com/mbredikhin/snippets/pkg/repository"
)

const (
	salt       = "j1jn2i0sf0q31ncl3"
	signingKey = "Z1-&hd1(hlqFcLpqG[me1&92#"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

// AuthService - Authorization service
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService - Authorization service constructor
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser - Create new user
func (s *AuthService) CreateUser(user snippets.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// GenerateToken - Generate new JWT token
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})
	return token.SignedString([]byte(signingKey))
}

// ParseToken - JWT token parser
func (s *AuthService) ParseToken(accessToken string) (int, int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserID, claims.ExpiresAt, nil
}

// generatePasswordHash - Password hash generator
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) BlacklistToken(token string, expiresAt int64) error {
	if err := s.repo.BlacklistToken(token, expiresAt); err != nil {
		return errors.New("unable to blacklist this token")
	}
	return nil
}

func (s *AuthService) CheckIfTokenBlacklisted(token string) bool {
	return s.repo.CheckIfTokenBlacklisted(token)
}
