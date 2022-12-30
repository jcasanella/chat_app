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
	SecretKey, err = GenerateRandomString(n)
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

func GenerateJWT(username string) (tokenString string, err error) {
	claims := model.UserClaims{
		Name: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Issuer:    "Chat App",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(SecretKey))
	return
}

// func ValidateTokenJWT(signedToken string) error {
// 	t, err := jwt.ParseWithClaims(token, func(token *jwt.Token)(interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigninMethodHMAC); !ok {
// 			return fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return nil
// 	})
// }
