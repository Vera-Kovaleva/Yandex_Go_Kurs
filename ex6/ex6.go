package main

import (
	"fmt"
)

func main() {
	l := 100
	s := make([]int, 0, l)

	for i := 1; i <= l; i++ {
		s = append(s, i)
	}

	s = append(s[:10], s[l-10:]...)
	l = len(s)
	for i := range s[:l/2] {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
	fmt.Println(s)
}
