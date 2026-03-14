package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	a := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	b := []string{"listen", "silent", "evil", "vile", "Astronomer", "Moon starer"}
	c := []string{"", "", "a", "a", "aaa", "aab", "abb", "eat", "TEA", "café", "éfac"}
	fmt.Println(GroupAnagrams(a))
	fmt.Println(GroupAnagrams(b))
	fmt.Println(GroupAnagrams(c))
}

func GroupAnagrams(words []string) map[string][]string {
	anagram := make(map[string][]string)
	for _, word := range words {
		trimmed := strings.TrimSpace(word)
		lower := strings.ToLower(trimmed)
		sortedWord := sortStrings(lower)
		if trimmed == "" {
			continue
		}
		anagram[sortedWord] = append(anagram[sortedWord], trimmed)
	}
	return anagram
}

func sortStrings(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
