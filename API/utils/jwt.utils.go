package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
    "journal/config"
)

// CreateToken generates a JWT token for a given user object and secret.
func CreateJwtToken(userObj interface{}, expiresIn *time.Duration) (string, error) {
    claims := jwt.MapClaims{}

    userMap, ok := userObj.(map[string]interface{})
    if !ok {
        return "", fmt.Errorf("userObj must be a map[string]interface{}")
    }

    for key, value := range userMap {
        claims[key] = value
    }
    
    if expiresIn != nil {
        claims["exp"] = time.Now().Add(*expiresIn).Unix()
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.JWT_SECRET_KEY))
}

// DecodeToken parses a JWT token string and returns the claims.
func DecodeToken(tokenStr string) (jwt.MapClaims, error) {
    claims := jwt.MapClaims{}
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(config.JWT_SECRET_KEY), nil
    })
    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    return claims, nil
}
