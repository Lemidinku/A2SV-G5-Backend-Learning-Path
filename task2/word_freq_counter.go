/* Write a Go function that takes a string as input and returns a dictionary containing the frequency of each
word in the string. Treat words in a case-insensitive manner and ignore punctuation marks.*/

package main

import (
	"regexp"
	"strings"
)


func wordFreqCounter(input string) map[string]int {
	words := strings.Fields(input)
	wordFreq := make(map[string]int)
	for _, word := range words {
		word = strings.ToLower(word)
		word = regexp.MustCompile(`[[:punct:]]`).ReplaceAllString(word, "")
		wordFreq[word]++
	}
	return wordFreq
}
