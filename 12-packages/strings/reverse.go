package strings

// Reverses a string
/*
	Since strings in Go are immutable, we first convert the string to a mutable array of runes ([]rune), 
	perform the reverse operation on that, and then re-cast to a string.
*/
func Reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}