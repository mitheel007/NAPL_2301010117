package strop

func Reverse(text string) string {
	characters := []rune(text)
	for left, right := 0, len(characters)-1; left < right; left, right = left+1, right-1 {
		characters[left], characters[right] = characters[right], characters[left]
	}
	return string(characters)
}
