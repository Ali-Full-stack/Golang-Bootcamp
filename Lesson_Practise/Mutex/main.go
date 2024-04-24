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
	var num int
	fmt.Print("Up to Number >> ")
	fmt.Scan(&num)
	
	counter := Counter{}

	wg.Add(5)

	go counter.Increment(num)
	go counter.Increment(num)
	go counter.Increment(num)
	go counter.Increment(num)
	go counter.Increment(num)
	
	wg.Wait()
	fmt.Println(counter.Numbers)

}
