package base64utils

import (
	"encoding/base64"
	"fmt"
)

func Encode(data string) string {
	var src []byte = []byte(data)
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(buf, src)
	return string(buf)
}

func Decode(data string) (string, error) {
	encoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("Error occurred while decoding to plaintext. Reason: %v", err)
	}
	return string(encoded), nil
}
