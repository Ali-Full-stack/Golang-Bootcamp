package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

type MyInt int

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	token, err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	data, err := getData(token, file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}
	return data, nil
}

func login(uid, pwd string) (string, error) {
	// world's worst login system
	if uid == "admin" && pwd == "admin" {
		return "user:admin", nil
	}
	return "", errors.New("bad user")
}

func getData(token, file string) ([]byte, error) {
	// world's worst access control
	if token == "user:admin" {
		switch file {
		case "secrets.txt":
			return []byte("passwords aplenty!"), nil
		case "payroll.csv":
			return []byte("everyone's salary"), nil
		}
	}
	return nil, os.ErrNotExist
}

func main() {
	// var i interface{}
	// var mine MyInt = 20
	// i = mine
	// i2 := i.(int)
	// fmt.Println(i2 + 1)

	// i2 := i.(string)
	// fmt.Println(i2)

	// i2, ok := i.(string)
	// if !ok {
	// 	fmt.Printf("unexpected type for %v", i2)
	// }
	// fmt.Println(i2 + "")

	// http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {

	// })

	// numerator := 20
	// denominator := 0
	// remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(remainder, mod)

	// result, err := doubleEven(1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(result)

	// result, err = doubleEven(2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(result)

	// data := []byte("This is not a zip file")
	// notAZipFile := bytes.NewReader(data)
	// _, err := zip.NewReader(notAZipFile, int64(len(data)))
	// if err == zip.ErrFormat {
	// 	fmt.Println("Told you so")
	// }

	data, err := LoginAndGetData("admin", "admin", "secrets.txt")
	fmt.Println(string(data), err)

	data, err = LoginAndGetData("admin", "admin", "chicken_recipe.txt")
	fmt.Println(string(data), err)

	data, err = LoginAndGetData("jon", "password", "secrets.txt")
	fmt.Println(string(data), err)
}
