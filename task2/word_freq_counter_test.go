package main

import (
	"reflect"
	"testing"
)

func TestWordFreqCounter(t *testing.T) {
	input := "Hello, world! This is a test. Hello, World!"
	expected := map[string]int{
		"hello": 2,
		"world": 2,
		"this":  1,
		"is":    1,
		"a":     1,
		"test":  1,
	}

	result := wordFreqCounter(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected word frequency %v, but got %v", expected, result)
	}
}

func TestWordFreqCounter_EmptyInput(t *testing.T) {
	input := ""
	expected := map[string]int{}

	result := wordFreqCounter(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected word frequency %v, but got %v", expected, result)
	}
}

func TestWordFreqCounter_SingleWord(t *testing.T) {
	input := "hello"
	expected := map[string]int{
		"hello": 1,
	}

	result := wordFreqCounter(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected word frequency %v, but got %v", expected, result)
	}
}