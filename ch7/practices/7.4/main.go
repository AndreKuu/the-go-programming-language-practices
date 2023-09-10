// For a complete implementation, you can refer to the built-in `strings.NewReader` function
package main

import (
	"fmt"
	"log"
)

type MyStringReader struct {
	s string
}

func (r *MyStringReader) Read(b []byte) (int, error) {
	n := copy(b, r.s)
	return n, nil
}

func NewReader(s string) *MyStringReader {
	return &MyStringReader{s}
}

func main() {
	s := "hello go programming language!"
	r := NewReader(s)

	p := make([]byte, len(s))
	if n, err := r.Read(p); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Read byte counts: %d, value: \"%s\"\n", n, string(p))
	}
}
