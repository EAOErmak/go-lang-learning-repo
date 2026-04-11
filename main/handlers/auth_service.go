package handlers

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"go-learn/main/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	authContextUsernameKey = "username"
	authContextRoleKey     = "role"
	defaultJWTSecret       = "secret"
	jwtLifetime            = time.Hour
)

var authUser = models.User{
	Username: "admin",
	Password: mustHashPassword("1234"),
	Role:     "admin",
}

type authClaims struct {
	Username string `json:"username"`
	Role     string `json:"role,omitempty"`
	jwt.RegisteredClaims
}

func authenticateUser(username, password string) (*models.User, error) {
	if strings.TrimSpace(username) != authUser.Username {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &authUser, nil
}

func generateToken(user models.User) (string, error) {
	now := time.Now().UTC()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims{
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(jwtLifetime)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	})

	return token.SignedString(jwtSecret())
}

func parseToken(tokenString string) (*authClaims, error) {
	claims := &authClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Method.Alg())
		}

		return jwtSecret(), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid || strings.TrimSpace(claims.Username) == "" {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func jwtSecret() []byte {
	secret := strings.TrimSpace(os.Getenv("JWT_SECRET"))
	if secret == "" {
		secret = defaultJWTSecret
	}

	return []byte(secret)
}

func mustHashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}
