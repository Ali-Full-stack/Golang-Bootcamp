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
