package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	const jsonStream = `[
	{"Name": "Ed", "Text": "Knock knock."},
	{"Name": "Sam", "Text": "Who's there?"},
	{"Name": "Ed", "Text": "Go fmt."},
	{"Name": "Sam", "Text": "Go fmt who?"},
	{"Name": "Ed", "Text": "Go fmt yourself!"}]
`
	type Message struct {
		Name, Text string
	}
	var slice []Message
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	err := dec.Decode(&slice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(slice)
}
