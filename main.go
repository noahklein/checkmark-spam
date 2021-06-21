package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

const commentsFile = "comments.txt"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	file, err := os.Open(commentsFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	comment, err := RandLine(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

// RandLine uniformly selects a random line from a reader.
func RandLine(r io.Reader) (string, error) {
	var (
		scanner = bufio.NewScanner(r)
		comment string
	)
	for i := 1; scanner.Scan(); i++ {
		// Each line has a 1/i chance of replacing the comment.
		// See https://en.wikipedia.org/wiki/Reservoir_sampling.
		if i == 1 || rand.Intn(i) == 0 {
			comment = scanner.Text()
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return comment, scanner.Err()
}
