package convert

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
)

func generateRandomASCIIString(length int) (string, error) {
	result := ""
	for {
		if len(result) >= length {
			return result, nil
		}
		num, err := rand.Int(rand.Reader, big.NewInt(int64(127)))
		if err != nil {
			return "", err
		}
		n := num.Int64()
		// Make sure that the number/byte/letter is inside
		// the range of printable ASCII characters (excluding space and DEL)
		if n > 64 && n < 127 {
			result += string(n)
		}
	}
}

func generateRandomBase64String(length int) (string, error) {
	result, err := generateRandomASCIIString(length)
	return base64.URLEncoding.EncodeToString([]byte(result))[:length], err

}
