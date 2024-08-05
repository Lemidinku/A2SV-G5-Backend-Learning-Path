package main

import (
	"fmt"
)


func main() {

	// test wordFreqCounter
	/*
	input := "Hello, my name is Lemi. I am Lemi. I am a software engineer. I am a backend engineer."
	freq := wordFreqCounter(input)
	fmt.Println(freq);*/

	// test isPalindrome
	s := "abba ?aBba_"
	isPalin := isPalindrome(s)
	if isPalin {
		fmt.Println(s, "is a palindrome")
	} else {
		fmt.Println(s, "is not a palindrome")
	}

}
