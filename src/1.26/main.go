package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	seen := make(map[rune]bool, len(s))
	s = strings.ToLower(s)

	for _, v := range s {
		if seen[v] {
			return false
		}
		seen[v] = true
	}
	return true
}

func main() {
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aA"))
}
