package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	str := "Hello world. I have no idea what I am boing here, but nontetheless am I writing this, whilst actually being pretty excited about F1. I can't wait for the qualifying to arrive, it'll be fun. Other than that, I feel great. Slowly working on my go knowledge and learning italian. And here I am, sleepless, writing this rn. I am worried for tomorrow, of making my SO wait."

	wordCounter(str)
}

func wordCounter(str string) {
	count := map[string]int{}
	lower := strings.ToLower(str)

	words := strings.FieldsFunc(lower, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	for _, word := range words {
		count[word]++
	}

	keys := []string{}
	for k := range count {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})

	for i := 0; i < 10; i++ {
		fmt.Printf("%s:%d\n", keys[i], count[keys[i]])
	}
}
