package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/assert"
) 

func Test_Search(t *testing.T)  {
	t.Run("Testing First", Test_first)
	t.Run("Testing Second", Test_second)
	t.Run("Testing Third", Test_third)
	t.Run("Testing Fourth", Test_fourth)
	t.Run("Testing Fifth", Test_fifth)
}

func Test_first(t *testing.T){
	result := Search("hello", "-n", []string{"input.txt", "greeting.txt"})
	result2 := []string{
		"input.txt:1 hello",
		"input.txt:3 hello again",
		"greeting.txt:2 hello is a file",
		"greeting.txt:3 with multiple hellos"}	
	assert.True(t, stringSlicesEqual(result, result2), "The string slices don't match")
}

func Test_second(t *testing.T){
	result := Search("hello", "-l", []string{"input.txt", "greeting.txt"})
	result2 := []string{"input.txt", "greeting.txt"}
	assert.True(t, stringSlicesEqual(result, result2), "The string slices don't match")
} 

func Test_third(t *testing.T){
	result := Search("hello", "-i", []string{"input.txt", "greeting.txt"})
	result2 := []string{
		"input.txt: hello",
		"input.txt: hello again",
		"greeting.txt: hello is a file",
		"greeting.txt: with multiple hellos"}
	assert.True(t, stringSlicesEqual(result, result2), "The string slices don't match")
}

func Test_fourth(t *testing.T){
	result := Search("hello", "-v", []string{"input.txt", "greeting.txt"})
	result2 := []string{
		"input.txt: world",
		"greeting.txt: Greetings reader"}
	assert.True(t, stringSlicesEqual(result, result2), "The string slices don't match")
}

func Test_fifth(t *testing.T){
	result := Search("hello", "-x", []string{"input.txt", "greeting.txt"})
	result2 := []string{"input.txt: hello"}
	assert.True(t, stringSlicesEqual(result, result2), "The string slices don't match")
}