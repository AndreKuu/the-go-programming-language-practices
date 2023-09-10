package main

import (
	"fmt"
	"io"
	"strings"
)

/*
# source
// A LimitedReader reads from R but limits the amount of
// data returned to just N bytes. Each call to Read
// updates N to reflect the new amount remaining.
// Read returns EOF when N <= 0 or when the underlying R returns EOF.

	type LimitedReader struct {
		R Reader // underlying reader
		N int64  // max bytes remaining
	}

	func (l *LimitedReader) Read(p []byte) (n int, err error) {
		if l.N <= 0 {
			return 0, EOF
		}
		if int64(len(p)) > l.N {
			p = p[0:l.N]
		}
		n, err = l.R.Read(p)
		l.N -= int64(n)
		return
	}
*/
type MyLimitReader struct {
	InnerReader io.Reader
	Limit       int64
	Index       int64
}

func (r *MyLimitReader) Read(b []byte) (int, error) {
	c := len(b)
	if c <= 0 {
		return 0, nil
	}
	remaining := r.Limit - r.Index
	if remaining == 0 {
		return 0, io.EOF
	}

	if int64(c) >= remaining {
		b = b[:remaining]
		n, err := r.InnerReader.Read(b)
		if err != nil && err != io.EOF {
			return n, err
		}
		r.Index = r.Limit
		return n, io.EOF
	}

	n, err := r.InnerReader.Read(b)
	r.Index += int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &MyLimitReader{r, n, 0}
}

func main() {
	stringReader := strings.NewReader("Hello go programming language")
	limitReader := LimitReader(stringReader, 7)

	for i := 0; i < 10; i++ {
		p := make([]byte, 2)
		n, err := limitReader.Read(p)
		fmt.Printf("Round: %d, read: \"%s\", bype count: %d, err: %v\n", i, string(p), n, err)
	}
}
