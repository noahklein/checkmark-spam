package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestRandLine(t *testing.T) {
	f, err := os.Open(commentsFile)
	if err != nil {
		t.Fatal(err)
	}
	lineCount, err := lineCounter(f)
	if err != nil {
		t.Fatal(err)
	}

	commentCount := map[string]int{}
	for i := 0; i < 10000; i++ {

		file, err := os.Open(commentsFile)
		if err != nil {
			t.Fatal(err)
		}
		line, err := RandLine(file)
		if err != nil {
			t.Fatal(err)
		}
		commentCount[line]++

		file.Close()
	}

	if len(commentCount) != lineCount {
		t.Errorf("Unexpected comment distribution: got %v (len = %d), want %d", commentCount, len(commentCount), lineCount)
	}
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
