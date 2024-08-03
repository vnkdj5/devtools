package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"
)

// Function to get the HMAC code for the input text
func GenerateHMAC(inputText, algorithm, secretKey string) (string, error) {
	var h hash.Hash
	switch strings.ToLower(algorithm) {
	case "md5":
		h = hmac.New(md5.New, []byte(secretKey))
	case "sha1":
		h = hmac.New(sha1.New, []byte(secretKey))
	case "sha256":
		h = hmac.New(sha256.New, []byte(secretKey))
	case "sha512":
		h = hmac.New(sha512.New, []byte(secretKey))
	default:
		return "", fmt.Errorf("unsupported algorithm: %s", algorithm)
	}
	h.Write([]byte(inputText))
	return hex.EncodeToString(h.Sum(nil)), nil
}
