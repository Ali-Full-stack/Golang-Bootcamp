package json

import (
	"encoding/json"
	"fmt"
	"log"
)
// marshal 
type MyJSON struct {
	Product `json:"gghgh"`
}

type Product struct {
	ID       uint64   `json:"id"`
	Name     string   `json:"name"`
	SKU      string   `json:"sku"`
	Category Category `json:"category"`
}

type Category struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
// to print result  change the package to main or from main import file

func main() {
	p := MyJSON{Product{ID: 42, Name: "Tea Pot", SKU: "TP12", Category: Category{ID: 2, Name: "Tea"}}}

	b, err := json.MarshalIndent(p, " ", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

// unmarshal 

type MyJson struct {
	Cat `json:"cat"`
}

type Cat struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}
// to print result  change the package to main or from main import file
func main() {
	myJson := []byte(`{"cat":{ "name":"Joey", "age":8 }}`)
	c := MyJson{}
	err := json.Unmarshal(myJson, &c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c.Name)
	fmt.Println(c.Cat)
}
