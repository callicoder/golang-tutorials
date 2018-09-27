package strings

// Reverses a string
/*
	Since strings in Go are immutable, we first convert the string to a mutable array of runes ([]rune), 
	perform the reverse operation on that, and then re-cast to a string.
*/
func Reverse(s string) string {
	runes := []rune(s)
	reversedRunes := reverseRunes(runes)
	return string(reversedRunes)
}