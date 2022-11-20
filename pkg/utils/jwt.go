package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type CustomClaim struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId int, role string) (string, error) {
	secret := viper.GetString("jwt.secretKey")
	jwtDuration := time.Duration(viper.GetInt("jwt.durationMinute")) * time.Minute

	now := time.Now()

	claims := CustomClaim{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    viper.GetString("appName"),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(jwtDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func CheckToken(input string) (int, string, error) {
	secret := viper.GetString("jwt.secretKey")

	token, err := jwt.ParseWithClaims(input, &CustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*CustomClaim); ok && token.Valid {
		return claims.UserId, claims.Role, nil
	} else {
		return 0, "", err
	}
}
