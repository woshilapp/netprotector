package global

import (
	"crypto/rand"
	"math/big"
)

type Token struct {
	Token string
}

type ApiUser struct {
	Username string
	Password string
}

func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		result[i] = charset[num.Int64()]
	}

	return string(result), nil
}
