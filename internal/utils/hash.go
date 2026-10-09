package utils

import "crypto/rand"

func GenerateShortCode(size int) (string, error) {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, size)

	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	for i, v := range b {
		b[i] = alphabet[v%byte(len(alphabet))]
	}
	
	return string(b), nil
}