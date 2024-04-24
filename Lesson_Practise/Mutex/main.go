package main

import (
	"fmt"
	"sync"
)

var count int
var wg sync.WaitGroup

type Counter struct {
	Numbers []int
	mutex   sync.Mutex
}

func (c *Counter) Increment(n int) {
	for i := 0; i < n; i++ {
		c.mutex.Lock()
		if count >= n {
			c.mutex.Unlock()
			break
		}
		count++
		c.Numbers = append(c.Numbers, count)
		c.mutex.Unlock()
	}
	wg.Done()
}

func main() {
	var userInput, num int
	fmt.Print("Up to Number >> ")
	fmt.Scan(&num)
	fmt.Print("Number of goroutines >> ")
	fmt.Scan(&userInput)
	

	counter := Counter{}

	wg.Add(userInput)

	for i := 0; i < userInput; i++ {
		go counter.Increment(num)
	}
	wg.Wait()
	fmt.Println(counter.Numbers)

}
