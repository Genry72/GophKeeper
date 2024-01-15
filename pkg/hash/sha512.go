package hash

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
)

// Sha512 Хеширование строки
func Sha512(src string) (string, error) {
	hash := sha512.New()

	if src == "" {
		return "", ErrEmptyPassword
	}

	_, err := io.WriteString(hash, src)
	if err != nil {
		return "", fmt.Errorf("write hash: %w", err)
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
