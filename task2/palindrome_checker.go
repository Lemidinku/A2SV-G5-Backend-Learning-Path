/* Write a Go function that takes a string as s and checks whether it is a palindrome or not. A palindrome 
is a word, phrase, number, or other sequence of characters that reads the same forward and backward
 (ignoring spaces, punctuation, and capitalization).*/


package main

import (
	"strings"
	"regexp"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = regexp.MustCompile(`[[:punct:]]`).ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, " ", "")
	
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}