package crypto

import (
	"crypto/rand"
	"errors"
	"math/big"
)

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
