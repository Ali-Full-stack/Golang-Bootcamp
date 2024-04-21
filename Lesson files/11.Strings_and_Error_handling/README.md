# Reja
* [Interfaces](#interfaces)


## Interfaces
* [A Quick Lesson on Interfaces](#a-quick-lesson-on-interfaces)
* [What is a type interface?](#what-is-a-type-interface)
* [Zero value of an interface type](#zero-value-of-an-interface-type)
* [Error interface](#error-interface)
* [fmt.Stringer interface](#fmtstringer-interface)
* [sort.Interface](#sortinterface)
* [PHP and JAVA](#php-and-java)
* [Interfaces Are Type-Safe Duck Typing](#interfaces-are-type-safe-duck-typing)
* [Embedding and Interfaces](#embedding-and-interfaces)
* [Interfaces and nil](#interfaces-and-nil)
* [Interfaces Are Comparable](#interfaces-are-comparable)


### A Quick Lesson on Interfaces
```
type Stringer interface { 
	String() string
}
```

```
type Incrementer interface { 
	Increment()
}
var myStringer fmt.Stringer 
var myIncrementer Incrementer 
pointerCounter := &Counter{} 
valueCounter := Counter{}
myStringer = pointerCounter
myStringer = valueCounter
myIncrementer = pointerCounter
myIncrementer = valueCounter
```

### What is a type interface?

* Interfeys - bu xatti-harakatlar to'plamini belgilaydigan shartnoma.

* Interfeyslar software dizaynning sof ob'ekti bo'lib, ular faqat xatti-harakatlar to'plamini (ya'ni, methodlarni) aniqlaydi, bu xatti-harakatlarning amalga oshirilishini ta'minlamaydi.

* Interfeys - bu methodlar to'plamini ularni amalga oshirmasdan aniqlaydigan turdagi

“Implement” = “Method kodini yozish" 

```
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

Methodni tanasi yo'q, amalga oshirish yo'q. 
Belgilangan yagona narsa - bu method nomi va uning signature (parametrlar va natijalar).


### Zero value of an interface type
```
var r io.Reader
fmt.Println(r)
```

```
type Human struct {
    Firstname string
    Lastname string
    Age int
    Country string
}

type DomesticAnimal interface {
    ReceiveAffection(from Human)
    GiveAffection(to Human)
}
```

DomesticAnimal - bu shartnoma.

It signals to the developer that to be a DomesticAnimal we need to have at least two behaviors: 
ReceiveAffection and GiveAffection

Bu dasturchiga DomesticAnimal bo'lish uchun kamida ikkita xatti-harakatga ega bo'lishimiz kerakligini bildiradi:


```
type Cat struct {
    Name string
}

type Dog struct {
    Name string
}

func (c Cat) ReceiveAffection(from Human) {
    fmt.Printf("The cat named %s has received affection from Human named %s\n", c.Name, from.Firstname)
}

func (c Cat) GiveAffection(to Human) {
    fmt.Printf("The cat named %s has given affection to Human named %s\n", c.Name, to.Firstname)
}

func (d Dog) ReceiveAffection(from Human) {
    fmt.Printf("The dog named %s has received affection from Human named %s\n", d.Name, from.Firstname)
}

func (d Dog) GiveAffection(to Human) {
    fmt.Printf("The dog named %s has given affection to Human named %s\n", d.Name, to.Firstname)
}
```

```
func Pet(animal DomesticAnimal, human Human) {
    animal.GiveAffection(human)
    animal.ReceiveAffection(human)
}

func main() {

    // Create the Human
    var john Human
    john.Firstname = "John"


    // Create a Cat
    var c Cat
    c.Name = "Maru"

    // then a dog
    var d Dog
    d.Name = "Medor"

    Pet(c, john)
    Pet(d,john)
}
```

Standart kutubxonadan ba'zi foydali (va mashhur) interfeyslar:
### Error interface

```
type error interface {
    Error() string
}
```

```
func (c *Communicator) SendEmailAsynchronously(email *Email) error {
    //...
}
```

### fmt.Stringer interface
```
type Stringer interface {
    String() string
}
```

```
type Human struct {
    Firstname string
    Lastname string
    Age int
    Country string
}

func (h Human) String() string {
    return fmt.Sprintf("human named %s %s of age %d living in %s",h.Firstname,h.Lastname,h.Age,h.Country)
}

func main() {
    var john Human
    john.Firstname = "John"
    john.Lastname = "Doe"
    john.Country = "USA"
    john.Age = 45

    fmt.Println(john)
}
```

### sort.Interface
```
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}


type Person struct {
    Age int
}
// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }


func main() {
    people := []Person{
        {"Bob", 31},
        {"John", 42},
        {"Michael", 17},
        {"Jenny", 26},
    }

    sort.Sort(ByAge(people))
}
```


Interfeyslar odatda "er" oxiri bilan nomlanadi. Siz allaqachon fmt.Stringer-ni ko'rgansiz, 
lekin yana ko'plari bor, jumladan io.Reader, io.Closer, io.ReadCloser, json.Marshaler va http.Handler

### PHP and JAVA
```
// JAVA
public class Cat implements DomesticAnimal{
    public void receiveAffection(){
        //...
    }
    public void giveAffection(){
        //..
    }
}
```
```
//PHP
<?php

class Cat implements DomesticAnimal {
    public function receiveAffection():void {
        // ...
    }
    public function giveAffection():void {
        // ...
    }
}
?>
```

### Interfaces Are Type-Safe Duck Typing

Go interfeyslarini o'ziga xos qiladigan narsa shundaki, ular bilvosita (yashirin) amalga oshiriladi.

Python, Ruby va JavaScript kabi dinamik tarzda terilgan tillarda interfeyslar mavjud emas. 
Buning o'rniga, bu ishlab chiquvchilar duck typing usulidan foydalanadilar, 
bu "Agar u o'rdak kabi yursa va o'rdak kabi chayqalsa, bu o'rdakdir" iborasiga asoslangan. 
Kontseptsiya shundan iboratki, funktsiya o'zi kutgan chaqiruv methodini topa olsa, 
siz funktsiyaga parametr sifatida turdagi namunani o'tkazishingiz mumkin:

```
class Logic:
	def process(self, data):
		# business logic

def program(logic):
# get data from somewhere 
logic.process(data)

logicToUse = Logic()
program(logicToUse)

```

```
type LogicProvider struct {}
func (lp LogicProvider) Process(data string) string { 
	// business logic
}

type Logic interface { 
	Process(data string) string
}

type Client struct{ 
	L Logic
}

func(c Client) Program() {
	// get data from somewhere 
	c.L.Process(data)
}

func main() {
    c := Client{
        L: LogicProvider{},
    }
    c.Program()
}
```


### Embedding and Interfaces
```
// the Stringer type interface from the standard library
type Stringer interface {
    String() string
}

// A homemade interface
type DomesticAnimal interface {
    ReceiveAffection(from Human)
    GiveAffection(to Human)
    // embed the interface Stringer into the DomesticAnimal interface
    Stringer
}
```


```

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadCloser interface { 
	Reader
	Closer
}

```

### Interfaces and nil
```
type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

type Incrementer interface { 
	Increment()
}


var pointerCounter *Counter 
fmt.Println(pointerCounter == nil) // prints true 
var incrementer Incrementer 
fmt.Println(incrementer == nil) // prints true 
incrementer = pointerCounter 
fmt.Println(incrementer == nil) // prints false
```

### Interfaces Are Comparable
```
type Doubler interface { 
	Double()
}

type DoubleInt int
func (d *DoubleInt) Double() { 
	*d=*d*2
}

type DoubleIntSlice []int

func (d DoubleIntSlice) Double() {
	for i := ranged {
		d[i] = d[i] * 2 
	}
}

func DoubleAndPrint(d Doubler) {
	d.Double()
	fmt.Println(d)
}

func DoublerCompare(d1, d2 Doubler) { 
	fmt.Println(d1 == d2)
}

func main() {
	var di DoubleInt = 10
	var di2 DoubleInt = 10
	var dis = DoubleIntSlice{1, 2, 3}
	var dis2 = DoubleIntSlice{1, 2, 3}
	// false because we are comparing pointers,
	// and they point to different values
	DoublerCompare(&di, &di2)
	// false because they have different underlying types
	DoublerCompare(&di, dis)
	// triggers a panic, because the underlying types
	// match, but are a non-comparable type
	DoublerCompare(dis, dis2)
}
```