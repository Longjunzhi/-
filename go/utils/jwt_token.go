package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"pxj/CloudTravelShopApi/go/constant"
	"time"
)

type Claims struct {
	UserId string `json:"userId,omitempty"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(userId string) (tokenStr string, err error) {
	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(4 * time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString([]byte(constant.APP_NAME))
	return tokenStr, err
}

func ValidToken(encodeToken string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(encodeToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.APP_NAME), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that is not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token not active yet")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if token != nil {
		if Claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return Claims, nil
		}
	}
	return nil, errors.New("couldn't handle this token")
}
