package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	nRead, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("bytes: could not read random bytes: %w", err)
	}
	if nRead < n {
		return nil, fmt.Errorf("bytes: could not read enough random bytes: %w", err)
	}
	return b, nil
}

func String(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("string: could not generate random string: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
