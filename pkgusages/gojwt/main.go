package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims) // claims is an empty map.
	claims["username"] = "john.doe"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	fmt.Println(claims) // claims has username and exp columns already.
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println("Error generating token:", err)
	}

	fmt.Println("JWT Token:", tokenString)
}
