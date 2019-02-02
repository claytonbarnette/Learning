// Package stringutil contains utility fuctions for working with strings
// This is a Library used by hello2.go

package stringutil


func Reverse(s string) string {

	r := []rune(s)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)

}
