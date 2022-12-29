package security

import (
	"crypto/rand"
	"errors"
	"log"
	"math/big"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jcasanella/chat_app/model"
)

var SecretKey string

func Init(n int) {
	var err error
	SecretKey, err = GenerateRandomString(64)
	if err != nil {
		log.Fatalf("error generateRandomString. %v", err)
	}
}

// Generate a securely random string
// Will return an error, if can not be generated
func GenerateRandomString(n int) (string, error) {
	if n <= 0 {
		return "", errors.New("invalid length")
	}

	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func GenerateJWT(username string) (string, error) {
	claims := model.UserClaims{
		Name: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			Issuer:    "Chat App",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "Signing Error", err
	}

	return tokenString, nil
}

// func ParseTokenJWT(token string) erro {
// 	t, err := jwt.Parse(token, func(token *jwt.Tok en)(interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigninMethodHMAC); !ok {
// 			return fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return nil
// 	})
// }
