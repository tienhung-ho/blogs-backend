package jwtcus

import (
	tokenhelper "blogs/internal/helpers/token"
	usersmodel "blogs/internal/model/users"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtServices struct {
	SecretKey string
	Issuer    string
}

type authJwtCusClaims struct {
	Payload tokenhelper.JwtPayload `json:"payload"`
	jwt.RegisteredClaims
}

func NewJwtServices(secretkey, issuer string) *JwtServices {
	return &JwtServices{
		SecretKey: secretkey,
		Issuer:    issuer,
	}
}

func (js *JwtServices) GenerateToken(user usersmodel.Users, duration time.Duration) (string, error) {

	claims := &authJwtCusClaims{
		Payload: tokenhelper.JwtPayload{
			UserId: user.ID,
			Role:   "user",
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Issuer:    js.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(js.SecretKey))
}

func (js *JwtServices) ValidateToken(tokenString string) (*tokenhelper.JwtPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authJwtCusClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(js.SecretKey), nil
		})

	log.Printf("Token valid: %v", token.Valid)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			log.Printf("Error parsing token: %v", err)
			return nil, jwt.ErrTokenExpired
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*authJwtCusClaims); ok && token.Valid {
		log.Printf("Claims: %v", claims)
		return &claims.Payload, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
