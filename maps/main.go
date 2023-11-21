package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// Makes map m using make function
	m := make(map[string]int)

	// Splits string on delimeter of " "
	str := strings.Split(s, " ")

	//Loops around map m and store word count of string in map
	for _, v := range str {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
