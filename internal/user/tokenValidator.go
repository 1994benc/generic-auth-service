package user

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type TokenValidator struct {
}

func (t *TokenValidator) ValidateToken(accessToken string) bool {
	log.Printf("Validating access token %s", accessToken)
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Replace this by loading in a private RSA cert for more security
	var mySigningKey = []byte(envs["AUTH_SECRET"])
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error parsing access token")
		}
		return mySigningKey, nil
	})

	if err != nil {
		log.Println(err)
		return false
	}

	return token.Valid
}
