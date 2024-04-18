package time

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

// This function take a file path as parametr and asks user to input a name and append it to given file
func AppendNewUserToFile(fPath string) {
	file, err := os.OpenFile(fPath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	var user_input string
	fmt.Print("Name >>> ")
	fmt.Scan(&user_input)

	n, err := fmt.Fprintln(file, user_input)

	if err != nil {
		fmt.Println("error appending file :", err)
		return
	}
	fmt.Println("Appended succesfully Total bytes :", n)
}

// This function takes a filepath as parametr and reads file and generate random lines from the file

func Randomizer(file string) string {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Sprintf("error :%w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var slice []string

	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}

	random := slice[rand.Intn(len(slice))]
	return random
}
