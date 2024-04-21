package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type MyJSON struct {
	Product `json:"product"`
}

type Product struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name"`
	SKU      string    `json:"sku,omitempty"`
	Category Category  `json:"category"`
	Time     time.Time `json:"time"`
}

type Category struct {
	ID   uint64 `json:"id"`
	Name string `json:"-"`
}

func main() {
	t := time.Now()
	sl := []MyJSON{
		MyJSON{Product{ID: 100, Name: "Tea Pot", Time: t, Category: Category{ID: 2, Name: "Tea"}}},
		MyJSON{Product{ID: 1, Name: "Tea Pot", Time: t, Category: Category{ID: 2, Name: "Tea"}}},
	}

	// p := MyJSON{Product{ID: 100, Name: "Tea Pot", Time: t, Category: Category{ID: 2, Name: "Tea"}}}

	// p := Product{ID: 32}
	// t := reflect.TypeOf(p)
	// for i := 0; i < t.NumField(); i++ {
	// 	fmt.Printf("field Name : %s\n", t.Field(i).Name)
	// 	fmt.Printf("field Tag : %s\n", t.Field(i).Tag)
	// }

	b, err := json.MarshalIndent(sl, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))

	f, err := os.OpenFile("test.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	// f, err := os.Create("test.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, string(b))

	// var m map[string]any
	// err = json.Unmarshal(b, &m)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(m)

	// b, err := json.MarshalIndent(p, "", " ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(p.Category)
	// fmt.Println(string(b))

}
