package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Create("test.csv")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// _ = f

	f, err := os.OpenFile("test.csv", os.O_RDWR|os.O_TRUNC|os.O_EXCL, 666)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(f)

	info, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(info.Mode())

	b, err := os.ReadFile("temp.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))

	err = os.WriteFile("temp.txt", []byte("new data"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

}
