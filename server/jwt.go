package main

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	UserName string
	jwt.StandardClaims
}

// get token
func GetToken(username string) (tokenString string, err error) {
	config := GetConfig()
	log := GetLogger()
	expiresDuration := time.Duration(config.JwtExpiresTime) * time.Second
	issuer := config.JwtIssuer
	secret := []byte(config.JwtSigningKey)

	cla := JwtClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiresDuration).Unix(), // 过期时间
			Issuer:    issuer,                                 // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	log.Info("Token", username, token)

	tokenString, err = token.SignedString(secret) // 进行签名生成对应的token
	if err != nil {
		return
	}

	return
}

// parse token
func ParseToken(tokenString string) (claims *JwtClaims, err error) {
	config := GetConfig()
	secret := []byte(config.JwtSigningKey)
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("登录过期！")
			} else {
				return nil, err
			}
		}
		return
	}

	claims, ok := token.Claims.(*JwtClaims)
	if ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}

}
