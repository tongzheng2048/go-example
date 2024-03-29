package chapter13

import (
	"fmt"
	"strings"
)

func Chapter13main12() {
	fmt.Println(
		// true
		strings.Contains("test", "es"),
		// 2
		strings.Count("test", "t"),
		// true
		strings.HasPrefix("test", "te"),
		// true
		strings.HasSuffix("test", "st"),
		// 1
		strings.Index("test", "e"),
		// "a-b"
		strings.Join([]string{"a", "b"}, "-"),
		// "aaaaa"
		strings.Repeat("a", 5),
		// "bbaa"
		strings.Replace("aaaa", "a", "b", 2),
		// []string{"a","b","c","d","e"}
		strings.Split("a-b-c-d-e", "-"),
		// "test"
		strings.ToLower("TEST"),
		strings.ToUpper("test"),
	)

	//arr := []byte("test")
	//str := string([]byte{'t', 'e', 's', 't'})
}
