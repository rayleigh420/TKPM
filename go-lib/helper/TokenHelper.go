package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type UserClaim struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	User_id string `json:"user_id"`
	Avatar  string	`json:"avatar"`
	jwt.StandardClaims
}

func GenerateToken(name string, email string, role string,user_id string,avatar string) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	claims := UserClaim{
		Name:  name,
		Email: email,
		Role:  role,
		User_id: user_id,
		Avatar: avatar,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}
	result, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return result, nil
}

func ValidateToken(tokenString string) (claim *UserClaim, msg string) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&UserClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*UserClaim)
	if !ok {
		msg = "the token is invalid"
		msg = err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token has expired"
		msg = err.Error()
		return
	}
	return claims, msg
}
