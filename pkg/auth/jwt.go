package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"ordent-test/config"
	"ordent-test/internal/domain/model"
	"time"
)

func CreateJWT(user *model.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"issuer":   user.Username,
			"sub":      user.SecureID,
			"exp":      time.Now().Add(time.Hour).Unix(),
			"username": user.Username,
			"email":    user.Email,
		})

	key := config.Get().JWTKey

	jwtString, err := token.SignedString([]byte(key))

	if err != nil {
		return nil, err
	}

	return &jwtString, nil
}

func ParseJWT(jwtToken string) (*model.User, error) {
	key := config.Get().JWTKey

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := &model.User{
			Username: claims["username"].(string),
			Email:    claims["email"].(string),
			SecureID: claims["sub"].(string),
		}

		return user, nil
	}

	return nil, fmt.Errorf("invalid token")
}
