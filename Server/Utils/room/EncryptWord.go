package roomUtils

import wordUtils "ishtaloo.io/Utils/word"

func EncryptWord(word []rune, guesses []rune) {
	visibleRunes := append(wordUtils.VisibleRunes(), guesses...)
Word:
	for wordIndex, wordValue := range word {
		for _, allowedValue := range visibleRunes {
			if wordValue == allowedValue {
				continue Word
			}
		}
		word[wordIndex] = '_'
	}
}
