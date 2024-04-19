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

func Add(a,b int,)int{
	return a + b 
}

func Multiply(a,b int)int{
	return a * b
}

func Power(a int)int{
	return a * a
}

func Collect(a,b,c int, ch chan int){

	ch <- a 
	ch <- b
	ch <- c
}
