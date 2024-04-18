package auth

import (
	"ambedo-api/src/config"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

// CreateToken returns a signed token with user permissions
func CreateToken(userID uint64) (string, error) {
	perms := jwt.MapClaims{}
	perms["authorized"] = true
	perms["exp"] = time.Now().Add(time.Hour * 72).Unix()
	perms["userID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perms)

	return token.SignedString([]byte(config.DefaultSecretKey))
}

// ValidateToken checks whether the received request token is valid
func ValidateToken(r *http.Request) error {
	return nil
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	
	if len(strings.Split(token, " ")) != 2 {
		return ""
	}

	return strings.Split(token, " ")[1]
}