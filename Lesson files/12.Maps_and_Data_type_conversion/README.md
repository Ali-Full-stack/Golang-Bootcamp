# Reja
* [Interfaces](#interfaces)
* [Errors](#errors)


## Interfaces
* [Type Assertions and Type Switches](#type-assertions-and-type-switches)
* [Use Type Assertions and Type Switches Sparingly](#use-type-assertions-and-type-switches-sparingly)
* [Function Types Are a Bridge to Interfaces](#function-types-are-a-bridge-to-interfaces)
* [Implicit Interfaces Make Dependency Injection Easier](#implicit-interfaces-make-dependency-injection-easier)

## Errors
* [How to Handle Errors: The Basics](#how-to-handle-errors-the-basics)
* [Use Strings for Simple Errors](#use-strings-for-simple-errors)
* [Sentinel Errors](#sentinel-errors)
* [Errors Are Values](#errors-are-values)
* [Wrapping Errors](#wrapping-errors)
* [Wrapping Multiple Errors](#wrapping-multiple-errors)


### Type Assertions and Type Switches
```
type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)

    // i2 := i.(string)
	// fmt.Println(i2)

	// i2, ok := i.(MyInt)
	// if !ok {
	// 	fmt.Printf("unexpected type for %v", i2)
	// }
	// fmt.Println(i2 + 1)
}
```

### Use Type Assertions and Type Switches Sparingly

Type assertion ning keng tarqalgan qo'llanilishidan biri interfeys ortidagi aniq turdagi boshqa interfeysni ham amalga oshirishini ko'rishdir. 
Bu ixtiyoriy interfeyslarni belgilash imkonini beradi

```
// copyBuffer is the actual implementation of Copy and CopyBuffer.
// if buf is nil, one is allocated.
func copyBuffer(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy. 
    // Avoids an allocation and a copy.
	if wt, ok := src.(io.WriterTo); ok {
		return wt.WriteTo(dst)
	}
	// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	if rt, ok := dst.(io.ReaderFrom); ok {
		return rt.ReadFrom(src)
	}
	// function continues...
}
```

### Function Types Are a Bridge to Interfaces
```
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) { 
    f(w, r)
}


```

### Implicit Interfaces Make Dependency Injection Easier

Dasturiy ta'minot muhandislari kodni ajratish haqida gapirishadi, shuning uchun dasturning turli qismlariga kiritilgan o'zgarishlar bir-biriga ta'sir qilmaydi.
Ajralishni osonlashtirish uchun ishlab chiqilgan usullardan biri Dependency Injection deb ataladi. 
Dependency Injection - bu sizning kodingiz o'z vazifasini bajarish uchun zarur bo'lgan funksionallikni aniq ko'rsatishi kerak bo'lgan tushunchadir.

```
func LogOutput(message string) { 
    fmt.Println(message)
}

type SimpleDataStore struct { 
    userData map[string]string
}


func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
    name, ok := sds.userData[userID]
    return name, ok
}

func NewSimpleDataStore() SimpleDataStore { 
    return SimpleDataStore{
        userData: map[string]string{
            "1": "Fred",
            "2": "Mary",
            "3": "Pat",
        }, 
    }
}

type DataStore interface { 
    UserNameForID(userID string) (string, bool)
}

type Logger interface { 
    Log(message string)
}


type LoggerAdapter func(message string)
func (lg LoggerAdapter) Log(message string) { 
    lg(message)
}

type SimpleLogic struct { 
    l Logger
    ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) { 
    sl.l.Log("in SayHello for " + userID)
    name, ok := sl.ds.UserNameForID(userID)
    if !ok {
        return "", errors.New("unknown user") 
    }
    return "Hello, " + name, nil 
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
    sl.l.Log("in SayGoodbye for " + userID) 
    name, ok := sl.ds.UserNameForID(userID)
    if !ok {
        return "", errors.New("unknown user") 
    }
    return "Goodbye, " + name, nil
}

// factory method
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic { 
    return SimpleLogic{
        l: l,
        ds: ds, 
    }
}

```

## Errors

### How to Handle Errors: The Basics
```
func calcRemainderAndMod(numerator, denominator int) (int, int, error) { 
    if denominator == 0 {
        return 0, 0, errors.New("denominator is 0") 
    }
    return numerator / denominator, numerator % denominator, nil 
}


func main() { 
    numerator := 20
    denominator := 3
    remainder, mod, err := calcRemainderAndMod(numerator, denominator) 
    if err! = nil{
        fmt.Println(err)
        os.Exit(1) 
    }
    fmt.Println(remainder, mod)
}
```

```
type error interface { 
    Error() string
}
```

### Use Strings for Simple Errors
```
func doubleEven(i int) (int, error) { 
    if i%2! = 0 {
        return 0, fmt.Errorf("%d isn't an even number", i) 
    }
    return i * 2, nil 
}

func main() {
    result, err := doubleEven(1) 
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(result)   
}
```

### Sentinel Errors

([archive/zip](https://pkg.go.dev/archive/zip#pkg-variables))
```
func main() {
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("Told you so")
	}
}
```

### Errors Are Values
```
type Status int
const (
    InvalidLogin Status = iota + 1 
    NotFound
)

type StatusErr struct { 
    Status Status 
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
	data, err := LoginAndGetData("admin", "admin", "secrets.txt")
	fmt.Println(string(data), err)

	data, err = LoginAndGetData("admin", "admin", "chicken_recipe.txt")
	fmt.Println(string(data), err)

	data, err = LoginAndGetData("jon", "password", "secrets.txt")
	fmt.Println(string(data), err)
}
```

### Wrapping Errors
```
func Filechecker(fileName string) error {
    f, err := os.Open(fileName)
    if err != nil {
        return fmt.Errorf("in Filechecker: %w", err)
    }
    f.Close()
    return nil
}

func main() {
    err := fileChecker("not_a_file.txt")
    if err != nil {
        fmt.Println(err)
        wrappedErr := errors.Unwrap(err)
        if wrapperErr != nil {
            fmt.Println(wrappedErr)
        }
    }
}
```

### Wrapping Multiple Errors
```
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func ValidatePerson(p Person) error {
	var errs []error
	if len(p.FirstName) == 0 {
		errs = append(errs, errors.New("field FirstName cannot be empty"))
	}
	if len(p.LastName) == 0 {
		errs = append(errs, errors.New("field LastName cannot be empty"))
	}
	if p.Age < 0 {
		errs = append(errs, errors.New("field Age cannot be negative"))
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}


func main() {
	person1 := Person{
		FirstName: "Hello",
		LastName:  "World",
		Age:       1,
	}

	person2 := Person{
		FirstName: "Hello",
		Age:       100,
	}

	person3 := Person{}

	people := []Person{person1, person2, person3}
	for _, person := range people {
		err := ValidatePerson(person)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Println()
	}
}
```







