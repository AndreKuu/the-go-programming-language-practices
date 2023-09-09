package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
)

type Word string

type WordAndRowCounter struct {
	WordsNum map[Word]int
	RowsNum  int
}

func (w *WordAndRowCounter) String() string {
	s := ""
	s += fmt.Sprintln("Number of rows:", w.RowsNum)
	s += fmt.Sprintln("Word counts:")
	for word := range w.WordsNum {
		s += fmt.Sprintf("  \"%s\": %d\n", word, w.WordsNum[word])
	}

	return s
}

func (w *WordAndRowCounter) Write(p []byte) (int, error) {
	l := len(p)

	// Set default value for RowsNum
	if l > 0 && w.RowsNum == 0 {
		w.RowsNum = 1
		w.WordsNum = map[Word]int{}
	}
	w.RowsNum += bytes.Count(p, []byte("\n"))

	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)

	// For simplification, ignore the complexity of punctuation
	for scanner.Scan() {
		word := Word(scanner.Text())
		w.WordsNum[word]++
	}
	return l, nil
}

const MOCK_CONTENT = `This is a practice of implement of the word and row counter
by interface of io.Writer with bufio.ScanWords.`

func main() {
	c := WordAndRowCounter{}

	for i := 0; i < 100; i++ {
		if _, err := fmt.Fprint(&c, MOCK_CONTENT); err != nil {
			log.Fatal(err)
		}
		fmt.Println(&c)
	}
}
