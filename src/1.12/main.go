package main

import "fmt"

func uniq(in []string) []string {
	m := make(map[string]struct{}, 100)
	out := make([]string, 0)

	for _, v := range in {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			out = append(out, v)
		}
	}

	return out
}

func main() {
	in := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(uniq(in))
}
