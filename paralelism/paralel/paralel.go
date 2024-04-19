package paralel

import (
	"fmt"
	"time"
)

func Say(s string){
	for i:=1; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(100 *time.Millisecond)
	}
}

func Sum(s []int, c chan int){
	total := 0
	for _,v :=range s{
		total +=v
	}
	c <- total
}

func Sum2(a, b int, ch chan int){
	ch <- a + b
}
