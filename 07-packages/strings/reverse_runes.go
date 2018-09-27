package strings

// Reverses an array of runes
// This function is not exported (It is only visible inside the `strings` package)
func reverseRunes(r []rune) []rune {
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}