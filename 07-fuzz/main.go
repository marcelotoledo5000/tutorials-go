package main

import "fmt"

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}

// Reverse func with bug
// func Reverse(s string) string {
// 	b := []byte(s)
// 	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
// 		b[i], b[j] = b[j], b[i]
// 	}
// 	return string(b)
// }

// Fixed the error
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
