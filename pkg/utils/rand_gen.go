package utils

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
)

// GenerateAccessKeyID генерирует уникальный access_key_id
// Формат: "admin_key_{16 случайных символов}"
// Возвращает: строку типа "admin_key_a7b3c9d2e4f51234"
func GenerateAccessKeyID() (string, error) {
	prefix := []byte("admin_key_")
	b := make([]byte, len(prefix)+16)
	copy(b, prefix)

	for i := len(prefix); i < len(b); i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
		if err != nil {
			return "", nil
		}
		b[i] = symbols[n.Int64()]
	}

	return string(b), nil
}

// GenerateSecretKey генерирует криптостойкий secret_key
// Возвращает:
// - string - base64 строку
func GenerateSecretKey() (string, error) {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}
