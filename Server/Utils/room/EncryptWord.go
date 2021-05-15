package roomUtils

func EncryptWord(word []rune, guesses []rune) {
	allowedrunes := []rune{'-', ' '}
	allowedrunes = append(allowedrunes, guesses...)
Word:
	for wordIndex, wordValue := range word {
		for _, allowedValue := range allowedrunes {
			if wordValue == allowedValue {
				continue Word
			}
		}
		word[wordIndex] = '_'
	}
}
