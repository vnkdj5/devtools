package stringutils

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
