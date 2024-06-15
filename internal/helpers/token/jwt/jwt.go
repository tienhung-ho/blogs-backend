package jwtcus

import (
	tokenhelper "blogs/internal/helpers/token"
	usersmodel "blogs/internal/model/users"
	"fmt"
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
			Role:   "",
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Issuer:    js.Issuer,
		},
		
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(js.SecretKey))
}

// ValidateToken validates a JWT token and returns the claims
func (s *JwtServices) ValidateToken(tokenString string) (*tokenhelper.JwtPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authJwtCusClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(s.SecretKey), nil
		})

	if claims, ok := token.Claims.(*authJwtCusClaims); ok && token.Valid {
		return &claims.Payload, nil
	} else {
		return nil, err
	}
}
