package stringutils

import (
	"encoding/json"
	"fmt"
)

// Reverse a string in place using two=pointer approach
func Reverse(s string) string {
	// Convert string to a slice of runes (Unicode code points) for proper handling
	// of multi-byte characters and preservation of encoding
	runes := []rune(s)

	// Reverse in-place using a two-pointer approach
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the reversed rune slice back to a string
	return string(runes)
}

// Helper function to convert data to string
func ConvertDataToString(data any) (string, error) {
	switch v := data.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case int:
		return fmt.Sprintf("%d", v), nil
	case float64:
		return fmt.Sprintf("%f", v), nil
	case bool:
		if v {
			return "true", nil
		}
		return "false", nil
	case map[string]interface{}:
		// Convert JSON object to string
		dataBytes, err := json.Marshal(v)
		if err != nil {
			return "", fmt.Errorf("error marshalling JSON object: %w", err)
		}
		return string(dataBytes), nil
	default:
		return "", fmt.Errorf("unsupported data type: %T", v)
	}
}
