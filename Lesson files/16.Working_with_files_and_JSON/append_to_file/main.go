package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Europe/Berlin")

	newTime, err := time.ParseInLocation(time.RFC3339, "Jul 9, 2012 at 5:02am (CEST)", loc)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(loc.String())
	fmt.Println(newTime)

	now := time.Now()
	fmt.Printf("%s\n", now)

	n := now.Format(time.DateTime)
	fmt.Println(n)
	// time.Sleep(2 * time.Second)

	f, err := os.OpenFile("lines.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")

	fmt.Println(time.Since(now))
}
