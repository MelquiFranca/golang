package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := map[string]int{}
	wordList := strings.Fields(s)
	for _, value := range wordList {
		if _, ok := m[value]; ok {
			m[value] += 1
		} else {
			m[value] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
