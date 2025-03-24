package services

import (
	"fmt"
	"itsware/configs"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CustomClaims defines custom token fields
type CustomClaims struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	TenantID int    `json:"tenant_id"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token with custom fields
func GenerateToken(userID, tenantID int, email, role string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		Email:    email,
		TenantID: tenantID,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(configs.AppSettings.AuthParams.JwtTtlMinutes)).Unix(),
			Issuer:    configs.AppSettings.AppParams.ServerName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

// ParseToken parses JWT token and returns custom fields
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Checking the token signature method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		fmt.Println("cannot parse token", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
