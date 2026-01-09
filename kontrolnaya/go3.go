package main

import (
	"fmt"
	"strings"
)

func encryptWord(word string) string {
	if len(word) <= 1 {
		return word
	}
	first := word[:1]
	rest := []rune(word[1:])
	for i, j := 0, len(rest)-1; i < j; i, j = i+1, j-1 {
		rest[i], rest[j] = rest[j], rest[i]
	}
	return first + string(rest)
}

func encryptPhrase(phrase string) string {
	words := strings.Fields(phrase)
	for i, w := range words {
		words[i] = encryptWord(w)
	}
	return strings.Join(words, " ")
}

func main() {
	phrases := []string{
		"пепе шнеле легендо",
		"ахах пепе шнеле какой я смешной и оригинальный",
		"максим безбородов лучший преподаватель",
	}

	for _, p := range phrases {
		fmt.Println(p, "->", encryptPhrase(p))
	}
}