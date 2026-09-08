package strop

func CountVowels(text string) int {
	count := 0
	for _, character := range text {
		switch character {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}
