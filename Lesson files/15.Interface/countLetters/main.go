package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("temps.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	fileCounts, err := countLetters(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("file counts: ", fileCounts)

	s := "The quick brown fox jumped over the lazy dog"
	// fmt.Println(len(s))
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(counts)
}

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	// fmt.Println("buf", buf)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}

	}
}
