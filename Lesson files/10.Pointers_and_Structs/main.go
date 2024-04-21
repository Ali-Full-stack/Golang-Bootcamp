package main

import (
	"fmt"
	"hello/person"
	"time"
)

func String(p person.Person) string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// func String(p Person) string {
// 	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
// }

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	// a := test()
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func test() string {
	return ""
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

type Score int

func (s Score) Add(val int) Score {
	return s + Score(val)
}

type HighScore Score

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// type Employee Person

type MailCategory int

const (
	Uncategorized MailCategory = iota << 3
	Personal
	Spam
	Social
	Advertisements
)

// const (
// 	Field1 = iota + 1
// 	Field2
// 	Field3
// 	Field4
// 	Field5
// 	Field6
// )

type BitField int

const (
	// 00000010
	Field1 BitField = 2 << iota
	Field2
	Field3
	Field4
	Field5
)

// 000000000 -> 0
// 000000100

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// do business logic
	return nil
}

// type Inner struct {
// 	X int
// 	Inner2
// }

// type Inner2 struct {
// 	X int
// }

// type Outer struct {
// 	Inner
// }

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	result := i.A * 2
	return i.IntPrinter(result)
}

type Outer struct {
	Inner
	S string
}

// func (o Outer) IntPrinter(val int) string {
// 	return fmt.Sprintf("Outer: %d", val)
// }

func main() {

	// p := person.Person{Age: 10, FirstName: "firstname", LastName: "lastname"}
	// p.Age = 11

	// fmt.Println(String(p))
	// var c Counter
	// fmt.Println(c.String())
	// c.Increment() is converted to (&c).Increment()
	// c.Increment()
	// fmt.Println(c.String())

	// c := &Counter{}
	// fmt.Println(c.String())
	// c.Increment()
	// c.String() is silently converted to (*c).String()
	// fmt.Println(c.String())

	// var c Counter
	// doUpdateWrong(c)
	// fmt.Println("in main:", c.String())
	// doUpdateRight(&c)
	// fmt.Println("in main:", c.String())

	// var it IntTree
	// it = it.Insert(5)
	// it = it.Insert(3)
	// it = it.Insert(10)
	// it = it.Insert(2)
	// fmt.Println(it.Contains(2))
	// fmt.Println(it.Contains(12))

	// myAdder := Adder{start: 10}
	// fmt.Println(myAdder.AddTo(5))

	// f1 := myAdder.AddTo
	// fmt.Println(f1(10))

	// f2 := Adder.AddTo
	// fmt.Println(f2(myAdder, 15))

	// var i int = 300
	// var s Score = 100
	// var hs HighScore = 200
	// hs = s
	// s = i
	// s = Score(i)
	// hs = HighScore(s)

	// var s Score = 50
	// scoreWithBonus := s + 100 // type of scoreWithBonus is Score
	// fmt.Println(scoreWithBonus)

	// hs += 2
	// fmt.Println(hs)
	// score := "sdsd"
	// var score Score = "hgfhgf"
	// fmt.Println(score.Add(10))

	// fmt.Println(Uncategorized)
	// fmt.Printf("%b\n", Uncategorized)
	// fmt.Println(Personal)
	// fmt.Printf("%b\n", Personal)
	// Personal = 1000
	// fmt.Println(Spam)
	// fmt.Printf("%b\n", Spam)
	// fmt.Println(Social)
	// fmt.Printf("%b\n", Social)
	// fmt.Println(Advertisements)
	// fmt.Printf("%b\n", Advertisements)

	// fmt.Println(Field1)
	// fmt.Println(Field2)
	// fmt.Println(Field3)
	// fmt.Println(Field4)
	// fmt.Println(Field5)

	// m := Manager{
	// 	Employee: Employee{
	// 		Name: "Bob Bobson",
	// 		ID:   "12345",
	// 	},
	// 	Reports: []Employee{},
	// }
	// fmt.Println(m.ID)
	// fmt.Println(m.Description())

	// o := Outer{
	// 	Inner: Inner{
	// 		X: 20,
	// 		Inner2: Inner2{
	// 			X: 10,
	// 		},
	// 	},
	// }
	// fmt.Println(o.X)
	// fmt.Println(o.Inner.Inner2.X)
	// m := Manager{
	// 	Employee: Employee{
	// 		Name: "Bob Bobson",
	// 		ID:   "12345",
	// 	},
	// 	Reports: []Employee{},
	// }

	// var eFail Employee = m // compilation error!
	// var eOK Employee = m.Employee

	// fmt.Println(eFail, eOK)

	// o := Outer{
	// 	Inner: Inner{
	// 		A: 10,
	// 	},
	// 	S: "Hello",
	// }
	// fmt.Println(o.Double())

	// fmt.Println(o.IntPrinter(10))
	// fmt.Println(o.Inner.IntPrinter(10))

	// var myStringer fmt.Stringer
	// var myIncrementer Incrementer

	// pointerCounter := &Counter{}
	// valueCounter := Counter{}

	// myStringer = pointerCounter
	// myStringer = valueCounter

	// var a MyInt
	// myStringer = a

	// MyInterface(a)
	// var i interface{
	NewInterface(1)
	NewInterface("asd")
	// NewInterface(a)

	// myIncrementer = pointerCounter
	// myIncrementer = valueCounter
	// fmt.Println(myStringer, myIncrementer)
}

func NewInterface(a interface{}) {
	fmt.Println(a)
}

func MyInterface(s fmt.Stringer) {
	fmt.Println("s", s.String())
}

type MyInt int

func (m MyInt) String() string {
	return string(m)
}

type Stringer interface {
	String() string
}

type Incrementer interface {
	Increment()
}
