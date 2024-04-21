package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
	timeformat :=now.Format("2006 Mon January 02")
	fmt.Println(timeformat)
}