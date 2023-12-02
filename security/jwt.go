package security

import (
	"fmt"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

// TODO, make this a env var
var Secret = []byte("Thisisthesupersecretsecret")

func CreateJwt(userType uint16, userEmail string, userId uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().UTC().Add(time.Hour).Unix()
	claims["userType"] = userType
	claims["email"] = userEmail
	claims["userId"] = userId
	tokenString, err := token.SignedString(Secret)
	if err != nil {
		return "", fmt.Errorf("CreateJWT: %w", err)
	}

	return tokenString, nil
}

func VerifyJWT(token string) error {
	verifiedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err != nil {
		return fmt.Errorf("unauthorized jwt")
	}

	_, ok := verifiedToken.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("unauthorized jwt")
	}

	return nil
}

func GetClaim(token string, claimName string) (string, error) {
	verifiedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return "", fmt.Errorf("invalid jwt")
	}

	claims, ok := verifiedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("could not retrieve claims")
	}

	return claims[claimName].(string), nil
}

func ExtractTokenFromHeader(header string) string {
	splitToken := strings.Split(header, "bearer")
	token := strings.Trim(splitToken[1], " ")

	return token
}
