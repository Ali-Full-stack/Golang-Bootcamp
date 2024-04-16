package main

import (
	"fmt"
	"strings"

	"github.com/Ali-Full-stack/Golang-Bootcamp/file"
)

func main() {
	s := "11 The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := file.CountLetters(sr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(counts)
}
