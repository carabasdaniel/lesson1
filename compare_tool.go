package main

import (
	//"fmt"

	v "github.com/carabasdaniel/lesson1/validator"
)

func funnelCompare(check v.Validator, word1 string, word2 string) bool {

	for i := 0; i < len(word1); i++ {
		compareWord := word1[0:i] + word1[i+1:]
		//fmt.Println(compare_word)
		if check.Compare(compareWord, word2) {
			return true
		}
	}

	return false
}
