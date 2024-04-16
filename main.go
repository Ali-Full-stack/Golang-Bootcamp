package main

import (
	"fmt"
	"os"

	"github.com/Ali-Full-stack/Golang-Bootcamp/file"
)

func main() {
	f, err := os.Open("file/text.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// s := "11 The quick brown fox jumped over the lazy dog"
	// sr := strings.NewReader(s)
	counts, err := file.CountLetters(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(counts)
}
