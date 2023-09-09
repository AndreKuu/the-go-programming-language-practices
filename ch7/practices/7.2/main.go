package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Value of pointer to int64 is the byte counts of new Writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	myWriter := MyWriter{w, 0}
	return &myWriter, &myWriter.ByteCounter
}

type MyWriter struct {
	InnerWriter io.Writer
	ByteCounter int64
}

func (w *MyWriter) Write(p []byte) (int, error) {
	l, err := w.InnerWriter.Write(p)
	if err != nil {
		return l, err
	}
	w.ByteCounter += int64(l)
	return l, nil
}

func main() {
	w := os.Stdout
	newWriter, byteCounter := CountingWriter(w)

	for i := 0; i < 100; i++ {
		if _, err := fmt.Fprintf(newWriter, "hello world"); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nByte counts: %d\n", *byteCounter)
	}
}
