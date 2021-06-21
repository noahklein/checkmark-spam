package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const commentsFile = "comments.txt"

func main() {
	rand.Seed(time.Now().Unix())

	file, err := os.Open(commentsFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var (
		scanner = bufio.NewScanner(file)
		comment string
		i       int
	)
	for scanner.Scan() {
		if i == 0 || rand.Intn(i) == 0 {
			comment = scanner.Text()
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(comment)
}
