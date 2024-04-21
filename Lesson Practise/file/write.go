package file

import (
	"fmt"
)
	// Create a file called user.txt
	file ,err := os.Create("user.txt")
	if err != nil {
		fmt.Println("error creating file :",err)
		return
	}
	defer file.Close()

	var userinput string 
	// take a user input 

	fmt.Print("Enter a word >>> ")
	fmt.Scan(&userinput)
	// write user input to the  file 
	n,err := file.WriteString(userinput)
	if err != nil {
		fmt.Println("error writing to file :",err)
		return
	}
	// read a file
	fi, err :=os.ReadFile(user.txt)
	if err != nil {
		fmt.Println("error reading file :",err)
		return
	}

	// display the content of  file to the terminal
	fmt.Printf("file content : %s\n",fi)

	// open new file in append|create|readwrite mode called text.txt
	file, err := os.OpenFile("text.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("error opening file:",err)
		return
	}
	defer file.Close()
	// declare variable and assign value to it
	line := "hello world\n"
	// append variable to the file using fprintf
	n, err :=fmt.Fprintln(file, line)
	// display how many bytes appended to file
	fmt.Println("Added succesfully : total bytes",n)
