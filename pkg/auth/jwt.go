package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"ordent-test/config"
	"time"
)

func CreateJWT(userID string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    "Ordent",
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		})

	key := config.Get().JWTKey

	jwtString, err := token.SignedString([]byte(key))

	if err != nil {
		return nil, err
	}

	return &jwtString, nil
}

func ParseJWT(jwtToken string) (*string, error) {

	key := config.Get().JWTKey

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["sub"].(string)
		return &id, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
